#!/usr/bin/env sh

BINARY=/mage/${BINARY:-mage}
ID=${ID:-0}
LOG=${LOG:-mage.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'mage'"
	exit 1
fi

BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"

if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export MAGEDHOME="/mage/node${ID}/mage"

if [ -d "$(dirname "${MAGEDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${MAGEDHOME}" "$@" | tee "${MAGEDHOME}/${LOG}"
else
  "${BINARY}" --home "${MAGEDHOME}" "$@"
fi
