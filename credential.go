package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IssueDG issue a new DeedGrain token
func (c *Core) IssueDG(opts *bind.TransactOpts, name string, symbol string, baseUri string, evidence []byte, transferable bool) (*types.Transaction, error) {
	if len(evidence) != 65 {
		return nil, ErrInvalidEvidence
	}
	return c.did.IssueDG(opts, name, symbol, baseUri, evidence, transferable)
}

// IssueNFT issue a new DeedGrain NFT
func (c *Core) IssueNFT(opts *bind.TransactOpts, name string, symbol string, baseUri string, evidence []byte, supply uint64) (*types.Transaction, error) {
	if len(evidence) != 65 {
		return nil, ErrInvalidEvidence
	}
	return c.did.IssueNFT(opts, name, symbol, baseUri, evidence, new(big.Int).SetUint64(supply))
}

// SetTokenSupply set every kind of token's supply
func (c *Core) SetTokenSupply(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, supply uint64) (*types.Transaction, error) {
	return c.did.SetTokenSupply(opts, DGAddr, tokenId, new(big.Int).SetUint64(supply))
}

// SetNFTSupply set NFT supply
func (c *Core) SetNFTSupply(opts *bind.TransactOpts, NFTAddr common.Address, supply uint64) (*types.Transaction, error) {
	return c.did.SetNFTSupply(opts, NFTAddr, new(big.Int).SetUint64(supply))
}
