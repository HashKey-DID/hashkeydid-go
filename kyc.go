package hashkeydid_go

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type KYCInfo struct {
	Status     bool
	UpdateTime *big.Int
	ExpireTime *big.Int
}

// AddKYC add KYC information for specific token id
func (c *Core) AddKYC(opts *bind.TransactOpts, tokenId uint64, KYCProvider common.Address, KYCId uint64, status bool, updateTime uint64, expireTime uint64, evidence []byte) (*types.Transaction, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	if new(big.Int).SetUint64(tokenId).Cmp(totalSupply) == 1 {
		return nil, ErrInvalidTokenId
	}
	if int64(expireTime) <= time.Now().Unix() {
		return nil, ErrTimeExpired
	}
	if len(evidence) != 65 {
		return nil, ErrInvalidEvidence
	}
	return c.did.AddKYC(opts, new(big.Int).SetUint64(tokenId), KYCProvider, new(big.Int).SetUint64(KYCId), status, new(big.Int).SetUint64(updateTime), new(big.Int).SetUint64(expireTime), evidence)
}

// GetKYCInfo returns the KYC information for specific token id from DID
func (c *Core) GetKYCInfo(opts *bind.CallOpts, tokenId uint64, KYCProvider common.Address, KYCId uint64) (*KYCInfo, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	if new(big.Int).SetUint64(tokenId).Cmp(totalSupply) == 1 {
		return nil, ErrInvalidTokenId
	}
	status, updatetime, expiretime, err := c.did.GetKYCInfo(opts, new(big.Int).SetUint64(tokenId), KYCProvider, new(big.Int).SetUint64(KYCId))
	return &KYCInfo{status, updatetime, expiretime}, err
}
