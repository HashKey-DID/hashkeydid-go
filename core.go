package hashkeydid_go

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hashkey-did/hashkeydid-go/contracts/did"
	"github.com/hashkey-did/hashkeydid-go/contracts/resolver"
)

// Core is structure for basic interaction with contracts
type Core struct {
	client   *ethclient.Client
	did      *did.Contract
	resolver *resolver.Contract
}

// NewDIDCore Creates a did-go core with give addresses
func NewDIDCore(rpcUrl string, didAddr string, resolverAddr string) (*Core, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("Connect with %s failed: %s\n", rpcUrl, err)
	}
	didInstance, err := did.NewContract(common.HexToAddress(didAddr), client)
	if err != nil {
		return nil, fmt.Errorf("New a did instance failed: %s\n", err)
	}
	resolverInstance, err := resolver.NewContract(common.HexToAddress(resolverAddr), client)
	if err != nil {
		return nil, fmt.Errorf("New a did resolver instance failed: %s\n", err)
	}
	return &Core{
		client:   client,
		did:      didInstance,
		resolver: resolverInstance,
	}, nil
}
