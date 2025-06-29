package contract 

import (
	"math/big"
	"context"
	"log"
	"fmt"
	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"crypto/ecdsa"
	"github.com/Vinoctis/ethereum-practice/contract/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)
//部署合约
func DeployContract(ethClient *svc.EthClient, privateKey string) {
	client := ethClient.Client
	privateECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//获取公钥生成的地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("owner address :", address.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	//设置gaslimit 获取 gasPrice
	gasLimit := uint64(210000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateECDSA,chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit  = uint64(gasLimit)
	auth.GasPrice = gasPrice

	contractAddr , tx, instance, err := gen.DeployTestDeployContract(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("合约地址hash：", contractAddr.Hex())
	fmt.Println("发布合约交易hash：", tx.Hash().Hex())
	//调用合约方法
	num, err := instance.MyNum(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract num:", num)
}

//加载合约
func LoadContract(ethClient *svc.EthClient, contractAddrHex string) *gen.TestDeployContract{
	client := ethClient.Client 
	contract, err := gen.NewTestDeployContract(common.HexToAddress(contractAddrHex), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract loaded:", contract)
	return contract
}