package svc 

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthClient struct {
	Client *ethclient.Client	
}
func NewEthClient() *EthClient {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/Xv7ZCu3-aOwIQSMKJ8g7EwID1YlrBA1J")
	if err!= nil {
		log.Fatal(err)
	}
	return &EthClient{Client: client}
}

func (eth *EthClient) Dial(url string) {
	client, err := ethclient.Dial(url)
	if err!= nil {
		log.Fatal(err)
	}
	eth.Client = client
}

func (eth *EthClient) GetDemoBlock() *types.Block {
	blockNumber := big.NewInt(5671744)
	block, err := eth.Client.BlockByNumber(context.Background(), blockNumber)
	if err!= nil {
		log.Fatal(err)
	}
	return block
}
