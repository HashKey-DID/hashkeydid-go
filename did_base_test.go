package hashkeydid_go

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_Base(t *testing.T) {
	core := initTestCore()
	fmt.Println(core.GetTokenIdByDid(nil, "gosdktest.key"))
	fmt.Println(core.GetDidByTokenId(nil, big.NewInt(13756)))
	fmt.Println(core.GetAddrByDIDName(nil, "gosdktest.key"))
	addr, _ := core.GetBlockChainAddress(nil, big.NewInt(13756), big.NewInt(1))
	fmt.Println(common.Bytes2Hex(addr))
	fmt.Println(core.VerifyDIDFormat(nil, "sssss=.key"))
}
