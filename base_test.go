package hashkeydid_go

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_Base(t *testing.T){
	core := initTestCore()
	fmt.Println(core.GetTokenIdByDid(nil, "gosdktest.key")) 
	fmt.Println(core.GetDidByTokenId(nil, "13756")) 
	fmt.Println(core.GetAddrByDIDName(nil, "gosdktest.key")) 
	addr,_ := core.GetBlockChainAddress(nil, "13756", 1)
	fmt.Println(common.Bytes2Hex(addr))
	fmt.Println(core.VerifyDIDFormat(nil, "sssss=.key"))
}

func Test_ClaimDID(t *testing.T){
	core := initTestCore()
	opt := GetOpts("78a957a02930610523e101f5", core.client)
	tx,_ := core.ClaimDID(opt, "xiaoming.key")
	fmt.Println(tx.Hash())
}