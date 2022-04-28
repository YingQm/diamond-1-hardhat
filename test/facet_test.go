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
	"math/big"
	"testing"
)

var (
	//txslog = log15.New("ethereum relayer", "ethtxs")
	GasLimit4Deploy  = uint64(0)
)

type DeployPara struct {
	DeployPrivateKey *ecdsa.PrivateKey
	Deployer         common.Address
	Operator         common.Address
}

type DeployResult struct {
	Address common.Address
	TxHash string
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
	auth.Nonce=big.NewInt(int64(nonce))

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

//// DeployDiamond
//func DeployDiamond(client ethinterface.EthClientSpec, privateKey *ecdsa.PrivateKey, deployer common.Address, operator common.Address, initValidators []common.Address, initPowers []*big.Int) (*generated.Valset, *ethtxs.DeployResult, error) {
//	auth, err := PrepareAuth(client, privateKey, deployer)
//	if nil != err {
//		return nil, nil, err
//	}
//
//	//部署合约
//	//func DeployDiamond(auth *bind.TransactOpts, backend bind.ContractBackend, _diamondCut []IDiamondCutFacetCut, _args DiamondDiamondArgs) (common.Address, *types.Transaction, *Diamond, error) {
//
//		addr, tx, valset, err :=generated.DeployDiamond(auth, client) //generated.DeployValset(auth, client, operator, initValidators, initPowers)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	deployResult := &DeployResult{
//		Address: addr,
//		TxHash:  tx.Hash().String(),
//	}
//
//	return valset, deployResult, nil
//}

func Test_DeployContracts(t *testing.T){
	//ctx := context.Background()
	sim, para := PrepareTestEnv()

	//sim, isSim := client.(*ethinterface.SimExtend)
	//if isSim {
	//	fmt.Println("Use the simulator")
	//} else {
	//	fmt.Println("Use the actual Ethereum")
	//}
	//
	//x2EthContracts.Valset, deployInfo.Valset, err = DeployValset(client, para.DeployPrivateKey, para.Deployer, para.Operator, para.InitValidators, para.InitPowers)


	diamondCutFacet, deployResult,err:=DeployDiamondCutFacet(sim,para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondCutFacet, deployResult,err)

	diamondLoupeFacet, deployResult,err:=DeployDiamondLoupeFacet(sim,para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondLoupeFacet, deployResult,err)

	//opts, _ := bind.NewKeyedTransactorWithChainID(para.DeployPrivateKey, big.NewInt(1337))
	//parsed, _ := abi.JSON(strings.NewReader(generated.DiamondMetaData.Bin))
	//contractAddr, _, _, _ := bind.DeployContract(opts, parsed, common.FromHex(generated.DiamondMetaData.Bin), sim)
	//sim.Commit()
	//
	//callMsg := ethereum.CallMsg{
	//	From: para.Deployer,
	//	To:   &contractAddr,
	//	Data: common.FromHex(generated.DiamondMetaData.Bin),
	//}
	//
	//_, err := sim.EstimateGas(ctx, callMsg)
	//if nil != err {
	//	panic("failed to estimate gas due to:" + err.Error())
	//}
	//x2EthContracts, x2EthDeployInfo, err := DeployAndInit(sim, para)
	//if nil != err {
	//	return nil, nil, nil, nil, err
	//}
	//sim.Commit()
	//
	//return para, sim, x2EthContracts, x2EthDeployInfo, nil
}

