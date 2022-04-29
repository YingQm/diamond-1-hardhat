package test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/YingQm/diamond-1-hardhat/contracts/generated"
	"github.com/YingQm/diamond-1-hardhat/test/ethinterface"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	//txslog = log15.New("ethereum relayer", "ethtxs")
	GasLimit4Deploy = uint64(0)
)

type DeployPara struct {
	DeployPrivateKey *ecdsa.PrivateKey
	Deployer         common.Address
	Operator         common.Address
}

type DeployResult struct {
	Address common.Address
	TxHash  string
}

//PrepareTestEnv ...
func PrepareTestEnv() (*ethinterface.SimExtend, *DeployPara) {
	genesiskey, _ := crypto.GenerateKey()
	alloc := make(core.GenesisAlloc)
	genesisAddr := crypto.PubkeyToAddress(genesiskey.PublicKey)
	genesisAccount := core.GenesisAccount{
		Balance:    big.NewInt(1000000000000 * 10000),
		PrivateKey: crypto.FromECDSA(genesiskey),
	}
	alloc[genesisAddr] = genesisAccount

	gasLimit := uint64(100000000)
	sim := new(ethinterface.SimExtend)
	sim.SimulatedBackend = backends.NewSimulatedBackend(alloc, gasLimit)

	para := &DeployPara{
		DeployPrivateKey: genesiskey,
		Deployer:         genesisAddr,
		Operator:         genesisAddr,
	}

	return sim, para
}

//PrepareAuth ...
func PrepareAuth(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, transactor common.Address) (*bind.TransactOpts, error) {
	if nil == privateKey || nil == client {
		//txslog.Error("PrepareAuth", "nil input parameter", "client", client, "privateKey", privateKey)
		return nil, errors.New("nil input parameter")
	}

	ctx := context.Background()
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		//txslog.Error("PrepareAuth", "Failed to SuggestGasPrice due to:", err.Error())
		return nil, errors.New("failed to get suggest gas price " + err.Error())
	}

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		//txslog.Error("PrepareAuth NetworkID", "err", err)
		return nil, err
	}

	_, isSim := client.(*ethinterface.SimExtend)
	if isSim {
		chainID = big.NewInt(1337)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		//txslog.Error("PrepareAuth NewKeyedTransactorWithChainID", "err", err, "chainID", chainID)
		return nil, err
	}
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = GasLimit4Deploy
	auth.GasPrice = gasPrice

	nonce, err := client.PendingNonceAt(context.Background(), transactor)
	if nil != err {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	return auth, nil
}

// DeployDiamondCutFacet
func DeployDiamondCutFacet(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, deployer common.Address) (*generated.DiamondCutFacet, *DeployResult, error) {
	auth, err := PrepareAuth(client, privateKey, deployer)
	if nil != err {
		return nil, nil, err
	}

	addr, tx, diamondCutFacet, err := generated.DeployDiamondCutFacet(auth, client)
	if err != nil {
		return nil, nil, err
	}

	deployResult := &DeployResult{
		Address: addr,
		TxHash:  tx.Hash().String(),
	}

	return diamondCutFacet, deployResult, nil
}

func DeployDiamond(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, deployer common.Address, diamondCutFacet common.Address) (*generated.Diamond, *DeployResult, error) {
	auth, err := PrepareAuth(client, privateKey, deployer)
	if nil != err {
		return nil, nil, err
	}

	addr, tx, diamond, err := generated.DeployDiamond(auth, client, deployer, diamondCutFacet)
	if err != nil {
		return nil, nil, err
	}

	deployResult := &DeployResult{
		Address: addr,
		TxHash:  tx.Hash().String(),
	}

	return diamond, deployResult, nil
}

func DeployDiamondInit(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, deployer common.Address) (*generated.DiamondInit, *DeployResult, error) {
	auth, err := PrepareAuth(client, privateKey, deployer)
	if nil != err {
		return nil, nil, err
	}

	addr, tx, diamondInit, err := generated.DeployDiamondInit(auth, client)
	if err != nil {
		return nil, nil, err
	}

	deployResult := &DeployResult{
		Address: addr,
		TxHash:  tx.Hash().String(),
	}

	return diamondInit, deployResult, nil
}

