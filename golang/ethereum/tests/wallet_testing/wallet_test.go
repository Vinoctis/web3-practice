package wallet

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Vinoctis/ethereum-practice/wallet"
	"github.com/Vinoctis/ethereum-practice/svc"
	// "math/big"
)

// func TestCrateWllet(t *testing.T) {
// 	privateKey := wallet.CreateWallet()
// 	assert.NotNil(t, privateKey)
// }

func TestTransfer(t *testing.T) {
	privateKeyECDSA := "b0adaf63abfaf4e6171f6de80ae3f657d0858db8adc8343768763c2e856532e1"
	privateKey := wallet.CreateWallet(privateKeyECDSA)
	assert.NotNil(t, privateKey)
	ethClient := svc.NewEthClient()
	assert.NotNil(t, ethClient)
	// toAddress := "0x3f8fBE733456e503c689451c4953E7B680B668d1"
	// value := new(big.Int)
	// value.SetString("100000000000000", 10) //10^-4 eth 
	// transactionHash := wallet.Transfer(ethClient, privateKey, toAddress, value)
	// assert.True(t, len(transactionHash) == 66 , "交易的 hash 不合法")
	// assert.True(t, transactionHash[:2] == "0x", "交易哈希缺少前缀0x")
}


func TestTransferToken(t *testing.T) {
	privateKeyECDSA := "b0adaf63abfaf4e6171f6de80ae3f657d0858db8adc8343768763c2e856532e1"
	privateKey := wallet.CreateWallet(privateKeyECDSA)
	assert.NotNil(t, privateKey)
	ethClient := svc.NewEthClient()
	assert.NotNil(t, ethClient)
	toAddress := "0x3f8fBE733456e503c689451c4953E7B680B668d1"
	tokenAddress := "0x90F7C3583dbB7dAb975541fD55D2988c4Cf23363"
	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10)//转了 100 个 token
	txHash := wallet.TransferToken(ethClient, privateKey, toAddress, tokenAddress, amount)
	assert.NotNil(t, txHash)
}



