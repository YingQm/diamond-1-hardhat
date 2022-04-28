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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// DiamondLoupeFacetMetaData contains all meta data concerning the DiamondLoupeFacet contract.
var DiamondLoupeFacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_facetFunctionSelectors\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cdffacc6": "facetAddress(bytes4)",
		"52ef6b2c": "facetAddresses()",
		"adfca15e": "facetFunctionSelectors(address)",
		"7a0ed627": "facets()",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610aa7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806301ffc9a71461005c57806352ef6b2c146100bd5780637a0ed627146100d2578063adfca15e146100e7578063cdffacc614610107575b600080fd5b6100a861006a36600461085b565b6001600160e01b03191660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e602052604090205460ff1690565b60405190151581526020015b60405180910390f35b6100c561015f565b6040516100b4919061088c565b6100da610300565b6040516100b4919061091e565b6100fa6100f536600461099b565b610712565b6040516100b491906109c4565b61014761011536600461085b565b6001600160e01b0319166000908152600080516020610a5283398151915260205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016100b4565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d54606090600080516020610a52833981519152908067ffffffffffffffff8111156101ad576101ad6109d7565b6040519080825280602002602001820160405280156101d6578160200160208202803683370190505b5092506000805b828110156102f65760008460010182815481106101fc576101fc6109ed565b6000918252602080832060088304015460079092166004026101000a90910460e01b6001600160e01b0319811683529087905260408220549092506001600160a01b031690805b858110156102985788818151811061025d5761025d6109ed565b60200260200101516001600160a01b0316836001600160a01b0316036102865760019150610298565b8061029081610a19565b915050610243565b5080156102a857506102e4915050565b818886815181106102bb576102bb6109ed565b6001600160a01b0390921660209283029190910190910152846102dd81610a19565b9550505050505b806102ee81610a19565b9150506101dd565b5080845250505090565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d54606090600080516020610a52833981519152908067ffffffffffffffff81111561034e5761034e6109d7565b60405190808252806020026020018201604052801561039457816020015b60408051808201909152600081526060602082015281526020019060019003908161036c5790505b50925060008167ffffffffffffffff8111156103b2576103b26109d7565b6040519080825280602002602001820160405280156103db578160200160208202803683370190505b5090506000805b838110156106a0576000856001018281548110610401576104016109ed565b6000918252602080832060088304015460079092166004026101000a90910460e01b6001600160e01b0319811683529088905260408220549092506001600160a01b031690805b8581101561056557826001600160a01b03168a828151811061046c5761046c6109ed565b6020026020010151600001516001600160a01b03160361055357838a8281518110610499576104996109ed565b6020026020010151602001518883815181106104b7576104b76109ed565b602002602001015160ff16815181106104d2576104d26109ed565b60200260200101906001600160e01b03191690816001600160e01b0319168152505060ff878281518110610508576105086109ed565b602002602001015160ff161061051d57600080fd5b86818151811061052f5761052f6109ed565b60200260200101805180919061054490610a32565b60ff1690525060019150610565565b8061055d81610a19565b915050610448565b508015610575575061068e915050565b81898681518110610588576105886109ed565b60209081029190910101516001600160a01b0390911690528667ffffffffffffffff8111156105b9576105b96109d7565b6040519080825280602002602001820160405280156105e2578160200160208202803683370190505b508986815181106105f5576105f56109ed565b60200260200101516020018190525082898681518110610617576106176109ed565b602002602001015160200151600081518110610635576106356109ed565b60200260200101906001600160e01b03191690816001600160e01b03191681525050600186868151811061066b5761066b6109ed565b60ff909216602092830291909101909101528461068781610a19565b9550505050505b8061069881610a19565b9150506103e2565b5060005b818110156107075760008382815181106106c0576106c06109ed565b602002602001015160ff16905060008783815181106106e1576106e16109ed565b6020026020010151602001519050818152505080806106ff90610a19565b9150506106a4565b508085525050505090565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d54606090600080516020610a528339815191529060008167ffffffffffffffff811115610762576107626109d7565b60405190808252806020026020018201604052801561078b578160200160208202803683370190505b50935060005b828110156108505760008460010182815481106107b0576107b06109ed565b6000918252602080832060088304015460079092166004026101000a90910460e01b6001600160e01b031981168352908790526040909120549091506001600160a01b0390811690881681900361083b5781878581518110610814576108146109ed565b6001600160e01b0319909216602092830291909101909101528361083781610a19565b9450505b5050808061084890610a19565b915050610791565b508352509092915050565b60006020828403121561086d57600080fd5b81356001600160e01b03198116811461088557600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156108cd5783516001600160a01b0316835292840192918401916001016108a8565b50909695505050505050565b600081518084526020808501945080840160005b838110156109135781516001600160e01b031916875295820195908201906001016108ed565b509495945050505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561098d57888303603f19018552815180516001600160a01b0316845287015187840187905261097a878501826108d9565b9588019593505090860190600101610945565b509098975050505050505050565b6000602082840312156109ad57600080fd5b81356001600160a01b038116811461088557600080fd5b60208152600061088560208301846108d9565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610a2b57610a2b610a03565b5060010190565b600060ff821660ff8103610a4857610a48610a03565b6001019291505056fec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131ca26469706673582212202d6b5d934c49ceb81029c2d4c6aaf93568325d7924821ea2d68cd48a379d421b64736f6c634300080d0033",
}

// DiamondLoupeFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondLoupeFacetMetaData.ABI instead.
var DiamondLoupeFacetABI = DiamondLoupeFacetMetaData.ABI

// Deprecated: Use DiamondLoupeFacetMetaData.Sigs instead.
// DiamondLoupeFacetFuncSigs maps the 4-byte function signature to its string representation.
var DiamondLoupeFacetFuncSigs = DiamondLoupeFacetMetaData.Sigs

// DiamondLoupeFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondLoupeFacetMetaData.Bin instead.
var DiamondLoupeFacetBin = DiamondLoupeFacetMetaData.Bin

// DeployDiamondLoupeFacet deploys a new Ethereum contract, binding an instance of DiamondLoupeFacet to it.
func DeployDiamondLoupeFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondLoupeFacet, error) {
	parsed, err := DiamondLoupeFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondLoupeFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// DiamondLoupeFacet is an auto generated Go binding around an Ethereum contract.
type DiamondLoupeFacet struct {
	DiamondLoupeFacetCaller     // Read-only binding to the contract
	DiamondLoupeFacetTransactor // Write-only binding to the contract
	DiamondLoupeFacetFilterer   // Log filterer for contract events
}

// DiamondLoupeFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondLoupeFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondLoupeFacetSession struct {
	Contract     *DiamondLoupeFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondLoupeFacetCallerSession struct {
	Contract *DiamondLoupeFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DiamondLoupeFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondLoupeFacetTransactorSession struct {
	Contract     *DiamondLoupeFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondLoupeFacetRaw struct {
	Contract *DiamondLoupeFacet // Generic contract binding to access the raw methods on
}

// DiamondLoupeFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCallerRaw struct {
	Contract *DiamondLoupeFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondLoupeFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactorRaw struct {
	Contract *DiamondLoupeFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondLoupeFacet creates a new instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacet(address common.Address, backend bind.ContractBackend) (*DiamondLoupeFacet, error) {
	contract, err := bindDiamondLoupeFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// NewDiamondLoupeFacetCaller creates a new read-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondLoupeFacetCaller, error) {
	contract, err := bindDiamondLoupeFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetCaller{contract: contract}, nil
}

// NewDiamondLoupeFacetTransactor creates a new write-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondLoupeFacetTransactor, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetTransactor{contract: contract}, nil
}

// NewDiamondLoupeFacetFilterer creates a new log filterer instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondLoupeFacetFilterer, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetFilterer{contract: contract}, nil
}

// bindDiamondLoupeFacet binds a generic wrapper to an already deployed contract.
func bindDiamondLoupeFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondLoupeFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}

// IDiamondLoupeMetaData contains all meta data concerning the IDiamondLoupe contract.
var IDiamondLoupeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cdffacc6": "facetAddress(bytes4)",
		"52ef6b2c": "facetAddresses()",
		"adfca15e": "facetFunctionSelectors(address)",
		"7a0ed627": "facets()",
	},
}

