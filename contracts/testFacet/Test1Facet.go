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

// Test1FacetMetaData contains all meta data concerning the Test1Facet contract.
var Test1FacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"}],\"name\":\"AddNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d3b374a1": "AddNum(uint256,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060e28061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063d3b374a114602d575b600080fd5b603c6038366004605f565b604e565b60405190815260200160405180910390f35b6000605882846080565b9392505050565b60008060408385031215607157600080fd5b50508035926020909101359150565b600081600019048311821515161560a757634e487b7160e01b600052601160045260246000fd5b50029056fea26469706673582212204245598c31894ac2812a32274f403a2943b7c1927afcb4523e53b39bb89cf98d64736f6c634300080d0033",
}

// Test1FacetABI is the input ABI used to generate the binding from.
// Deprecated: Use Test1FacetMetaData.ABI instead.
var Test1FacetABI = Test1FacetMetaData.ABI

// Deprecated: Use Test1FacetMetaData.Sigs instead.
// Test1FacetFuncSigs maps the 4-byte function signature to its string representation.
var Test1FacetFuncSigs = Test1FacetMetaData.Sigs

// Test1FacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Test1FacetMetaData.Bin instead.
var Test1FacetBin = Test1FacetMetaData.Bin

// DeployTest1Facet deploys a new Ethereum contract, binding an instance of Test1Facet to it.
func DeployTest1Facet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test1Facet, error) {
	parsed, err := Test1FacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Test1FacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test1Facet{Test1FacetCaller: Test1FacetCaller{contract: contract}, Test1FacetTransactor: Test1FacetTransactor{contract: contract}, Test1FacetFilterer: Test1FacetFilterer{contract: contract}}, nil
}

// Test1Facet is an auto generated Go binding around an Ethereum contract.
type Test1Facet struct {
	Test1FacetCaller     // Read-only binding to the contract
	Test1FacetTransactor // Write-only binding to the contract
	Test1FacetFilterer   // Log filterer for contract events
}

// Test1FacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type Test1FacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test1FacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Test1FacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test1FacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Test1FacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Test1FacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Test1FacetSession struct {
	Contract     *Test1Facet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Test1FacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Test1FacetCallerSession struct {
	Contract *Test1FacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Test1FacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Test1FacetTransactorSession struct {
	Contract     *Test1FacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Test1FacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type Test1FacetRaw struct {
	Contract *Test1Facet // Generic contract binding to access the raw methods on
}

// Test1FacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Test1FacetCallerRaw struct {
	Contract *Test1FacetCaller // Generic read-only contract binding to access the raw methods on
}

// Test1FacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Test1FacetTransactorRaw struct {
	Contract *Test1FacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest1Facet creates a new instance of Test1Facet, bound to a specific deployed contract.
func NewTest1Facet(address common.Address, backend bind.ContractBackend) (*Test1Facet, error) {
	contract, err := bindTest1Facet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test1Facet{Test1FacetCaller: Test1FacetCaller{contract: contract}, Test1FacetTransactor: Test1FacetTransactor{contract: contract}, Test1FacetFilterer: Test1FacetFilterer{contract: contract}}, nil
}

// NewTest1FacetCaller creates a new read-only instance of Test1Facet, bound to a specific deployed contract.
func NewTest1FacetCaller(address common.Address, caller bind.ContractCaller) (*Test1FacetCaller, error) {
	contract, err := bindTest1Facet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Test1FacetCaller{contract: contract}, nil
}

// NewTest1FacetTransactor creates a new write-only instance of Test1Facet, bound to a specific deployed contract.
func NewTest1FacetTransactor(address common.Address, transactor bind.ContractTransactor) (*Test1FacetTransactor, error) {
	contract, err := bindTest1Facet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Test1FacetTransactor{contract: contract}, nil
}

// NewTest1FacetFilterer creates a new log filterer instance of Test1Facet, bound to a specific deployed contract.
func NewTest1FacetFilterer(address common.Address, filterer bind.ContractFilterer) (*Test1FacetFilterer, error) {
	contract, err := bindTest1Facet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Test1FacetFilterer{contract: contract}, nil
}

// bindTest1Facet binds a generic wrapper to an already deployed contract.
func bindTest1Facet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Test1FacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test1Facet *Test1FacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test1Facet.Contract.Test1FacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test1Facet *Test1FacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1FacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test1Facet *Test1FacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test1Facet.Contract.Test1FacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test1Facet *Test1FacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test1Facet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test1Facet *Test1FacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test1Facet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test1Facet *Test1FacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test1Facet.Contract.contract.Transact(opts, method, params...)
}

// AddNum is a free data retrieval call binding the contract method 0xd3b374a1.
//
// Solidity: function AddNum(uint256 v1, uint256 v2) view returns(uint256)
func (_Test1Facet *Test1FacetCaller) AddNum(opts *bind.CallOpts, v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "AddNum", v1, v2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddNum is a free data retrieval call binding the contract method 0xd3b374a1.
//
// Solidity: function AddNum(uint256 v1, uint256 v2) view returns(uint256)
func (_Test1Facet *Test1FacetSession) AddNum(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _Test1Facet.Contract.AddNum(&_Test1Facet.CallOpts, v1, v2)
}

// AddNum is a free data retrieval call binding the contract method 0xd3b374a1.
//
// Solidity: function AddNum(uint256 v1, uint256 v2) view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) AddNum(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _Test1Facet.Contract.AddNum(&_Test1Facet.CallOpts, v1, v2)
}
