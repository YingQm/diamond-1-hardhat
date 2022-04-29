package test

import (
	"context"
	"fmt"
	"github.com/YingQm/diamond-1-hardhat/contracts/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func Test_DeployContracts(t *testing.T) {
	//ctx := context.Background()
	sim, para := PrepareTestEnv()
	diamondCutFacet, deployDiamondCutFacetResult, err := DeployDiamondCutFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondCutFacet, deployDiamondCutFacetResult, err)

	sim.Commit()

	diamond, deployDiamondResult, err := DeployDiamond(sim, para.DeployPrivateKey, para.Deployer, deployDiamondCutFacetResult.Address)
	fmt.Println(diamond, deployDiamondResult, err)

	sim.Commit()

	diamondInit, deployDiamondInitResult, err := DeployDiamondInit(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondInit, deployDiamondInitResult, err)

	sim.Commit()

	var _diamondCut []generated.IDiamondCutFacetCut

	diamondLoupeFacet, deployDiamondLoupeFacetResult, err := DeployDiamondLoupeFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(diamondLoupeFacet, deployDiamondLoupeFacetResult, err)
	_diamondCut = append(_diamondCut, generated.IDiamondCutFacetCut{FacetAddress: deployDiamondLoupeFacetResult.Address, Action: 0, FunctionSelectors: GetSelectors(generated.DiamondLoupeFacetMetaData.Sigs)})

	sim.Commit()

	ownershipFacet, deployOwnershipFacetResult, err := DeployOwnershipFacet(sim, para.DeployPrivateKey, para.Deployer)
	fmt.Println(ownershipFacet, deployOwnershipFacetResult, err)
	_diamondCut = append(_diamondCut, generated.IDiamondCutFacetCut{FacetAddress: deployDiamondLoupeFacetResult.Address, Action: 0, FunctionSelectors: GetSelectors(generated.OwnershipFacetMetaData.Sigs)})

	sim.Commit()

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
		From: deployDiamondResult.Address,
		//From:    para.Deployer,
		Context: context.Background(),
	}

	{
		auth, err := PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
		if nil != err {
			require.Nil(t, err)
		}

		gasLimit := uint64(21100)

		amount := big.NewInt(1)
		amount, _ = amount.SetString(TrimZeroAndDot("1000000"), 10)

		tx := types.NewTx(&types.LegacyTx{
			Nonce:    uint64(auth.Nonce.Int64()),
			To:       &deployDiamondResult.Address,
			Value:    amount,
			Gas:      gasLimit, //auth.GasLimit,
			GasPrice: auth.GasPrice,
			//Data:     data,
		})

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1337)), para.DeployPrivateKey)
		if err != nil {
			require.Nil(t, err)
		}

		err = sim.SendTransaction(context.Background(), signedTx)
		if err != nil {
			require.Nil(t, err)
		}

		sim.Commit()
	}

	tx, err := diamondCut.DiamondCut(opts, _diamondCut, deployDiamondInitResult.Address, functionCall)
	if nil != err {
		require.Nil(t, err)
	}
	sim.Commit()

	fmt.Println(tx)
}
