// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolver

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"ContentHashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"contentHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"DIDAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_addr\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"}],\"name\":\"setContentHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isReverse\",\"type\":\"bool\"}],\"name\":\"setReverse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Addr is a free data retrieval call binding the contract method 0x724474cd.
//
// Solidity: function addr(uint256 tokenId, uint256 coinType) view returns(bytes)
func (_Contract *ContractCaller) Addr(opts *bind.CallOpts, tokenId *big.Int, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addr", tokenId, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x724474cd.
//
// Solidity: function addr(uint256 tokenId, uint256 coinType) view returns(bytes)
func (_Contract *ContractSession) Addr(tokenId *big.Int, coinType *big.Int) ([]byte, error) {
	return _Contract.Contract.Addr(&_Contract.CallOpts, tokenId, coinType)
}

// Addr is a free data retrieval call binding the contract method 0x724474cd.
//
// Solidity: function addr(uint256 tokenId, uint256 coinType) view returns(bytes)
func (_Contract *ContractCallerSession) Addr(tokenId *big.Int, coinType *big.Int) ([]byte, error) {
	return _Contract.Contract.Addr(&_Contract.CallOpts, tokenId, coinType)
}

// ContentHash is a free data retrieval call binding the contract method 0xf3e0c290.
//
// Solidity: function contentHash(uint256 tokenId) view returns(bytes)
func (_Contract *ContractCaller) ContentHash(opts *bind.CallOpts, tokenId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "contentHash", tokenId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ContentHash is a free data retrieval call binding the contract method 0xf3e0c290.
//
// Solidity: function contentHash(uint256 tokenId) view returns(bytes)
func (_Contract *ContractSession) ContentHash(tokenId *big.Int) ([]byte, error) {
	return _Contract.Contract.ContentHash(&_Contract.CallOpts, tokenId)
}

// ContentHash is a free data retrieval call binding the contract method 0xf3e0c290.
//
// Solidity: function contentHash(uint256 tokenId) view returns(bytes)
func (_Contract *ContractCallerSession) ContentHash(tokenId *big.Int) ([]byte, error) {
	return _Contract.Contract.ContentHash(&_Contract.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x01984892.
//
// Solidity: function name(address _addr) view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts, _addr common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name", _addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x01984892.
//
// Solidity: function name(address _addr) view returns(string)
func (_Contract *ContractSession) Name(_addr common.Address) (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x01984892.
//
// Solidity: function name(address _addr) view returns(string)
func (_Contract *ContractCallerSession) Name(_addr common.Address) (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts, _addr)
}

// Pubkey is a free data retrieval call binding the contract method 0x2cdfb6ae.
//
// Solidity: function pubkey(uint256 tokenId) view returns(bytes32 x, bytes32 y)
func (_Contract *ContractCaller) Pubkey(opts *bind.CallOpts, tokenId *big.Int) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pubkey", tokenId)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0x2cdfb6ae.
//
// Solidity: function pubkey(uint256 tokenId) view returns(bytes32 x, bytes32 y)
func (_Contract *ContractSession) Pubkey(tokenId *big.Int) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Contract.Contract.Pubkey(&_Contract.CallOpts, tokenId)
}

// Pubkey is a free data retrieval call binding the contract method 0x2cdfb6ae.
//
// Solidity: function pubkey(uint256 tokenId) view returns(bytes32 x, bytes32 y)
func (_Contract *ContractCallerSession) Pubkey(tokenId *big.Int) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Contract.Contract.Pubkey(&_Contract.CallOpts, tokenId)
}

// Text is a free data retrieval call binding the contract method 0x308e3386.
//
// Solidity: function text(uint256 tokenId, string key) view returns(string)
func (_Contract *ContractCaller) Text(opts *bind.CallOpts, tokenId *big.Int, key string) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "text", tokenId, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x308e3386.
//
// Solidity: function text(uint256 tokenId, string key) view returns(string)
func (_Contract *ContractSession) Text(tokenId *big.Int, key string) (string, error) {
	return _Contract.Contract.Text(&_Contract.CallOpts, tokenId, key)
}

// Text is a free data retrieval call binding the contract method 0x308e3386.
//
// Solidity: function text(uint256 tokenId, string key) view returns(string)
func (_Contract *ContractCallerSession) Text(tokenId *big.Int, key string) (string, error) {
	return _Contract.Contract.Text(&_Contract.CallOpts, tokenId, key)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address DIDAddr) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, DIDAddr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", DIDAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address DIDAddr) returns()
func (_Contract *ContractSession) Initialize(DIDAddr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, DIDAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address DIDAddr) returns()
func (_Contract *ContractTransactorSession) Initialize(DIDAddr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, DIDAddr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xf0e44e75.
//
// Solidity: function setAddr(uint256 tokenId, uint256 coinType, bytes _addr) returns()
func (_Contract *ContractTransactor) SetAddr(opts *bind.TransactOpts, tokenId *big.Int, coinType *big.Int, _addr []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAddr", tokenId, coinType, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xf0e44e75.
//
// Solidity: function setAddr(uint256 tokenId, uint256 coinType, bytes _addr) returns()
func (_Contract *ContractSession) SetAddr(tokenId *big.Int, coinType *big.Int, _addr []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr(&_Contract.TransactOpts, tokenId, coinType, _addr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xf0e44e75.
//
// Solidity: function setAddr(uint256 tokenId, uint256 coinType, bytes _addr) returns()
func (_Contract *ContractTransactorSession) SetAddr(tokenId *big.Int, coinType *big.Int, _addr []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr(&_Contract.TransactOpts, tokenId, coinType, _addr)
}

// SetContentHash is a paid mutator transaction binding the contract method 0x18d977f8.
//
// Solidity: function setContentHash(uint256 tokenId, bytes cid) returns()
func (_Contract *ContractTransactor) SetContentHash(opts *bind.TransactOpts, tokenId *big.Int, cid []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setContentHash", tokenId, cid)
}

// SetContentHash is a paid mutator transaction binding the contract method 0x18d977f8.
//
// Solidity: function setContentHash(uint256 tokenId, bytes cid) returns()
func (_Contract *ContractSession) SetContentHash(tokenId *big.Int, cid []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetContentHash(&_Contract.TransactOpts, tokenId, cid)
}

// SetContentHash is a paid mutator transaction binding the contract method 0x18d977f8.
//
// Solidity: function setContentHash(uint256 tokenId, bytes cid) returns()
func (_Contract *ContractTransactorSession) SetContentHash(tokenId *big.Int, cid []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetContentHash(&_Contract.TransactOpts, tokenId, cid)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x0c1906ec.
//
// Solidity: function setPubkey(uint256 tokenId, bytes32 x, bytes32 y) returns()
func (_Contract *ContractTransactor) SetPubkey(opts *bind.TransactOpts, tokenId *big.Int, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPubkey", tokenId, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x0c1906ec.
//
// Solidity: function setPubkey(uint256 tokenId, bytes32 x, bytes32 y) returns()
func (_Contract *ContractSession) SetPubkey(tokenId *big.Int, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetPubkey(&_Contract.TransactOpts, tokenId, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x0c1906ec.
//
// Solidity: function setPubkey(uint256 tokenId, bytes32 x, bytes32 y) returns()
func (_Contract *ContractTransactorSession) SetPubkey(tokenId *big.Int, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetPubkey(&_Contract.TransactOpts, tokenId, x, y)
}

// SetReverse is a paid mutator transaction binding the contract method 0xac8682ca.
//
// Solidity: function setReverse(bool isReverse) returns()
func (_Contract *ContractTransactor) SetReverse(opts *bind.TransactOpts, isReverse bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setReverse", isReverse)
}

// SetReverse is a paid mutator transaction binding the contract method 0xac8682ca.
//
// Solidity: function setReverse(bool isReverse) returns()
func (_Contract *ContractSession) SetReverse(isReverse bool) (*types.Transaction, error) {
	return _Contract.Contract.SetReverse(&_Contract.TransactOpts, isReverse)
}

// SetReverse is a paid mutator transaction binding the contract method 0xac8682ca.
//
// Solidity: function setReverse(bool isReverse) returns()
func (_Contract *ContractTransactorSession) SetReverse(isReverse bool) (*types.Transaction, error) {
	return _Contract.Contract.SetReverse(&_Contract.TransactOpts, isReverse)
}

// SetText is a paid mutator transaction binding the contract method 0x3fb24782.
//
// Solidity: function setText(uint256 tokenId, string key, string value) returns()
func (_Contract *ContractTransactor) SetText(opts *bind.TransactOpts, tokenId *big.Int, key string, value string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setText", tokenId, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x3fb24782.
//
// Solidity: function setText(uint256 tokenId, string key, string value) returns()
func (_Contract *ContractSession) SetText(tokenId *big.Int, key string, value string) (*types.Transaction, error) {
	return _Contract.Contract.SetText(&_Contract.TransactOpts, tokenId, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x3fb24782.
//
// Solidity: function setText(uint256 tokenId, string key, string value) returns()
func (_Contract *ContractTransactorSession) SetText(tokenId *big.Int, key string, value string) (*types.Transaction, error) {
	return _Contract.Contract.SetText(&_Contract.TransactOpts, tokenId, key, value)
}

// ContractAddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the Contract contract.
type ContractAddressChangedIterator struct {
	Event *ContractAddressChanged // Event containing the contract specifics and raw log

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
func (it *ContractAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddressChanged)
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
		it.Event = new(ContractAddressChanged)
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
func (it *ContractAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddressChanged represents a AddressChanged event raised by the Contract contract.
type ContractAddressChanged struct {
	TokenId    *big.Int
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x88c6f12b5d80135a1513a91ef6bea8f1e2f7e880e6512ac5ca1e06d8c4cd9455.
//
// Solidity: event AddressChanged(uint256 tokenId, uint256 coinType, bytes newAddress)
func (_Contract *ContractFilterer) FilterAddressChanged(opts *bind.FilterOpts) (*ContractAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractAddressChangedIterator{contract: _Contract.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x88c6f12b5d80135a1513a91ef6bea8f1e2f7e880e6512ac5ca1e06d8c4cd9455.
//
// Solidity: event AddressChanged(uint256 tokenId, uint256 coinType, bytes newAddress)
func (_Contract *ContractFilterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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

// ParseAddressChanged is a log parse operation binding the contract event 0x88c6f12b5d80135a1513a91ef6bea8f1e2f7e880e6512ac5ca1e06d8c4cd9455.
//
// Solidity: event AddressChanged(uint256 tokenId, uint256 coinType, bytes newAddress)
func (_Contract *ContractFilterer) ParseAddressChanged(log types.Log) (*ContractAddressChanged, error) {
	event := new(ContractAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractContentHashChangedIterator is returned from FilterContentHashChanged and is used to iterate over the raw logs and unpacked data for ContentHashChanged events raised by the Contract contract.
type ContractContentHashChangedIterator struct {
	Event *ContractContentHashChanged // Event containing the contract specifics and raw log

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
func (it *ContractContentHashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractContentHashChanged)
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
		it.Event = new(ContractContentHashChanged)
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
func (it *ContractContentHashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractContentHashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractContentHashChanged represents a ContentHashChanged event raised by the Contract contract.
type ContractContentHashChanged struct {
	TokenId *big.Int
	Cid     []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContentHashChanged is a free log retrieval operation binding the contract event 0x2cfa0a96c5f362560b63aae3d9d8822b14c57d16b78c940ffc27774fd2d4e850.
//
// Solidity: event ContentHashChanged(uint256 indexed tokenId, bytes cid)
func (_Contract *ContractFilterer) FilterContentHashChanged(opts *bind.FilterOpts, tokenId []*big.Int) (*ContractContentHashChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ContentHashChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractContentHashChangedIterator{contract: _Contract.contract, event: "ContentHashChanged", logs: logs, sub: sub}, nil
}

// WatchContentHashChanged is a free log subscription operation binding the contract event 0x2cfa0a96c5f362560b63aae3d9d8822b14c57d16b78c940ffc27774fd2d4e850.
//
// Solidity: event ContentHashChanged(uint256 indexed tokenId, bytes cid)
func (_Contract *ContractFilterer) WatchContentHashChanged(opts *bind.WatchOpts, sink chan<- *ContractContentHashChanged, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ContentHashChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractContentHashChanged)
				if err := _Contract.contract.UnpackLog(event, "ContentHashChanged", log); err != nil {
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

// ParseContentHashChanged is a log parse operation binding the contract event 0x2cfa0a96c5f362560b63aae3d9d8822b14c57d16b78c940ffc27774fd2d4e850.
//
// Solidity: event ContentHashChanged(uint256 indexed tokenId, bytes cid)
func (_Contract *ContractFilterer) ParseContentHashChanged(log types.Log) (*ContractContentHashChanged, error) {
	event := new(ContractContentHashChanged)
	if err := _Contract.contract.UnpackLog(event, "ContentHashChanged", log); err != nil {
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

// ContractPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the Contract contract.
type ContractPubkeyChangedIterator struct {
	Event *ContractPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ContractPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPubkeyChanged)
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
		it.Event = new(ContractPubkeyChanged)
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
func (it *ContractPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPubkeyChanged represents a PubkeyChanged event raised by the Contract contract.
type ContractPubkeyChanged struct {
	TokenId *big.Int
	X       [32]byte
	Y       [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x8186dbd5b7b59a23cc021ac204ef3930b9d607bc737eb3559fa463d1c50bb149.
//
// Solidity: event PubkeyChanged(uint256 indexed tokenId, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, tokenId []*big.Int) (*ContractPubkeyChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PubkeyChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractPubkeyChangedIterator{contract: _Contract.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x8186dbd5b7b59a23cc021ac204ef3930b9d607bc737eb3559fa463d1c50bb149.
//
// Solidity: event PubkeyChanged(uint256 indexed tokenId, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ContractPubkeyChanged, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PubkeyChanged", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPubkeyChanged)
				if err := _Contract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x8186dbd5b7b59a23cc021ac204ef3930b9d607bc737eb3559fa463d1c50bb149.
//
// Solidity: event PubkeyChanged(uint256 indexed tokenId, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) ParsePubkeyChanged(log types.Log) (*ContractPubkeyChanged, error) {
	event := new(ContractPubkeyChanged)
	if err := _Contract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the Contract contract.
type ContractTextChangedIterator struct {
	Event *ContractTextChanged // Event containing the contract specifics and raw log

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
func (it *ContractTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTextChanged)
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
		it.Event = new(ContractTextChanged)
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
func (it *ContractTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTextChanged represents a TextChanged event raised by the Contract contract.
type ContractTextChanged struct {
	TokenId    *big.Int
	IndexedKey common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xe219c694b6b58e5263cea71424d10e93e7dc7f2ec3a0291aa27009084fd05a8b.
//
// Solidity: event TextChanged(uint256 indexed tokenId, string indexed indexedKey, string key, string value)
func (_Contract *ContractFilterer) FilterTextChanged(opts *bind.FilterOpts, tokenId []*big.Int, indexedKey []string) (*ContractTextChangedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TextChanged", tokenIdRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ContractTextChangedIterator{contract: _Contract.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xe219c694b6b58e5263cea71424d10e93e7dc7f2ec3a0291aa27009084fd05a8b.
//
// Solidity: event TextChanged(uint256 indexed tokenId, string indexed indexedKey, string key, string value)
func (_Contract *ContractFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ContractTextChanged, tokenId []*big.Int, indexedKey []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TextChanged", tokenIdRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTextChanged)
				if err := _Contract.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xe219c694b6b58e5263cea71424d10e93e7dc7f2ec3a0291aa27009084fd05a8b.
//
// Solidity: event TextChanged(uint256 indexed tokenId, string indexed indexedKey, string key, string value)
func (_Contract *ContractFilterer) ParseTextChanged(log types.Log) (*ContractTextChanged, error) {
	event := new(ContractTextChanged)
	if err := _Contract.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
