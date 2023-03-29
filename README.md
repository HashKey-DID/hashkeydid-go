# HashKey DID Go SDK
[![Tag](https://img.shields.io/badge/tags-v2.0.0-blue)](https://github.com/hashkeydid/hashkeydid-go/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/hashkeydid/hashkeydid-go.svg)](https://pkg.go.dev/github.com/hashkeydid/hashkeydid-go)
[![License](https://img.shields.io/badge/License-MIT-yellow)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hashkeydid/hashkeydid-go)](https://goreportcard.com/report/github.com/hashkeydid/hashkeydid-go)

Go module to work with HashKey DID Protocol.

## Installation

`go get github.com/hashkeydid/hashkeydid-go`

## Usage

`hashkeydid-go` provides simple access to the [HashKey DID](https://hashkey.id) Contracts.

### Core

`hashkeydid-go.Core` is the core struct of this project to use this SDK function.

```go
core, err := hashkeydid.NewDIDCore("rpc url")
```

### Opts
To query the historical info of DID, SDK provides an option arg called `opts`. 
User can query the status with a custom block height.
```go
// query the status on Block 40069811
opts := &bind.CallOpts{BlockNumber: big.NewInt(40069811)}

```

### Sample1-GetName
```go
// after user set reverse
name, err := core.GetDIDNameByAddr(opts, address)
// force
name, err := core.GetDIDNameByAddrForce(opts, address)
```

### Sample2-GetAvatar
In [HashKey DID](https://hashkey.id), `avatarUrl` supports many forms ([Detail]()).

`chainList` is a map which includes information of chains.
SDK provides [a default chainList](https://github.com/hashkeydid/hashkeydid-go/blob/main/default.go) for user to query on-chain status.
User can provide a list with custom chainIds and RPCUrls.

`chainList` in args can be nil.
```go
// query by DID name
avatar, err := core.GetAvatarByDIDName(opts, name)

// query by tokenId
avatar, err := core.GetAvatarByTokenId(opts, tokenId)
```
