package wallet

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
	"crypto/ecdsa"
	"log"
	"fmt"
)

func CreateWallet(hex string) *ecdsa.PrivateKey {
	var err error
	var privateKey *ecdsa.PrivateKey
	if hex != "" {
		privateKey, err = crypto.HexToECDSA(hex)
		if err!= nil {
			log.Fatal(err)
		}
	} else {
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Println("private key: ", hexutil.Encode(privateKeyBytes)[2:]) //去掉Ox前缀
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("public key: ", hexutil.Encode(publicKeyBytes)[4:]) //去掉Ox前缀和04前缀
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex() //公钥转地址,也即是后面20位，40字节
	fmt.Println("address: ", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("clip:", hexutil.Encode(hash.Sum(nil)[12:])) //取后20位,原长32位

	return privateKey
}