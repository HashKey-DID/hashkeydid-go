package hashkeydid_go

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetTokenIdByDid returns DID's tokenId by name
func (c *Core) GetTokenIdByDid(opts *bind.CallOpts, didName string) (int64, error) {
	tokenId, err := c.did.Did2TokenId(opts, didName)
	if err != nil {
		return 0, err
	}
	return tokenId.Int64(), nil
}

// GetDidByTokenId returns DID's name by tokenId
func (c *Core) GetDidByTokenId(opts *bind.CallOpts, tokenId int64) (string, error) {
	didName, err := c.did.TokenId2Did(opts, big.NewInt(tokenId))
	if err != nil {
		return "", err
	}
	return didName, nil
}

// GetAddrByDIDName returns DID's address by name
func (c *Core) GetAddrByDIDName(opts *bind.CallOpts, didName string) (common.Address, error) {
	tokenId, err := c.GetTokenIdByDid(opts, didName)
	if err != nil {
		return common.HexToAddress("0"), err
	}
	addr, err := c.did.OwnerOf(opts, big.NewInt(tokenId))
	if err != nil {
		return common.HexToAddress("0"), err
	}
	return addr, nil
}

// GetBlockChainAddress returns DID's binding addresses according to coinType
func (c *Core) GetBlockChainAddress(opts *bind.CallOpts, tokenId, coinType int64) ([]byte, error){
	addr, err := c.resolver.Addr(opts, big.NewInt(tokenId), big.NewInt(coinType))
	if err != nil {
		return addr, err
	}
	return addr, nil
}
