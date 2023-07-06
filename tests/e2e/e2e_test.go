package e2e_test

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	emtypes "github.com/evmos/ethermint/types"

	"github.com/merlin-network/nemo/app"
	"github.com/merlin-network/nemo/tests/e2e/testutil"
	"github.com/merlin-network/nemo/tests/util"
)

var (
	minEvmGasPrice = big.NewInt(1e10) // afury
)

func ufury(amt int64) sdk.Coin {
	return sdk.NewCoin("ufury", sdkmath.NewInt(amt))
}

type IntegrationTestSuite struct {
	testutil.E2eTestSuite
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

// example test that queries nemo via SDK and EVM
func (suite *IntegrationTestSuite) TestChainID() {
	expectedEvmNetworkId, err := emtypes.ParseChainID(suite.Nemo.ChainId)
	suite.NoError(err)

	// EVM query
	evmNetworkId, err := suite.Nemo.EvmClient.NetworkID(context.Background())
	suite.NoError(err)
	suite.Equal(expectedEvmNetworkId, evmNetworkId)

	// SDK query
	nodeInfo, err := suite.Nemo.Tm.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	suite.NoError(err)
	suite.Equal(suite.Nemo.ChainId, nodeInfo.DefaultNodeInfo.Network)
}

// example test that funds a new account & queries its balance
func (suite *IntegrationTestSuite) TestFundedAccount() {
	funds := ufury(1e3)
	acc := suite.Nemo.NewFundedAccount("example-acc", sdk.NewCoins(funds))

	// check that the sdk & evm signers are for the same account
	suite.Equal(acc.SdkAddress.String(), util.EvmToSdkAddress(acc.EvmAddress).String())
	suite.Equal(acc.EvmAddress.Hex(), util.SdkToEvmAddress(acc.SdkAddress).Hex())

	// check balance via SDK query
	res, err := suite.Nemo.Bank.Balance(context.Background(), banktypes.NewQueryBalanceRequest(
		acc.SdkAddress, "ufury",
	))
	suite.NoError(err)
	suite.Equal(funds, *res.Balance)

	// check balance via EVM query
	afuryBal, err := suite.Nemo.EvmClient.BalanceAt(context.Background(), acc.EvmAddress, nil)
	suite.NoError(err)
	suite.Equal(funds.Amount.MulRaw(1e12).BigInt(), afuryBal)
}

// example test that signs & broadcasts an EVM tx
func (suite *IntegrationTestSuite) TestTransferOverEVM() {
	// fund an account that can perform the transfer
	initialFunds := ufury(1e6) // 1 NEMO
	acc := suite.Nemo.NewFundedAccount("evm-test-transfer", sdk.NewCoins(initialFunds))

	// get a rando account to send nemo to
	randomAddr := app.RandomAddress()
	to := util.SdkToEvmAddress(randomAddr)

	// example fetching of nonce (account sequence)
	nonce, err := suite.Nemo.EvmClient.PendingNonceAt(context.Background(), acc.EvmAddress)
	suite.NoError(err)
	suite.Equal(uint64(0), nonce) // sanity check. the account should have no prior txs

	// transfer nemo over EVM
	nemoToTransfer := big.NewInt(1e17) // .1 NEMO; afury has 18 decimals.
	req := util.EvmTxRequest{
		Tx:   ethtypes.NewTransaction(nonce, to, nemoToTransfer, 1e5, minEvmGasPrice, nil),
		Data: "any ol' data to track this through the system",
	}
	res := acc.SignAndBroadcastEvmTx(req)
	suite.NoError(res.Err)
	suite.Equal(ethtypes.ReceiptStatusSuccessful, res.Receipt.Status)

	// evm txs refund unused gas. so to know the expected balance we need to know how much gas was used.
	ufuryUsedForGas := sdkmath.NewIntFromBigInt(minEvmGasPrice).
		Mul(sdkmath.NewIntFromUint64(res.Receipt.GasUsed)).
		QuoRaw(1e12) // convert afury to ufury

	// expect (9 - gas used) NEMO remaining in account.
	balance := suite.Nemo.QuerySdkForBalances(acc.SdkAddress)
	suite.Equal(sdkmath.NewInt(9e5).Sub(ufuryUsedForGas), balance.AmountOf("ufury"))
}

// TestIbcTransfer transfers NEMO from the primary nemo chain (suite.Nemo) to the ibc chain (suite.Ibc).
// Note that because the IBC chain also runs nemo's binary, this tests both the sending & receiving.
func (suite *IntegrationTestSuite) TestIbcTransfer() {
	suite.SkipIfIbcDisabled()

	// ARRANGE
	// setup nemo account
	funds := ufury(1e5) // .1 NEMO
	nemoAcc := suite.Nemo.NewFundedAccount("ibc-transfer-nemo-side", sdk.NewCoins(funds))
	// setup ibc account
	ibcAcc := suite.Ibc.NewFundedAccount("ibc-transfer-ibc-side", sdk.NewCoins())

	gasLimit := int64(2e5)
	fee := ufury(200)

	fundsToSend := ufury(5e4) // .005 NEMO
	transferMsg := ibctypes.NewMsgTransfer(
		testutil.IbcPort,
		testutil.IbcChannel,
		fundsToSend,
		nemoAcc.SdkAddress.String(),
		ibcAcc.SdkAddress.String(),
		ibcclienttypes.NewHeight(0, 0), // timeout height disabled when 0
		uint64(time.Now().Add(30*time.Second).UnixNano()),
		"",
	)
	// initial - sent - fee
	expectedSrcBalance := funds.Sub(fundsToSend).Sub(fee)

	// ACT
	// IBC transfer from nemo -> ibc
	transferTo := util.NemoMsgRequest{
		Msgs:      []sdk.Msg{transferMsg},
		GasLimit:  uint64(gasLimit),
		FeeAmount: sdk.NewCoins(fee),
		Memo:      "sent from Nemo!",
	}
	res := nemoAcc.SignAndBroadcastNemoTx(transferTo)

	// ASSERT
	suite.NoError(res.Err)

	// the balance should be deducted from nemo account
	suite.Eventually(func() bool {
		balance := suite.Nemo.QuerySdkForBalances(nemoAcc.SdkAddress)
		return balance.AmountOf("ufury").Equal(expectedSrcBalance.Amount)
	}, 10*time.Second, 1*time.Second)

	// expect the balance to be transferred to the ibc chain!
	suite.Eventually(func() bool {
		balance := suite.Ibc.QuerySdkForBalances(ibcAcc.SdkAddress)
		found := false
		for _, c := range balance {
			// find the ibc denom coin
			if strings.HasPrefix(c.Denom, "ibc/") {
				suite.Equal(fundsToSend.Amount, c.Amount)
				found = true
			}
		}
		return found
	}, 15*time.Second, 1*time.Second)
}
