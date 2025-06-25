package query 

import (
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/Vinoctis/ethereum-practice/svc"
	"fmt"
	"context"
)

func GetBlockInfo(client *svc.EthClient, blockNum int64, cprint bool) *types.Block {
	block, err := client.Client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	if err != nil {
		log.Fatal(err)
	}

	if cprint {
		fmt.Println("\n****** 区块信息 *******\n")
		fmt.Println("区块高度：", block.Number())
		fmt.Println("区块时间戳：", block.Time())
		fmt.Println("出块难度：", block.Difficulty())  
		fmt.Println("区块交易数：", len(block.Transactions()))
		count, _ := client.Client.TransactionCount(context.Background(), block.Hash())
		fmt.Println("区块交易数：", count)
		fmt.Println("区块Hash：", block.Hash().Hex())
		fmt.Println("区块父Hash：", block.ParentHash().Hex())
		
	}
	
	return block
}

// 订阅最新区块
func DescribeLatestBlock(ethClient *svc.EthClient) {
	client := ethClient.RpcClient
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <- sub.Err():
			log.Fatal(err)
		case header := <- headers:
			fmt.Println("\n****** 区块信息 *******\n")
			fmt.Println("最新区块Hash：", header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("最新区块时间戳：", block.Time())
			fmt.Println("最新区块交易数：", len(block.Transactions()))
			fmt.Println("区块高度：", block.Number())
			fmt.Println("区块时间戳：", block.Time())
			fmt.Println("出块难度：", block.Difficulty())  
			fmt.Println("区块交易数：", len(block.Transactions()))
			count, _ := client.TransactionCount(context.Background(), block.Hash())
			fmt.Println("区块交易数：", count)
			fmt.Println("区块父Hash：", block.ParentHash().Hex())
		}
	}
}							