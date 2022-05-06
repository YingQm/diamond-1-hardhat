package test

import (
	"context"
	"fmt"
	"github.com/YingQm/diamond-1-hardhat/contracts/generated"
	"github.com/YingQm/diamond-1-hardhat/contracts/testFacet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func Test_DeployContracts(t *testing.T) {
	// 部署合约
	sim, para := PrepareTestEnv()

	// deploy DiamondCutFacet
	diamondCutFacet, deployDiamondCutFacetResult, err := DeployDiamondCutFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondCutFacet, deployDiamondCutFacetResult, err)
	sim.Commit()

	// deploy Diamond
	diamond, deployDiamondResult, err := DeployDiamond(sim, para.DeployPrivateKey, para.Deployer, deployDiamondCutFacetResult.Address)
	fmt.Println(diamond, deployDiamondResult, err)
	sim.Commit()

	// deploy DiamondInit
	// DiamondInit provides a function that is called when the diamond is upgraded to initialize state variables
	// Read about how the diamondCut function works here: https://eips.ethereum.org/EIPS/eip-2535#addingreplacingremoving-functions
	diamondInit, deployDiamondInitResult, err := DeployDiamondInit(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondInit, deployDiamondInitResult, err)
	sim.Commit()

	// deploy facets
	var _diamondCut []generated.IDiamondCutFacetCut
	diamondLoupeFacet, deployDiamondLoupeFacetResult, err := DeployDiamondLoupeFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondLoupeFacet, deployDiamondLoupeFacetResult, err)
	_diamondCut = append(_diamondCut, generated.IDiamondCutFacetCut{FacetAddress: deployDiamondLoupeFacetResult.Address, Action: 0, FunctionSelectors: GetSelectors(generated.DiamondLoupeFacetMetaData.Sigs)})
	sim.Commit()

	ownershipFacet, deployOwnershipFacetResult, err := DeployOwnershipFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(ownershipFacet, deployOwnershipFacetResult, err)
	_diamondCut = append(_diamondCut, generated.IDiamondCutFacetCut{FacetAddress: deployDiamondLoupeFacetResult.Address, Action: 0, FunctionSelectors: GetSelectors(generated.OwnershipFacetMetaData.Sigs)})
	sim.Commit()

	// upgrade diamond with facets
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

	auth, err := PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}

	// call to init function
	tx, err := diamondCut.DiamondCut(auth, _diamondCut, deployDiamondInitResult.Address, functionCall)
	if nil != err {
		require.Nil(t, err.Error())
	}
	sim.Commit()

	fmt.Println(tx)

	// test
	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}

	addrFacet1, _, test1Facet, err := testFacet.DeployTest1Facet(auth, sim)
	if nil != err {
		require.Nil(t, err.Error())
	}
	fmt.Println(addrFacet1, test1Facet, err)
	sim.Commit()

	opts := &bind.CallOpts{
		Pending: true,
		From:    addrFacet1,
		Context: context.Background(),
	}

	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}

	_, err = test1Facet.SetNumA(auth, big.NewInt(2))
	if nil != err {
		require.Nil(t, err.Error())
	}

	numA, err := test1Facet.Test1FacetCaller.NumA(opts)
	fmt.Println(numA, err)

	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}

	_, err = test1Facet.SetNumB(auth, big.NewInt(3))
	if nil != err {
		require.Nil(t, err.Error())
	}

	numB, err := test1Facet.Test1FacetCaller.NumB(opts)
	fmt.Println(numB, err)

	ret, err := test1Facet.Test1FacetCaller.AddNum(opts, numA, numB)
	if nil != err {
		require.Nil(t, err.Error())
	}
	fmt.Println(ret)

	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}

	_, err = test1Facet.SetNumC(auth, ret)
	if nil != err {
		require.Nil(t, err.Error())
	}

	numC, err := test1Facet.Test1FacetCaller.NumC(opts)
	fmt.Println(numC, err)

	var _diamondCutTest []generated.IDiamondCutFacetCut
	_diamondCutTest = append(_diamondCutTest, generated.IDiamondCutFacetCut{FacetAddress: addrFacet1, Action: 0, FunctionSelectors: GetSelectors(testFacet.Test1FacetMetaData.Sigs)})

	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}
	diamondCutFacet.DiamondCut(auth, _diamondCutTest, common.HexToAddress(nullAddress), []byte(""))



}
