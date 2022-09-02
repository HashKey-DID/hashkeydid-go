package hashkeydid_go

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func initTestCore() *Core {
	core, err := NewDIDCore(DefaultPlatONUrl, DefaultDIDContractAddr, DefaultDIDResolverContractAddr)
	if err != nil {
		log.Fatalf("init test core err: %s\n", err)
	}
	return core
}

func TestCore_GetDIDNameByAddr(t *testing.T) {
	core := initTestCore()
	// the did is claimed at 40069812
	opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(40069811)}
	_, err := core.GetDIDNameByAddr(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	assert.Equal(t, ErrAddrNotClaimed, err)
	opts.BlockNumber = new(big.Int).SetUint64(40069812)
	_, err = core.GetDIDNameByAddr(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	assert.Equal(t, ErrAddrNotSetReverse, err)
	// the did is set reverse at 40070952
	opts.BlockNumber = new(big.Int).SetUint64(40070952)
	name, err := core.GetDIDNameByAddr(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	assert.Nil(t, err)
	assert.Equal(t, "gosdktest.key", name)
}

func TestCore_GetDIDNameByAddrForce(t *testing.T) {
	core := initTestCore()
	opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(40069812)}
	name, err := core.GetDIDNameByAddrForce(opts, common.HexToAddress("0x617E266FFA5c2B168fB6B6cE1Bee9CA2E461DD58"))
	assert.Nil(t, err)
	assert.Equal(t, "gosdktest.key", name)
}
