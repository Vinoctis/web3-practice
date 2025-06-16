package query 

import (
	"math"
	"math/big"
	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/ethereum/go-ethereum/common"
	"context"
	"log"
	"fmt"
)

func GetBalance(ethClient *svc.EthClient, address string) {
	
	client := ethClient.Client
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance: ", balance)
	
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println("balanceAt: ", balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue: ", ethValue)
	fmt.Println("ethValue: ", ethValue.Text('f', 3))
	penddingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("penddingBalance: ", penddingBalance)
}
