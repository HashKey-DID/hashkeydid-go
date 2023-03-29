package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// GetDIDNameByAddr returns the did name by address when user set reverse
func (c *Core) GetDIDNameByAddr(opts *bind.CallOpts, address common.Address) (string, error) {
	return c.resolver.Name(opts, address)
}

// GetAddr Check address on different chain
func (c *Core) GetAddr(opts *bind.CallOpts, tokenId *big.Int, coinType uint64) ([]byte, error) {
	return c.resolver.Addr(opts, tokenId, new(big.Int).SetUint64(coinType))
}

// SetAddr Sets DID different address on different chain
func (c *Core) SetAddr(opts *bind.TransactOpts, tokenId *big.Int, coinType uint64, addr []byte) (*types.Transaction, error) {
	return c.resolver.SetAddr(opts, tokenId, new(big.Int).SetUint64(coinType), addr)
}

// GetContentHash Check tokenId's content hash(ipfs/CID)
func (c *Core) GetContentHash(opts *bind.CallOpts, tokenId *big.Int) ([]byte, error) {
	return c.resolver.ContentHash(opts, tokenId)
}

// SetContentHash Sets tokenId's content hash(ipfs/CID)
func (c *Core) SetContentHash(opts *bind.TransactOpts, tokenId *big.Int, cid []byte) (*types.Transaction, error) {
	return c.resolver.SetContentHash(opts, tokenId, cid)
}

// GetPubkey Gets the SECP256k1 public key associated with a did
func (c *Core) GetPubkey(opts *bind.CallOpts, tokenId *big.Int) ([32]byte, [32]byte, error) {
	pub, err := c.resolver.Pubkey(opts, tokenId)
	if err != nil {
		return [32]byte{}, [32]byte{}, err
	}
	return pub.X, pub.Y, nil
}

// SetPubkey Sets the SECP256k1 public key associated with a did
func (c *Core) SetPubkey(opts *bind.TransactOpts, tokenId *big.Int, x, y [32]byte) (*types.Transaction, error) {
	return c.resolver.SetPubkey(opts, tokenId, x, y)
}

// GetText Gets the text information associated with a did
func (c *Core) GetText(opts *bind.CallOpts, tokenId *big.Int, key string) (string, error) {
	return c.resolver.Text(opts, tokenId, key)
}

// SetText Sets the text information associated with a did
func (c *Core) SetText(opts *bind.TransactOpts, tokenId *big.Int, key, value string) (*types.Transaction, error) {
	return c.resolver.SetText(opts, tokenId, key, value)
}
