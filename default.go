package hashkeydid_go

import "math/big"

// Default on-chain data

const DefaultDIDContractAddr = "0x7fDd3f96cBDE51737A9E24b461E7E92A057C3BBf"
const DefaultDIDResolverContractAddr = "0x606729294604A1c71f4BFc001894E4f8095Ec2eF"
const DefaultPlatONUrl = "https://openapi2.platon.network/rpc"

// Const args

var (
	Big0 = new(big.Int).SetUint64(0)
)