// IDiamondLoupeABI is the input ABI used to generate the binding from.
// Deprecated: Use IDiamondLoupeMetaData.ABI instead.
var IDiamondLoupeABI = IDiamondLoupeMetaData.ABI

// Deprecated: Use IDiamondLoupeMetaData.Sigs instead.
// IDiamondLoupeFuncSigs maps the 4-byte function signature to its string representation.
var IDiamondLoupeFuncSigs = IDiamondLoupeMetaData.Sigs

// IDiamondLoupe is an auto generated Go binding around an Ethereum contract.
type IDiamondLoupe struct {
	IDiamondLoupeCaller     // Read-only binding to the contract
	IDiamondLoupeTransactor // Write-only binding to the contract
	IDiamondLoupeFilterer   // Log filterer for contract events
}

// IDiamondLoupeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDiamondLoupeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDiamondLoupeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDiamondLoupeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondLoupeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDiamondLoupeSession struct {
	Contract     *IDiamondLoupe    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDiamondLoupeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDiamondLoupeCallerSession struct {
	Contract *IDiamondLoupeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IDiamondLoupeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDiamondLoupeTransactorSession struct {
	Contract     *IDiamondLoupeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IDiamondLoupeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDiamondLoupeRaw struct {
	Contract *IDiamondLoupe // Generic contract binding to access the raw methods on
}

// IDiamondLoupeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDiamondLoupeCallerRaw struct {
	Contract *IDiamondLoupeCaller // Generic read-only contract binding to access the raw methods on
}

// IDiamondLoupeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDiamondLoupeTransactorRaw struct {
	Contract *IDiamondLoupeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDiamondLoupe creates a new instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupe(address common.Address, backend bind.ContractBackend) (*IDiamondLoupe, error) {
	contract, err := bindIDiamondLoupe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupe{IDiamondLoupeCaller: IDiamondLoupeCaller{contract: contract}, IDiamondLoupeTransactor: IDiamondLoupeTransactor{contract: contract}, IDiamondLoupeFilterer: IDiamondLoupeFilterer{contract: contract}}, nil
}

// NewIDiamondLoupeCaller creates a new read-only instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeCaller(address common.Address, caller bind.ContractCaller) (*IDiamondLoupeCaller, error) {
	contract, err := bindIDiamondLoupe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeCaller{contract: contract}, nil
}

// NewIDiamondLoupeTransactor creates a new write-only instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeTransactor(address common.Address, transactor bind.ContractTransactor) (*IDiamondLoupeTransactor, error) {
	contract, err := bindIDiamondLoupe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeTransactor{contract: contract}, nil
}

// NewIDiamondLoupeFilterer creates a new log filterer instance of IDiamondLoupe, bound to a specific deployed contract.
func NewIDiamondLoupeFilterer(address common.Address, filterer bind.ContractFilterer) (*IDiamondLoupeFilterer, error) {
	contract, err := bindIDiamondLoupe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDiamondLoupeFilterer{contract: contract}, nil
}

// bindIDiamondLoupe binds a generic wrapper to an already deployed contract.
func bindIDiamondLoupe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDiamondLoupeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondLoupe *IDiamondLoupeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondLoupe.Contract.IDiamondLoupeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondLoupe *IDiamondLoupeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.IDiamondLoupeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondLoupe *IDiamondLoupeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.IDiamondLoupeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondLoupe *IDiamondLoupeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondLoupe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondLoupe *IDiamondLoupeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondLoupe *IDiamondLoupeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondLoupe.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddress(&_IDiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddress(&_IDiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetAddresses() ([]common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddresses(&_IDiamondLoupe.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetAddresses() ([]common.Address, error) {
	return _IDiamondLoupe.Contract.FacetAddresses(&_IDiamondLoupe.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IDiamondLoupe.Contract.FacetFunctionSelectors(&_IDiamondLoupe.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IDiamondLoupe.Contract.FacetFunctionSelectors(&_IDiamondLoupe.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _IDiamondLoupe.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _IDiamondLoupe.Contract.Facets(&_IDiamondLoupe.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_IDiamondLoupe *IDiamondLoupeCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _IDiamondLoupe.Contract.Facets(&_IDiamondLoupe.CallOpts)
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}
