package hashkeydid_go

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestCore_GetAuthorizedAddrs(t *testing.T) {
	core := initTestCore()
	tokenId := big.NewInt(13756)
	addrs, err := core.GetAuthorizedAddrs(nil, tokenId)
	assert.Nil(t, err)
	assert.Equal(t, len(addrs), 2)
}

func TestCore_IsAddrAuthorized(t *testing.T) {
	core := initTestCore()
	tokenId := big.NewInt(13756)
	flag, err := core.IsAddrAuthorized(nil, tokenId, common.HexToAddress("0xc6642B7980A5a702732B243b0C21655e82e80181"))
	assert.Nil(t, err)
	assert.Equal(t, flag, false)
}
