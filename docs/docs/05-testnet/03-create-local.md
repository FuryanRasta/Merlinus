---
id: create-local
title: Create a local testnet
sidebar_label: Create local testent
slug: create-local
---

# Create a local testnet
:::warning Required Mage executables  
You need to [install `mage`](02-install.md) before going further.  
:::

There are two types of local testnets:

- [Single node testnet](#creating-a-single-node-testnet), which allows you to have a faster testnet with only one validator running on your machine. 

- [Multi-node testnet](#creating-a-multi-node-testnet), which requires you to have [Docker](https://docker.io) installed to run 4 validator nodes locally on your machine. 

## Creating a single node testnet
To create a single node local testnet, run the following commands:

1. Create a local key. Replace `<your-key-name>` with whatever name you prefer.
   ```bash 
   mage keys add <your-key-name>
   ```

   You will be required to input a password. Please make sure you use one that you will remember later. You should now see an output like

   ```bash
   $ mage keys add jack --dry-run
   
   - name: jack
     type: local
     address: mage1qdv08q76fmfwwzrxcqs78z6pzfxe88cgc5a3tk
     pubkey: magepub1addwnpepq2j9a35spphh6q529y2thg8tjw9l2c32hck98fnmu99sxpw9a9aegugm6xs
     mnemonic: ""
     threshold: 0
     pubkeys: []
   
   
   **Important** write this mnemonic phrase in a safe place.
   It is the only way to recover your account if you ever forget your password.
   
   glory discover erosion mention grow prosper supreme term nephew venue pear eternal budget rely outdoor lobster strong sign space make soccer medal tuition patrol
   ```
   
   Make sure you save the shown mnemonic phrase in some safe place as it might return useful in the future. 
   
2. Initialize the testnet
   ```bash
   mage init testnet --chain-id testnet
   mage add-genesis-account <your-key-name> 100000000000stake
   mage gentx <your-key-name> 1000000000stake --chain-id testnet
   mage collect-gentxs
   ``` 
   
   During the procedure you will be asked to input the same key password you have set inside point 1. 
   
3. Start the testnet.  
   Once you have completed all the steps, you are ready to start your local testnet by running: 
   ```bash
   mage start
   ```

## Creating a multi node testnet 
To create a local multi node testnet, you can simply run the following command: 

```bash
make localnet-start
```

This command creates a 4-node network using the `magelabs/magenode` image. The ports for each node are found in this
table:

| Node ID       | P2P Port | RPC Port |
|---------------|----------|----------|
| `magenode0` | `26656`  | `26657`  |
| `magenode1` | `26659`  | `26660`  |
| `magenode2` | `26661`  | `26662`  |
| `magenode3` | `26663`  | `26664`  |

To update the binary, just rebuild it and restart the nodes:

```
make build-linux localnet-start
```

#### Configuration

The `make localnet-start` creates files for a 4-node testnet in `./build` by calling the `mage testnet` command. This outputs a handful of files in the `./build` directory:

```bash
$ tree -L 2 build/
build/
├── mage
├── gentxs
│   ├── node0.json
│   ├── node1.json
│   ├── node2.json
│   └── node3.json
├── node0
│   ├── mage
│   │   ├── key_seed.json
│   │   └── keys
│   └── mage
│       ├── ${LOG:-mage.log}
│       ├── config
│       └── data
├── node1
│   ├── mage
│   │   └── key_seed.json
│   └── mage
│       ├── ${LOG:-mage.log}
│       ├── config
│       └── data
├── node2
│   ├── mage
│   │   └── key_seed.json
│   └── mage
│       ├── ${LOG:-mage.log}
│       ├── config
│       └── data
└── node3
    ├── mage
    │   └── key_seed.json
    └── mage
        ├── ${LOG:-mage.log}
        ├── config
        └── data
```

Each `./build/nodeN` directory is mounted to the `/mage` directory in each container.

#### Logging

Logs are saved under each `./build/nodeN/mage/mage.log`. You can also watch logs directly via Docker, for example:

```
docker logs -f magenode0
```

#### Keys & Accounts

To interact with `mage` and start querying state or creating txs, you use the
`mage` directory of any given node as your `home`, for example:

```bash
mage keys list --home ./build/node0/mage
```

Now that accounts exists, you may create new accounts and send those accounts funds!

**Note**: Each node's seed is located at `./build/nodeN/mage/key_seed.json` and can be restored to the CLI using the `mage keys add --restore` command
