package hashkeydid_go

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestCore_AddAuth(t *testing.T) {
	core := initTestCore()
	opt := GetOpts("xxx", core.client)
	tokenId := big.NewInt(13756)
	tx, err := core.AddAuth(opt, tokenId, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80111"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx.Hash())
}

func TestCore_RemoveAuth(t *testing.T) {
	core := initTestCore()
	opt := GetOpts("xxx", core.client)
	tokenId := big.NewInt(13756)
	tx, err := core.RemoveAuth(opt, tokenId, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80111"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tx.Hash())
}

func TestCore_GetAuthorizedAddrs(t *testing.T) {
	core := initTestCore()
	tokenId := big.NewInt(13756)
	addrs, err := core.GetAuthorizedAddrs(nil, tokenId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addrs)
}

func TestCore_IsAddrAuthorized(t *testing.T) {
	core := initTestCore()
	tokenId := big.NewInt(13756)
	flag, err := core.IsAddrAuthorized(nil, tokenId, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80181"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(flag)
}
