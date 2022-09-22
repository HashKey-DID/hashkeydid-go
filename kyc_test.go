package hashkeydid_go

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test_GetKYCInfo(t *testing.T) {
	core := initTestCore()
	var tokenId uint64 = 5405
	var KYCId uint64 = 3
	opts := &bind.CallOpts{BlockNumber: new(big.Int).SetUint64(41591302)}
	name, err := core.GetKYCInfo(opts, tokenId, common.HexToAddress("0x0FC1021d0B7111f2170d1183367AAcaC26c68888"), KYCId)
	assert.Nil(t, err)
	assert.Equal(t, &KYCInfo{Status: true, UpdateTime: new(big.Int).SetUint64(1641325440), ExpireTime: new(big.Int).SetUint64(1704397440)}, name)
}
