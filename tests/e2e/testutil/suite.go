package testutil

import (
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlin-network/nemo/app"
	"github.com/merlin-network/nemo/tests/e2e/runner"
	"github.com/merlin-network/nemo/tests/util"
)

const (
	FundedAccountName = "whale"
	// use coin type 60 so we are compatible with accounts from `nemo add keys --eth <name>`
	// these accounts use the ethsecp256k1 signing algorithm that allows the signing client
	// to manage both sdk & evm txs.
	Bip44CoinType = 60

	IbcPort    = "transfer"
	IbcChannel = "channel-0"
)

// DeployedErc20 is a type that wraps the details of the pre-deployed erc20 used by the e2e test suite.
// The Address comes from SuiteConfig.NemoErc20Address
// The CosmosDenom is fetched from the EnabledConversionPairs param of x/evmutil.
// The tests expect the following:
// - the funded account has a nonzero balance of the erc20
// - the erc20 is enabled for conversion to sdk.Coin
// - the corresponding sdk.Coin is enabled as an earn vault denom
// These requirements are checked in InitNemoEvmData().
type DeployedErc20 struct {
	Address     common.Address
	CosmosDenom string
}

// E2eTestSuite is a testify test suite for running end-to-end integration tests on Nemo.
type E2eTestSuite struct {
	suite.Suite

	config SuiteConfig
	runner runner.NodeRunner

	Nemo *Chain
	Ibc  *Chain

	UpgradeHeight int64
	DeployedErc20 DeployedErc20

	cost          costSummary
	enableRefunds bool
}

// costSummary wraps info about what funds get irrecoverably spent by the test suite run
type costSummary struct {
	sdkAddress string
	evmAddress string

	erc20BalanceBefore *big.Int
	erc20BalanceAfter  *big.Int

	sdkBalanceBefore sdk.Coins
	sdkBalanceAfter  sdk.Coins
}

// String implements fmt.Stringer
func (s costSummary) String() string {
	before := sdk.NewCoins(s.sdkBalanceBefore...).Add(sdk.NewCoin("erc20", sdkmath.NewIntFromBigInt(s.erc20BalanceBefore)))
	after := sdk.NewCoins(s.sdkBalanceAfter...).Add(sdk.NewCoin("erc20", sdkmath.NewIntFromBigInt(s.erc20BalanceAfter)))
	cost, _ := before.SafeSub(after...)
	return fmt.Sprintf("Cost Summary for %s (%s):\nbefore:\n%s\nafter:\n%s\ncost:\n%s\n",
		s.sdkAddress, s.evmAddress, util.PrettyPrintCoins(before), util.PrettyPrintCoins(after), util.PrettyPrintCoins(cost),
	)
}

// SetupSuite is run before all tests. It initializes chain connections and sets up the
// account used for funding accounts in the tests.
func (suite *E2eTestSuite) SetupSuite() {
	var err error
	fmt.Println("setting up test suite.")
	app.SetSDKConfig()

	suiteConfig := ParseSuiteConfig()
	suite.config = suiteConfig
	suite.DeployedErc20 = DeployedErc20{
		Address: common.HexToAddress(suiteConfig.NemoErc20Address),
		// Denom is fetched in InitNemoEvmData()
	}

	// setup the correct NodeRunner for the given config
	if suiteConfig.Nmtool != nil {
		suite.runner = suite.SetupNmtoolNodeRunner()
	} else if suiteConfig.LiveNetwork != nil {
		suite.runner = suite.SetupLiveNetworkNodeRunner()
	} else {
		panic("expected either nmtool or live network configs to be defined")
	}

	chains := suite.runner.StartChains()
	nemochain := chains.MustGetChain("nemo")
	suite.Nemo, err = NewChain(suite.T(), nemochain, suiteConfig.FundedAccountMnemonic)
	if err != nil {
		suite.runner.Shutdown()
		suite.T().Fatalf("failed to create nemo chain querier: %s", err)
	}

	if suiteConfig.IncludeIbcTests {
		ibcchain := chains.MustGetChain("ibc")
		suite.Ibc, err = NewChain(suite.T(), ibcchain, suiteConfig.FundedAccountMnemonic)
		if err != nil {
			suite.runner.Shutdown()
			suite.T().Fatalf("failed to create ibc chain querier: %s", err)
		}
	}

	suite.InitNemoEvmData()

	whale := suite.Nemo.GetAccount(FundedAccountName)
	suite.cost = costSummary{
		sdkAddress:         whale.SdkAddress.String(),
		evmAddress:         whale.EvmAddress.Hex(),
		sdkBalanceBefore:   suite.Nemo.QuerySdkForBalances(whale.SdkAddress),
		erc20BalanceBefore: suite.Nemo.GetErc20Balance(suite.DeployedErc20.Address, whale.EvmAddress),
	}
}

