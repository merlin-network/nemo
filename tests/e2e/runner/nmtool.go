package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type NmtoolRunnerConfig struct {
	NemoConfigTemplate string

	ImageTag   string
	IncludeIBC bool

	EnableAutomatedUpgrade  bool
	NemoUpgradeName         string
	NemoUpgradeHeight       int64
	NemoUpgradeBaseImageTag string

	SkipShutdown bool
}

// NmtoolRunner implements a NodeRunner that spins up local chains with nmtool.
// It has support for the following:
// - running a Nemo node
// - optionally, running an IBC node with a channel opened to the Nemo node
// - optionally, start the Nemo node on one version and upgrade to another
type NmtoolRunner struct {
	config NmtoolRunnerConfig
}

var _ NodeRunner = &NmtoolRunner{}

// NewNmtoolRunner creates a new NmtoolRunner.
func NewNmtoolRunner(config NmtoolRunnerConfig) *NmtoolRunner {
	return &NmtoolRunner{
		config: config,
	}
}

// StartChains implements NodeRunner.
// For NmtoolRunner, it sets up, runs, and connects to a local chain via nmtool.
func (k *NmtoolRunner) StartChains() Chains {
	// install nmtool if not already installed
	installNmtoolCmd := exec.Command("./scripts/install-nmtool.sh")
	installNmtoolCmd.Stdout = os.Stdout
	installNmtoolCmd.Stderr = os.Stderr
	if err := installNmtoolCmd.Run(); err != nil {
		panic(fmt.Sprintf("failed to install nmtool: %s", err.Error()))
	}

	// start local test network with nmtool
	log.Println("starting nemo node")
	nmtoolArgs := []string{"testnet", "bootstrap", "--nemo.configTemplate", k.config.NemoConfigTemplate}
	// include an ibc chain if desired
	if k.config.IncludeIBC {
		nmtoolArgs = append(nmtoolArgs, "--ibc")
	}
	// handle automated upgrade functionality, if defined
	if k.config.EnableAutomatedUpgrade {
		nmtoolArgs = append(nmtoolArgs,
			"--upgrade-name", k.config.NemoUpgradeName,
			"--upgrade-height", fmt.Sprint(k.config.NemoUpgradeHeight),
			"--upgrade-base-image-tag", k.config.NemoUpgradeBaseImageTag,
		)
	}
	// start the chain
	startNemoCmd := exec.Command("nmtool", nmtoolArgs...)
	startNemoCmd.Env = os.Environ()
	startNemoCmd.Env = append(startNemoCmd.Env, fmt.Sprintf("NEMO_TAG=%s", k.config.ImageTag))
	startNemoCmd.Stdout = os.Stdout
	startNemoCmd.Stderr = os.Stderr
	log.Println(startNemoCmd.String())
	if err := startNemoCmd.Run(); err != nil {
		panic(fmt.Sprintf("failed to start nemo: %s", err.Error()))
	}

	// wait for chain to be live.
	// if an upgrade is defined, this waits for the upgrade to be completed.
	if err := waitForChainStart(nmtoolNemoChain); err != nil {
		k.Shutdown()
		panic(err)
	}
	log.Println("nemo is started!")

	chains := NewChains()
	chains.Register("nemo", &nmtoolNemoChain)
	if k.config.IncludeIBC {
		chains.Register("ibc", &nmtoolIbcChain)
	}
	return chains
}

// Shutdown implements NodeRunner.
// For NmtoolRunner, it shuts down the local nmtool network.
// To prevent shutting down the chain (eg. to preserve logs or examine post-test state)
// use the `SkipShutdown` option on the config.
func (k *NmtoolRunner) Shutdown() {
	if k.config.SkipShutdown {
		log.Printf("would shut down but SkipShutdown is true")
		return
	}
	log.Println("shutting down nemo node")
	shutdownNemoCmd := exec.Command("nmtool", "testnet", "down")
	shutdownNemoCmd.Stdout = os.Stdout
	shutdownNemoCmd.Stderr = os.Stderr
	if err := shutdownNemoCmd.Run(); err != nil {
		panic(fmt.Sprintf("failed to shutdown nmtool: %s", err.Error()))
	}
}
