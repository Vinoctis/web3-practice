package wallet 

import (
	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"crypto/ecdsa"
	"context"
	"log"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum"
)

func TransferToken(ethClient *svc.EthClient, privateKey *ecdsa.PrivateKey, to string, token string,  amount *big.Int) string {
	publicKey := privateKey.Public()
	publicECDSA, ok  := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(ok)
	}

	client := ethClient.Client 

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice is :", gasPrice)

	fromAddress := crypto.PubkeyToAddress(*publicECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(token)
	// 函数名将是传递函数的名称，相当于行数选择器
	// 即 ERC-20 规范中的 transfer 和参数类型。 第一个参数类型是 address（令牌的接收者），第二个类型是 uint256（要发送的代币数量）
	//不需要没有空格和参数名称。 我们还需要用字节切片格式。
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	//生成函数签名的 Keccak256 哈希。 然后我们只使用前 4 个字节来获取方法 ID。
	methodID := hash.Sum(nil)[:4]
	fmt.Println("methodID: ", hexutil.Encode(methodID))// 0xa9059cbb
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32) // 地址左填充至32字节
	fmt.Println("paddedAddress: ",hexutil.Encode(paddedAddress)) // 数值左填充至32字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println("paddedAmount: " , hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit , err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data, 
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasLimit: ", gasLimit)
	//因为是转账代币所以 value传 0 
	value := big.NewInt(0)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID , err  := client.NetworkID(context.Background())
	if err!= nil {
		log.Fatal(err)
	}
	signedTx , err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err =  client.SendTransaction(context.Background(), signedTx)
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx.Hash().Hex()
}