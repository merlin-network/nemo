package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/incubus-network/nemo/app"
	"github.com/incubus-network/nemo/x/liquid/types"
	"github.com/stretchr/testify/require"
)

func TestParseLiquidStakingTokenDenom(t *testing.T) {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	tests := []struct {
		name        string
		giveDenom   string
		wantAddress sdk.ValAddress
		wantErr     error
	}{
		{
			name:        "valid denom",
			giveDenom:   "bfury-furyvaloper1ze7y9qwdddejmy7jlw4cymqqlt2wh05yh6yyun",
			wantAddress: mustValAddressFromBech32("furyvaloper1ze7y9qwdddejmy7jlw4cymqqlt2wh05yh6yyun"),
			wantErr:     nil,
		},
		{
			name:        "invalid prefix",
			giveDenom:   "ufury-furyvaloper1ze7y9qwdddejmy7jlw4cymqqlt2wh05yh6yyun",
			wantAddress: mustValAddressFromBech32("furyvaloper1ze7y9qwdddejmy7jlw4cymqqlt2wh05yh6yyun"),
			wantErr:     fmt.Errorf("invalid denom prefix, expected %s, got %s", types.DefaultDerivativeDenom, "ufury"),
		},
		{
			name:        "invalid validator address",
			giveDenom:   "bfury-||| furyvaloper1ze7y9qw",
			wantAddress: sdk.ValAddress{},
			wantErr:     fmt.Errorf("invalid denom validator address: decoding bech32 failed: invalid checksum"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := types.ParseLiquidStakingTokenDenom(tt.giveDenom)

			if tt.wantErr != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantAddress, addr)
			}
		})
	}
}
