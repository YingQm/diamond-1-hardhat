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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"}],\"name\":\"AddNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setNumA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setNumB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setNumC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d3b374a1": "AddNum(uint256,uint256)",
		"8782c019": "getNumA()",
		"9bf5cdb2": "getNumB()",
		"20fde452": "getNumC()",
		"a558a2f4": "numA()",
		"c6159d44": "numB()",
		"21dd8f3c": "numC()",
		"0719a4ed": "setNumA(uint256)",
		"6b3c1706": "setNumB(uint256)",
		"b67f2424": "setNumC(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506101de806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80639bf5cdb2116100665780639bf5cdb2146100f2578063a558a2f4146100fa578063b67f242414610103578063c6159d4414610116578063d3b374a11461011f57600080fd5b80630719a4ed146100a357806320fde452146100b857806321dd8f3c146100ce5780636b3c1706146100d75780638782c019146100ea575b600080fd5b6100b66100b1366004610140565b600055565b005b6002545b60405190815260200160405180910390f35b6100bc60025481565b6100b66100e5366004610140565b600155565b6000546100bc565b6001546100bc565b6100bc60005481565b6100b6610111366004610140565b600255565b6100bc60015481565b6100bc61012d366004610159565b6000610139828461017b565b9392505050565b60006020828403121561015257600080fd5b5035919050565b6000806040838503121561016c57600080fd5b50508035926020909101359150565b60008160001904831182151516156101a357634e487b7160e01b600052601160045260246000fd5b50029056fea264697066735822122006e73ceb54dc0ef325900cca7d2d71c2f8466febadd1c98fcb262439de3508ba64736f6c634300080d0033",
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

// GetNumA is a free data retrieval call binding the contract method 0x8782c019.
//
// Solidity: function getNumA() view returns(uint256)
func (_Test1Facet *Test1FacetCaller) GetNumA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "getNumA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumA is a free data retrieval call binding the contract method 0x8782c019.
//
// Solidity: function getNumA() view returns(uint256)
func (_Test1Facet *Test1FacetSession) GetNumA() (*big.Int, error) {
	return _Test1Facet.Contract.GetNumA(&_Test1Facet.CallOpts)
}

// GetNumA is a free data retrieval call binding the contract method 0x8782c019.
//
// Solidity: function getNumA() view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) GetNumA() (*big.Int, error) {
	return _Test1Facet.Contract.GetNumA(&_Test1Facet.CallOpts)
}

// GetNumB is a free data retrieval call binding the contract method 0x9bf5cdb2.
//
// Solidity: function getNumB() view returns(uint256)
func (_Test1Facet *Test1FacetCaller) GetNumB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "getNumB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumB is a free data retrieval call binding the contract method 0x9bf5cdb2.
//
// Solidity: function getNumB() view returns(uint256)
func (_Test1Facet *Test1FacetSession) GetNumB() (*big.Int, error) {
	return _Test1Facet.Contract.GetNumB(&_Test1Facet.CallOpts)
}

// GetNumB is a free data retrieval call binding the contract method 0x9bf5cdb2.
//
// Solidity: function getNumB() view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) GetNumB() (*big.Int, error) {
	return _Test1Facet.Contract.GetNumB(&_Test1Facet.CallOpts)
}

// GetNumC is a free data retrieval call binding the contract method 0x20fde452.
//
// Solidity: function getNumC() view returns(uint256)
func (_Test1Facet *Test1FacetCaller) GetNumC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "getNumC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumC is a free data retrieval call binding the contract method 0x20fde452.
//
// Solidity: function getNumC() view returns(uint256)
func (_Test1Facet *Test1FacetSession) GetNumC() (*big.Int, error) {
	return _Test1Facet.Contract.GetNumC(&_Test1Facet.CallOpts)
}

// GetNumC is a free data retrieval call binding the contract method 0x20fde452.
//
// Solidity: function getNumC() view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) GetNumC() (*big.Int, error) {
	return _Test1Facet.Contract.GetNumC(&_Test1Facet.CallOpts)
}

// NumA is a free data retrieval call binding the contract method 0xa558a2f4.
//
// Solidity: function numA() view returns(uint256)
func (_Test1Facet *Test1FacetCaller) NumA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "numA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumA is a free data retrieval call binding the contract method 0xa558a2f4.
//
// Solidity: function numA() view returns(uint256)
func (_Test1Facet *Test1FacetSession) NumA() (*big.Int, error) {
	return _Test1Facet.Contract.NumA(&_Test1Facet.CallOpts)
}

