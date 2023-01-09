#!/usr/bin/env sh

BINARY=cosmovisor
ID=${ID:-0}
LOG=${LOG:-mage.log}

export MAGEDHOME="/mage/node${ID}/mage"

# Set environment variables
export DAEMON_NAME=mage
export DAEMON_HOME="$MAGEDHOME"
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true

# Setup Cosmovisor
COSMOVISOR_GENESIS="$MAGEDHOME/cosmovisor/genesis/bin"
if [ ! -d "$COSMOVISOR_GENESIS" ]; then
  mkdir -p $COSMOVISOR_GENESIS
  cp $(which mage) "$COSMOVISOR_GENESIS/mage"
fi

# Run the command
if [ -d "$(dirname "${MAGEDHOME}"/"${LOG}")" ]; then
  "${BINARY}" "$@" --home "${MAGEDHOME}" | tee "${MAGEDHOME}/${LOG}"
else
  "${BINARY}" "$@" --home "${MAGEDHOME}"
fi