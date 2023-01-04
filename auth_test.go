package hashkeydid_go

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestCore_AddAuth(t *testing.T){
	core := initTestCore()
	opt := GetOpts("xxx", core.client)
	tx, err := core.AddAuth(opt, 13756, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80111"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx.Hash())
}

func TestCore_RemoveAuth(t *testing.T){
	core := initTestCore()
	opt := GetOpts("xxx", core.client)
	tx, err := core.RemoveAuth(opt, 13756, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80111"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx.Hash())
}

func TestCore_GetAuthorizedAddrs(t *testing.T){
	core := initTestCore()
	addrs, err := core.GetAuthorizedAddrs(nil, 13756)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addrs)
}


func TestCore_IsAddrAuthorized(t *testing.T){
	core := initTestCore()
	flag, err := core.IsAddrAuthorized(nil, 13756, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80181"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(flag)
}
