package util_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/merlin-network/nemo/app"
	"github.com/merlin-network/nemo/tests/util"
)

func TestAddressConversion(t *testing.T) {
	app.SetSDKConfig()
	bech32Addr := sdk.MustAccAddressFromBech32("fury17d2wax0zhjrrecvaszuyxdf5wcu5a0p4vpmrwu")
	hexAddr := common.HexToAddress("0xf354ee99e2bc863cE19d80b843353476394EbC35")
	require.Equal(t, bech32Addr, util.EvmToSdkAddress(hexAddr))
	require.Equal(t, hexAddr, util.SdkToEvmAddress(bech32Addr))
}
