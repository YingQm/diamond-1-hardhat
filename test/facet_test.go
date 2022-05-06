package test

import (
	"context"
	"fmt"
	"github.com/YingQm/diamond-1-hardhat/contracts/generated"
	"github.com/YingQm/diamond-1-hardhat/contracts/testFacet"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

type DiamondCut struct {
	_diamondCut []generated.IDiamondCutFacetCut
	_init       common.Address
	_calldata   []byte
	//_diamondCut IDiamondCut.FacetCut[]
}

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

	// 订阅事件
	ctx := context.Background()
	eventName := "DiamondCut"
	libDiamondABI, err := generated.LibDiamondMetaData.GetAbi()
	logNewLibDiamondSig := libDiamondABI.Events[eventName].ID.Hex()
	query := ethereum.FilterQuery{
		Addresses: []common.Address{deployDiamondResult.Address},
		FromBlock: big.NewInt(int64(1)),
	}
	// We will check logs for new events
	logs := make(chan types.Log, 10)
	// Filter by contract and event, write results to logs
	sub, err := sim.SubscribeFilterLogs(ctx, query, logs)
	require.Nil(t, err)

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
	//_, err = parsed.Pack("init")
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
	//tx, err := diamondCut.DiamondCut(auth, _diamondCut, common.HexToAddress(nullAddress), []byte(""))
	if nil != err {
		require.Nil(t, err.Error())
	}
	fmt.Println("tx", tx.Hash().String(), "_diamondCut", len(_diamondCut))
	sim.Commit()

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

	//_, err = test1Facet.SetNumA(auth, big.NewInt(2))
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//
	//numA, err := test1Facet.Test1FacetCaller.NumA(opts)
	//fmt.Println(numA, err)
	//
	//auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//
	//_, err = test1Facet.SetNumB(auth, big.NewInt(3))
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//
	//numB, err := test1Facet.Test1FacetCaller.NumB(opts)
	//fmt.Println(numB, err)

	//ret, err := test1Facet.Test1FacetCaller.AddNum(opts, numA, numB)
	ret, err := test1Facet.Test1FacetCaller.AddNum(opts, big.NewInt(2), big.NewInt(3))
	if nil != err {
		require.Nil(t, err.Error())
	}
	fmt.Println(ret)

	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}

	//_, err = test1Facet.SetNumC(auth, ret)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//
	//numC, err := test1Facet.Test1FacetCaller.NumC(opts)
	//fmt.Println(numC, err)

	var _diamondCutTest1 []generated.IDiamondCutFacetCut
	_diamondCutTest1 = append(_diamondCutTest1, generated.IDiamondCutFacetCut{FacetAddress: addrFacet1, Action: 0, FunctionSelectors: GetSelectors(testFacet.Test1FacetMetaData.Sigs)})

	auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	if nil != err {
		require.Nil(t, err.Error())
	}
	tx, err = diamondCut.DiamondCut(auth, _diamondCutTest1, common.HexToAddress(nullAddress), []byte(""))
	if nil != err {
		require.Nil(t, err.Error())
	}
	sim.Commit()

	parsedTest1, err := testFacet.Test1FacetMetaData.GetAbi()
	if nil != err {
		require.Nil(t, err)
	}
	functionCall1, err := parsedTest1.Pack("AddNum", big.NewInt(2), big.NewInt(3))
	if nil != err {
		require.Nil(t, err)
	}

	cx := context.Background()
	var msg ethereum.CallMsg
	msg.To = &addrFacet1 //合约地址
	msg.Data = functionCall1
	result, err := sim.CallContract(cx, msg, nil)
	if nil != err {
		require.Nil(t, err)
	}
	fmt.Println(result)

	//parsedLoupeFacet, err := generated.DiamondLoupeFacetMetaData.GetAbi()
	//if nil != err {
	//	require.Nil(t, err)
	//}
	//_, err = parsedLoupeFacet.Pack("facetAddresses")
	//if nil != err {
	//	require.Nil(t, err)
	//}
	//
	//msg.To = &deployDiamondLoupeFacetResult.Address //合约地址
	//msg.Data = functionCall1
	//result, err = sim.CallContract(cx, msg, nil)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//fmt.Println(result)
	//

	//diamondCutLoupeFacet, err := generated.NewIDiamondLoupe(deployDiamondResult.Address, sim)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//
	//auth, err = PrepareAuth(sim, para.DeployPrivateKey, para.Deployer)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//
	//opts = &bind.CallOpts{
	//	Pending: true,
	//	//From:    deployDiamondLoupeFacetResult.Address,
	//	From: para.Deployer,
	//	Context: context.Background(),
	//}
	//
	//addrs, err := diamondCutLoupeFacet.FacetAddresses(opts)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//fmt.Println(addrs)

	//
	//facets0, err := diamondCutLoupeFacet.FacetAddresses(opts)
	//if nil != err {
	//	require.Nil(t, err.Error())
	//}
	//fmt.Println(facets0)
	//
	//facets, err := diamondLoupeFacet.FacetAddresses(opts)
	//if nil != err {
	//	require.Nil(t, err)
	//}
	//fmt.Println(facets)

	timer := time.NewTimer(20 * time.Second)
	for {
		select {
		case <-timer.C:
			t.Fatal("failed due to timeout")
		// Handle any errors
		case err := <-sub.Err():
			t.Fatalf("sub error:%s", err.Error())
		// vLog is raw event data
		case vLog := <-logs:
			// Check if the event is a 'LogLock' event
			if vLog.Topics[0].Hex() == logNewLibDiamondSig {
				t.Logf("Witnessed new event:%s, Block number:%d, Tx hash:%s", eventName, vLog.BlockNumber, vLog.TxHash.Hex())
				logEvent := &DiamondCut{}
				//err = bridgeBankABI.Unpack(logEvent, eventName, vLog.Data)
				_, err = libDiamondABI.Unpack(eventName, vLog.Data)
				require.Nil(t, err)
				fmt.Println("logEvent", logEvent._init, logEvent._diamondCut, logEvent._calldata)
				//return
			}
		}
	}
}
