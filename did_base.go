package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetTokenIdByDid returns DID's tokenId by name
func (c *Core) GetTokenIdByDid(opts *bind.CallOpts, didName string) (*big.Int, error) {
	tokenId, err := c.did.Did2TokenId(opts, didName)
	if err != nil {
		return Big0, err
	}

	if tokenId.Int64() == 0 {
		return Big0, ErrDidNotClaimed
	}
	return tokenId, nil
}

// GetDidByTokenId returns DID's name by tokenId
func (c *Core) GetDidByTokenId(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	didName, err := c.did.TokenId2Did(opts, tokenId)
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
	return c.did.OwnerOf(opts, tokenId)
}

// GetBlockChainAddress returns DID's binding addresses according to coinType
func (c *Core) GetBlockChainAddress(opts *bind.CallOpts, tokenId *big.Int, coinType *big.Int) ([]byte, error) {
	return c.resolver.Addr(opts, tokenId, coinType)
}

// VerifyDIDFormat returns Verify did format
func (c *Core) VerifyDIDFormat(opts *bind.CallOpts, did string) (bool, error) {
	return c.did.VerifyDIDFormat(opts, did)
}
