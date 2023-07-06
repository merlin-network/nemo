package types

// Querier routes for the validator vesting module
const (
	QueryCirculatingSupply     = "circulating-supply"
	QueryTotalSupply           = "total-supply"
	QueryCirculatingSupplyJINX = "circulating-supply-jinx"
	QueryCirculatingSupplyUSDX = "circulating-supply-usdx"
	QueryCirculatingSupplySWP  = "circulating-supply-swp"
	QueryTotalSupplyJINX       = "total-supply-jinx"
	QueryTotalSupplyUSDX       = "total-supply-usdx"
)

// BaseQueryParams defines the parameters necessary for querying for all Evidence.
type BaseQueryParams struct {
	Page  int `json:"page" yaml:"page"`
	Limit int `json:"limit" yaml:"limit"`
}

// NewBaseQueryParams returns a new BaseQueryParams
func NewBaseQueryParams(page, limit int) BaseQueryParams {
	return BaseQueryParams{
		Page:  page,
		Limit: limit,
	}
}
