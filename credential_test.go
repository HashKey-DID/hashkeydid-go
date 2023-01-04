package hashkeydid_go

import (
	"context"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func Test_IssueDG(t *testing.T) {
	core := initTestCore()
	var name string = "hashkey"
	var symbol string = "key"
	var baseUri string = "hashkey.com/"
	var evidence []byte = common.Hex2Bytes("0x9863bf0ef8ed5f7199307c80d2d9bf3d7ba46905a7e969153229df11c5128591073413291e0223cd7774d42fdef6f56f615d4a71b5a43f1f3a5affc9b056eee601")
	var transferable bool = true
	pri, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	auth, _ := bind.NewKeyedTransactorWithChainID(pri, big.NewInt(210425))
	gasPrice, err := core.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	_, err = core.IssueDG(auth, name, symbol, baseUri, evidence, transferable)
	assert.Nil(t, err)
}

func Test_IssueNFT(t *testing.T) {
	core := initTestCore()
	var name string = "hashkey"
	var symbol string = "key"
	var baseUri string = "hashkey.com/"
	var evidence []byte = common.Hex2Bytes("0x878e0f542d468ec4040e477148c95bebc8780960d8348a09da0e5059482434ab240ccbe04c9b38140b35847439aec2afe2bf9e15930862b99f78f10e2cd74b4700")
	var supply uint64 = 100
	pri, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	auth, _ := bind.NewKeyedTransactorWithChainID(pri, big.NewInt(210425))
	gasPrice, err := core.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	_, err = core.IssueNFT(auth, name, symbol, baseUri, evidence, supply)
	assert.Nil(t, err)
}
