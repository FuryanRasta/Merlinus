#!/usr/bin/env bash

NODES=$1
GENESIS_VERSION=$2
GENESIS_URL=$3
UPGRADE_NAME=$4

BUILDDIR=$(pwd)/build
CONTRIBFOLDER=$(pwd)/contrib
TESTNETDIR=$CONTRIBFOLDER/upgrade_testnet

# Remove the build folder
echo "===> Removing build folder"
rm -r -f $BUILDDIR

# Create the 4 nodes folders with the correct denom
echo "===> Creating $NODES nodes localnet"
make setup-localnet COIN_DENOM="ughost" NODES=$NODES > /dev/null > /dev/null

# Run the Python script to setup the genesis
echo "===> Setting up the genesis file"
docker run --rm --user $UID:$GID \
  -v $TESTNETDIR:/usr/src/app \
  -v $BUILDDIR:/mage:Z \
  desmoslabs/desmos-python python setup_genesis.py /mage $NODES $GENESIS_URL > /dev/null

# Build the new Mage-Cosmovisor image
echo "===> Building the new Mage-Cosmovisor image"
make -C $CONTRIBFOLDER/images mage-cosmovisor MAGE_VERSION=$GENESIS_VERSION > /dev/null

# Set the correct Mage image version inside the docker compose file
echo "===> Setting up the Docker compose file"
sed -i "s|image: \".*\"|image: \"desmoslabs/desmos-cosmovisor:$GENESIS_VERSION\"|g" $TESTNETDIR/docker-compose.yml

# Build the current code using Alpine to make sure it's later compatible with the devnet
echo "===> Building Mage"
make build-alpine > /dev/null

# Copy the Mage binary into the proper folders
UPGRADE_FOLDER="$BUILDDIR/node0/mage/cosmovisor/upgrades/$UPGRADE_NAME/bin"
if [ ! -d "$UPGRADE_FOLDER" ]; then
  echo "===> Setting up upgrade binary"

  for ((i = 0; i < $NODES; i++)); do
    echo "====> Node $i"
    mkdir -p "$BUILDDIR/node$i/mage/cosmovisor/upgrades/$UPGRADE_NAME/bin"
    cp "$BUILDDIR/mage" "$BUILDDIR/node$i/mage/cosmovisor/upgrades/$UPGRADE_NAME/bin/mage"
  done
fi

# Start the devnet
echo "===> Starting the devnet"
docker-compose -f $TESTNETDIR/docker-compose.yml up -d
