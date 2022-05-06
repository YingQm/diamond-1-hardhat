// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testFacet

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

// Test2FacetMetaData contains all meta data concerning the Test2Facet contract.
var Test2FacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"}],\"name\":\"AddNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d3b374a1": "AddNum(uint256,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060db8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063d3b374a114602d575b600080fd5b603c6038366004605f565b604e565b60405190815260200160405180910390f35b6000605882846080565b9392505050565b60008060408385031215607157600080fd5b50508035926020909101359150565b6000821982111560a057634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220edb9f1570857727c1c99c92636cbb7aecd126a7059034d6ced2f4bd1add13a1c64736f6c634300080d0033",
}

// Test2FacetABI is the input ABI used to generate the binding from.
// Deprecated: Use Test2FacetMetaData.ABI instead.
var Test2FacetABI = Test2FacetMetaData.ABI

// Deprecated: Use Test2FacetMetaData.Sigs instead.
// Test2FacetFuncSigs maps the 4-byte function signature to its string representation.
var Test2FacetFuncSigs = Test2FacetMetaData.Sigs

// Test2FacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Test2FacetMetaData.Bin instead.
var Test2FacetBin = Test2FacetMetaData.Bin

// DeployTest2Facet deploys a new Ethereum contract, binding an instance of Test2Facet to it.
func DeployTest2Facet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test2Facet, error) {
	parsed, err := Test2FacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Test2FacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test2Facet{Test2FacetCaller: Test2FacetCaller{contract: contract}, Test2FacetTransactor: Test2FacetTransactor{contract: contract}, Test2FacetFilterer: Test2FacetFilterer{contract: contract}}, nil
}

// Test2Facet is an auto generated Go binding around an Ethereum contract.
type Test2Facet struct {
	Test2FacetCaller     // Read-only binding to the contract
	Test2FacetTransactor // Write-only binding to the contract
	Test2FacetFilterer   // Log filterer for contract events
}

// Test2FacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type Test2FacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2FacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Test2FacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2FacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Test2FacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test2FacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Test2FacetSession struct {
	Contract     *Test2Facet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Test2FacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Test2FacetCallerSession struct {
	Contract *Test2FacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Test2FacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Test2FacetTransactorSession struct {
	Contract     *Test2FacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Test2FacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type Test2FacetRaw struct {
	Contract *Test2Facet // Generic contract binding to access the raw methods on
}

// Test2FacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Test2FacetCallerRaw struct {
	Contract *Test2FacetCaller // Generic read-only contract binding to access the raw methods on
}

// Test2FacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Test2FacetTransactorRaw struct {
	Contract *Test2FacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest2Facet creates a new instance of Test2Facet, bound to a specific deployed contract.
func NewTest2Facet(address common.Address, backend bind.ContractBackend) (*Test2Facet, error) {
	contract, err := bindTest2Facet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test2Facet{Test2FacetCaller: Test2FacetCaller{contract: contract}, Test2FacetTransactor: Test2FacetTransactor{contract: contract}, Test2FacetFilterer: Test2FacetFilterer{contract: contract}}, nil
}

// NewTest2FacetCaller creates a new read-only instance of Test2Facet, bound to a specific deployed contract.
func NewTest2FacetCaller(address common.Address, caller bind.ContractCaller) (*Test2FacetCaller, error) {
	contract, err := bindTest2Facet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Test2FacetCaller{contract: contract}, nil
}

// NewTest2FacetTransactor creates a new write-only instance of Test2Facet, bound to a specific deployed contract.
func NewTest2FacetTransactor(address common.Address, transactor bind.ContractTransactor) (*Test2FacetTransactor, error) {
	contract, err := bindTest2Facet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Test2FacetTransactor{contract: contract}, nil
}

// NewTest2FacetFilterer creates a new log filterer instance of Test2Facet, bound to a specific deployed contract.
func NewTest2FacetFilterer(address common.Address, filterer bind.ContractFilterer) (*Test2FacetFilterer, error) {
	contract, err := bindTest2Facet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Test2FacetFilterer{contract: contract}, nil
}

// bindTest2Facet binds a generic wrapper to an already deployed contract.
func bindTest2Facet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Test2FacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test2Facet *Test2FacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test2Facet.Contract.Test2FacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test2Facet *Test2FacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2FacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test2Facet *Test2FacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test2Facet.Contract.Test2FacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test2Facet *Test2FacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test2Facet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test2Facet *Test2FacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test2Facet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test2Facet *Test2FacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test2Facet.Contract.contract.Transact(opts, method, params...)
}

// AddNum is a free data retrieval call binding the contract method 0xd3b374a1.
//
// Solidity: function AddNum(uint256 v1, uint256 v2) view returns(uint256)
func (_Test2Facet *Test2FacetCaller) AddNum(opts *bind.CallOpts, v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Test2Facet.contract.Call(opts, &out, "AddNum", v1, v2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddNum is a free data retrieval call binding the contract method 0xd3b374a1.
//
// Solidity: function AddNum(uint256 v1, uint256 v2) view returns(uint256)
func (_Test2Facet *Test2FacetSession) AddNum(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _Test2Facet.Contract.AddNum(&_Test2Facet.CallOpts, v1, v2)
}

// AddNum is a free data retrieval call binding the contract method 0xd3b374a1.
//
// Solidity: function AddNum(uint256 v1, uint256 v2) view returns(uint256)
func (_Test2Facet *Test2FacetCallerSession) AddNum(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _Test2Facet.Contract.AddNum(&_Test2Facet.CallOpts, v1, v2)
}
