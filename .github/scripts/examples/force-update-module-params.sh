#!/bin/bash
# Example script for using the god committee to update
# the params of a module using dynamic values provided
# via environment variables or set during script execution


# stop the script on first error encountered
# and enable debug logging of all script steps
set -ex

######## EXAMPLE VALUES #########

# setup values used by the script, in github actions these
# are set as environment variables when running seed scripts
CHAIN_API_URL=https://rpc.app.internal.testnet.us-east.production.nemo.io:443
CHAIN_ID=highbury_710-17000

MULTICHAIN_wBTC_CONTRACT_ADDRESS=0x288429bc07B8d030ba418bb30F170327F9fBE502
MULTICHAIN_USDC_CONTRACT_ADDRESS=0xBb304f44b7EFD865361F2AD973d8ebA433893ABC
MULTICHAIN_USDT_CONTRACT_ADDRESS=0xA637F4CECbA91Ad19075bA3d330cd95f694B1707
AXL_USDC_CONTRACT_ADDRESS=0x7a5DBf8e6ac1F6aCCF14f5B4E88b21EAA04c983d
AXL_WBTC_CONTRACT_ADDRESS=0x7d2Ee2914324d5D4dC33A5c295E720659D5F3fA7
wETH_CONTRACT_ADDRESS=0x5d6D67a665C9F169B0f9436E05B11108C1606043
######## EXAMPLE VALUES #########



# configure nemo binary to talk to the desired chain endpoint
nemo config node "${CHAIN_API_URL}"
nemo config chain-id "${CHAIN_ID}"

# use the test keyring to allow scriptng key generation
nemo config keyring-backend test

# wait for transactions to be committed per CLI command
nemo config broadcast-mode block

# setup god's wallet
echo "${NEMO_TESTNET_GOD_MNEMONIC}" | nemo keys add --recover god

# create template string for the proposal we want to enact
# https://merlin-network.atlassian.net/wiki/spaces/ENG/pages/1228537857/Submitting+Governance+Proposals+WIP
PARAM_CHANGE_PROP_TEMPLATE=$(cat <<'END_HEREDOC'
{
    "@type": "/cosmos.params.v1beta1.ParameterChangeProposal",
    "title": "Set Initial ERC-20 Contracts",
    "description": "Set Initial ERC-20 Contracts",
    "changes": [
        {
            "subspace": "evmutil",
            "key": "EnabledConversionPairs",
            "value": "[{\"nemo_erc20_address\":\"MULTICHAIN_USDC_CONTRACT_ADDRESS\",\"denom\":\"erc20/multichain/usdc\"},{\"nemo_erc20_address\":\"MULTICHAIN_USDT_CONTRACT_ADDRESS\",\"denom\":\"erc20/multichain/usdt\"},{\"nemo_erc20_address\":\"MULTICHAIN_wBTC_CONTRACT_ADDRESS\",\"denom\":\"erc20/multichain/btc\"},{\"nemo_erc20_address\":\"AXL_USDC_CONTRACT_ADDRESS\",\"denom\":\"erc20/axelar/usdc\"},{\"nemo_erc20_address\":\"wBTC_CONTRACT_ADDRESS\",\"denom\":\"erc20/axelar/wbtc\"},{\"nemo_erc20_address\":\"wETH_CONTRACT_ADDRESS\",\"denom\":\"erc20/axelar/eth\"}]"
        }
    ]
}
END_HEREDOC
)

# substitute freshly deployed contract addresses
finalProposal=$PARAM_CHANGE_PROP_TEMPLATE

finalProposal="${finalProposal/MULTICHAIN_USDC_CONTRACT_ADDRESS/$MULTICHAIN_USDC_CONTRACT_ADDRESS}"
finalProposal="${finalProposal/MULTICHAIN_USDT_CONTRACT_ADDRESS/$MULTICHAIN_USDT_CONTRACT_ADDRESS}"
finalProposal="${finalProposal/MULTICHAIN_wBTC_CONTRACT_ADDRESS/$MULTICHAIN_wBTC_CONTRACT_ADDRESS}"
finalProposal="${finalProposal/AXL_USDC_CONTRACT_ADDRESS/$AXL_USDC_CONTRACT_ADDRESS}"
finalProposal="${finalProposal/AXL_WBTC_CONTRACT_ADDRESS/$AXL_WBTC_CONTRACT_ADDRESS}"
finalProposal="${finalProposal/wETH_CONTRACT_ADDRESS/$wETH_CONTRACT_ADDRESS}"

# create unique proposal filename
proposalFileName="$(date +%s)-proposal.json"
touch $proposalFileName

# save proposal as file to disk
echo "$finalProposal" > $proposalFileName

# snapshot original module params
originalEvmUtilParams=$(curl https://api.app.internal.testnet.us-east.production.nemo.io/nemo/evmutil/v1beta1/params)
printf "original evm util module params\n %s" , "$originalEvmUtilParams"

# change the params of the chain like a god - make it so 🖖🏽
# make sure to update god committee member permissions for the module
# and params being updated (see below for example)
# https://github.com/Merlin-Network/nemo/pull/1556/files#diff-0bd6043650c708661f37bbe6fa5b29b52149e0ec0069103c3954168fc9f12612R900-R903
nemo tx committee submit-proposal 1 "$proposalFileName" --gas 2000000 --gas-prices 0.01ufury --from god -y

# fetch current module params
updatedEvmUtilParams=$(curl https://api.app.internal.testnet.us-east.production.nemo.io/nemo/evmutil/v1beta1/params)
printf "updated evm util module params\n %s" , "$updatedEvmUtilParams"
