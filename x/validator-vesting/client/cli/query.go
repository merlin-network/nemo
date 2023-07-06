package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/merlin-network/nemo/x/validator-vesting/types"
)

// GetQueryCmd returns the cli query commands for the nemodist module
func GetQueryCmd() *cobra.Command {
	valVestingQueryCmd := &cobra.Command{
		Use:   types.QueryPath,
		Short: "Querying commands for the validator vesting module",
	}

	cmds := []*cobra.Command{
		queryCirculatingSupply(),
		queryTotalSupply(),
		queryCirculatingSupplyJINX(),
		queryCirculatingSupplyUSDF(),
		queryCirculatingSupplySWP(),
		queryTotalSupplyJINX(),
		queryTotalSupplyUSDF(),
	}

	for _, cmd := range cmds {
		flags.AddQueryFlagsToCmd(cmd)
	}

	valVestingQueryCmd.AddCommand(cmds...)
	return valVestingQueryCmd
}

func queryCirculatingSupply() *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply",
		Short: "Get circulating supply",
		Long:  "Get the current circulating supply of nemo tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupply), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}

func queryTotalSupply() *cobra.Command {
	return &cobra.Command{
		Use:   "total-supply",
		Short: "Get total supply",
		Long:  "Get the current total supply of nemo tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupply), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}

func queryCirculatingSupplyJINX() *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply-jinx",
		Short: "Get JINX circulating supply",
		Long:  "Get the current circulating supply of JINX tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupplyJINX), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}

func queryCirculatingSupplyUSDF() *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply-usdf",
		Short: "Get USDF circulating supply",
		Long:  "Get the current circulating supply of USDF tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupplyUSDF), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}

func queryCirculatingSupplySWP() *cobra.Command {
	return &cobra.Command{
		Use:   "circulating-supply-swp",
		Short: "Get SWP circulating supply",
		Long:  "Get the current circulating supply of SWP tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCirculatingSupplySWP), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}

func queryTotalSupplyJINX() *cobra.Command {
	return &cobra.Command{
		Use:   "total-supply-jinx",
		Short: "Get JINX total supply",
		Long:  "Get the current total supply of JINX tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupplyJINX), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}

func queryTotalSupplyUSDF() *cobra.Command {
	return &cobra.Command{
		Use:   "total-supply-usdf",
		Short: "Get USDF total supply",
		Long:  "Get the current total supply of USDF tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Query
			res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupplyUSDF), nil)
			if err != nil {
				return err
			}
			cliCtx = cliCtx.WithHeight(height)

			// Decode and print results
			var out int64
			if err := cliCtx.LegacyAmino.UnmarshalJSON(res, &out); err != nil {
				return fmt.Errorf("failed to unmarshal supply: %w", err)
			}
			return cliCtx.PrintObjectLegacy(out)
		},
	}
}
