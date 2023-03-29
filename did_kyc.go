package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type KYCInfo struct {
	Status     bool
	UpdateTime uint64
	ExpireTime uint64
}

// GetKYCInfo returns the KYC information for specific token id from DID
func (c *Core) GetKYCInfo(opts *bind.CallOpts, tokenId *big.Int, KYCProvider common.Address, KYCId *big.Int) (*KYCInfo, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	if tokenId.Cmp(totalSupply) == 1 {
		return nil, ErrInvalidTokenId
	}
	status, updatetime, expiretime, err := c.did.GetKYCInfo(opts, tokenId, KYCProvider, KYCId)
	return &KYCInfo{status, updatetime.Uint64(), expiretime.Uint64()}, err
}
