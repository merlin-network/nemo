package types_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/incubus-network/nemo/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"
)

func TestParams_UnmarshalJSON(t *testing.T) {
	pools := types.NewAllowedPools(
		types.NewAllowedPool("jinx", "ufury"),
		types.NewAllowedPool("jinx", "usdx"),
	)
	poolData, err := json.Marshal(pools)
	require.NoError(t, err)

	fee, err := sdk.NewDecFromStr("0.5")
	require.NoError(t, err)
	feeData, err := json.Marshal(fee)
	require.NoError(t, err)

	data := fmt.Sprintf(`{
	"allowed_pools": %s,
	"swap_fee": %s
}`, string(poolData), string(feeData))

	var params types.Params
	err = json.Unmarshal([]byte(data), &params)
	require.NoError(t, err)

	assert.Equal(t, pools, params.AllowedPools)
	assert.Equal(t, fee, params.SwapFee)
}

func TestParams_MarshalYAML(t *testing.T) {
	pools := types.NewAllowedPools(
		types.NewAllowedPool("jinx", "ufury"),
		types.NewAllowedPool("jinx", "usdx"),
	)
	fee, err := sdk.NewDecFromStr("0.5")
	require.NoError(t, err)

	p := types.Params{
		AllowedPools: pools,
		SwapFee:      fee,
	}

	data, err := yaml.Marshal(p)
	require.NoError(t, err)

	var params map[string]interface{}
	err = yaml.Unmarshal(data, &params)
	require.NoError(t, err)

	_, ok := params["allowed_pools"]
	require.True(t, ok)
	_, ok = params["swap_fee"]
	require.True(t, ok)
}

func TestParams_Default(t *testing.T) {
	defaultParams := types.DefaultParams()

	require.NoError(t, defaultParams.Validate())

	assert.Equal(t, types.DefaultAllowedPools, defaultParams.AllowedPools)
	assert.Equal(t, types.DefaultSwapFee, defaultParams.SwapFee)

	assert.Equal(t, 0, len(defaultParams.AllowedPools))
	assert.Equal(t, sdk.ZeroDec(), defaultParams.SwapFee)
}

func TestParams_ParamSetPairs_AllowedPools(t *testing.T) {
	assert.Equal(t, []byte("AllowedPools"), types.KeyAllowedPools)
	defaultParams := types.DefaultParams()

	var paramSetPair *paramstypes.ParamSetPair
	for _, pair := range defaultParams.ParamSetPairs() {
		if bytes.Equal(pair.Key, types.KeyAllowedPools) {
			paramSetPair = &pair
			break
		}
	}
	require.NotNil(t, paramSetPair)

	pairs, ok := paramSetPair.Value.(*types.AllowedPools)
	require.True(t, ok)
	assert.Equal(t, pairs, &defaultParams.AllowedPools)

	assert.Nil(t, paramSetPair.ValidatorFn(*pairs))
	assert.EqualError(t, paramSetPair.ValidatorFn(struct{}{}), "invalid parameter type: struct {}")
}

func TestParams_ParamSetPairs_SwapFee(t *testing.T) {
	assert.Equal(t, []byte("SwapFee"), types.KeySwapFee)
	defaultParams := types.DefaultParams()

	var paramSetPair *paramstypes.ParamSetPair
	for _, pair := range defaultParams.ParamSetPairs() {
		if bytes.Equal(pair.Key, types.KeySwapFee) {
			paramSetPair = &pair
			break
		}
	}
	require.NotNil(t, paramSetPair)

	swapFee, ok := paramSetPair.Value.(*sdk.Dec)
	require.True(t, ok)
	assert.Equal(t, swapFee, &defaultParams.SwapFee)

	assert.Nil(t, paramSetPair.ValidatorFn(*swapFee))
	assert.EqualError(t, paramSetPair.ValidatorFn(struct{}{}), "invalid parameter type: struct {}")
}

func TestParams_Validation(t *testing.T) {
	testCases := []struct {
		name        string
		key         []byte
		testFn      func(params *types.Params)
		expectedErr string
	}{
		{
			name: "duplicate pools",
			key:  types.KeyAllowedPools,
			testFn: func(params *types.Params) {
				params.AllowedPools = types.NewAllowedPools(types.NewAllowedPool("ufury", "ufury"))
			},
			expectedErr: "pool cannot have two tokens of the same type, received 'ufury' and 'ufury'",
		},
		{
			name: "nil swap fee",
			key:  types.KeySwapFee,
			testFn: func(params *types.Params) {
				params.SwapFee = sdk.Dec{}
			},
			expectedErr: "invalid swap fee: <nil>",
		},
		{
			name: "negative swap fee",
			key:  types.KeySwapFee,
			testFn: func(params *types.Params) {
				params.SwapFee = sdk.NewDec(-1)
			},
			expectedErr: "invalid swap fee: -1.000000000000000000",
		},
		{
			name: "swap fee greater than 1",
			key:  types.KeySwapFee,
			testFn: func(params *types.Params) {
				params.SwapFee = sdk.MustNewDecFromStr("1.000000000000000001")
			},
			expectedErr: "invalid swap fee: 1.000000000000000001",
		},
		{
			name: "0 swap fee",
			key:  types.KeySwapFee,
			testFn: func(params *types.Params) {
				params.SwapFee = sdk.ZeroDec()
			},
			expectedErr: "",
		},
		{
			name: "1 swap fee",
			key:  types.KeySwapFee,
			testFn: func(params *types.Params) {
				params.SwapFee = sdk.OneDec()
			},
			expectedErr: "invalid swap fee: 1.000000000000000000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := types.DefaultParams()
			tc.testFn(&params)

			err := params.Validate()

			if tc.expectedErr == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr)
			}

			var paramSetPair *paramstypes.ParamSetPair
			for _, pair := range params.ParamSetPairs() {
				if bytes.Equal(pair.Key, tc.key) {
					paramSetPair = &pair
					break
				}
			}
			require.NotNil(t, paramSetPair)
			value := reflect.ValueOf(paramSetPair.Value).Elem().Interface()

			// assert validation error is same as param set validation
			assert.Equal(t, err, paramSetPair.ValidatorFn(value))
		})
	}
}

