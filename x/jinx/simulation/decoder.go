package simulation

import (
	"bytes"
	"fmt"

	"github.com/tendermint/tendermint/libs/kv"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/merlin-network/nemo/x/jinx/types"
)

// DecodeStore unmarshals the KVPair's Value to the corresponding jinx type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.Equal(kvA.Key[:1], types.DepositsKeyPrefix):
		var depA, depB types.Deposit
		cdc.MustUnmarshalBinaryBare(kvA.Value, &depA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &depB)
		return fmt.Sprintf("%s\n%s", depA, depB)
	default:
		panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
	}
}