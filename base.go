package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetTokenIdByDid returns DID's tokenId by name
func (c *Core) GetTokenIdByDid(opts *bind.CallOpts, didName string) (string, error) {
	tokenId, err := c.did.Did2TokenId(opts, didName)
	if err != nil {
		return "", err
	}
	if tokenId != nil && tokenId.Int64() == 0 {
		return "", ErrDidNotClaimed
	}
	return tokenId.String(), nil
}

// GetDidByTokenId returns DID's name by tokenId
func (c *Core) GetDidByTokenId(opts *bind.CallOpts, tokenId string) (string, error) {
	id, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return "", ErrInvalidTokenId
	}
	didName, err := c.did.TokenId2Did(opts, id)
	if err != nil {
		return "", err
	}
	if didName == "" {
		return "", ErrInvalidTokenId
	}
	return didName, nil
}

// GetAddrByDIDName returns DID's address by name
func (c *Core) GetAddrByDIDName(opts *bind.CallOpts, didName string) (common.Address, error) {
	tokenId, err := c.GetTokenIdByDid(opts, didName)
	if err != nil {
		return common.Address{}, err
	}
	id, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return common.Address{}, ErrInvalidTokenId
	}
	return c.did.OwnerOf(opts, id)
}

// GetBlockChainAddress returns DID's binding addresses according to coinType
func (c *Core) GetBlockChainAddress(opts *bind.CallOpts, tokenId string, coinType int64) ([]byte, error) {
	id, ok := new(big.Int).SetString(tokenId, 10)
	if !ok {
		return []byte{}, ErrInvalidTokenId
	}
	return c.resolver.Addr(opts, id, big.NewInt(coinType))
}
