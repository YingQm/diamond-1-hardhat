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

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondMetaData contains all meta data concerning the Diamond contract.
var DiamondMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_diamondCutFacet\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052604051620024ed380380620024ed833981016040819052620000269162000fed565b6200003c826200015660201b620000b61760201c565b604080516001808252818301909252600091816020015b60408051606080820183526000808352602083015291810191909152815260200190600190039081620000535750506040805160018082528183019092529192506000919060208083019080368337019050509050631f931c1c60e01b81600081518110620000c657620000c662001025565b6001600160e01b031990921660209283029190910182015260408051606081019091526001600160a01b038516815290810160008152602001828152508260008151811062000119576200011962001025565b60200260200101819052506200014c82600060405180602001604052806000815250620001da60201b620001391760201c565b5050505062001275565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f80546001600160a01b031981166001600160a01b0384811691821790935560405160008051602062002419833981519152939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60005b8351811015620003e6576000848281518110620001fe57620001fe62001025565b6020026020010151602001519050600060028111156200022257620002226200103b565b8160028111156200023757620002376200103b565b0362000295576200028f85838151811062000256576200025662001025565b60200260200101516000015186848151811062000277576200027762001025565b6020026020010151604001516200043560201b60201c565b620003d0565b6001816002811115620002ac57620002ac6200103b565b0362000304576200028f858381518110620002cb57620002cb62001025565b602002602001015160000151868481518110620002ec57620002ec62001025565b602002602001015160400151620006d860201b60201c565b60028160028111156200031b576200031b6200103b565b0362000373576200028f8583815181106200033a576200033a62001025565b6020026020010151600001518684815181106200035b576200035b62001025565b602002602001015160400151620009c560201b60201c565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b60648201526084015b60405180910390fd5b5080620003dd8162001067565b915050620001dd565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb6738383836040516200041c93929190620010e0565b60405180910390a162000430828262000d8d565b505050565b60008151116200048b5760405162461bcd60e51b815260206004820152602b6024820152600080516020620024cd83398151915260448201526a1858d95d081d1bc818dd5d60aa1b6064820152608401620003c7565b600080516020620024858339815191525460008051602062002419833981519152906001600160a01b0384166200051a5760405162461bcd60e51b815260206004820152602c60248201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260448201526b65206164647265737328302960a01b6064820152608401620003c7565b6200053f84604051806060016040528060248152602001620024396024913962000fac565b60005b8351811015620006d157600084828151811062000563576200056362001025565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156200060b5760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f60448201527f6e207468617420616c72656164792065786973747300000000000000000000006064820152608401620003c7565b6040805180820182526001600160a01b03808a16825261ffff80881660208085019182526001600160e01b0319881660009081528b8252958620945185549251909316600160a01b026001600160b01b0319909216929093169190911717909155600180880180549182018155835291206008820401805460e085901c60046007909416939093026101000a92830263ffffffff909302191691909117905583620006b681620011e7565b94505050508080620006c89062001067565b91505062000542565b5050505050565b60008151116200072e5760405162461bcd60e51b815260206004820152602b6024820152600080516020620024cd83398151915260448201526a1858d95d081d1bc818dd5d60aa1b6064820152608401620003c7565b600080516020620024198339815191526001600160a01b038316620007af5760405162461bcd60e51b815260206004820152603060248201527f4c69624469616d6f6e644375743a205265706c6163652066616365742063616e60448201526f2774206265206164647265737328302960801b6064820152608401620003c7565b620007d483604051806060016040528060288152602001620024a56028913962000fac565b60005b8251811015620009bf576000838281518110620007f857620007f862001025565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b0316308103620008935760405162461bcd60e51b815260206004820152602f60248201527f4c69624469616d6f6e644375743a2043616e2774207265706c61636520696d6d60448201526e3aba30b1363290333ab731ba34b7b760891b6064820152608401620003c7565b856001600160a01b0316816001600160a01b0316036200090b5760405162461bcd60e51b81526020600482015260386024820152600080516020620023f983398151915260448201527f6374696f6e20776974682073616d652066756e6374696f6e00000000000000006064820152608401620003c7565b6001600160a01b038116620009785760405162461bcd60e51b81526020600482015260386024820152600080516020620023f983398151915260448201527f6374696f6e207468617420646f65736e277420657869737400000000000000006064820152608401620003c7565b506001600160e01b031916600090815260208390526040902080546001600160a01b0319166001600160a01b03861617905580620009b68162001067565b915050620007d7565b50505050565b600081511162000a1b5760405162461bcd60e51b815260206004820152602b6024820152600080516020620024cd83398151915260448201526a1858d95d081d1bc818dd5d60aa1b6064820152608401620003c7565b600080516020620024858339815191525460008051602062002419833981519152906001600160a01b0384161562000abc5760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f7665206661636574206164647260448201527f657373206d7573742062652061646472657373283029000000000000000000006064820152608401620003c7565b60005b8351811015620006d157600084828151811062000ae05762000ae062001025565b6020908102919091018101516001600160e01b0319811660009081528683526040908190208151808301909252546001600160a01b038116808352600160a01b90910461ffff16938201939093529092509062000ba65760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e27742065786973740000000000000000006064820152608401620003c7565b8051306001600160a01b039091160362000c1b5760405162461bcd60e51b815260206004820152602f60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526e3a30b1363290333ab731ba34b7b71760891b6064820152608401620003c7565b8362000c27816200120b565b94505083816020015161ffff161462000d0c57600085600101858154811062000c545762000c5462001025565b90600052602060002090600891828204019190066004029054906101000a900460e01b90508086600101836020015161ffff168154811062000c9a5762000c9a62001025565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c92909202939093179055838201516001600160e01b03199390931681529087905260409020805461ffff60a01b1916600160a01b61ffff909316929092029190911790555b8460010180548062000d225762000d2262001225565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b0319909316815291859052506040902080546001600160b01b03191690558062000d848162001067565b91505062000abf565b6001600160a01b03821662000e175780511562000e135760405162461bcd60e51b815260206004820152603c60248201527f4c69624469616d6f6e644375743a205f696e697420697320616464726573732860448201527f3029206275745f63616c6c64617461206973206e6f7420656d707479000000006064820152608401620003c7565b5050565b600081511162000e905760405162461bcd60e51b815260206004820152603d60248201527f4c69624469616d6f6e644375743a205f63616c6c6461746120697320656d707460448201527f7920627574205f696e6974206973206e6f7420616464726573732830290000006064820152608401620003c7565b6001600160a01b038216301462000ec65762000ec6826040518060600160405280602881526020016200245d6028913962000fac565b600080836001600160a01b03168360405162000ee391906200123b565b600060405180830381855af49150503d806000811462000f20576040519150601f19603f3d011682016040523d82523d6000602084013e62000f25565b606091505b509150915081620009bf5780511562000f54578060405162461bcd60e51b8152600401620003c7919062001259565b60405162461bcd60e51b815260206004820152602660248201527f4c69624469616d6f6e644375743a205f696e69742066756e6374696f6e2072656044820152651d995c9d195960d21b6064820152608401620003c7565b813b8181620009bf5760405162461bcd60e51b8152600401620003c7919062001259565b80516001600160a01b038116811462000fe857600080fd5b919050565b600080604083850312156200100157600080fd5b6200100c8362000fd0565b91506200101c6020840162000fd0565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016200107c576200107c62001051565b5060010190565b60005b83811015620010a057818101518382015260200162001086565b83811115620009bf5750506000910152565b60008151808452620010cc81602086016020860162001083565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b84811015620011b557898403607f19018652815180516001600160a01b031685528381015189860190600381106200115157634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b808310156200119f5783516001600160e01b031916825292860192600192909201919086019062001173565b5097850197955050509082019060010162001109565b50506001600160a01b038a16908801528681036040880152620011d98189620010b2565b9a9950505050505050505050565b600061ffff80831681810362001201576200120162001051565b6001019392505050565b6000816200121d576200121d62001051565b506000190190565b634e487b7160e01b600052603160045260246000fd5b600082516200124f81846020870162001083565b9190910192915050565b6020815260006200126e6020830184620010b2565b9392505050565b61117480620012856000396000f3fe60806040523661000b57005b600080356001600160e01b03191681526000805160206110ab833981519152602081905260409091205481906001600160a01b0316806100925760405162461bcd60e51b815260206004820181905260248201527f4469616d6f6e643a2046756e6374696f6e20646f6573206e6f7420657869737460448201526064015b60405180910390fd5b3660008037600080366000845af43d6000803e8080156100b1573d6000f35b3d6000fd5b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f80546001600160a01b031981166001600160a01b038481169182179093556040516000805160206110ab833981519152939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60005b83518110156102ff57600084828151811061015957610159610e28565b60200260200101516020015190506000600281111561017a5761017a610e3e565b81600281111561018c5761018c610e3e565b036101da576101d58583815181106101a6576101a6610e28565b6020026020010151600001518684815181106101c4576101c4610e28565b60200260200101516040015161034a565b6102ec565b60018160028111156101ee576101ee610e3e565b03610237576101d585838151811061020857610208610e28565b60200260200101516000015186848151811061022657610226610e28565b6020026020010151604001516105b0565b600281600281111561024b5761024b610e3e565b03610294576101d585838151811061026557610265610e28565b60200260200101516000015186848151811061028357610283610e28565b602002602001015160400151610877565b60405162461bcd60e51b815260206004820152602760248201527f4c69624469616d6f6e644375743a20496e636f727265637420466163657443756044820152663a20b1ba34b7b760c91b6064820152608401610089565b50806102f781610e6a565b91505061013c565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb67383838360405161033393929190610edb565b60405180910390a16103458282610bfa565b505050565b600081511161036b5760405162461bcd60e51b815260040161008990610fdb565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d546000805160206110ab833981519152906001600160a01b0384166104085760405162461bcd60e51b815260206004820152602c60248201527f4c69624469616d6f6e644375743a204164642066616365742063616e2774206260448201526b65206164647265737328302960a01b6064820152608401610089565b61042a846040518060600160405280602481526020016110cb60249139610e07565b60005b83518110156105a957600084828151811061044a5761044a610e28565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156104e85760405162461bcd60e51b815260206004820152603560248201527f4c69624469616d6f6e644375743a2043616e2774206164642066756e6374696f6044820152746e207468617420616c72656164792065786973747360581b6064820152608401610089565b6040805180820182526001600160a01b03808a16825261ffff80881660208085019182526001600160e01b0319881660009081528b8252958620945185549251909316600160a01b026001600160b01b0319909216929093169190911717909155600180880180549182018155835291206008820401805460e085901c60046007909416939093026101000a92830263ffffffff90930219169190911790558361059181611026565b945050505080806105a190610e6a565b91505061042d565b5050505050565b60008151116105d15760405162461bcd60e51b815260040161008990610fdb565b6000805160206110ab8339815191526001600160a01b03831661064f5760405162461bcd60e51b815260206004820152603060248201527f4c69624469616d6f6e644375743a205265706c6163652066616365742063616e60448201526f2774206265206164647265737328302960801b6064820152608401610089565b6106718360405180606001604052806028815260200161111760289139610e07565b60005b825181101561087157600083828151811061069157610691610e28565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b031630810361072a5760405162461bcd60e51b815260206004820152602f60248201527f4c69624469616d6f6e644375743a2043616e2774207265706c61636520696d6d60448201526e3aba30b1363290333ab731ba34b7b760891b6064820152608401610089565b856001600160a01b0316816001600160a01b0316036107b15760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e20776974682073616d652066756e6374696f6e00000000000000006064820152608401610089565b6001600160a01b03811661082d5760405162461bcd60e51b815260206004820152603860248201527f4c69624469616d6f6e644375743a2043616e2774207265706c6163652066756e60448201527f6374696f6e207468617420646f65736e277420657869737400000000000000006064820152608401610089565b506001600160e01b031916600090815260208390526040902080546001600160a01b0319166001600160a01b0386161790558061086981610e6a565b915050610674565b50505050565b60008151116108985760405162461bcd60e51b815260040161008990610fdb565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d546000805160206110ab833981519152906001600160a01b038416156109405760405162461bcd60e51b815260206004820152603660248201527f4c69624469616d6f6e644375743a2052656d6f76652066616365742061646472604482015275657373206d757374206265206164647265737328302960501b6064820152608401610089565b60005b83518110156105a957600084828151811061096057610960610e28565b6020908102919091018101516001600160e01b0319811660009081528683526040908190208151808301909252546001600160a01b038116808352600160a01b90910461ffff169382019390935290925090610a245760405162461bcd60e51b815260206004820152603760248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f76652066756e6360448201527f74696f6e207468617420646f65736e27742065786973740000000000000000006064820152608401610089565b8051306001600160a01b0390911603610a975760405162461bcd60e51b815260206004820152602f60248201527f4c69624469616d6f6e644375743a2043616e27742072656d6f766520696d6d7560448201526e3a30b1363290333ab731ba34b7b71760891b6064820152608401610089565b83610aa181611047565b94505083816020015161ffff1614610b7f576000856001018581548110610aca57610aca610e28565b90600052602060002090600891828204019190066004029054906101000a900460e01b90508086600101836020015161ffff1681548110610b0d57610b0d610e28565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c92909202939093179055838201516001600160e01b03199390931681529087905260409020805461ffff60a01b1916600160a01b61ffff909316929092029190911790555b84600101805480610b9257610b9261105e565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b0319909316815291859052506040902080546001600160b01b031916905580610bf281610e6a565b915050610943565b6001600160a01b038216610c8157805115610c7d5760405162461bcd60e51b815260206004820152603c60248201527f4c69624469616d6f6e644375743a205f696e697420697320616464726573732860448201527f3029206275745f63616c6c64617461206973206e6f7420656d707479000000006064820152608401610089565b5050565b6000815111610cf85760405162461bcd60e51b815260206004820152603d60248201527f4c69624469616d6f6e644375743a205f63616c6c6461746120697320656d707460448201527f7920627574205f696e6974206973206e6f7420616464726573732830290000006064820152608401610089565b6001600160a01b0382163014610d2a57610d2a826040518060600160405280602881526020016110ef60289139610e07565b600080836001600160a01b031683604051610d459190611074565b600060405180830381855af49150503d8060008114610d80576040519150601f19603f3d011682016040523d82523d6000602084013e610d85565b606091505b50915091508161087157805115610db0578060405162461bcd60e51b81526004016100899190611090565b60405162461bcd60e51b815260206004820152602660248201527f4c69624469616d6f6e644375743a205f696e69742066756e6374696f6e2072656044820152651d995c9d195960d21b6064820152608401610089565b813b81816108715760405162461bcd60e51b81526004016100899190611090565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610e7c57610e7c610e54565b5060010190565b60005b83811015610e9e578181015183820152602001610e86565b838111156108715750506000910152565b60008151808452610ec7816020860160208601610e83565b601f01601f19169290920160200192915050565b60006060808301818452808751808352608092508286019150828160051b8701016020808b0160005b84811015610fab57898403607f19018652815180516001600160a01b03168552838101518986019060038110610f4a57634e487b7160e01b600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b80831015610f965783516001600160e01b0319168252928601926001929092019190860190610f6c565b50978501979550505090820190600101610f04565b50506001600160a01b038a16908801528681036040880152610fcd8189610eaf565b9a9950505050505050505050565b6020808252602b908201527f4c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e206660408201526a1858d95d081d1bc818dd5d60aa1b606082015260800190565b600061ffff80831681810361103d5761103d610e54565b6001019392505050565b60008161105657611056610e54565b506000190190565b634e487b7160e01b600052603160045260246000fd5b60008251611086818460208701610e83565b9190910192915050565b6020815260006110a36020830184610eaf565b939250505056fec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a2041646420666163657420686173206e6f20636f64654c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a205265706c61636520666163657420686173206e6f20636f6465a2646970667358221220dd3162a1b1eba6c33228ecb733b14b7c5000d004b000f8c17d199e55b5366baf64736f6c634300080d00334c69624469616d6f6e644375743a2043616e2774207265706c6163652066756ec8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a2041646420666163657420686173206e6f20636f64654c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f6465c8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d4c69624469616d6f6e644375743a205265706c61636520666163657420686173206e6f20636f64654c69624469616d6f6e644375743a204e6f2073656c6563746f727320696e2066",
}

// DiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondMetaData.ABI instead.
var DiamondABI = DiamondMetaData.ABI

// DiamondBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondMetaData.Bin instead.
var DiamondBin = DiamondMetaData.Bin

// DeployDiamond deploys a new Ethereum contract, binding an instance of Diamond to it.
func DeployDiamond(auth *bind.TransactOpts, backend bind.ContractBackend, _contractOwner common.Address, _diamondCutFacet common.Address) (common.Address, *types.Transaction, *Diamond, error) {
	parsed, err := DiamondMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondBin), backend, _contractOwner, _diamondCutFacet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// Diamond is an auto generated Go binding around an Ethereum contract.
type Diamond struct {
	DiamondCaller     // Read-only binding to the contract
	DiamondTransactor // Write-only binding to the contract
	DiamondFilterer   // Log filterer for contract events
}

// DiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondSession struct {
	Contract     *Diamond          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCallerSession struct {
	Contract *DiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondTransactorSession struct {
	Contract     *DiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondRaw struct {
	Contract *Diamond // Generic contract binding to access the raw methods on
}

// DiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCallerRaw struct {
	Contract *DiamondCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondTransactorRaw struct {
	Contract *DiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamond creates a new instance of Diamond, bound to a specific deployed contract.
func NewDiamond(address common.Address, backend bind.ContractBackend) (*Diamond, error) {
	contract, err := bindDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// NewDiamondCaller creates a new read-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondCaller(address common.Address, caller bind.ContractCaller) (*DiamondCaller, error) {
	contract, err := bindDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCaller{contract: contract}, nil
}

// NewDiamondTransactor creates a new write-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondTransactor, error) {
	contract, err := bindDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondTransactor{contract: contract}, nil
}

// NewDiamondFilterer creates a new log filterer instance of Diamond, bound to a specific deployed contract.
func NewDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondFilterer, error) {
	contract, err := bindDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondFilterer{contract: contract}, nil
}

// bindDiamond binds a generic wrapper to an already deployed contract.
func bindDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.DiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondSession) Receive() (*types.Transaction, error) {
	return _Diamond.Contract.Receive(&_Diamond.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondTransactorSession) Receive() (*types.Transaction, error) {
	return _Diamond.Contract.Receive(&_Diamond.TransactOpts)
}

// IDiamondCutMetaData contains all meta data concerning the IDiamondCut contract.
var IDiamondCutMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1f931c1c": "diamondCut((address,uint8,bytes4[])[],address,bytes)",
	},
}

