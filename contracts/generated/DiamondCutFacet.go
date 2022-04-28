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

// DiamondCutFacetMetaData contains all meta data concerning the DiamondCutFacet contract.
var DiamondCutFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1f931c1c": "diamondCut((address,uint8,bytes4[])[],address,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061143d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631f931c1c14610030575b600080fd5b61004361003e366004610e67565b610045565b005b61004d61009e565b61009761005a8587610fad565b8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061011a92505050565b5050505050565b600080516020611374833981519152600301546001600160a01b031633146101185760405162461bcd60e51b815260206004820152602260248201527f4c69624469616d6f6e643a204d75737420626520636f6e7472616374206f776e60448201526132b960f11b60648201526084015b60405180910390fd5b565b60005b83518110156102e057600084828151811061013a5761013a6110f1565b60200260200101516020015190506000600281111561015b5761015b611107565b81600281111561016d5761016d611107565b036101bb576101b6858381518110610187576101876110f1565b6020026020010151600001518684815181106101a5576101a56110f1565b60200260200101516040015161032b565b6102cd565b60018160028111156101cf576101cf611107565b03610218576101b68583815181106101e9576101e96110f1565b602002602001015160000151868481518110610207576102076110f1565b60200260200101516040015161058a565b600281600281111561022c5761022c611107565b03610275576101b6858381518110610246576102466110f1565b602002602001015160000151868481518110610264576102646110f1565b602002602001015160400151610851565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b606482015260840161010f565b50806102d881611133565b91505061011d565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673838383604051610314939291906111a4565b60405180910390a16103268282610bd4565b505050565b600081511161034c5760405162461bcd60e51b815260040161010f906112a4565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d54600080516020611374833981519152906001600160a01b0384166103e95760405162461bcd60e51b815260206004820152602c60248201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260448201526b65206164647265737328302960a01b606482015260840161010f565b61040b8460405180606001604052806024815260200161139460249139610de1565b60005b835181101561009757600084828151811061042b5761042b6110f1565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156104c95760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f6044820152746e207468617420616c72656164792065786973747360581b606482015260840161010f565b6040805180820182526001600160a01b03808a16825261ffff80881660208085019182526001600160e01b0319881660009081528b8252958620945185549251909316600160a01b026001600160b01b0319909216929093169190911717909155600180880180549182018155835291206008820401805460e085901c60046007909416939093026101000a92830263ffffffff909302191691909117905583610572816112ef565b9450505050808061058290611133565b91505061040e565b60008151116105ab5760405162461bcd60e51b815260040161010f906112a4565b6000805160206113748339815191526001600160a01b0383166106295760405162461bcd60e51b815260206004820152603060248201527f4c69624469616d6f6e644375743a205265706c6163652066616365742063616e60448201526f2774206265206164647265737328302960801b606482015260840161010f565b61064b836040518060600160405280602881526020016113e060289139610de1565b60005b825181101561084b57600083828151811061066b5761066b6110f1565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b03163081036107045760405162461bcd60e51b815260206004820152602f60248201527f4c69624469616d6f6e644375743a2043616e2774207265706c61636520696d6d60448201526e3aba30b1363290333ab731ba34b7b760891b606482015260840161010f565b856001600160a01b0316816001600160a01b03160361078b5760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e0000000000000000606482015260840161010f565b6001600160a01b0381166108075760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e207468617420646f65736e27742065786973740000000000000000606482015260840161010f565b506001600160e01b031916600090815260208390526040902080546001600160a01b0319166001600160a01b0386161790558061084381611133565b91505061064e565b50505050565b60008151116108725760405162461bcd60e51b815260040161010f906112a4565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d54600080516020611374833981519152906001600160a01b0384161561091a5760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f76652066616365742061646472604482015275657373206d757374206265206164647265737328302960501b606482015260840161010f565b60005b835181101561009757600084828151811061093a5761093a6110f1565b6020908102919091018101516001600160e01b0319811660009081528683526040908190208151808301909252546001600160a01b038116808352600160a01b90910461ffff1693820193909352909250906109fe5760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e2774206578697374000000000000000000606482015260840161010f565b8051306001600160a01b0390911603610a715760405162461bcd60e51b815260206004820152602f60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526e3a30b1363290333ab731ba34b7b71760891b606482015260840161010f565b83610a7b81611310565b94505083816020015161ffff1614610b59576000856001018581548110610aa457610aa46110f1565b90600052602060002090600891828204019190066004029054906101000a900460e01b90508086600101836020015161ffff1681548110610ae757610ae76110f1565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c92909202939093179055838201516001600160e01b03199390931681529087905260409020805461ffff60a01b1916600160a01b61ffff909316929092029190911790555b84600101805480610b6c57610b6c611327565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b0319909316815291859052506040902080546001600160b01b031916905580610bcc81611133565b91505061091d565b6001600160a01b038216610c5b57805115610c575760405162461bcd60e51b815260206004820152603c60248201527f4c69624469616d6f6e644375743a205f696e697420697320616464726573732860448201527f3029206275745f63616c6c64617461206973206e6f7420656d70747900000000606482015260840161010f565b5050565b6000815111610cd25760405162461bcd60e51b815260206004820152603d60248201527f4c69624469616d6f6e644375743a205f63616c6c6461746120697320656d707460448201527f7920627574205f696e6974206973206e6f742061646472657373283029000000606482015260840161010f565b6001600160a01b0382163014610d0457610d04826040518060600160405280602881526020016113b860289139610de1565b600080836001600160a01b031683604051610d1f919061133d565b600060405180830381855af49150503d8060008114610d5a576040519150601f19603f3d011682016040523d82523d6000602084013e610d5f565b606091505b50915091508161084b57805115610d8a578060405162461bcd60e51b815260040161010f9190611359565b60405162461bcd60e51b815260206004820152602660248201527f4c69624469616d6f6e644375743a205f696e69742066756e6374696f6e2072656044820152651d995c9d195960d21b606482015260840161010f565b813b818161084b5760405162461bcd60e51b815260040161010f9190611359565b80356001600160a01b0381168114610e1957600080fd5b919050565b60008083601f840112610e3057600080fd5b50813567ffffffffffffffff811115610e4857600080fd5b602083019150836020828501011115610e6057600080fd5b9250929050565b600080600080600060608688031215610e7f57600080fd5b853567ffffffffffffffff80821115610e9757600080fd5b818801915088601f830112610eab57600080fd5b813581811115610eba57600080fd5b8960208260051b8501011115610ecf57600080fd5b60208301975080965050610ee560208901610e02565b94506040880135915080821115610efb57600080fd5b50610f0888828901610e1e565b969995985093965092949392505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610f5257610f52610f19565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610f8157610f81610f19565b604052919050565b600067ffffffffffffffff821115610fa357610fa3610f19565b5060051b60200190565b6000610fc0610fbb84610f89565b610f58565b83815260208082019190600586811b860136811115610fde57600080fd5b865b818110156110e457803567ffffffffffffffff808211156110015760008081fd5b818a019150606082360312156110175760008081fd5b61101f610f2f565b61102883610e02565b8152868301356003811061103c5760008081fd5b81880152604083810135838111156110545760008081fd5b939093019236601f85011261106b57600092508283fd5b8335925061107b610fbb84610f89565b83815292871b840188019288810190368511156110985760008081fd5b948901945b848610156110cd5785356001600160e01b0319811681146110be5760008081fd5b8252948901949089019061109d565b918301919091525088525050948301948301610fe0565b5092979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016111455761114561111d565b5060010190565b60005b8381101561116757818101518382015260200161114f565b8381111561084b5750506000910152565b6000815180845261119081602086016020860161114c565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b8481101561127457898403607f19018652815180516001600160a01b0316855283810151898601906003811061121357634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b8083101561125f5783516001600160e01b0319168252928601926001929092019190860190611235565b509785019795505050908201906001016111cd565b50506001600160a01b038a169088015286810360408801526112968189611178565b9a9950505050505050505050565b6020808252602b908201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660408201526a1858d95d081d1bc818dd5d60aa1b606082015260800190565b600061ffff8083168181036113065761130661111d565b6001019392505050565b60008161131f5761131f61111d565b506000190190565b634e487b7160e01b600052603160045260246000fd5b6000825161134f81846020870161114c565b9190910192915050565b60208152600061136c6020830184611178565b939250505056fec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a2041646420666163657420686173206e6f20636f64654c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a205265706c61636520666163657420686173206e6f20636f6465a26469706673582212201aa3d827cf0b7a1bbaaf72bfba15042d087a59bac8c5d5cb2e635c409a80be5a64736f6c634300080d0033",
}

// DiamondCutFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondCutFacetMetaData.ABI instead.
var DiamondCutFacetABI = DiamondCutFacetMetaData.ABI

// Deprecated: Use DiamondCutFacetMetaData.Sigs instead.
// DiamondCutFacetFuncSigs maps the 4-byte function signature to its string representation.
var DiamondCutFacetFuncSigs = DiamondCutFacetMetaData.Sigs

// DiamondCutFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondCutFacetMetaData.Bin instead.
var DiamondCutFacetBin = DiamondCutFacetMetaData.Bin

// DeployDiamondCutFacet deploys a new Ethereum contract, binding an instance of DiamondCutFacet to it.
func DeployDiamondCutFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondCutFacet, error) {
	parsed, err := DiamondCutFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondCutFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// DiamondCutFacet is an auto generated Go binding around an Ethereum contract.
type DiamondCutFacet struct {
	DiamondCutFacetCaller     // Read-only binding to the contract
	DiamondCutFacetTransactor // Write-only binding to the contract
	DiamondCutFacetFilterer   // Log filterer for contract events
}

// DiamondCutFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCutFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondCutFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondCutFacetSession struct {
	Contract     *DiamondCutFacet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCutFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCutFacetCallerSession struct {
	Contract *DiamondCutFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DiamondCutFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondCutFacetTransactorSession struct {
	Contract     *DiamondCutFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DiamondCutFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondCutFacetRaw struct {
	Contract *DiamondCutFacet // Generic contract binding to access the raw methods on
}

// DiamondCutFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCutFacetCallerRaw struct {
	Contract *DiamondCutFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondCutFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactorRaw struct {
	Contract *DiamondCutFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondCutFacet creates a new instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacet(address common.Address, backend bind.ContractBackend) (*DiamondCutFacet, error) {
	contract, err := bindDiamondCutFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// NewDiamondCutFacetCaller creates a new read-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondCutFacetCaller, error) {
	contract, err := bindDiamondCutFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetCaller{contract: contract}, nil
}

// NewDiamondCutFacetTransactor creates a new write-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondCutFacetTransactor, error) {
	contract, err := bindDiamondCutFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetTransactor{contract: contract}, nil
}

// NewDiamondCutFacetFilterer creates a new log filterer instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondCutFacetFilterer, error) {
	contract, err := bindDiamondCutFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetFilterer{contract: contract}, nil
}

// bindDiamondCutFacet binds a generic wrapper to an already deployed contract.
func bindDiamondCutFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondCutFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.DiamondCutFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCutFacetDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCutIterator struct {
	Event *DiamondCutFacetDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondCutFacetDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutFacetDiamondCut)
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
		it.Event = new(DiamondCutFacetDiamondCut)
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
func (it *DiamondCutFacetDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutFacetDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutFacetDiamondCut represents a DiamondCut event raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondCutFacetDiamondCutIterator, error) {

	logs, sub, err := _DiamondCutFacet.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetDiamondCutIterator{contract: _DiamondCutFacet.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondCutFacetDiamondCut) (event.Subscription, error) {

	logs, sub, err := _DiamondCutFacet.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutFacetDiamondCut)
				if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) ParseDiamondCut(log types.Log) (*DiamondCutFacetDiamondCut, error) {
	event := new(DiamondCutFacetDiamondCut)
	if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
