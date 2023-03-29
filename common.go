package hashkeydid_go

import "math/big"

var (
	Big0 = new(big.Int).SetUint64(0)
)

type ChainRpcInfo struct {
	ChainName string
	ChainId   uint64
	RPCUrl    string
}

var ChainRpcMap = map[string]*ChainRpcInfo{
	"1":    {"Ethereum", 1, "https://eth-mainnet.nodereal.io/v1/1659dfb40aa24bbb8153a677b98064d7"},
	"137":  {"Polygon", 137, "https://matic-mainnet-archive-rpc.bwarelabs.com"},
	"1284": {"Moonbeam", 1284, "https://moonbeam.public.blastapi.io"},
	//"8217":   {"Klaytn", 8217, "https://klaytn01.fandom.finance"},
	"210425": {"PlatON", 210425, "https://openapi2.platon.network/rpc"},
}
