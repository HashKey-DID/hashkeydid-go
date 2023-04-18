package hashkeydid_go

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/hashkeydid/hashkeydid-go/contracts/did"
	"github.com/hashkeydid/hashkeydid-go/contracts/resolver"
)

// Core is structure for basic interaction with contracts
type Core struct {
	client *ethclient.Client

	did      *did.Contract
	resolver *resolver.Contract
}

// NewDIDCore Creates a did-go core with give rpc url
func NewDIDCore(rpcUrl string) (*Core, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("Connect with %s failed: %s\n", rpcUrl, err)
	}

	chainIdBig, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	chainInfo, ok := ChainList[chainIdBig.String()]
	if !ok {
		return nil, ErrNotSupportYet
	}

	didInstance, err := did.NewContract(common.HexToAddress(chainInfo.DIDContract), client)
	if err != nil {
		return nil, fmt.Errorf("New a did instance failed: %s\n", err)
	}

	resolverInstance, err := resolver.NewContract(common.HexToAddress(chainInfo.ResolveContract), client)
	if err != nil {
		return nil, fmt.Errorf("New a did resolver instance failed: %s\n", err)
	}

	chainCore := &Core{
		client:   client,
		did:      didInstance,
		resolver: resolverInstance,
	}

	return chainCore, nil
}
