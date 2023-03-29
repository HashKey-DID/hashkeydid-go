package hashkeydid_go

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCore_GetAddr(t *testing.T) {
	core := initTestCore()
	data, err := core.GetAddr(nil, big.NewInt(13756), 2203)
	fmt.Println(data, string(data), err)
}

func TestCore_SetAddr(t *testing.T) {
	core := initTestCore()
	opt := GetOpts("xxx", core.client)
	tx, err := core.SetAddr(opt, big.NewInt(13756), 2203, []byte("3QhSgXSkW6wTTN8JsAjqRtNtHGv8qwc897"))
	fmt.Println(tx.Hash(), err)
}

func TestCore_SetText(t *testing.T) {
	core := initTestCore()
	opt := GetOpts("xxx", core.client)
	tx, err := core.SetText(opt, big.NewInt(13756), "name", "咚咚咚")
	fmt.Println(tx.Hash(), err)
	fmt.Println(tx.Gas(), err)
	fmt.Println(tx.GasPrice(), err)
}

func TestCore_GetText(t *testing.T) {
	core := initTestCore()
	name, err := core.GetText(nil, big.NewInt(13756), "name")
	fmt.Println(name, err)
}
