package testutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

func init() {
	// read the .env file, if present
	gotenv.Load()
}

// SuiteConfig wraps configuration details for running the end-to-end test suite.
type SuiteConfig struct {
	// A funded account used to fnd all other accounts.
	FundedAccountMnemonic string

	// A config for using nmtool local networks for the test run
	Nmtool *NmtoolConfig
	// A config for connecting to a running network
	LiveNetwork *LiveNetworkConfig

	// Whether or not to start an IBC chain. Use `suite.SkipIfIbcDisabled()` in IBC tests in IBC tests.
	IncludeIbcTests bool

	// The contract address of a deployed ERC-20 token
	NemoErc20Address string

	// When true, the chains will remain running after tests complete (pass or fail)
	SkipShutdown bool
}

// NmtoolConfig wraps configuration options for running the end-to-end test suite against
// a locally running chain. This config must be defined if E2E_RUN_NMTOOL_NETWORKS is true.
type NmtoolConfig struct {
	// The nemo.configTemplate flag to be passed to nmtool, usually "master".
	// This allows one to change the base genesis used to start the chain.
	NemoConfigTemplate string

	// Whether or not to run a chain upgrade & run post-upgrade tests. Use `suite.SkipIfUpgradeDisabled()` in post-upgrade tests.
	IncludeAutomatedUpgrade bool
	// Name of the upgrade, if upgrade is enabled.
	NemoUpgradeName string
	// Height upgrade will be applied to the test chain, if upgrade is enabled.
	NemoUpgradeHeight int64
	// Tag of nemo docker image that will be upgraded to the current image before tests are run, if upgrade is enabled.
	NemoUpgradeBaseImageTag string
}

// LiveNetworkConfig wraps configuration options for running the end-to-end test suite
// against a live network. It must be defined if E2E_RUN_NMTOOL_NETWORKS is false.
type LiveNetworkConfig struct {
	NemoRpcUrl    string
	NemoGrpcUrl   string
	NemoEvmRpcUrl string
}

// ParseSuiteConfig builds a SuiteConfig from environment variables.
func ParseSuiteConfig() SuiteConfig {
	config := SuiteConfig{
		// this mnemonic is expected to be a funded account that can seed the funds for all
		// new accounts created during tests. it will be available under Accounts["whale"]
		FundedAccountMnemonic: nonemptyStringEnv("E2E_NEMO_FUNDED_ACCOUNT_MNEMONIC"),
		NemoErc20Address:      nonemptyStringEnv("E2E_NEMO_ERC20_ADDRESS"),
		IncludeIbcTests:       mustParseBool("E2E_INCLUDE_IBC_TESTS"),
	}

	skipShutdownEnv := os.Getenv("E2E_SKIP_SHUTDOWN")
	if skipShutdownEnv != "" {
		config.SkipShutdown = mustParseBool("E2E_SKIP_SHUTDOWN")
	}

	useNmtoolNetworks := mustParseBool("E2E_RUN_NMTOOL_NETWORKS")
	if useNmtoolNetworks {
		nmtoolConfig := ParseNmtoolConfig()
		config.Nmtool = &nmtoolConfig
	} else {
		liveNetworkConfig := ParseLiveNetworkConfig()
		config.LiveNetwork = &liveNetworkConfig
	}

	return config
}

// ParseNmtoolConfig builds a NmtoolConfig from environment variables.
func ParseNmtoolConfig() NmtoolConfig {
	config := NmtoolConfig{
		NemoConfigTemplate:      nonemptyStringEnv("E2E_NMTOOL_NEMO_CONFIG_TEMPLATE"),
		IncludeAutomatedUpgrade: mustParseBool("E2E_INCLUDE_AUTOMATED_UPGRADE"),
	}

	if config.IncludeAutomatedUpgrade {
		config.NemoUpgradeName = nonemptyStringEnv("E2E_NEMO_UPGRADE_NAME")
		config.NemoUpgradeBaseImageTag = nonemptyStringEnv("E2E_NEMO_UPGRADE_BASE_IMAGE_TAG")
		upgradeHeight, err := strconv.ParseInt(nonemptyStringEnv("E2E_NEMO_UPGRADE_HEIGHT"), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("E2E_NEMO_UPGRADE_HEIGHT must be a number: %s", err))
		}
		config.NemoUpgradeHeight = upgradeHeight
	}

	return config
}

// ParseLiveNetworkConfig builds a LiveNetworkConfig from environment variables.
func ParseLiveNetworkConfig() LiveNetworkConfig {
	return LiveNetworkConfig{
		NemoRpcUrl:    nonemptyStringEnv("E2E_NEMO_RPC_URL"),
		NemoGrpcUrl:   nonemptyStringEnv("E2E_NEMO_GRPC_URL"),
		NemoEvmRpcUrl: nonemptyStringEnv("E2E_NEMO_EVM_RPC_URL"),
	}
}

// mustParseBool is a helper method that panics if the env variable `name`
// cannot be parsed to a boolean
func mustParseBool(name string) bool {
	envValue := os.Getenv(name)
	if envValue == "" {
		panic(fmt.Sprintf("%s is unset but expected a bool", name))
	}
	value, err := strconv.ParseBool(envValue)
	if err != nil {
		panic(fmt.Sprintf("%s (%s) cannot be parsed to a bool: %s", name, envValue, err))
	}
	return value
}

// nonemptyStringEnv is a helper method that panics if the env variable `name`
// is empty or undefined.
func nonemptyStringEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(fmt.Sprintf("no %s env variable provided", name))
	}
	return value
}
