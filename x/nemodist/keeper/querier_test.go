package keeper_test

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	abci "github.com/tendermint/tendermint/abci/types"

// 	"github.com/incubus-network/nemo/x/nemodist/keeper"
// 	"github.com/incubus-network/nemo/x/nemodist/types"
// )

// func (suite *KeeperTestSuite) TestQuerierGetParams() {
// 	querier := keeper.NewQuerier(suite.keeper)
// 	bz, err := querier(suite.ctx, []string{types.QueryGetParams}, abci.RequestQuery{})
// 	suite.Require().NoError(err)
// 	suite.NotNil(bz)

// 	testParams := types.NewParams(true, testPeriods)
// 	var p types.Params
// 	suite.Nil(types.ModuleCdc.UnmarshalJSON(bz, &p))
// 	suite.Require().Equal(testParams, p)
// }

// func (suite *KeeperTestSuite) TestQuerierGetBalance() {
// 	sk := suite.supplyKeeper

// 	sk.MintCoins(suite.ctx, types.NemoDistMacc, sdk.NewCoins(sdk.NewCoin("unemo", sdkmath.NewInt(100e6))))

// 	querier := keeper.NewQuerier(suite.keeper)
// 	bz, err := querier(suite.ctx, []string{types.QueryGetBalance}, abci.RequestQuery{})
// 	suite.Require().NoError(err)
// 	suite.Require().NotNil(bz)

// 	var coins sdk.Coins
// 	types.ModuleCdc.UnmarshalJSON(bz, &coins)
// 	suite.Require().Equal(sdkmath.NewInt(100e6), coins.AmountOf("unemo"))
// }