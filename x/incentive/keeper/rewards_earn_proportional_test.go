package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/incubus-network/nemo/x/incentive/keeper"
	"github.com/incubus-network/nemo/x/incentive/types"
	"github.com/stretchr/testify/require"
)

func TestGetProportionalRewardPeriod(t *testing.T) {
	tests := []struct {
		name                  string
		giveRewardPeriod      types.MultiRewardPeriod
		giveTotalBnemoSupply  sdkmath.Int
		giveSingleBnemoSupply sdkmath.Int
		wantRewardsPerSecond  sdk.DecCoins
	}{
		{
			"full amount",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("unemo", 100), c("hard", 200)),
			),
			i(100),
			i(100),
			toDcs(c("unemo", 100), c("hard", 200)),
		},
		{
			"3/4 amount",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("unemo", 100), c("hard", 200)),
			),
			i(10_000000),
			i(7_500000),
			toDcs(c("unemo", 75), c("hard", 150)),
		},
		{
			"half amount",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("unemo", 100), c("hard", 200)),
			),
			i(100),
			i(50),
			toDcs(c("unemo", 50), c("hard", 100)),
		},
		{
			"under 1 unit",
			types.NewMultiRewardPeriod(
				true,
				"",
				time.Time{},
				time.Time{},
				cs(c("unemo", 100), c("hard", 200)),
			),
			i(1000), // total bnemo
			i(1),    // bnemo supply of this specific vault
			dcs(dc("unemo", "0.1"), dc("hard", "0.2")), // rewards per second rounded to 0 if under 1unemo/1hard
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rewardsPerSecond := keeper.GetProportionalRewardsPerSecond(
				tt.giveRewardPeriod,
				tt.giveTotalBnemoSupply,
				tt.giveSingleBnemoSupply,
			)

			require.Equal(t, tt.wantRewardsPerSecond, rewardsPerSecond)
		})
	}
}