// IDiamondCutABI is the input ABI used to generate the binding from.
// Deprecated: Use IDiamondCutMetaData.ABI instead.
var IDiamondCutABI = IDiamondCutMetaData.ABI

// Deprecated: Use IDiamondCutMetaData.Sigs instead.
// IDiamondCutFuncSigs maps the 4-byte function signature to its string representation.
var IDiamondCutFuncSigs = IDiamondCutMetaData.Sigs

// IDiamondCut is an auto generated Go binding around an Ethereum contract.
type IDiamondCut struct {
	IDiamondCutCaller     // Read-only binding to the contract
	IDiamondCutTransactor // Write-only binding to the contract
	IDiamondCutFilterer   // Log filterer for contract events
}

// IDiamondCutCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDiamondCutCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDiamondCutTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDiamondCutFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDiamondCutSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDiamondCutSession struct {
	Contract     *IDiamondCut      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDiamondCutCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDiamondCutCallerSession struct {
	Contract *IDiamondCutCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDiamondCutTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDiamondCutTransactorSession struct {
	Contract     *IDiamondCutTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDiamondCutRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDiamondCutRaw struct {
	Contract *IDiamondCut // Generic contract binding to access the raw methods on
}

// IDiamondCutCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDiamondCutCallerRaw struct {
	Contract *IDiamondCutCaller // Generic read-only contract binding to access the raw methods on
}

// IDiamondCutTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDiamondCutTransactorRaw struct {
	Contract *IDiamondCutTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDiamondCut creates a new instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCut(address common.Address, backend bind.ContractBackend) (*IDiamondCut, error) {
	contract, err := bindIDiamondCut(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDiamondCut{IDiamondCutCaller: IDiamondCutCaller{contract: contract}, IDiamondCutTransactor: IDiamondCutTransactor{contract: contract}, IDiamondCutFilterer: IDiamondCutFilterer{contract: contract}}, nil
}

// NewIDiamondCutCaller creates a new read-only instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutCaller(address common.Address, caller bind.ContractCaller) (*IDiamondCutCaller, error) {
	contract, err := bindIDiamondCut(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutCaller{contract: contract}, nil
}

// NewIDiamondCutTransactor creates a new write-only instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutTransactor(address common.Address, transactor bind.ContractTransactor) (*IDiamondCutTransactor, error) {
	contract, err := bindIDiamondCut(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutTransactor{contract: contract}, nil
}

// NewIDiamondCutFilterer creates a new log filterer instance of IDiamondCut, bound to a specific deployed contract.
func NewIDiamondCutFilterer(address common.Address, filterer bind.ContractFilterer) (*IDiamondCutFilterer, error) {
	contract, err := bindIDiamondCut(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDiamondCutFilterer{contract: contract}, nil
}

// bindIDiamondCut binds a generic wrapper to an already deployed contract.
func bindIDiamondCut(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDiamondCutABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondCut *IDiamondCutRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondCut.Contract.IDiamondCutCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondCut *IDiamondCutRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondCut.Contract.IDiamondCutTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondCut *IDiamondCutRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondCut.Contract.IDiamondCutTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDiamondCut *IDiamondCutCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDiamondCut.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDiamondCut *IDiamondCutTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDiamondCut.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDiamondCut *IDiamondCutTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDiamondCut.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.Contract.DiamondCut(&_IDiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_IDiamondCut *IDiamondCutTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _IDiamondCut.Contract.DiamondCut(&_IDiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// IDiamondCutDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the IDiamondCut contract.
type IDiamondCutDiamondCutIterator struct {
	Event *IDiamondCutDiamondCut // Event containing the contract specifics and raw log

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
func (it *IDiamondCutDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDiamondCutDiamondCut)
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
		it.Event = new(IDiamondCutDiamondCut)
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
func (it *IDiamondCutDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDiamondCutDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDiamondCutDiamondCut represents a DiamondCut event raised by the IDiamondCut contract.
type IDiamondCutDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*IDiamondCutDiamondCutIterator, error) {

	logs, sub, err := _IDiamondCut.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &IDiamondCutDiamondCutIterator{contract: _IDiamondCut.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_IDiamondCut *IDiamondCutFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *IDiamondCutDiamondCut) (event.Subscription, error) {

	logs, sub, err := _IDiamondCut.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDiamondCutDiamondCut)
				if err := _IDiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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
func (_IDiamondCut *IDiamondCutFilterer) ParseDiamondCut(log types.Log) (*IDiamondCutDiamondCut, error) {
	event := new(IDiamondCutDiamondCut)
	if err := _IDiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibDiamondMetaData contains all meta data concerning the LibDiamond contract.
var LibDiamondMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c9cb3a8f9ef9d6004d3b01967eaf2a7295c2757ed016d27385eacadd9ec6edfa64736f6c634300080d0033",
}

// LibDiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use LibDiamondMetaData.ABI instead.
var LibDiamondABI = LibDiamondMetaData.ABI

// LibDiamondBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LibDiamondMetaData.Bin instead.
var LibDiamondBin = LibDiamondMetaData.Bin

// DeployLibDiamond deploys a new Ethereum contract, binding an instance of LibDiamond to it.
func DeployLibDiamond(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibDiamond, error) {
	parsed, err := LibDiamondMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LibDiamondBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibDiamond{LibDiamondCaller: LibDiamondCaller{contract: contract}, LibDiamondTransactor: LibDiamondTransactor{contract: contract}, LibDiamondFilterer: LibDiamondFilterer{contract: contract}}, nil
}

// LibDiamond is an auto generated Go binding around an Ethereum contract.
type LibDiamond struct {
	LibDiamondCaller     // Read-only binding to the contract
	LibDiamondTransactor // Write-only binding to the contract
	LibDiamondFilterer   // Log filterer for contract events
}

// LibDiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibDiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibDiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibDiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibDiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibDiamondSession struct {
	Contract     *LibDiamond       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibDiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibDiamondCallerSession struct {
	Contract *LibDiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LibDiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibDiamondTransactorSession struct {
	Contract     *LibDiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LibDiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibDiamondRaw struct {
	Contract *LibDiamond // Generic contract binding to access the raw methods on
}

// LibDiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibDiamondCallerRaw struct {
	Contract *LibDiamondCaller // Generic read-only contract binding to access the raw methods on
}

// LibDiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibDiamondTransactorRaw struct {
	Contract *LibDiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibDiamond creates a new instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamond(address common.Address, backend bind.ContractBackend) (*LibDiamond, error) {
	contract, err := bindLibDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibDiamond{LibDiamondCaller: LibDiamondCaller{contract: contract}, LibDiamondTransactor: LibDiamondTransactor{contract: contract}, LibDiamondFilterer: LibDiamondFilterer{contract: contract}}, nil
}

// NewLibDiamondCaller creates a new read-only instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamondCaller(address common.Address, caller bind.ContractCaller) (*LibDiamondCaller, error) {
	contract, err := bindLibDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibDiamondCaller{contract: contract}, nil
}

// NewLibDiamondTransactor creates a new write-only instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*LibDiamondTransactor, error) {
	contract, err := bindLibDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibDiamondTransactor{contract: contract}, nil
}

// NewLibDiamondFilterer creates a new log filterer instance of LibDiamond, bound to a specific deployed contract.
func NewLibDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*LibDiamondFilterer, error) {
	contract, err := bindLibDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibDiamondFilterer{contract: contract}, nil
}

// bindLibDiamond binds a generic wrapper to an already deployed contract.
func bindLibDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibDiamondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDiamond *LibDiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDiamond.Contract.LibDiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDiamond *LibDiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDiamond.Contract.LibDiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDiamond *LibDiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDiamond.Contract.LibDiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibDiamond *LibDiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LibDiamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibDiamond *LibDiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibDiamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibDiamond *LibDiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibDiamond.Contract.contract.Transact(opts, method, params...)
}

// LibDiamondDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the LibDiamond contract.
type LibDiamondDiamondCutIterator struct {
	Event *LibDiamondDiamondCut // Event containing the contract specifics and raw log

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
func (it *LibDiamondDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibDiamondDiamondCut)
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
		it.Event = new(LibDiamondDiamondCut)
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
func (it *LibDiamondDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibDiamondDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibDiamondDiamondCut represents a DiamondCut event raised by the LibDiamond contract.
type LibDiamondDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_LibDiamond *LibDiamondFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*LibDiamondDiamondCutIterator, error) {

	logs, sub, err := _LibDiamond.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &LibDiamondDiamondCutIterator{contract: _LibDiamond.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_LibDiamond *LibDiamondFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *LibDiamondDiamondCut) (event.Subscription, error) {

	logs, sub, err := _LibDiamond.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibDiamondDiamondCut)
				if err := _LibDiamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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
func (_LibDiamond *LibDiamondFilterer) ParseDiamondCut(log types.Log) (*LibDiamondDiamondCut, error) {
	event := new(LibDiamondDiamondCut)
	if err := _LibDiamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LibDiamondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LibDiamond contract.
type LibDiamondOwnershipTransferredIterator struct {
	Event *LibDiamondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LibDiamondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LibDiamondOwnershipTransferred)
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
		it.Event = new(LibDiamondOwnershipTransferred)
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
func (it *LibDiamondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LibDiamondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LibDiamondOwnershipTransferred represents a OwnershipTransferred event raised by the LibDiamond contract.
type LibDiamondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibDiamond *LibDiamondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LibDiamondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LibDiamond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LibDiamondOwnershipTransferredIterator{contract: _LibDiamond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibDiamond *LibDiamondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LibDiamondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LibDiamond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LibDiamondOwnershipTransferred)
				if err := _LibDiamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LibDiamond *LibDiamondFilterer) ParseOwnershipTransferred(log types.Log) (*LibDiamondOwnershipTransferred, error) {
	event := new(LibDiamondOwnershipTransferred)
	if err := _LibDiamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
