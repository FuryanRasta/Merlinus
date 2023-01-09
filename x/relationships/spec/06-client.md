---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client 

## CLI

A user can query and interact with the `relationships` module using the CLI. 

### Query 

The `query` commands allow users to query the `relationships` state. 

```bash
mage query relationships --help
```

#### relationships
The `relationships` command allows user to query the relationships inside a subspace having a specific id, providing an optional creator and counterparty. 

```bash
mage query relationships [subspace-id] [[creator]] [[counterparty]] [flags]
```

Example: 
```bash
mage query relationships relationships 1 mage1... mage1...
```

Example Output: 
```yaml
pagination:
  next_key: null
  total: "0"
relationships:
- counterparty: mage1tamzg6rfj9wlmqhthgfmn9awq0d8ssgfr8fjns
  creator: mage13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu
  subspace_id: "1"
```

#### blocks 
The `blocks` command allows to query the user blocks stored inside a subspace having a given id, providing an optional blocker and blocked addresses.

```bash
mage query relationships blocks [subspace-id] [[blocker]] [[blocked]] [flags]
```

Example: 
```bash
mage query relationships blocks 1 mage1... mage1...
```

Example Output: 
```yaml
blocks:
  - blocked: mage1tamzg6rfj9wlmqhthgfmn9awq0d8ssgfr8fjns
    blocker: mage13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu
    reason: ""
    subspace_id: "1"
pagination:
  next_key: null
  total: "0"
```

### Transactions
The `tx` commands allow users to interact with the `relationships` module. 

```bash
mage tx relationships --help
```

#### create-relationship
The `create-relationship` allows users to create a relationship with another user inside a specific subspace.

```bash
mage tx relationships create-relationship [counterparty] [subspace-id] [flags]
```

Example:
```bash
mage tx relationships create-relationship mage1... 1
```

#### delete-relationship
The `delete-relationship` allows users to delete an existing relationship. 

```bash
mage tx relationships delete-relationship [counterparty] [subspace-id] [flags]
```

Example:
```bash
mage tx relationships delete-relationship mage1... 1
```

#### block
The `block` command allows users to block another user inside a specific subspace, optionally providing a reason. 

```bash
mage tx relationships block [address] [subspace] [[reason]] [flags]
```

Example:
```bash
mage tx relationships block mage1... 1 "My reason"
```

#### unblock
The `unblock` command allows users to unblock a previously blocked user.

```bash
mage tx relationships unblock [address] [subspace] [flags]
```

Example:
```bash
mage tx relationships unblock mage1... 1
```


## gRPC
A user can query the `relationships` module gRPC endpoints.

### Relationships
The `Relationships` endpoint allows users to query for the relationships inside a subspace having a given id, optionally providing user and counterparty addresses.

```bash
mage.relationships.v1.Query/Relationships
```

Example:
```bash
grpcurl -plaintext \
  -d '{"subspace_id": "1"}' localhost:9090 mage.relationships.v1.Query/Relationships
```

Example Output:
```json
{
  "relationships": [
    {
      "creator": "mage13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu",
      "counterparty": "mage1tamzg6rfj9wlmqhthgfmn9awq0d8ssgfr8fjns",
      "subspaceId": "1"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

### Blocks
The `Blocks` endpoint allows users to query for the user blocks inside a subspace having a given id, optionally providing user and counterparty addresses.

```bash
mage.relationships.v1.Query/Blocks
```

Example:
```bash
grpcurl -plaintext \
  -d '{"subspace_id": "1"}' localhost:9090 mage.relationships.v1.Query/Blocks
```

Example Output:
```json
{
  "blocks": [
    {
      "blocker": "mage13yp2fq3tslq6mmtq4628q38xzj75ethzela9uu",
      "blocked": "mage1tamzg6rfj9wlmqhthgfmn9awq0d8ssgfr8fjns",
      "subspaceId": "1"
    }
  ],
  "pagination": {
    "total": "1"
  }
}
```

## REST
A user can query the `relationships` module using REST endpoints. 

## Relationships
The `relationships` endpoint allows users to query for the relationships inside a subspace having a given id, optionally providing user and counterparty addresses.

```
/mage/relationships/v1/relationships?subspace_id={subspaceID}&user={userAddress}&counterparty={counterpartyAddress}
```

## Blocks
The `blocks` endpoint allows users to query for the user blocks inside a subspace having a given id, optionally providing user and counterparty addresses.

```
/mage/relationships/v1/blocks?subspace_id={subspaceID}&blocker={blockerAddress}&blocked={blockedAddress}
```