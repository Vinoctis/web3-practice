package query

import (
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
)

func GetTransactionInfo(block *types.Block, num int) types.Transactions {
	transactions := block.Transactions() 

	if len(transactions) == 0 {
		return nil
	}

	var count int
	if num == 0 {
		count = len(transactions)
	} else {
		count = num
	}
	fmt.Println("\n****** 交易信息 *******\n")

	for _, v := range transactions {
		if count == 0 {
			break
		} else {
			count--
		}
		fmt.Println("交易Hash：", v.Hash().Hex())
		fmt.Println("交易Nonce：", v.Nonce())
		fmt.Println("交易金额：", v.Value())
		fmt.Println("交易GasLimit：", v.Gas())
	}

	return transactions
}