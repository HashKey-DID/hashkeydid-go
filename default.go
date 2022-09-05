package hashkeydid_go

import "math/big"

type ChainInfo struct {
	ChainName string
	RPCUrl    string
}

// ChainList is a map contains the infos of chain by the index of chainId
var ChainList = map[string]*ChainInfo{
	"1":      {"Ethereum", "https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7"},
	"137":    {"Polygon", "https://matic-mainnet-archive-rpc.bwarelabs.com"},
	"8217":   {"Klaytn", "https://klaytn01.fandom.finance"},
	"210425": {"PlatON", "https://openapi2.platon.network/rpc"},
}

// Default on-chain data

const DefaultDIDContractAddr = "0x7fDd3f96cBDE51737A9E24b461E7E92A057C3BBf"
const DefaultDIDResolverContractAddr = "0x606729294604A1c71f4BFc001894E4f8095Ec2eF"
const DefaultPlatONUrl = "https://openapi2.platon.network/rpc"

// Const args

var (
	Big0 = new(big.Int).SetUint64(0)
)
