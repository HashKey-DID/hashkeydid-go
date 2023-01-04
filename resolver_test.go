package hashkeydid_go

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/common"
)

func TestCore_GetAddr(t *testing.T){
	core := initTestCore()
	data, err := core.GetAddr(nil, big.NewInt(13756),1)
	// fmt.Println(common.Bytes2Hex(data), err)
	fmt.Println(string(data), err)
}

func TestCore_SetAddr(t *testing.T){
	core := initTestCore()
	opt := GetOpts("b3db6526e98e79c7bd1dcadfa15a01e1de5c7293669608f90b9230581047cbc4", core.client)
	tx, err := core.SetAddr(opt, big.NewInt(13756),1, []byte("aaac5eac26af321dd9c02693418976f52e1219b6"))
	fmt.Println(tx.Hash(), err)
}

func TestCore_SetText(t *testing.T){
	core := initTestCore()
	opt := GetOpts("b3db6526e98e79c7bd1dcadfa15a01e1de5c7293669608f90b9230581047cbc4", core.client)
	tx, err := core.SetText(opt, big.NewInt(13756),"name", "毕小")
	fmt.Println(tx.Hash(), err)
}

func TestCore_GetText(t *testing.T){
	core := initTestCore()
	name, err := core.GetText(nil, big.NewInt(13756),"name")
	fmt.Println(name, err)
}

func TestCore_SetReverse(t *testing.T){
	core := initTestCore()
	opt := GetOpts("b3db6526e98e79c7bd1dcadfa15a01e1de5c7293669608f90b9230581047cbc4", core.client)

	tx, err := core.SetReverse(opt, common.HexToAddress("0xB45c5Eac26AF321dd9C02693418976F52E1219b6"), false)
	fmt.Println(tx.Hash(), err)
}