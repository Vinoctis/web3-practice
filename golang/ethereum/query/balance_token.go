package query

import (
	"math/big"
	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/Vinoctis/ethereum-practice/contract/gen"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"fmt"
	"math"
)

func GetTokenBalance(ethClient *svc.EthClient, addressHex string, tokenHex string) *big.Int {
	client := ethClient.Client 
	address := common.HexToAddress(addressHex)
	tokenAddress := common.HexToAddress(tokenHex)
	instance, err := contract.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	//这边是bind.CallOpts 是只读调用的配置项，也可以直接传nil
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals)
	fmt.Printf("balance wei: %s\n", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance value(token): %f\n", value)
	return bal
}