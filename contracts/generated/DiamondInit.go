// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// DiamondInitMetaData contains all meta data concerning the DiamondInit contract.
var DiamondInitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e1c7392a": "init()",
	},
	Bin: "0x608060405234801561001057600080fd5b50610149806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e1c7392a14610030575b600080fd5b6101117fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e6020527f673a26ab9c976db950bbe987aa80c5e387f329563bb0afe093ddccc970489e318054600160ff1991821681179092557f9bed265332efc30fa7643cc339edc91cb284a0f6566818a5788922af58c86b5080548216831790557f795db15802e151b19272d3e7b72ebd9d0cedc282cc23a6e937c8c3c90d9e213780548216831790556307f5828d60e41b6000527fe616bea4664e595328e525b24998219caecea2090de91847473acfb3efaa8aad80549091169091179055565b00fea26469706673582212204bb1cbe242e2ba1d9d358373d60f05229e9416caa54a3d93d435d6960d00025c64736f6c634300080d0033",
}

// DiamondInitABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondInitMetaData.ABI instead.
var DiamondInitABI = DiamondInitMetaData.ABI

// Deprecated: Use DiamondInitMetaData.Sigs instead.
// DiamondInitFuncSigs maps the 4-byte function signature to its string representation.
var DiamondInitFuncSigs = DiamondInitMetaData.Sigs

// DiamondInitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondInitMetaData.Bin instead.
var DiamondInitBin = DiamondInitMetaData.Bin

// DeployDiamondInit deploys a new Ethereum contract, binding an instance of DiamondInit to it.
func DeployDiamondInit(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondInit, error) {
	parsed, err := DiamondInitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondInitBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondInit{DiamondInitCaller: DiamondInitCaller{contract: contract}, DiamondInitTransactor: DiamondInitTransactor{contract: contract}, DiamondInitFilterer: DiamondInitFilterer{contract: contract}}, nil
}

// DiamondInit is an auto generated Go binding around an Ethereum contract.
type DiamondInit struct {
	DiamondInitCaller     // Read-only binding to the contract
	DiamondInitTransactor // Write-only binding to the contract
	DiamondInitFilterer   // Log filterer for contract events
}

// DiamondInitCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondInitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondInitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondInitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondInitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondInitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondInitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondInitSession struct {
	Contract     *DiamondInit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondInitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondInitCallerSession struct {
	Contract *DiamondInitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DiamondInitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondInitTransactorSession struct {
	Contract     *DiamondInitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DiamondInitRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondInitRaw struct {
	Contract *DiamondInit // Generic contract binding to access the raw methods on
}

// DiamondInitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondInitCallerRaw struct {
	Contract *DiamondInitCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondInitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondInitTransactorRaw struct {
	Contract *DiamondInitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondInit creates a new instance of DiamondInit, bound to a specific deployed contract.
func NewDiamondInit(address common.Address, backend bind.ContractBackend) (*DiamondInit, error) {
	contract, err := bindDiamondInit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondInit{DiamondInitCaller: DiamondInitCaller{contract: contract}, DiamondInitTransactor: DiamondInitTransactor{contract: contract}, DiamondInitFilterer: DiamondInitFilterer{contract: contract}}, nil
}

// NewDiamondInitCaller creates a new read-only instance of DiamondInit, bound to a specific deployed contract.
func NewDiamondInitCaller(address common.Address, caller bind.ContractCaller) (*DiamondInitCaller, error) {
	contract, err := bindDiamondInit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondInitCaller{contract: contract}, nil
}

// NewDiamondInitTransactor creates a new write-only instance of DiamondInit, bound to a specific deployed contract.
func NewDiamondInitTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondInitTransactor, error) {
	contract, err := bindDiamondInit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondInitTransactor{contract: contract}, nil
}

// NewDiamondInitFilterer creates a new log filterer instance of DiamondInit, bound to a specific deployed contract.
func NewDiamondInitFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondInitFilterer, error) {
	contract, err := bindDiamondInit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondInitFilterer{contract: contract}, nil
}

// bindDiamondInit binds a generic wrapper to an already deployed contract.
func bindDiamondInit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondInitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondInit *DiamondInitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondInit.Contract.DiamondInitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondInit *DiamondInitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondInit.Contract.DiamondInitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondInit *DiamondInitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondInit.Contract.DiamondInitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondInit *DiamondInitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondInit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondInit *DiamondInitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondInit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondInit *DiamondInitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondInit.Contract.contract.Transact(opts, method, params...)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_DiamondInit *DiamondInitTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondInit.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_DiamondInit *DiamondInitSession) Init() (*types.Transaction, error) {
	return _DiamondInit.Contract.Init(&_DiamondInit.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_DiamondInit *DiamondInitTransactorSession) Init() (*types.Transaction, error) {
	return _DiamondInit.Contract.Init(&_DiamondInit.TransactOpts)
}