func DeployDiamondLoupeFacet(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, deployer common.Address) (*generated.DiamondLoupeFacet, *DeployResult, error) {
	auth, err := PrepareAuth(client, privateKey, deployer)
	if nil != err {
		return nil, nil, err
	}

	addr, tx, diamondLoupeFacet, err := generated.DeployDiamondLoupeFacet(auth, client)
	if err != nil {
		return nil, nil, err
	}

	deployResult := &DeployResult{
		Address: addr,
		TxHash:  tx.Hash().String(),
	}

	return diamondLoupeFacet, deployResult, nil
}

func DeployOwnershipFacet(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, deployer common.Address) (*generated.OwnershipFacet, *DeployResult, error) {
	auth, err := PrepareAuth(client, privateKey, deployer)
	if nil != err {
		return nil, nil, err
	}

	addr, tx, ownershipFacet, err := generated.DeployOwnershipFacet(auth, client)
	if err != nil {
		return nil, nil, err
	}

	deployResult := &DeployResult{
		Address: addr,
		TxHash:  tx.Hash().String(),
	}

	return ownershipFacet, deployResult, nil
}

func Test_DeployContracts(t *testing.T) {
	//ctx := context.Background()
	sim, para := PrepareTestEnv()
	diamondCutFacet, deployDiamondCutFacetResult, err := DeployDiamondCutFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondCutFacet, deployDiamondCutFacetResult, err)

	diamond, deployDiamondResult, err := DeployDiamond(sim, para.DeployPrivateKey, para.Deployer, deployDiamondCutFacetResult.Address)
	fmt.Println(diamond, deployDiamondResult, err)

	diamondInit, deployDiamondInitResult, err := DeployDiamondInit(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondInit, deployDiamondInitResult, err)

	var _diamondCut []generated.IDiamondCutFacetCut

	diamondLoupeFacet, deployDiamondLoupeFacetResult, err := DeployDiamondLoupeFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondLoupeFacet, deployDiamondLoupeFacetResult, err)

	//_diamondCut=append(_diamondCut, generated.IDiamondCutFacetCut{FacetAddress:deployDiamondLoupeFacetResult.Address, Action:0,FunctionSelectors:nil})

	ownershipFacet, deployOwnershipFacetResult, err := DeployOwnershipFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(ownershipFacet, deployOwnershipFacetResult, err)

	//cut.push({
	//facetAddress: facet.address,
	//	action: FacetCutAction.Add,
	//		functionSelectors: getSelectors(facet)
	//})

	//{// get function selectors from ABI
	//	function getSelectors (contract) {
	//	const signatures = Object.keys(contract.interface.functions)
	//	const selectors = signatures.reduce((acc, val) => {
	//	if (val !== 'init(bytes)') {
	//	acc.push(contract.interface.getSighash(val))
	//	}
	//	return acc
	//	}, [])
	//selectors.contract = contract
	//selectors.remove = remove
	//selectors.get = get
	//	return selectors
	//}}
	parsed, err := generated.DiamondInitMetaData.GetAbi()
	if nil != err {
		require.Nil(t, err)
	}
	functionCall, err := parsed.Pack("init")
	if nil != err {
		require.Nil(t, err)
	}

	diamondCut, err := generated.NewIDiamondCut(deployDiamondResult.Address, sim)
	if nil != err {
		require.Nil(t, err)
	}
	opts := &bind.TransactOpts{
		From:    deployDiamondResult.Address,
		Context: context.Background(),
	}
	//  DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte)
	tx, err := diamondCut.DiamondCut(opts, _diamondCut, deployDiamondInitResult.Address, functionCall)
	if nil != err {
		require.Nil(t, err)
	}
	sim.Commit()

	fmt.Println(tx)

}

//func GetSelectors(metaData *bind.MetaData) [][4]byte {
//	var functionSelectors [][4]byte
//
//	parsed, err := metaData.GetAbi()
//	if nil != err {
//		fmt.Println("getabi err:",err)
//		return functionSelectors
//		//require.Nil(t, err)
//	}
//
//	for k, v := range metaData.Sigs {
//		fmt.Println(k, v)
//		method, exist := parsed.Methods[v]
//		if exist {
//			functionSelectors = append(functionSelectors, [4]byte(method.ID))
//		}
//	}
//
//return functionSelectors
//}
