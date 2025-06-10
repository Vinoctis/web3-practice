package wallet

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"fmt"
)

func CreateWallet() *ecdsa.PrivateKey {
	privateKey, err   := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey)
	return privateKey
}