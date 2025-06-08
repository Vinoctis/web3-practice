package query 

import (
	"github.com/Vinoctis/ethereum-practice/svc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"fmt"
	"context"
	"log"
)
//通过区块查看票据信息
func GetReceiptsInfoByBlock(ethClient *svc.EthClient, block *types.Block, num int) (types.Receipts, error) {
	// use block.Hash() or block.Number().Int64()
	receipts, err := ethClient.Client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(block.Hash(), false))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var count int
	if num == 0 {
		count = len(receipts)
	} else {
		count = num
	}

	fmt.Println("\n****** 通过区块查询的票据信息 *******\n")

	for _, receipt := range receipts  {
		if count == 0 {
			break
		} else {
			count--
		}
		fmt.Println("票据状态：", receipt.Status)
		fmt.Println("票据日志：", receipt.Logs)
		fmt.Println("区块hash:", receipt.BlockHash.Hex())
		fmt.Println("订单hash:", receipt.TxHash.Hex())
		fmt.Println("交易索引：", receipt.TransactionIndex)
		fmt.Println("合约地址", receipt.ContractAddress.Hex())

		
	}
	return receipts, nil
}

//通过交易查看票据信息
func GetReceiptsInfoByTx(ethClient *svc.EthClient, tx types.Transaction ) (*types.Receipt, error) {
	// use block.Hash() or block.Number()
	receipt, err := ethClient.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("\n****** 通过单条交易查询的票据信息 *******\n")
	fmt.Println("票据状态：", receipt.Status)
	fmt.Println("票据日志：", receipt.Logs)
	fmt.Println("区块hash:", receipt.BlockHash.Hex())
	fmt.Println("订单hash:", receipt.TxHash.Hex())
	fmt.Println("交易索引：", receipt.TransactionIndex)
	fmt.Println("合约地址：", receipt.ContractAddress.Hex())
		
	
	return receipt, nil
}

