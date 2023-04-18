package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AddAuth Did add address authorization
func (c *Core) AddAuth(opts *bind.TransactOpts, tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return c.did.AddAuth(opts, tokenId, addr)
}

// RemoveAuth Did cancel address authorization
func (c *Core) RemoveAuth(opts *bind.TransactOpts, tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return c.did.RemoveAuth(opts, tokenId, addr)
}

// GetAuthorizedAddrs Get all authorized addresses of did
func (c *Core) GetAuthorizedAddrs(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	return c.did.GetAuthorizedAddrs(opts, tokenId)
}

// IsAddrAuthorized Get the address is authorized
func (c *Core) IsAddrAuthorized(opts *bind.CallOpts, tokenId *big.Int, addr common.Address) (bool, error) {
	return c.did.IsAddrAuthorized(opts, tokenId, addr)
}
