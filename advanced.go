package hashkeydid_go

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetDIDNameByAddr returns the did name by address when user set reverse
func (c *Core) GetDIDNameByAddr(opts *bind.CallOpts, address common.Address) (string, error) {
	claimed, err := c.did.AddrClaimed(opts, address)
	if err != nil {
		return "", err
	}
	if !claimed {
		return "", ErrAddrNotClaimed
	}
	name, err := c.resolver.Name(opts, address)
	if err != nil {
		if strings.Contains(err.Error(), "this addr has not set reverse record") {
			return "", ErrAddrNotSetReverse
		}
		return "", err
	}
	return name, nil
}

// GetDIDNameByAddrForce returns the did name by address
func (c *Core) GetDIDNameByAddrForce(opts *bind.CallOpts, address common.Address) (string, error) {
	claimed, err := c.did.AddrClaimed(opts, address)
	if err != nil {
		return "", err
	}
	if !claimed {
		return "", ErrAddrNotClaimed
	}
	tokenId, err := c.did.TokenOfOwnerByIndex(opts, address, Big0)
	if err != nil {
		return "", err
	}
	name, err := c.did.TokenId2Did(opts, tokenId)
	if err != nil {
		return "", err
	}
	return name, nil
}
