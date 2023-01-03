package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AddAuth Did add address authorization
func (c *Core) AddAuth(opts *bind.TransactOpts, tokenId uint64, addr common.Address) (*types.Transaction, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	tokenIdBig := new(big.Int).SetUint64(tokenId)
	if tokenIdBig.Cmp(totalSupply) == 1 {
		return nil, ErrInvalidTokenId
	}
	return c.did.AddAuth(opts, tokenIdBig, addr)
}

// RemoveAuth Did cancel address authorization
func (c *Core) RemoveAuth(opts *bind.TransactOpts, tokenId uint64, addr common.Address) (*types.Transaction, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	tokenIdBig := new(big.Int).SetUint64(tokenId)
	if tokenIdBig.Cmp(totalSupply) == 1 {
		return nil, ErrInvalidTokenId
	}
	return c.did.RemoveAuth(opts, tokenIdBig, addr)
}

// GetAuthorizedAddrs Get all authorized addresses of did
func (c *Core) GetAuthorizedAddrs(opts *bind.CallOpts, tokenId uint64) ([]common.Address, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return nil, err
	}
	tokenIdBig := new(big.Int).SetUint64(tokenId)
	if tokenIdBig.Cmp(totalSupply) == 1 {
		return nil, ErrInvalidTokenId
	}
	return c.did.GetAuthorizedAddrs(opts, tokenIdBig)
}

// IsAddrAuthorized Get the address is authorized
func (c *Core) IsAddrAuthorized(opts *bind.CallOpts, tokenId uint64, addr common.Address) (bool, error) {
	totalSupply, err := c.did.TotalSupply(nil)
	if err != nil {
		return false, err
	}
	tokenIdBig := new(big.Int).SetUint64(tokenId)
	if tokenIdBig.Cmp(totalSupply) == 1 {
		return false, ErrInvalidTokenId
	}
	return c.did.IsAddrAuthorized(opts, tokenIdBig, addr)
}
