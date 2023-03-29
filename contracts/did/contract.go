// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package did

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DidV2StorageKYCInfo is an auto generated low-level Go binding around an user-defined struct.
type DidV2StorageKYCInfo struct {
	Status     bool
	UpdateTime *big.Int
	ExpireTime *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"AddAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"KYCProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"KYCId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"evidence\",\"type\":\"bytes\"}],\"name\":\"AddKYC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IssueDG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"IssueNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"RemoveAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"KYCProviders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"KYCIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"internalType\":\"structDidV2Storage.KYCInfo[]\",\"name\":\"KYCInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"evidences\",\"type\":\"bytes[]\"}],\"name\":\"addKYCs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI_\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiredTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"evidence\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"DGAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"evidence\",\"type\":\"bytes\"}],\"name\":\"claimDG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"NFTAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"evidence\",\"type\":\"bytes\"}],\"name\":\"claimDGNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deedGrainAddrToIssuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dgFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dgMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"did2TokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"didClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"didMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"didSync\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuthorizedAddrs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"KYCProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"KYCId\",\"type\":\"uint256\"}],\"name\":\"getKYCInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAddrAuthorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_evidence\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"issueDG\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_evidence\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"}],\"name\":\"issueNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiredTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"evidence\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"NFTAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sid\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"mintDGNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"DGAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"mintDGV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"DGAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintDGV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"avatar\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"KYCProviders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"KYCIds\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"updateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTime\",\"type\":\"uint256\"}],\"internalType\":\"structDidV2Storage.KYCInfo[]\",\"name\":\"KYCInfos\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"evidences\",\"type\":\"bytes[]\"}],\"name\":\"mintDidLZ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"setDGFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"setDGMinterAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"setDidMinterAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_didSync\",\"type\":\"address\"}],\"name\":\"setDidSync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"NFTAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"setNFTBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"NFTAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"setNFTSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolverAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"DGAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"}],\"name\":\"setTokenBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"DGAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"setTokenSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenId2Did\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"verifyDIDFormat\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AddrClaimed is a free data retrieval call binding the contract method 0xfce1554c.
//
// Solidity: function addrClaimed(address ) view returns(bool)
func (_Contract *ContractCaller) AddrClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addrClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddrClaimed is a free data retrieval call binding the contract method 0xfce1554c.
//
// Solidity: function addrClaimed(address ) view returns(bool)
func (_Contract *ContractSession) AddrClaimed(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AddrClaimed(&_Contract.CallOpts, arg0)
}

// AddrClaimed is a free data retrieval call binding the contract method 0xfce1554c.
//
// Solidity: function addrClaimed(address ) view returns(bool)
func (_Contract *ContractCallerSession) AddrClaimed(arg0 common.Address) (bool, error) {
	return _Contract.Contract.AddrClaimed(&_Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0xf259a29e.
//
// Solidity: function baseURI_() view returns(string)
func (_Contract *ContractCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "baseURI_")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0xf259a29e.
//
// Solidity: function baseURI_() view returns(string)
func (_Contract *ContractSession) BaseURI() (string, error) {
	return _Contract.Contract.BaseURI(&_Contract.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0xf259a29e.
//
// Solidity: function baseURI_() view returns(string)
func (_Contract *ContractCallerSession) BaseURI() (string, error) {
	return _Contract.Contract.BaseURI(&_Contract.CallOpts)
}

// DeedGrainAddrToIssuer is a free data retrieval call binding the contract method 0x4f9a1c8c.
//
// Solidity: function deedGrainAddrToIssuer(address ) view returns(address)
func (_Contract *ContractCaller) DeedGrainAddrToIssuer(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "deedGrainAddrToIssuer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeedGrainAddrToIssuer is a free data retrieval call binding the contract method 0x4f9a1c8c.
//
// Solidity: function deedGrainAddrToIssuer(address ) view returns(address)
func (_Contract *ContractSession) DeedGrainAddrToIssuer(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.DeedGrainAddrToIssuer(&_Contract.CallOpts, arg0)
}

// DeedGrainAddrToIssuer is a free data retrieval call binding the contract method 0x4f9a1c8c.
//
// Solidity: function deedGrainAddrToIssuer(address ) view returns(address)
func (_Contract *ContractCallerSession) DeedGrainAddrToIssuer(arg0 common.Address) (common.Address, error) {
	return _Contract.Contract.DeedGrainAddrToIssuer(&_Contract.CallOpts, arg0)
}

// DgFactory is a free data retrieval call binding the contract method 0x03919352.
//
// Solidity: function dgFactory() view returns(address)
func (_Contract *ContractCaller) DgFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dgFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DgFactory is a free data retrieval call binding the contract method 0x03919352.
//
// Solidity: function dgFactory() view returns(address)
func (_Contract *ContractSession) DgFactory() (common.Address, error) {
	return _Contract.Contract.DgFactory(&_Contract.CallOpts)
}

// DgFactory is a free data retrieval call binding the contract method 0x03919352.
//
// Solidity: function dgFactory() view returns(address)
func (_Contract *ContractCallerSession) DgFactory() (common.Address, error) {
	return _Contract.Contract.DgFactory(&_Contract.CallOpts)
}

// DgMinter is a free data retrieval call binding the contract method 0xaf457dde.
//
// Solidity: function dgMinter() view returns(address)
func (_Contract *ContractCaller) DgMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dgMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DgMinter is a free data retrieval call binding the contract method 0xaf457dde.
//
// Solidity: function dgMinter() view returns(address)
func (_Contract *ContractSession) DgMinter() (common.Address, error) {
	return _Contract.Contract.DgMinter(&_Contract.CallOpts)
}

// DgMinter is a free data retrieval call binding the contract method 0xaf457dde.
//
// Solidity: function dgMinter() view returns(address)
func (_Contract *ContractCallerSession) DgMinter() (common.Address, error) {
	return _Contract.Contract.DgMinter(&_Contract.CallOpts)
}

// Did2TokenId is a free data retrieval call binding the contract method 0x4bb903bd.
//
// Solidity: function did2TokenId(string ) view returns(uint256)
func (_Contract *ContractCaller) Did2TokenId(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "did2TokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Did2TokenId is a free data retrieval call binding the contract method 0x4bb903bd.
//
// Solidity: function did2TokenId(string ) view returns(uint256)
func (_Contract *ContractSession) Did2TokenId(arg0 string) (*big.Int, error) {
	return _Contract.Contract.Did2TokenId(&_Contract.CallOpts, arg0)
}

// Did2TokenId is a free data retrieval call binding the contract method 0x4bb903bd.
//
// Solidity: function did2TokenId(string ) view returns(uint256)
func (_Contract *ContractCallerSession) Did2TokenId(arg0 string) (*big.Int, error) {
	return _Contract.Contract.Did2TokenId(&_Contract.CallOpts, arg0)
}

// DidClaimed is a free data retrieval call binding the contract method 0x284d08ca.
//
// Solidity: function didClaimed(string ) view returns(bool)
func (_Contract *ContractCaller) DidClaimed(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "didClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidClaimed is a free data retrieval call binding the contract method 0x284d08ca.
//
// Solidity: function didClaimed(string ) view returns(bool)
func (_Contract *ContractSession) DidClaimed(arg0 string) (bool, error) {
	return _Contract.Contract.DidClaimed(&_Contract.CallOpts, arg0)
}

// DidClaimed is a free data retrieval call binding the contract method 0x284d08ca.
//
// Solidity: function didClaimed(string ) view returns(bool)
func (_Contract *ContractCallerSession) DidClaimed(arg0 string) (bool, error) {
	return _Contract.Contract.DidClaimed(&_Contract.CallOpts, arg0)
}

// DidMinter is a free data retrieval call binding the contract method 0x1d2bbf39.
//
// Solidity: function didMinter() view returns(address)
func (_Contract *ContractCaller) DidMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "didMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DidMinter is a free data retrieval call binding the contract method 0x1d2bbf39.
//
// Solidity: function didMinter() view returns(address)
func (_Contract *ContractSession) DidMinter() (common.Address, error) {
	return _Contract.Contract.DidMinter(&_Contract.CallOpts)
}

// DidMinter is a free data retrieval call binding the contract method 0x1d2bbf39.
//
// Solidity: function didMinter() view returns(address)
func (_Contract *ContractCallerSession) DidMinter() (common.Address, error) {
	return _Contract.Contract.DidMinter(&_Contract.CallOpts)
}

// DidSync is a free data retrieval call binding the contract method 0xdfcbbe5b.
//
// Solidity: function didSync() view returns(address)
func (_Contract *ContractCaller) DidSync(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "didSync")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DidSync is a free data retrieval call binding the contract method 0xdfcbbe5b.
//
// Solidity: function didSync() view returns(address)
func (_Contract *ContractSession) DidSync() (common.Address, error) {
	return _Contract.Contract.DidSync(&_Contract.CallOpts)
}

// DidSync is a free data retrieval call binding the contract method 0xdfcbbe5b.
//
// Solidity: function didSync() view returns(address)
func (_Contract *ContractCallerSession) DidSync() (common.Address, error) {
	return _Contract.Contract.DidSync(&_Contract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetAuthorizedAddrs is a free data retrieval call binding the contract method 0xfd5a6780.
//
// Solidity: function getAuthorizedAddrs(uint256 tokenId) view returns(address[])
func (_Contract *ContractCaller) GetAuthorizedAddrs(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAuthorizedAddrs", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedAddrs is a free data retrieval call binding the contract method 0xfd5a6780.
//
// Solidity: function getAuthorizedAddrs(uint256 tokenId) view returns(address[])
func (_Contract *ContractSession) GetAuthorizedAddrs(tokenId *big.Int) ([]common.Address, error) {
	return _Contract.Contract.GetAuthorizedAddrs(&_Contract.CallOpts, tokenId)
}

// GetAuthorizedAddrs is a free data retrieval call binding the contract method 0xfd5a6780.
//
// Solidity: function getAuthorizedAddrs(uint256 tokenId) view returns(address[])
func (_Contract *ContractCallerSession) GetAuthorizedAddrs(tokenId *big.Int) ([]common.Address, error) {
	return _Contract.Contract.GetAuthorizedAddrs(&_Contract.CallOpts, tokenId)
}

// GetKYCInfo is a free data retrieval call binding the contract method 0x7626884d.
//
// Solidity: function getKYCInfo(uint256 tokenId, address KYCProvider, uint256 KYCId) view returns(bool, uint256, uint256)
func (_Contract *ContractCaller) GetKYCInfo(opts *bind.CallOpts, tokenId *big.Int, KYCProvider common.Address, KYCId *big.Int) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getKYCInfo", tokenId, KYCProvider, KYCId)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetKYCInfo is a free data retrieval call binding the contract method 0x7626884d.
//
// Solidity: function getKYCInfo(uint256 tokenId, address KYCProvider, uint256 KYCId) view returns(bool, uint256, uint256)
func (_Contract *ContractSession) GetKYCInfo(tokenId *big.Int, KYCProvider common.Address, KYCId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetKYCInfo(&_Contract.CallOpts, tokenId, KYCProvider, KYCId)
}

// GetKYCInfo is a free data retrieval call binding the contract method 0x7626884d.
//
// Solidity: function getKYCInfo(uint256 tokenId, address KYCProvider, uint256 KYCId) view returns(bool, uint256, uint256)
func (_Contract *ContractCallerSession) GetKYCInfo(tokenId *big.Int, KYCProvider common.Address, KYCId *big.Int) (bool, *big.Int, *big.Int, error) {
	return _Contract.Contract.GetKYCInfo(&_Contract.CallOpts, tokenId, KYCProvider, KYCId)
}

// IsAddrAuthorized is a free data retrieval call binding the contract method 0xd08f28a6.
//
// Solidity: function isAddrAuthorized(uint256 tokenId, address addr) view returns(bool)
func (_Contract *ContractCaller) IsAddrAuthorized(opts *bind.CallOpts, tokenId *big.Int, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isAddrAuthorized", tokenId, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAddrAuthorized is a free data retrieval call binding the contract method 0xd08f28a6.
//
// Solidity: function isAddrAuthorized(uint256 tokenId, address addr) view returns(bool)
func (_Contract *ContractSession) IsAddrAuthorized(tokenId *big.Int, addr common.Address) (bool, error) {
	return _Contract.Contract.IsAddrAuthorized(&_Contract.CallOpts, tokenId, addr)
}

// IsAddrAuthorized is a free data retrieval call binding the contract method 0xd08f28a6.
//
// Solidity: function isAddrAuthorized(uint256 tokenId, address addr) view returns(bool)
func (_Contract *ContractCallerSession) IsAddrAuthorized(tokenId *big.Int, addr common.Address) (bool, error) {
	return _Contract.Contract.IsAddrAuthorized(&_Contract.CallOpts, tokenId, addr)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractSession) Resolver() (common.Address, error) {
	return _Contract.Contract.Resolver(&_Contract.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contract *ContractCallerSession) Resolver() (common.Address, error) {
	return _Contract.Contract.Resolver(&_Contract.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Contract *ContractCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Contract *ContractSession) Signer() (common.Address, error) {
	return _Contract.Contract.Signer(&_Contract.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Contract *ContractCallerSession) Signer() (common.Address, error) {
	return _Contract.Contract.Signer(&_Contract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, index)
}

// TokenId2Did is a free data retrieval call binding the contract method 0xc78826b7.
//
// Solidity: function tokenId2Did(uint256 ) view returns(string)
func (_Contract *ContractCaller) TokenId2Did(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenId2Did", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenId2Did is a free data retrieval call binding the contract method 0xc78826b7.
//
// Solidity: function tokenId2Did(uint256 ) view returns(string)
func (_Contract *ContractSession) TokenId2Did(arg0 *big.Int) (string, error) {
	return _Contract.Contract.TokenId2Did(&_Contract.CallOpts, arg0)
}

// TokenId2Did is a free data retrieval call binding the contract method 0xc78826b7.
//
// Solidity: function tokenId2Did(uint256 ) view returns(string)
func (_Contract *ContractCallerSession) TokenId2Did(arg0 *big.Int) (string, error) {
	return _Contract.Contract.TokenId2Did(&_Contract.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contract *ContractCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// VerifyDIDFormat is a free data retrieval call binding the contract method 0xf9251132.
//
// Solidity: function verifyDIDFormat(string did) pure returns(bool)
func (_Contract *ContractCaller) VerifyDIDFormat(opts *bind.CallOpts, did string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifyDIDFormat", did)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDIDFormat is a free data retrieval call binding the contract method 0xf9251132.
//
// Solidity: function verifyDIDFormat(string did) pure returns(bool)
func (_Contract *ContractSession) VerifyDIDFormat(did string) (bool, error) {
	return _Contract.Contract.VerifyDIDFormat(&_Contract.CallOpts, did)
}

// VerifyDIDFormat is a free data retrieval call binding the contract method 0xf9251132.
//
// Solidity: function verifyDIDFormat(string did) pure returns(bool)
func (_Contract *ContractCallerSession) VerifyDIDFormat(did string) (bool, error) {
	return _Contract.Contract.VerifyDIDFormat(&_Contract.CallOpts, did)
}

// AddAuth is a paid mutator transaction binding the contract method 0x9e9e864d.
//
// Solidity: function addAuth(uint256 tokenId, address addr) returns()
func (_Contract *ContractTransactor) AddAuth(opts *bind.TransactOpts, tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addAuth", tokenId, addr)
}

// AddAuth is a paid mutator transaction binding the contract method 0x9e9e864d.
//
// Solidity: function addAuth(uint256 tokenId, address addr) returns()
func (_Contract *ContractSession) AddAuth(tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddAuth(&_Contract.TransactOpts, tokenId, addr)
}

// AddAuth is a paid mutator transaction binding the contract method 0x9e9e864d.
//
// Solidity: function addAuth(uint256 tokenId, address addr) returns()
func (_Contract *ContractTransactorSession) AddAuth(tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddAuth(&_Contract.TransactOpts, tokenId, addr)
}

// AddKYCs is a paid mutator transaction binding the contract method 0xf6f38845.
//
// Solidity: function addKYCs(uint256 tokenId, address[] KYCProviders, uint256[] KYCIds, (bool,uint256,uint256)[] KYCInfos, bytes[] evidences) returns()
func (_Contract *ContractTransactor) AddKYCs(opts *bind.TransactOpts, tokenId *big.Int, KYCProviders []common.Address, KYCIds []*big.Int, KYCInfos []DidV2StorageKYCInfo, evidences [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addKYCs", tokenId, KYCProviders, KYCIds, KYCInfos, evidences)
}

// AddKYCs is a paid mutator transaction binding the contract method 0xf6f38845.
//
// Solidity: function addKYCs(uint256 tokenId, address[] KYCProviders, uint256[] KYCIds, (bool,uint256,uint256)[] KYCInfos, bytes[] evidences) returns()
func (_Contract *ContractSession) AddKYCs(tokenId *big.Int, KYCProviders []common.Address, KYCIds []*big.Int, KYCInfos []DidV2StorageKYCInfo, evidences [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.AddKYCs(&_Contract.TransactOpts, tokenId, KYCProviders, KYCIds, KYCInfos, evidences)
}

// AddKYCs is a paid mutator transaction binding the contract method 0xf6f38845.
//
// Solidity: function addKYCs(uint256 tokenId, address[] KYCProviders, uint256[] KYCIds, (bool,uint256,uint256)[] KYCInfos, bytes[] evidences) returns()
func (_Contract *ContractTransactorSession) AddKYCs(tokenId *big.Int, KYCProviders []common.Address, KYCIds []*big.Int, KYCInfos []DidV2StorageKYCInfo, evidences [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.AddKYCs(&_Contract.TransactOpts, tokenId, KYCProviders, KYCIds, KYCInfos, evidences)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0xccd3f749.
//
// Solidity: function claim(uint256 expiredTimestamp, string did, bytes evidence, string avatar) payable returns()
func (_Contract *ContractTransactor) Claim(opts *bind.TransactOpts, expiredTimestamp *big.Int, did string, evidence []byte, avatar string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claim", expiredTimestamp, did, evidence, avatar)
}

// Claim is a paid mutator transaction binding the contract method 0xccd3f749.
//
// Solidity: function claim(uint256 expiredTimestamp, string did, bytes evidence, string avatar) payable returns()
func (_Contract *ContractSession) Claim(expiredTimestamp *big.Int, did string, evidence []byte, avatar string) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, expiredTimestamp, did, evidence, avatar)
}

// Claim is a paid mutator transaction binding the contract method 0xccd3f749.
//
// Solidity: function claim(uint256 expiredTimestamp, string did, bytes evidence, string avatar) payable returns()
func (_Contract *ContractTransactorSession) Claim(expiredTimestamp *big.Int, did string, evidence []byte, avatar string) (*types.Transaction, error) {
	return _Contract.Contract.Claim(&_Contract.TransactOpts, expiredTimestamp, did, evidence, avatar)
}

// ClaimDG is a paid mutator transaction binding the contract method 0x8103f461.
//
// Solidity: function claimDG(address DGAddr, uint256 tokenId, bytes data, bytes evidence) returns()
func (_Contract *ContractTransactor) ClaimDG(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, data []byte, evidence []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDG", DGAddr, tokenId, data, evidence)
}

// ClaimDG is a paid mutator transaction binding the contract method 0x8103f461.
//
// Solidity: function claimDG(address DGAddr, uint256 tokenId, bytes data, bytes evidence) returns()
func (_Contract *ContractSession) ClaimDG(DGAddr common.Address, tokenId *big.Int, data []byte, evidence []byte) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDG(&_Contract.TransactOpts, DGAddr, tokenId, data, evidence)
}

// ClaimDG is a paid mutator transaction binding the contract method 0x8103f461.
//
// Solidity: function claimDG(address DGAddr, uint256 tokenId, bytes data, bytes evidence) returns()
func (_Contract *ContractTransactorSession) ClaimDG(DGAddr common.Address, tokenId *big.Int, data []byte, evidence []byte) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDG(&_Contract.TransactOpts, DGAddr, tokenId, data, evidence)
}

// ClaimDGNFT is a paid mutator transaction binding the contract method 0x79c3919f.
//
// Solidity: function claimDGNFT(address NFTAddr, uint256 sid, bytes evidence) returns()
func (_Contract *ContractTransactor) ClaimDGNFT(opts *bind.TransactOpts, NFTAddr common.Address, sid *big.Int, evidence []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimDGNFT", NFTAddr, sid, evidence)
}

// ClaimDGNFT is a paid mutator transaction binding the contract method 0x79c3919f.
//
// Solidity: function claimDGNFT(address NFTAddr, uint256 sid, bytes evidence) returns()
func (_Contract *ContractSession) ClaimDGNFT(NFTAddr common.Address, sid *big.Int, evidence []byte) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDGNFT(&_Contract.TransactOpts, NFTAddr, sid, evidence)
}

// ClaimDGNFT is a paid mutator transaction binding the contract method 0x79c3919f.
//
// Solidity: function claimDGNFT(address NFTAddr, uint256 sid, bytes evidence) returns()
func (_Contract *ContractTransactorSession) ClaimDGNFT(NFTAddr common.Address, sid *big.Int, evidence []byte) (*types.Transaction, error) {
	return _Contract.Contract.ClaimDGNFT(&_Contract.TransactOpts, NFTAddr, sid, evidence)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string _baseTokenURI, address _owner) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _baseTokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _name, _symbol, _baseTokenURI, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string _baseTokenURI, address _owner) returns()
func (_Contract *ContractSession) Initialize(_name string, _symbol string, _baseTokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _baseTokenURI, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x5c6d8da1.
//
// Solidity: function initialize(string _name, string _symbol, string _baseTokenURI, address _owner) returns()
func (_Contract *ContractTransactorSession) Initialize(_name string, _symbol string, _baseTokenURI string, _owner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _name, _symbol, _baseTokenURI, _owner)
}

// IssueDG is a paid mutator transaction binding the contract method 0x46a0e536.
//
// Solidity: function issueDG(string _name, string _symbol, string _baseUri, bytes _evidence, bool _transferable) returns()
func (_Contract *ContractTransactor) IssueDG(opts *bind.TransactOpts, _name string, _symbol string, _baseUri string, _evidence []byte, _transferable bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "issueDG", _name, _symbol, _baseUri, _evidence, _transferable)
}

// IssueDG is a paid mutator transaction binding the contract method 0x46a0e536.
//
// Solidity: function issueDG(string _name, string _symbol, string _baseUri, bytes _evidence, bool _transferable) returns()
func (_Contract *ContractSession) IssueDG(_name string, _symbol string, _baseUri string, _evidence []byte, _transferable bool) (*types.Transaction, error) {
	return _Contract.Contract.IssueDG(&_Contract.TransactOpts, _name, _symbol, _baseUri, _evidence, _transferable)
}

// IssueDG is a paid mutator transaction binding the contract method 0x46a0e536.
//
// Solidity: function issueDG(string _name, string _symbol, string _baseUri, bytes _evidence, bool _transferable) returns()
func (_Contract *ContractTransactorSession) IssueDG(_name string, _symbol string, _baseUri string, _evidence []byte, _transferable bool) (*types.Transaction, error) {
	return _Contract.Contract.IssueDG(&_Contract.TransactOpts, _name, _symbol, _baseUri, _evidence, _transferable)
}

// IssueNFT is a paid mutator transaction binding the contract method 0xe4b8f11f.
//
// Solidity: function issueNFT(string _name, string _symbol, string _baseUri, bytes _evidence, uint256 _supply) returns()
func (_Contract *ContractTransactor) IssueNFT(opts *bind.TransactOpts, _name string, _symbol string, _baseUri string, _evidence []byte, _supply *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "issueNFT", _name, _symbol, _baseUri, _evidence, _supply)
}

// IssueNFT is a paid mutator transaction binding the contract method 0xe4b8f11f.
//
// Solidity: function issueNFT(string _name, string _symbol, string _baseUri, bytes _evidence, uint256 _supply) returns()
func (_Contract *ContractSession) IssueNFT(_name string, _symbol string, _baseUri string, _evidence []byte, _supply *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IssueNFT(&_Contract.TransactOpts, _name, _symbol, _baseUri, _evidence, _supply)
}

// IssueNFT is a paid mutator transaction binding the contract method 0xe4b8f11f.
//
// Solidity: function issueNFT(string _name, string _symbol, string _baseUri, bytes _evidence, uint256 _supply) returns()
func (_Contract *ContractTransactorSession) IssueNFT(_name string, _symbol string, _baseUri string, _evidence []byte, _supply *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IssueNFT(&_Contract.TransactOpts, _name, _symbol, _baseUri, _evidence, _supply)
}

// Mint is a paid mutator transaction binding the contract method 0xa1c15c94.
//
// Solidity: function mint(address to, uint256 expiredTimestamp, string did, bytes evidence, string avatar) returns()
func (_Contract *ContractTransactor) Mint(opts *bind.TransactOpts, to common.Address, expiredTimestamp *big.Int, did string, evidence []byte, avatar string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mint", to, expiredTimestamp, did, evidence, avatar)
}

// Mint is a paid mutator transaction binding the contract method 0xa1c15c94.
//
// Solidity: function mint(address to, uint256 expiredTimestamp, string did, bytes evidence, string avatar) returns()
func (_Contract *ContractSession) Mint(to common.Address, expiredTimestamp *big.Int, did string, evidence []byte, avatar string) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, to, expiredTimestamp, did, evidence, avatar)
}

// Mint is a paid mutator transaction binding the contract method 0xa1c15c94.
//
// Solidity: function mint(address to, uint256 expiredTimestamp, string did, bytes evidence, string avatar) returns()
func (_Contract *ContractTransactorSession) Mint(to common.Address, expiredTimestamp *big.Int, did string, evidence []byte, avatar string) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, to, expiredTimestamp, did, evidence, avatar)
}

// MintDGNFT is a paid mutator transaction binding the contract method 0x258bd726.
//
// Solidity: function mintDGNFT(address NFTAddr, uint256 sid, address[] addrs) returns()
func (_Contract *ContractTransactor) MintDGNFT(opts *bind.TransactOpts, NFTAddr common.Address, sid *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintDGNFT", NFTAddr, sid, addrs)
}

// MintDGNFT is a paid mutator transaction binding the contract method 0x258bd726.
//
// Solidity: function mintDGNFT(address NFTAddr, uint256 sid, address[] addrs) returns()
func (_Contract *ContractSession) MintDGNFT(NFTAddr common.Address, sid *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MintDGNFT(&_Contract.TransactOpts, NFTAddr, sid, addrs)
}

// MintDGNFT is a paid mutator transaction binding the contract method 0x258bd726.
//
// Solidity: function mintDGNFT(address NFTAddr, uint256 sid, address[] addrs) returns()
func (_Contract *ContractTransactorSession) MintDGNFT(NFTAddr common.Address, sid *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MintDGNFT(&_Contract.TransactOpts, NFTAddr, sid, addrs)
}

// MintDGV1 is a paid mutator transaction binding the contract method 0x5f8c8d54.
//
// Solidity: function mintDGV1(address DGAddr, uint256 tokenId, address[] addrs) returns()
func (_Contract *ContractTransactor) MintDGV1(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintDGV1", DGAddr, tokenId, addrs)
}

// MintDGV1 is a paid mutator transaction binding the contract method 0x5f8c8d54.
//
// Solidity: function mintDGV1(address DGAddr, uint256 tokenId, address[] addrs) returns()
func (_Contract *ContractSession) MintDGV1(DGAddr common.Address, tokenId *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MintDGV1(&_Contract.TransactOpts, DGAddr, tokenId, addrs)
}

// MintDGV1 is a paid mutator transaction binding the contract method 0x5f8c8d54.
//
// Solidity: function mintDGV1(address DGAddr, uint256 tokenId, address[] addrs) returns()
func (_Contract *ContractTransactorSession) MintDGV1(DGAddr common.Address, tokenId *big.Int, addrs []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.MintDGV1(&_Contract.TransactOpts, DGAddr, tokenId, addrs)
}

// MintDGV2 is a paid mutator transaction binding the contract method 0xa80c59b6.
//
// Solidity: function mintDGV2(address DGAddr, uint256 tokenId, address[] addrs, bytes data) returns()
func (_Contract *ContractTransactor) MintDGV2(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, addrs []common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintDGV2", DGAddr, tokenId, addrs, data)
}

// MintDGV2 is a paid mutator transaction binding the contract method 0xa80c59b6.
//
// Solidity: function mintDGV2(address DGAddr, uint256 tokenId, address[] addrs, bytes data) returns()
func (_Contract *ContractSession) MintDGV2(DGAddr common.Address, tokenId *big.Int, addrs []common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.MintDGV2(&_Contract.TransactOpts, DGAddr, tokenId, addrs, data)
}

// MintDGV2 is a paid mutator transaction binding the contract method 0xa80c59b6.
//
// Solidity: function mintDGV2(address DGAddr, uint256 tokenId, address[] addrs, bytes data) returns()
func (_Contract *ContractTransactorSession) MintDGV2(DGAddr common.Address, tokenId *big.Int, addrs []common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.MintDGV2(&_Contract.TransactOpts, DGAddr, tokenId, addrs, data)
}

// MintDidLZ is a paid mutator transaction binding the contract method 0xd65f1da8.
//
// Solidity: function mintDidLZ(uint256 tokenId, address user, string did, string avatar, address[] KYCProviders, uint256[] KYCIds, (bool,uint256,uint256)[] KYCInfos, bytes[] evidences) returns()
func (_Contract *ContractTransactor) MintDidLZ(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, did string, avatar string, KYCProviders []common.Address, KYCIds []*big.Int, KYCInfos []DidV2StorageKYCInfo, evidences [][]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintDidLZ", tokenId, user, did, avatar, KYCProviders, KYCIds, KYCInfos, evidences)
}

// MintDidLZ is a paid mutator transaction binding the contract method 0xd65f1da8.
//
// Solidity: function mintDidLZ(uint256 tokenId, address user, string did, string avatar, address[] KYCProviders, uint256[] KYCIds, (bool,uint256,uint256)[] KYCInfos, bytes[] evidences) returns()
func (_Contract *ContractSession) MintDidLZ(tokenId *big.Int, user common.Address, did string, avatar string, KYCProviders []common.Address, KYCIds []*big.Int, KYCInfos []DidV2StorageKYCInfo, evidences [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.MintDidLZ(&_Contract.TransactOpts, tokenId, user, did, avatar, KYCProviders, KYCIds, KYCInfos, evidences)
}

// MintDidLZ is a paid mutator transaction binding the contract method 0xd65f1da8.
//
// Solidity: function mintDidLZ(uint256 tokenId, address user, string did, string avatar, address[] KYCProviders, uint256[] KYCIds, (bool,uint256,uint256)[] KYCInfos, bytes[] evidences) returns()
func (_Contract *ContractTransactorSession) MintDidLZ(tokenId *big.Int, user common.Address, did string, avatar string, KYCProviders []common.Address, KYCIds []*big.Int, KYCInfos []DidV2StorageKYCInfo, evidences [][]byte) (*types.Transaction, error) {
	return _Contract.Contract.MintDidLZ(&_Contract.TransactOpts, tokenId, user, did, avatar, KYCProviders, KYCIds, KYCInfos, evidences)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0x467034df.
//
// Solidity: function removeAuth(uint256 tokenId, address addr) returns()
func (_Contract *ContractTransactor) RemoveAuth(opts *bind.TransactOpts, tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeAuth", tokenId, addr)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0x467034df.
//
// Solidity: function removeAuth(uint256 tokenId, address addr) returns()
func (_Contract *ContractSession) RemoveAuth(tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAuth(&_Contract.TransactOpts, tokenId, addr)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0x467034df.
//
// Solidity: function removeAuth(uint256 tokenId, address addr) returns()
func (_Contract *ContractTransactorSession) RemoveAuth(tokenId *big.Int, addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAuth(&_Contract.TransactOpts, tokenId, addr)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Contract *ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Contract *ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom0(&_Contract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetDGFactory is a paid mutator transaction binding the contract method 0x381c5336.
//
// Solidity: function setDGFactory(address factory) returns()
func (_Contract *ContractTransactor) SetDGFactory(opts *bind.TransactOpts, factory common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDGFactory", factory)
}

// SetDGFactory is a paid mutator transaction binding the contract method 0x381c5336.
//
// Solidity: function setDGFactory(address factory) returns()
func (_Contract *ContractSession) SetDGFactory(factory common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDGFactory(&_Contract.TransactOpts, factory)
}

// SetDGFactory is a paid mutator transaction binding the contract method 0x381c5336.
//
// Solidity: function setDGFactory(address factory) returns()
func (_Contract *ContractTransactorSession) SetDGFactory(factory common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDGFactory(&_Contract.TransactOpts, factory)
}

// SetDGMinterAddr is a paid mutator transaction binding the contract method 0x03a19ae4.
//
// Solidity: function setDGMinterAddr(address minter) returns()
func (_Contract *ContractTransactor) SetDGMinterAddr(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDGMinterAddr", minter)
}

// SetDGMinterAddr is a paid mutator transaction binding the contract method 0x03a19ae4.
//
// Solidity: function setDGMinterAddr(address minter) returns()
func (_Contract *ContractSession) SetDGMinterAddr(minter common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDGMinterAddr(&_Contract.TransactOpts, minter)
}

// SetDGMinterAddr is a paid mutator transaction binding the contract method 0x03a19ae4.
//
// Solidity: function setDGMinterAddr(address minter) returns()
func (_Contract *ContractTransactorSession) SetDGMinterAddr(minter common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDGMinterAddr(&_Contract.TransactOpts, minter)
}

// SetDidMinterAddr is a paid mutator transaction binding the contract method 0x8fd0f98e.
//
// Solidity: function setDidMinterAddr(address minter) returns()
func (_Contract *ContractTransactor) SetDidMinterAddr(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDidMinterAddr", minter)
}

// SetDidMinterAddr is a paid mutator transaction binding the contract method 0x8fd0f98e.
//
// Solidity: function setDidMinterAddr(address minter) returns()
func (_Contract *ContractSession) SetDidMinterAddr(minter common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDidMinterAddr(&_Contract.TransactOpts, minter)
}

// SetDidMinterAddr is a paid mutator transaction binding the contract method 0x8fd0f98e.
//
// Solidity: function setDidMinterAddr(address minter) returns()
func (_Contract *ContractTransactorSession) SetDidMinterAddr(minter common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDidMinterAddr(&_Contract.TransactOpts, minter)
}

// SetDidSync is a paid mutator transaction binding the contract method 0x6423ef18.
//
// Solidity: function setDidSync(address _didSync) returns()
func (_Contract *ContractTransactor) SetDidSync(opts *bind.TransactOpts, _didSync common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDidSync", _didSync)
}

// SetDidSync is a paid mutator transaction binding the contract method 0x6423ef18.
//
// Solidity: function setDidSync(address _didSync) returns()
func (_Contract *ContractSession) SetDidSync(_didSync common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDidSync(&_Contract.TransactOpts, _didSync)
}

// SetDidSync is a paid mutator transaction binding the contract method 0x6423ef18.
//
// Solidity: function setDidSync(address _didSync) returns()
func (_Contract *ContractTransactorSession) SetDidSync(_didSync common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetDidSync(&_Contract.TransactOpts, _didSync)
}

// SetNFTBaseUri is a paid mutator transaction binding the contract method 0xa135b302.
//
// Solidity: function setNFTBaseUri(address NFTAddr, string baseUri) returns()
func (_Contract *ContractTransactor) SetNFTBaseUri(opts *bind.TransactOpts, NFTAddr common.Address, baseUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNFTBaseUri", NFTAddr, baseUri)
}

// SetNFTBaseUri is a paid mutator transaction binding the contract method 0xa135b302.
//
// Solidity: function setNFTBaseUri(address NFTAddr, string baseUri) returns()
func (_Contract *ContractSession) SetNFTBaseUri(NFTAddr common.Address, baseUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetNFTBaseUri(&_Contract.TransactOpts, NFTAddr, baseUri)
}

// SetNFTBaseUri is a paid mutator transaction binding the contract method 0xa135b302.
//
// Solidity: function setNFTBaseUri(address NFTAddr, string baseUri) returns()
func (_Contract *ContractTransactorSession) SetNFTBaseUri(NFTAddr common.Address, baseUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetNFTBaseUri(&_Contract.TransactOpts, NFTAddr, baseUri)
}

// SetNFTSupply is a paid mutator transaction binding the contract method 0x7a9e526b.
//
// Solidity: function setNFTSupply(address NFTAddr, uint256 supply) returns()
func (_Contract *ContractTransactor) SetNFTSupply(opts *bind.TransactOpts, NFTAddr common.Address, supply *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNFTSupply", NFTAddr, supply)
}

// SetNFTSupply is a paid mutator transaction binding the contract method 0x7a9e526b.
//
// Solidity: function setNFTSupply(address NFTAddr, uint256 supply) returns()
func (_Contract *ContractSession) SetNFTSupply(NFTAddr common.Address, supply *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetNFTSupply(&_Contract.TransactOpts, NFTAddr, supply)
}

// SetNFTSupply is a paid mutator transaction binding the contract method 0x7a9e526b.
//
// Solidity: function setNFTSupply(address NFTAddr, uint256 supply) returns()
func (_Contract *ContractTransactorSession) SetNFTSupply(NFTAddr common.Address, supply *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetNFTSupply(&_Contract.TransactOpts, NFTAddr, supply)
}

// SetResolverAddr is a paid mutator transaction binding the contract method 0xee544878.
//
// Solidity: function setResolverAddr(address _resolver) returns()
func (_Contract *ContractTransactor) SetResolverAddr(opts *bind.TransactOpts, _resolver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setResolverAddr", _resolver)
}

// SetResolverAddr is a paid mutator transaction binding the contract method 0xee544878.
//
// Solidity: function setResolverAddr(address _resolver) returns()
func (_Contract *ContractSession) SetResolverAddr(_resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetResolverAddr(&_Contract.TransactOpts, _resolver)
}

// SetResolverAddr is a paid mutator transaction binding the contract method 0xee544878.
//
// Solidity: function setResolverAddr(address _resolver) returns()
func (_Contract *ContractTransactorSession) SetResolverAddr(_resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetResolverAddr(&_Contract.TransactOpts, _resolver)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_Contract *ContractTransactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_Contract *ContractSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSigner(&_Contract.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_Contract *ContractTransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSigner(&_Contract.TransactOpts, signer_)
}

// SetTokenBaseUri is a paid mutator transaction binding the contract method 0xc3535aaa.
//
// Solidity: function setTokenBaseUri(address DGAddr, string baseUri) returns()
func (_Contract *ContractTransactor) SetTokenBaseUri(opts *bind.TransactOpts, DGAddr common.Address, baseUri string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTokenBaseUri", DGAddr, baseUri)
}

// SetTokenBaseUri is a paid mutator transaction binding the contract method 0xc3535aaa.
//
// Solidity: function setTokenBaseUri(address DGAddr, string baseUri) returns()
func (_Contract *ContractSession) SetTokenBaseUri(DGAddr common.Address, baseUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetTokenBaseUri(&_Contract.TransactOpts, DGAddr, baseUri)
}

// SetTokenBaseUri is a paid mutator transaction binding the contract method 0xc3535aaa.
//
// Solidity: function setTokenBaseUri(address DGAddr, string baseUri) returns()
func (_Contract *ContractTransactorSession) SetTokenBaseUri(DGAddr common.Address, baseUri string) (*types.Transaction, error) {
	return _Contract.Contract.SetTokenBaseUri(&_Contract.TransactOpts, DGAddr, baseUri)
}

// SetTokenSupply is a paid mutator transaction binding the contract method 0x7dea57b9.
//
// Solidity: function setTokenSupply(address DGAddr, uint256 tokenId, uint256 supply) returns()
func (_Contract *ContractTransactor) SetTokenSupply(opts *bind.TransactOpts, DGAddr common.Address, tokenId *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTokenSupply", DGAddr, tokenId, supply)
}

// SetTokenSupply is a paid mutator transaction binding the contract method 0x7dea57b9.
//
// Solidity: function setTokenSupply(address DGAddr, uint256 tokenId, uint256 supply) returns()
func (_Contract *ContractSession) SetTokenSupply(DGAddr common.Address, tokenId *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetTokenSupply(&_Contract.TransactOpts, DGAddr, tokenId, supply)
}

// SetTokenSupply is a paid mutator transaction binding the contract method 0x7dea57b9.
//
// Solidity: function setTokenSupply(address DGAddr, uint256 tokenId, uint256 supply) returns()
func (_Contract *ContractTransactorSession) SetTokenSupply(DGAddr common.Address, tokenId *big.Int, supply *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetTokenSupply(&_Contract.TransactOpts, DGAddr, tokenId, supply)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractAddAuthIterator is returned from FilterAddAuth and is used to iterate over the raw logs and unpacked data for AddAuth events raised by the Contract contract.
type ContractAddAuthIterator struct {
	Event *ContractAddAuth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAddAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddAuth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAddAuth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAddAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddAuth represents a AddAuth event raised by the Contract contract.
type ContractAddAuth struct {
	Did      string
	Addr     common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddAuth is a free log retrieval operation binding the contract event 0x6c2c4ba0077a62001b35851f5d6d20db1bd91dd851a976751bda124fe5f0bade.
//
// Solidity: event AddAuth(string did, address indexed addr, address indexed operator)
func (_Contract *ContractFilterer) FilterAddAuth(opts *bind.FilterOpts, addr []common.Address, operator []common.Address) (*ContractAddAuthIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddAuth", addrRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddAuthIterator{contract: _Contract.contract, event: "AddAuth", logs: logs, sub: sub}, nil
}

// WatchAddAuth is a free log subscription operation binding the contract event 0x6c2c4ba0077a62001b35851f5d6d20db1bd91dd851a976751bda124fe5f0bade.
//
// Solidity: event AddAuth(string did, address indexed addr, address indexed operator)
func (_Contract *ContractFilterer) WatchAddAuth(opts *bind.WatchOpts, sink chan<- *ContractAddAuth, addr []common.Address, operator []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddAuth", addrRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddAuth)
				if err := _Contract.contract.UnpackLog(event, "AddAuth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddAuth is a log parse operation binding the contract event 0x6c2c4ba0077a62001b35851f5d6d20db1bd91dd851a976751bda124fe5f0bade.
//
// Solidity: event AddAuth(string did, address indexed addr, address indexed operator)
func (_Contract *ContractFilterer) ParseAddAuth(log types.Log) (*ContractAddAuth, error) {
	event := new(ContractAddAuth)
	if err := _Contract.contract.UnpackLog(event, "AddAuth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAddKYCIterator is returned from FilterAddKYC and is used to iterate over the raw logs and unpacked data for AddKYC events raised by the Contract contract.
type ContractAddKYCIterator struct {
	Event *ContractAddKYC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAddKYCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddKYC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAddKYC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAddKYCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddKYCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddKYC represents a AddKYC event raised by the Contract contract.
type ContractAddKYC struct {
	TokenId     *big.Int
	KYCProvider common.Address
	KYCId       *big.Int
	Status      bool
	UpdateTime  *big.Int
	ExpireTime  *big.Int
	Evidence    []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddKYC is a free log retrieval operation binding the contract event 0xb3220829e8a54b0d9795c82f6d25b6695376dcc7bda1031d0bdbe648ff7f5bb8.
//
// Solidity: event AddKYC(uint256 tokenId, address KYCProvider, uint256 KYCId, bool status, uint256 updateTime, uint256 expireTime, bytes evidence)
func (_Contract *ContractFilterer) FilterAddKYC(opts *bind.FilterOpts) (*ContractAddKYCIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddKYC")
	if err != nil {
		return nil, err
	}
	return &ContractAddKYCIterator{contract: _Contract.contract, event: "AddKYC", logs: logs, sub: sub}, nil
}

// WatchAddKYC is a free log subscription operation binding the contract event 0xb3220829e8a54b0d9795c82f6d25b6695376dcc7bda1031d0bdbe648ff7f5bb8.
//
// Solidity: event AddKYC(uint256 tokenId, address KYCProvider, uint256 KYCId, bool status, uint256 updateTime, uint256 expireTime, bytes evidence)
func (_Contract *ContractFilterer) WatchAddKYC(opts *bind.WatchOpts, sink chan<- *ContractAddKYC) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddKYC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddKYC)
				if err := _Contract.contract.UnpackLog(event, "AddKYC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddKYC is a log parse operation binding the contract event 0xb3220829e8a54b0d9795c82f6d25b6695376dcc7bda1031d0bdbe648ff7f5bb8.
//
// Solidity: event AddKYC(uint256 tokenId, address KYCProvider, uint256 KYCId, bool status, uint256 updateTime, uint256 expireTime, bytes evidence)
func (_Contract *ContractFilterer) ParseAddKYC(log types.Log) (*ContractAddKYC, error) {
	event := new(ContractAddKYC)
	if err := _Contract.contract.UnpackLog(event, "AddKYC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contract *ContractFilterer) ParseApprovalForAll(log types.Log) (*ContractApprovalForAll, error) {
	event := new(ContractApprovalForAll)
	if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIssueDGIterator is returned from FilterIssueDG and is used to iterate over the raw logs and unpacked data for IssueDG events raised by the Contract contract.
type ContractIssueDGIterator struct {
	Event *ContractIssueDG // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIssueDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIssueDG)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIssueDG)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIssueDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIssueDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIssueDG represents a IssueDG event raised by the Contract contract.
type ContractIssueDG struct {
	Arg0 common.Address
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIssueDG is a free log retrieval operation binding the contract event 0xc05872623594b1c2574e0531d0cc06b56ceb48baddce03b13163aa822ddfd52c.
//
// Solidity: event IssueDG(address indexed arg0, address arg1)
func (_Contract *ContractFilterer) FilterIssueDG(opts *bind.FilterOpts, arg0 []common.Address) (*ContractIssueDGIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "IssueDG", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &ContractIssueDGIterator{contract: _Contract.contract, event: "IssueDG", logs: logs, sub: sub}, nil
}

// WatchIssueDG is a free log subscription operation binding the contract event 0xc05872623594b1c2574e0531d0cc06b56ceb48baddce03b13163aa822ddfd52c.
//
// Solidity: event IssueDG(address indexed arg0, address arg1)
func (_Contract *ContractFilterer) WatchIssueDG(opts *bind.WatchOpts, sink chan<- *ContractIssueDG, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "IssueDG", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIssueDG)
				if err := _Contract.contract.UnpackLog(event, "IssueDG", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssueDG is a log parse operation binding the contract event 0xc05872623594b1c2574e0531d0cc06b56ceb48baddce03b13163aa822ddfd52c.
//
// Solidity: event IssueDG(address indexed arg0, address arg1)
func (_Contract *ContractFilterer) ParseIssueDG(log types.Log) (*ContractIssueDG, error) {
	event := new(ContractIssueDG)
	if err := _Contract.contract.UnpackLog(event, "IssueDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIssueNFTIterator is returned from FilterIssueNFT and is used to iterate over the raw logs and unpacked data for IssueNFT events raised by the Contract contract.
type ContractIssueNFTIterator struct {
	Event *ContractIssueNFT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIssueNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIssueNFT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIssueNFT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIssueNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIssueNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIssueNFT represents a IssueNFT event raised by the Contract contract.
type ContractIssueNFT struct {
	Arg0 common.Address
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIssueNFT is a free log retrieval operation binding the contract event 0xf9d4b55952c081f85101e9a2f8c6d5843afd8c96692b343ae78aa9f653090c39.
//
// Solidity: event IssueNFT(address indexed arg0, address arg1)
func (_Contract *ContractFilterer) FilterIssueNFT(opts *bind.FilterOpts, arg0 []common.Address) (*ContractIssueNFTIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "IssueNFT", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &ContractIssueNFTIterator{contract: _Contract.contract, event: "IssueNFT", logs: logs, sub: sub}, nil
}

// WatchIssueNFT is a free log subscription operation binding the contract event 0xf9d4b55952c081f85101e9a2f8c6d5843afd8c96692b343ae78aa9f653090c39.
//
// Solidity: event IssueNFT(address indexed arg0, address arg1)
func (_Contract *ContractFilterer) WatchIssueNFT(opts *bind.WatchOpts, sink chan<- *ContractIssueNFT, arg0 []common.Address) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "IssueNFT", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIssueNFT)
				if err := _Contract.contract.UnpackLog(event, "IssueNFT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssueNFT is a log parse operation binding the contract event 0xf9d4b55952c081f85101e9a2f8c6d5843afd8c96692b343ae78aa9f653090c39.
//
// Solidity: event IssueNFT(address indexed arg0, address arg1)
func (_Contract *ContractFilterer) ParseIssueNFT(log types.Log) (*ContractIssueNFT, error) {
	event := new(ContractIssueNFT)
	if err := _Contract.contract.UnpackLog(event, "IssueNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Contract contract.
type ContractOwnerChangedIterator struct {
	Event *ContractOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnerChanged represents a OwnerChanged event raised by the Contract contract.
type ContractOwnerChanged struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address previousOwner, address newOwner)
func (_Contract *ContractFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ContractOwnerChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ContractOwnerChangedIterator{contract: _Contract.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address previousOwner, address newOwner)
func (_Contract *ContractFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ContractOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnerChanged)
				if err := _Contract.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address previousOwner, address newOwner)
func (_Contract *ContractFilterer) ParseOwnerChanged(log types.Log) (*ContractOwnerChanged, error) {
	event := new(ContractOwnerChanged)
	if err := _Contract.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRemoveAuthIterator is returned from FilterRemoveAuth and is used to iterate over the raw logs and unpacked data for RemoveAuth events raised by the Contract contract.
type ContractRemoveAuthIterator struct {
	Event *ContractRemoveAuth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRemoveAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRemoveAuth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRemoveAuth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRemoveAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRemoveAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRemoveAuth represents a RemoveAuth event raised by the Contract contract.
type ContractRemoveAuth struct {
	Did      string
	Addr     common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuth is a free log retrieval operation binding the contract event 0x33b3e038c0e28aca0b92d455e5141d87b1801617ba321b7ccce20b09f836a65c.
//
// Solidity: event RemoveAuth(string did, address indexed addr, address indexed operator)
func (_Contract *ContractFilterer) FilterRemoveAuth(opts *bind.FilterOpts, addr []common.Address, operator []common.Address) (*ContractRemoveAuthIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RemoveAuth", addrRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractRemoveAuthIterator{contract: _Contract.contract, event: "RemoveAuth", logs: logs, sub: sub}, nil
}

// WatchRemoveAuth is a free log subscription operation binding the contract event 0x33b3e038c0e28aca0b92d455e5141d87b1801617ba321b7ccce20b09f836a65c.
//
// Solidity: event RemoveAuth(string did, address indexed addr, address indexed operator)
func (_Contract *ContractFilterer) WatchRemoveAuth(opts *bind.WatchOpts, sink chan<- *ContractRemoveAuth, addr []common.Address, operator []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RemoveAuth", addrRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRemoveAuth)
				if err := _Contract.contract.UnpackLog(event, "RemoveAuth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveAuth is a log parse operation binding the contract event 0x33b3e038c0e28aca0b92d455e5141d87b1801617ba321b7ccce20b09f836a65c.
//
// Solidity: event RemoveAuth(string did, address indexed addr, address indexed operator)
func (_Contract *ContractFilterer) ParseRemoveAuth(log types.Log) (*ContractRemoveAuth, error) {
	event := new(ContractRemoveAuth)
	if err := _Contract.contract.UnpackLog(event, "RemoveAuth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