// NumA is a free data retrieval call binding the contract method 0xa558a2f4.
//
// Solidity: function numA() view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) NumA() (*big.Int, error) {
	return _Test1Facet.Contract.NumA(&_Test1Facet.CallOpts)
}

// NumB is a free data retrieval call binding the contract method 0xc6159d44.
//
// Solidity: function numB() view returns(uint256)
func (_Test1Facet *Test1FacetCaller) NumB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "numB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumB is a free data retrieval call binding the contract method 0xc6159d44.
//
// Solidity: function numB() view returns(uint256)
func (_Test1Facet *Test1FacetSession) NumB() (*big.Int, error) {
	return _Test1Facet.Contract.NumB(&_Test1Facet.CallOpts)
}

// NumB is a free data retrieval call binding the contract method 0xc6159d44.
//
// Solidity: function numB() view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) NumB() (*big.Int, error) {
	return _Test1Facet.Contract.NumB(&_Test1Facet.CallOpts)
}

// NumC is a free data retrieval call binding the contract method 0x21dd8f3c.
//
// Solidity: function numC() view returns(uint256)
func (_Test1Facet *Test1FacetCaller) NumC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test1Facet.contract.Call(opts, &out, "numC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumC is a free data retrieval call binding the contract method 0x21dd8f3c.
//
// Solidity: function numC() view returns(uint256)
func (_Test1Facet *Test1FacetSession) NumC() (*big.Int, error) {
	return _Test1Facet.Contract.NumC(&_Test1Facet.CallOpts)
}

// NumC is a free data retrieval call binding the contract method 0x21dd8f3c.
//
// Solidity: function numC() view returns(uint256)
func (_Test1Facet *Test1FacetCallerSession) NumC() (*big.Int, error) {
	return _Test1Facet.Contract.NumC(&_Test1Facet.CallOpts)
}

// SetNumA is a paid mutator transaction binding the contract method 0x0719a4ed.
//
// Solidity: function setNumA(uint256 v) returns()
func (_Test1Facet *Test1FacetTransactor) SetNumA(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "setNumA", v)
}

// SetNumA is a paid mutator transaction binding the contract method 0x0719a4ed.
//
// Solidity: function setNumA(uint256 v) returns()
func (_Test1Facet *Test1FacetSession) SetNumA(v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.Contract.SetNumA(&_Test1Facet.TransactOpts, v)
}

// SetNumA is a paid mutator transaction binding the contract method 0x0719a4ed.
//
// Solidity: function setNumA(uint256 v) returns()
func (_Test1Facet *Test1FacetTransactorSession) SetNumA(v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.Contract.SetNumA(&_Test1Facet.TransactOpts, v)
}

// SetNumB is a paid mutator transaction binding the contract method 0x6b3c1706.
//
// Solidity: function setNumB(uint256 v) returns()
func (_Test1Facet *Test1FacetTransactor) SetNumB(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "setNumB", v)
}

// SetNumB is a paid mutator transaction binding the contract method 0x6b3c1706.
//
// Solidity: function setNumB(uint256 v) returns()
func (_Test1Facet *Test1FacetSession) SetNumB(v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.Contract.SetNumB(&_Test1Facet.TransactOpts, v)
}

// SetNumB is a paid mutator transaction binding the contract method 0x6b3c1706.
//
// Solidity: function setNumB(uint256 v) returns()
func (_Test1Facet *Test1FacetTransactorSession) SetNumB(v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.Contract.SetNumB(&_Test1Facet.TransactOpts, v)
}

// SetNumC is a paid mutator transaction binding the contract method 0xb67f2424.
//
// Solidity: function setNumC(uint256 v) returns()
func (_Test1Facet *Test1FacetTransactor) SetNumC(opts *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.contract.Transact(opts, "setNumC", v)
}

// SetNumC is a paid mutator transaction binding the contract method 0xb67f2424.
//
// Solidity: function setNumC(uint256 v) returns()
func (_Test1Facet *Test1FacetSession) SetNumC(v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.Contract.SetNumC(&_Test1Facet.TransactOpts, v)
}

// SetNumC is a paid mutator transaction binding the contract method 0xb67f2424.
//
// Solidity: function setNumC(uint256 v) returns()
func (_Test1Facet *Test1FacetTransactorSession) SetNumC(v *big.Int) (*types.Transaction, error) {
	return _Test1Facet.Contract.SetNumC(&_Test1Facet.TransactOpts, v)
}
