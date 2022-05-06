package test

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
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
	"strings"
)

var (
	//txslog = log15.New("ethereum relayer", "ethtxs")
	GasLimit4Deploy = uint64(0)
	nullAddress     = "0x0000000000000000000000000000000000000000"
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
		Balance:    big.NewInt(100000000000000 * 10000),
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

	//chainID, err := client.NetworkID(ctx)
	//if err != nil {
	//	//txslog.Error("PrepareAuth NetworkID", "err", err)
	//	return nil, err
	//}
	//
	//_, isSim := client.(*ethinterface.SimExtend)
	//if isSim {
	//	chainID = big.NewInt(1337)
	//}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
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

//checksum: first four bytes of double-SHA256.
func checksum(input []byte) (cksum [4]byte) {
	h := sha256.New()
	_, err := h.Write(input)
	if err != nil {
		return
	}
	intermediateHash := h.Sum(nil)
	h.Reset()
	_, err = h.Write(intermediateHash)
	if err != nil {
		return
	}
	finalHash := h.Sum(nil)
	copy(cksum[:], finalHash[:])
	return
}

func GetSelectors(sigs map[string]string) [][4]byte {
	var functionSelectors [][4]byte
	for k, _ := range sigs {
		data := checksum([]byte(k))
		functionSelectors = append(functionSelectors, data)
	}

	fmt.Println("abi sigs:", sigs)
	fmt.Println("functionSelectors:", functionSelectors)

	return functionSelectors
}

//TrimZeroAndDot ...
func TrimZeroAndDot(s string) string {
	if strings.Contains(s, ".") {
		var trimDotStr string
		trimZeroStr := strings.TrimRight(s, "0")
		trimDotStr = strings.TrimRight(trimZeroStr, ".")
		return trimDotStr
	}

	return s
}

//{
//	auth, err := PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
//	if nil != err {
//		require.Nil(t, err)
//	}
//
//	gasLimit := uint64(21100)
//
//	amount := big.NewInt(1)
//	amount, _ = amount.SetString(TrimZeroAndDot("1000000"), 10)
//
//	tx := types.NewTx(&types.LegacyTx{
//		Nonce:    uint64(auth.Nonce.Int64()),
//		To:       &deployDiamondResult.Address,
//		Value:    amount,
//		Gas:      gasLimit, //auth.GasLimit,
//		GasPrice: auth.GasPrice,
//		//Data:     data,
//	})
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1337)), para.DeployPrivateKey)
//	if err != nil {
//		require.Nil(t, err)
//	}
//
//	err = sim.SendTransaction(context.Background(), signedTx)
//	if err != nil {
//		require.Nil(t, err)
//	}
//
//	sim.Commit()
//}
