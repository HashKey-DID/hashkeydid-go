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

// SetNFTBaseUri Only issuer can set NFT's baseuri
func (c *Core) SetNFTBaseUri(opts *bind.TransactOpts, NFTAddr common.Address, baseUri string) (*types.Transaction, error) {
	return c.did.SetNFTBaseUri(opts, NFTAddr, baseUri)
}

// MintDGV1 Only issuer can airdrop the nft for dg version 1
func (c *Core) MintDGV1(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return c.did.MintDGV1(opts, DGAddr, tokenId, addrs)
}

// MintDGV2 Only issuer can airdrop the nft for dg version 2
func (c *Core) MintDGV2(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, addrs []common.Address, data []byte) (*types.Transaction, error) {
	return c.did.MintDGV2(opts, DGAddr, tokenId, addrs, data)
}

// ClaimDG User claim the nft
func (c *Core) ClaimDG(opts *bind.TransactOpts, addr common.Address, tokenId *big.Int, data, evidence []byte) (*types.Transaction, error) {
	return c.did.ClaimDG(opts, addr, tokenId, data, evidence)
}

// MintDGNFT Only issuer can airdrop the nft
func (c *Core) MintDGNFT(opts *bind.TransactOpts, NFTAddr common.Address, sid *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return c.did.MintDGNFT(opts, NFTAddr, sid, addrs)
}

// ClaimDGNFT User claim the nft
func (c *Core) ClaimDGNFT(opts *bind.TransactOpts, NFTAddr common.Address, sid *big.Int, evidence []byte) (*types.Transaction, error) {
	return c.did.ClaimDGNFT(opts, NFTAddr, sid, evidence)
}
