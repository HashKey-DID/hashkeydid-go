package hashkeydid_go

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

func TestCore_SetReverse(t *testing.T) {
	core := initTestCore()
	opts := GetOpts("xxxx", core.client)
	tx, _ := core.SetReverse(opts, false)
	fmt.Println(tx.Hash())
}

func Test_GetDIDNameByAddrFalse(t *testing.T){
	core := initTestCore()
	fmt.Println(core.GetDIDNameByAddr(nil, common.HexToAddress("0xa060C1C3807059027Ca141EFb63f19E12e0cBF0c")))
	fmt.Println(core.GetDIDNameByAddrForce(nil, common.HexToAddress("0xa060C1C3807059027Ca141EFb63f19E12e0cBF0c")))
}

func GetOpts(privateKeyStr string, client *ethclient.Client) *bind.TransactOpts {
	privateKey, _ := crypto.HexToECDSA(privateKeyStr)
	chainID, _ := client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	auth.Nonce = big.NewInt(int64(nonce))
	fmt.Println(chainID, gasPrice, nonce, err)
	auth.GasPrice = big.NewInt(50000000000)
	auth.GasLimit = 3000000
	// auth.GasLimit = uint64(0)
	auth.Value = big.NewInt(0) // in wei
	return auth
}
