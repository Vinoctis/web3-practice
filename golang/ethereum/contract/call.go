package contract 

import (
	"math/big"
	"context"
	"fmt"
	"log"
	"io/ioutil"
	"strings"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/Vinoctis/ethereum-practice/contract/gen"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
)

func CallContract(ethClient *svc.EthClient, contract *gen.TestDeployContract, privateKey string) {
	callOpt := &bind.CallOpts{Context: context.Background()}
	num, err := contract.MyNum(nil)
	if err != nil {
		log.Fatal(num)
	}
	fmt.Printf("TestDeployContract initial MyNum is: %d\n", num)
	
	client := ethClient.Client
	privateECDSA,err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	chainID,err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateECDSA, chainID)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(int64(789))
	tx, err := contract.SetNum(auth, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	//等待交易确认
	fmt.Println("等待交易确认...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("交易确认成功")
	} else {
		log.Fatal("交易失败，请前往 ethereum.io 查看")
	}

	newNum , err := contract.MyNum(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TestDeployContract new MyNum is: %d\n", newNum)
}

func CallContractWithABI(ethClient *svc.EthClient, contractAddr string, abiPath string, privateKey string) {
	privateKeyECDSA,err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	
	//获取pending 的nonce PendingNonceAt
	client := ethClient.Client
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	//获取gas price 
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(210000)

	//获取合约实例
	abiBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		log.Fatal(err)
	}
	parsedABI , err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatal(err)
	}
	methodName := "setNum"
	newNum := big.NewInt(1008)
	input , err := parsedABI.Pack(methodName, newNum)
	if err != nil {
		log.Fatal(err)
	}

	//创建交易并签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(0)
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), value, gasLimit, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	//等待交易确认
	fmt.Println("等待交易确认...")
	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("交易确认成功")
	} else {
		log.Fatal("交易失败，请前往 ethereum.io 查看")
	}

	// 创建合约实例, 新的方式获取 合约实例
    contract := bind.NewBoundContract(common.HexToAddress(contractAddr), parsedABI, client, client, client)
    
    // 调用只读方法
	var myNum *big.Int
	result := []interface{}{&myNum}
	callOpts := &bind.CallOpts{Context: context.Background()}
	err = contract.Call(callOpts, &result, "myNum")				
    if err != nil {
        log.Fatal("查看myNum失败")
    }
	fmt.Printf("合约 myNum 当前值: %s\n", myNum.String())
}