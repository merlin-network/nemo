package nemodist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlin-network/nemo/x/nemodist/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}
