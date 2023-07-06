#!/usr/bin/env sh

BINARY=/nmd/linux/${BINARY:-nmd}
echo "binary: ${BINARY}"
ID=${ID:-0}
LOG=${LOG:-nmd.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'nmd' E.g.: -e BINARY=nmd_my_test_version"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export NMDHOME="/nmd/node${ID}/nmd"

if [ -d "$(dirname "${NMDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${NMDHOME}" "$@" | tee "${NMDHOME}/${LOG}"
else
  "${BINARY}" --home "${NMDHOME}" "$@"
fi