func TestParams_String(t *testing.T) {
	params := types.NewParams(
		types.NewAllowedPools(
			types.NewAllowedPool("jinx", "ufury"),
			types.NewAllowedPool("ufury", "usdx"),
		),
		sdk.MustNewDecFromStr("0.5"),
	)

	require.NoError(t, params.Validate())

	output := params.String()
	assert.Contains(t, output, types.PoolID("jinx", "ufury"))
	assert.Contains(t, output, types.PoolID("ufury", "usdx"))
	assert.Contains(t, output, "0.5")
}

func TestAllowedPool_Validation(t *testing.T) {
	testCases := []struct {
		name        string
		allowedPool types.AllowedPool
		expectedErr string
	}{
		{
			name:        "blank token a",
			allowedPool: types.NewAllowedPool("", "ufury"),
			expectedErr: "invalid denom: ",
		},
		{
			name:        "blank token b",
			allowedPool: types.NewAllowedPool("ufury", ""),
			expectedErr: "invalid denom: ",
		},
		{
			name:        "invalid token a",
			allowedPool: types.NewAllowedPool("1ufury", "ufury"),
			expectedErr: "invalid denom: 1ufury",
		},
		{
			name:        "invalid token b",
			allowedPool: types.NewAllowedPool("ufury", "1ufury"),
			expectedErr: "invalid denom: 1ufury",
		},
		{
			name:        "matching tokens",
			allowedPool: types.NewAllowedPool("ufury", "ufury"),
			expectedErr: "pool cannot have two tokens of the same type, received 'ufury' and 'ufury'",
		},
		{
			name:        "invalid token order",
			allowedPool: types.NewAllowedPool("usdx", "ufury"),
			expectedErr: "invalid token order: 'ufury' must come before 'usdx'",
		},
		{
			name:        "invalid token order due to capitalization",
			allowedPool: types.NewAllowedPool("ufury", "UNEMO"),
			expectedErr: "invalid token order: 'UNEMO' must come before 'ufury'",
		},
		{
			name:        "invalid token a with colon",
			allowedPool: types.NewAllowedPool("test:denom", "ufury"),
			expectedErr: "tokenA cannot have colons in the denom: test:denom",
		},
		{
			name:        "invalid token b with colon",
			allowedPool: types.NewAllowedPool("ufury", "u:nemo"),
			expectedErr: "tokenB cannot have colons in the denom: u:nemo",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.allowedPool.Validate()
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}

func TestAllowedPool_TokenMatch_CaseSensitive(t *testing.T) {
	allowedPool := types.NewAllowedPool("UNEMO", "ufury")
	err := allowedPool.Validate()
	assert.NoError(t, err)

	allowedPool = types.NewAllowedPool("haRd", "jinx")
	err = allowedPool.Validate()
	assert.NoError(t, err)

	allowedPool = types.NewAllowedPool("Usdx", "uSdX")
	err = allowedPool.Validate()
	assert.NoError(t, err)
}

func TestAllowedPool_String(t *testing.T) {
	allowedPool := types.NewAllowedPool("jinx", "ufury")
	require.NoError(t, allowedPool.Validate())

	output := `AllowedPool:
  Name: jinx:ufury
	Token A: jinx
	Token B: ufury
`
	assert.Equal(t, output, allowedPool.String())
}

func TestAllowedPool_Name(t *testing.T) {
	testCases := []struct {
		tokens string
		name   string
	}{
		{
			tokens: "atoken btoken",
			name:   "atoken:btoken",
		},
		{
			tokens: "aaa aaaa",
			name:   "aaa:aaaa",
		},
		{
			tokens: "aaaa aaab",
			name:   "aaaa:aaab",
		},
		{
			tokens: "a001 a002",
			name:   "a001:a002",
		},
		{
			tokens: "jinx ufury",
			name:   "jinx:ufury",
		},
		{
			tokens: "bnb jinx",
			name:   "bnb:jinx",
		},
		{
			tokens: "bnb xrpb",
			name:   "bnb:xrpb",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.tokens, func(t *testing.T) {
			tokens := strings.Split(tc.tokens, " ")
			require.Equal(t, 2, len(tokens))

			allowedPool := types.NewAllowedPool(tokens[0], tokens[1])
			require.NoError(t, allowedPool.Validate())

			assert.Equal(t, tc.name, allowedPool.Name())
		})
	}
}

func TestAllowedPools_Validate(t *testing.T) {
	testCases := []struct {
		name         string
		allowedPools types.AllowedPools
		expectedErr  string
	}{
		{
			name: "duplicate pool",
			allowedPools: types.NewAllowedPools(
				types.NewAllowedPool("jinx", "ufury"),
				types.NewAllowedPool("jinx", "ufury"),
			),
			expectedErr: "duplicate pool: jinx:ufury",
		},
		{
			name: "duplicate pools",
			allowedPools: types.NewAllowedPools(
				types.NewAllowedPool("jinx", "ufury"),
				types.NewAllowedPool("bnb", "usdx"),
				types.NewAllowedPool("btcb", "xrpb"),
				types.NewAllowedPool("bnb", "usdx"),
			),
			expectedErr: "duplicate pool: bnb:usdx",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.allowedPools.Validate()
			assert.EqualError(t, err, tc.expectedErr)
		})
	}
}
