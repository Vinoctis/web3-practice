package wallet

import (
	"math/big"

	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"context"
	"crypto/ecdsa"
	"log"
	"fmt"
)

func Transfer(ethClient *svc.EthClient, privateKey *ecdsa.PrivateKey, to string, value *big.Int) string{
	publicKey := privateKey.Public()
	publicECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	client := ethClient.Client
	fromAddress := crypto.PubkeyToAddress(*publicECDSA)
	nounce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err!= nil {
		log.Fatal(err)
	}

	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err!= nil  {
		log.Fatal(err)	
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTransaction(nounce, toAddress, value, gasLimit, gasPrice, data)
``
	chainID, err := client.NetworkID(context.Background())
	if err!= nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return signedTx.Hash().Hex()
}