// TearDownSuite is run after all tests have run.
// In the event of a panic during the tests, it is run after testify recovers.
func (suite *E2eTestSuite) TearDownSuite() {
	fmt.Println("tearing down test suite.")

	whale := suite.Nemo.GetAccount(FundedAccountName)

	if suite.enableRefunds {
		suite.cost.sdkBalanceAfter = suite.Nemo.QuerySdkForBalances(whale.SdkAddress)
		suite.cost.erc20BalanceAfter = suite.Nemo.GetErc20Balance(suite.DeployedErc20.Address, whale.EvmAddress)
		fmt.Println("==BEFORE REFUNDS==")
		fmt.Println(suite.cost)

		fmt.Println("attempting to return all unused funds")
		suite.Nemo.ReturnAllFunds()

		fmt.Println("==AFTER REFUNDS==")
	}

	// calculate & output cost summary for funded account
	suite.cost.sdkBalanceAfter = suite.Nemo.QuerySdkForBalances(whale.SdkAddress)
	suite.cost.erc20BalanceAfter = suite.Nemo.GetErc20Balance(suite.DeployedErc20.Address, whale.EvmAddress)
	fmt.Println(suite.cost)

	// TODO: track asset denoms & then return all funds to initial funding account.

	// close all account request channels
	suite.Nemo.Shutdown()
	if suite.Ibc != nil {
		suite.Ibc.Shutdown()
	}
	// gracefully shutdown docker container(s)
	suite.runner.Shutdown()
}

// SetupNmtoolNodeRunner is a helper method for building a NmtoolRunnerConfig from the suite config.
func (suite *E2eTestSuite) SetupNmtoolNodeRunner() *runner.NmtoolRunner {
	// upgrade tests are only supported on nmtool networks
	suite.UpgradeHeight = suite.config.Nmtool.NemoUpgradeHeight
	suite.enableRefunds = false

	runnerConfig := runner.NmtoolRunnerConfig{
		NemoConfigTemplate: suite.config.Nmtool.NemoConfigTemplate,

		IncludeIBC: suite.config.IncludeIbcTests,
		ImageTag:   "local",

		EnableAutomatedUpgrade:  suite.config.Nmtool.IncludeAutomatedUpgrade,
		NemoUpgradeName:         suite.config.Nmtool.NemoUpgradeName,
		NemoUpgradeHeight:       suite.config.Nmtool.NemoUpgradeHeight,
		NemoUpgradeBaseImageTag: suite.config.Nmtool.NemoUpgradeBaseImageTag,

		SkipShutdown: suite.config.SkipShutdown,
	}

	return runner.NewNmtoolRunner(runnerConfig)
}

// SetupLiveNetworkNodeRunner is a helper method for building a LiveNodeRunner from the suite config.
func (suite *E2eTestSuite) SetupLiveNetworkNodeRunner() *runner.LiveNodeRunner {
	// live network setup doesn't presently support ibc
	if suite.config.IncludeIbcTests {
		panic("ibc tests not supported for live network configuration")
	}
	suite.enableRefunds = true

	runnerConfig := runner.LiveNodeRunnerConfig{
		NemoRpcUrl:    suite.config.LiveNetwork.NemoRpcUrl,
		NemoGrpcUrl:   suite.config.LiveNetwork.NemoGrpcUrl,
		NemoEvmRpcUrl: suite.config.LiveNetwork.NemoEvmRpcUrl,
	}

	return runner.NewLiveNodeRunner(runnerConfig)
}

// SkipIfIbcDisabled should be called at the start of tests that require IBC.
// It gracefully skips the current test if IBC tests are disabled.
func (suite *E2eTestSuite) SkipIfIbcDisabled() {
	if !suite.config.IncludeIbcTests {
		suite.T().SkipNow()
	}
}

// SkipIfUpgradeDisabled should be called at the start of tests that require automated upgrades.
// It gracefully skips the current test if upgrades are dissabled.
// Note: automated upgrade tests are currently only enabled for Nmtool suite runs.
func (suite *E2eTestSuite) SkipIfUpgradeDisabled() {
	if suite.config.Nmtool != nil && suite.config.Nmtool.IncludeAutomatedUpgrade {
		suite.T().SkipNow()
	}
}

// NemoHomePath returns the OS-specific filepath for the nemo home directory
// Assumes network is running with nmtool installed from the sub-repository in tests/e2e/nmtool
func (suite *E2eTestSuite) NemoHomePath() string {
	return filepath.Join("nmtool", "full_configs", "generated", "nemo", "initstate", ".nemo")
}

// BigIntsEqual is a helper method for comparing the equality of two big ints
func (suite *E2eTestSuite) BigIntsEqual(expected *big.Int, actual *big.Int, msg string) {
	suite.Truef(expected.Cmp(actual) == 0, "%s (expected: %s, actual: %s)", msg, expected.String(), actual.String())
}
