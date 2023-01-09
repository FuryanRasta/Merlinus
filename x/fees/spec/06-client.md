---
id: client
title: Client
sidebar_label: Client
slug: client
---

# Client 

## CLI

A user can query the `fees` module using the CLI. 

### Query

The `query` commands allow users to query the `fees` state.

```
mage query fees --help
```

#### parameters
The `parameters` command allows users to get hte currently set parameters.

```bash
mage query fees parameters [flags]
```

Example:
```bash
mage query fees parameters
```

Example Output:
```yaml
params:
  min_fees:
    - message_type: /mage.profiles.v2.SaveProfile
      amount: 
        - amount: "100"
          denom: "stake"
```

## gRPC 
A user can query the `fees` module gRPC endpoints. 

### Params
The `Params` endpoint allows users to query the current params of the `fees` module.

```bash
mage.fees.v1.Query/Params
```

Example:
```bash
grpcurl localhost:9090 mage.fees.v1.Query/Params
```

Example Output: 
```json
{
  "params": {
    "min_fees": [
      {
        "message_type": "/mage.profiles.v2.SaveProfile",
        "amount": [
          {
            "amount": "100",
            "denom": "stake"
          }
        ]
      }
    ]
  }
}
```

## REST 
A user can query the `fees` module using REST endpoints.

### Params
The `params` endpoint allows users to query for the module parameters.

```bash
/mage/fees/v1/params
```

Example: 
```bash
curl localhost:1317/mage/fees/v1/params
```

Example Output:
```json
{
  "params": {
    "min_fees": [
      {
        "message_type": "/mage.profiles.v2.SaveProfile",
        "amount": [
          {
            "amount": "100",
            "denom": "stake"
          }
        ]
      }
    ]
  }
}
```
