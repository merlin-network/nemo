<p align="center">
  <img src="./nemo-logo.svg" width="300">
</p>

<div align="center">

[![version](https://img.shields.io/github/tag/merlin-network/nemo.svg)](https://github.com/merlin-network/nemo/releases/latest)
[![CircleCI](https://circleci.com/gh/Merlin-Network/nemo/tree/master.svg?style=shield)](https://circleci.com/gh/Merlin-Network/nemo/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/merlin-network/nemo)](https://goreportcard.com/report/github.com/merlin-network/nemo)
[![API Reference](https://godoc.org/github.com/Merlin-Network/nemo?status.svg)](https://godoc.org/github.com/Merlin-Network/nemo)
[![GitHub](https://img.shields.io/github/license/merlin-network/nemo.svg)](https://github.com/Merlin-Network/nemo/blob/master/LICENSE.md)
[![Twitter Follow](https://img.shields.io/twitter/follow/NEMO_CHAIN.svg?label=Follow&style=social)](https://twitter.com/NEMO_CHAIN)
[![Discord Chat](https://img.shields.io/discord/704389840614981673.svg)](https://discord.com/invite/kQzh3Uv)

</div>

<div align="center">

### [Telegram](https://t.me/nemolabs) | [Medium](https://medium.com/merlin-network) | [Discord](https://discord.gg/JJYnuCx)

</div>

Reference implementation of Nemo, a blockchain for cross-chain DeFi. Built using the [cosmos-sdk](https://github.com/cosmos/cosmos-sdk).

## Mainnet

The current recommended version of the software for mainnet is [v0.23.0](https://github.com/Merlin-Network/nemo/releases/tag/v0.23.0). The master branch of this repository often contains considerable development work since the last mainnet release and is __not__ runnable on mainnet.

### Installation and Setup
For detailed instructions see [the Nemo docs](https://docs.nemo.io/docs/participate/validator-node).

```bash
git checkout v0.23.0
make install
```

End-to-end tests of Nemo use a tool for generating networks with different configurations: [nmtool](https://github.com/Merlin-Network/nmtool).
This is included as a git submodule at [`tests/e2e/nmtool`](tests/e2e/nmtool/).
When first cloning the repository, if you intend to run the e2e integration tests, you must also
clone the submodules:
```bash
git clone --recurse-submodules https://github.com/Merlin-Network/nemo.git
```

Or, if you have already cloned the repo: `git submodule update --init`

## Testnet

For further information on joining the testnet, head over to the [testnet repo](https://github.com/Merlin-Network/nemo-testnets).

## Docs

Nemo protocol and client documentation can be found in the [Nemo docs](https://docs.nemo.io).

If you have technical questions or concerns, ask a developer or community member in the [Nemo discord](https://discord.com/invite/kQzh3Uv).

## Security

If you find a security issue, please report it to security [at] nemo.io. Depending on the verification and severity, a bug bounty may be available.

## License

Copyright © Nemo Labs, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).
