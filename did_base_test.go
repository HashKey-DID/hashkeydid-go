package hashkeydid_go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_Base_GetTokenIdByDid(t *testing.T) {
	core := initTestCore()
	tokenid, err := core.GetTokenIdByDid(nil, "gosdktest.key")
	assert.Nil(t, err)
	assert.Equal(t, tokenid, big.NewInt(15921))
	fmt.Println(core.VerifyDIDFormat("sssss=.key"))
}

func Test_Base_GetDidByTokenId(t *testing.T) {
	core := initTestCore()
	did, err := core.GetDidByTokenId(nil, big.NewInt(13756))
	assert.Nil(t, err)
	assert.Equal(t, did, "herro.key")
}

func Test_Base_GetAddrByDIDName(t *testing.T) {
	core := initTestCore()
	addr, err := core.GetAddrByDIDName(nil, "gosdktest.key")
	assert.Nil(t, err)
	assert.Equal(t, addr, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
}

func Test_Base_GetBlockChainAddress(t *testing.T) {
	core := initTestCore()
	addr, err := core.GetBlockChainAddress(nil, big.NewInt(13756), big.NewInt(1))
	assert.Nil(t, err)
	assert.Equal(t, common.Bytes2Hex(addr), "b45c5eac26af321dd9c02693418976f52e1219b6")
}

func Test_Base_VerifyDIDFormat(t *testing.T) {
	core := initTestCore()
	format, err := core.VerifyDIDFormat("sssss=.key")
	assert.Nil(t, err)
	assert.Equal(t, format, false)
}
