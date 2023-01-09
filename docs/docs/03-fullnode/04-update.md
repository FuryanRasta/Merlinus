---
id: update
title: Update
sidebar_label: Update
slug: update
---

# Upgrade your Mage full node
These instructions are for full nodes that are running on previous versions of Mage and need to update to the latest version of the Mage software.

## Manual upgrade
The following instructions explain how to **manually upgrade** the node:

1. Stop your node:  
   ```bash sudo systemctl stop maged```

2. Backup your validator files:
   ```bash 
   cp ~/.mage/config/priv_validator_key.json ~/priv_validator_key.json
   cp ~/.mage/config/node_key.json ~/node_key.json
   cp ~/.mage/data/priv_validator_state.json ~/priv_validator_state.json
   ```
   
3. Go into the directory in which you have installed `mage`. If you have followed
the installation instructions and didn't change the path, it should be `~/mage`:
    ```bash
    cd <installation-path> 

    # e.g.
    # cd ~/mage
    ```

4. Now, update the `mage` software:
    ```bash
    git fetch --tags
    git checkout tags/$(git describe --tags `git rev-list --tags --max-count=1`)
    make build && make install
    ```

:::tip Select the version you need  
The above commands checks out the latest release that has been tagged on our repository. If you wish to check out a specific version instead, use the following commands:

1. List all the tags  
   ```bash
   git tags --list
   ```
   
2. Checkout the tag you want 
   ```bash
   git checkout tags/<tag>
   # Example: git checkout tags/v4.1.0
   ```
:::

:::tip Note   
If you have issues at this step, please check that you have the [latest stable version](https://golang.org/dl/) of Go installed.  
:::

### Cosmovisor

:::caution 
If your node is using cosmovisor, and you've followed the above procedure to manually upgrade, don't forget to move the upgraded binary inside the cosmovisor folder by typing the following command:

```bash
cp build/mage ~/.mage/cosmovisor/current/bin/mage
```

Then check if the version of cosmovisor matches with the latest mage version by running:
```bash
cosmovisor version
```
:::

## Automatic upgrade (with Cosmovisor)
Here below it is explained how to prepare your node to be able to **automatically upgrade** itself.

1.Cosmovisor handles the automatic upgrades that happens after the _upgrade governance proposal_ passes.
If during an upgrade your node doesn't have enough space left or if the cosmovisor backup it is taking too much
  time, you can do the following:
   1. Open your `maged` editor:
   ```bash
   sudo systemctl edit maged --full
   ``` 
   1. Add the following line after the last `Environment` line:
   ```bash
    Environment="UNSAFE_SKIP_BACKUP=true"
   ```
