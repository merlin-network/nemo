package runner

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// LiveNodeRunnerConfig implements NodeRunner.
// It connects to a running network via the RPC, GRPC, and EVM urls.
type LiveNodeRunnerConfig struct {
	NemoRpcUrl    string
	NemoGrpcUrl   string
	NemoEvmRpcUrl string
}

// LiveNodeRunner implements NodeRunner for an already-running chain.
// If a LiveNodeRunner is used, end-to-end tests are run against a live chain.
type LiveNodeRunner struct {
	config LiveNodeRunnerConfig
}

var _ NodeRunner = LiveNodeRunner{}

// NewLiveNodeRunner creates a new LiveNodeRunner.
func NewLiveNodeRunner(config LiveNodeRunnerConfig) *LiveNodeRunner {
	return &LiveNodeRunner{config}
}

// StartChains implements NodeRunner.
// It initializes connections to the chain based on parameters.
// It attempts to ping the necessary endpoints and panics if they cannot be reached.
func (r LiveNodeRunner) StartChains() Chains {
	fmt.Println("establishing connection to live nemo network")
	chains := NewChains()

	nemoChain := ChainDetails{
		RpcUrl:    r.config.NemoRpcUrl,
		GrpcUrl:   r.config.NemoGrpcUrl,
		EvmRpcUrl: r.config.NemoEvmRpcUrl,
	}

	if err := waitForChainStart(nemoChain); err != nil {
		panic(fmt.Sprintf("failed to ping chain: %s", err))
	}

	// determine chain id
	grpc, err := nemoChain.GrpcConn()
	if err != nil {
		panic(fmt.Sprintf("failed to establish grpc conn to %s: %s", r.config.NemoGrpcUrl, err))
	}
	tm := tmservice.NewServiceClient(grpc)
	nodeInfo, err := tm.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	if err != nil {
		panic(fmt.Sprintf("failed to fetch nemo node info: %s", err))
	}
	nemoChain.ChainId = nodeInfo.DefaultNodeInfo.Network

	// determine staking denom
	staking := stakingtypes.NewQueryClient(grpc)
	stakingParams, err := staking.Params(context.Background(), &stakingtypes.QueryParamsRequest{})
	if err != nil {
		panic(fmt.Sprintf("failed to fetch nemo staking params: %s", err))
	}
	nemoChain.StakingDenom = stakingParams.Params.BondDenom

	chains.Register("nemo", &nemoChain)

	fmt.Printf("successfully connected to live network %+v\n", nemoChain)

	return chains
}

// Shutdown implements NodeRunner.
// As the chains are externally operated, this is a no-op.
func (LiveNodeRunner) Shutdown() {
	fmt.Println("shutting down e2e test connections.")
}
