package query_testing

import (
	"testing"
	"github.com/Vinoctis/ethereum-practice/svc"  
	"github.com/Vinoctis/ethereum-practice/query" 
	"math/big"
	"github.com/stretchr/testify/assert"
)

var blockNum int64 = big.NewInt(5671744).Int64()

func TestGetBlockInfo(t *testing.T) {
	client := svc.NewEthClient()
	block  := query.GetBlockInfo(client, blockNum, true)
	assert.NotNil(t, block)
}

func TestGetTransactionInfo(t *testing.T) {
	client := svc.NewEthClient()
	block  := query.GetBlockInfo(client, blockNum, false)
	assert.NotNil(t, block)
	transactions := query.GetTransactionInfo(block, 1)
	assert.NotNil(t, transactions)
}

func TestGetReceiptsInfoByBlock(t *testing.T) {
	client := svc.NewEthClient()
	block  := query.GetBlockInfo(client, blockNum, false)
	assert.NotNil(t, block)
	receipts,_ := query.GetReceiptsInfoByBlock(client, block, 1)
	assert.NotNil(t, receipts)
}

func TestGetReceiptsInfoByTx(t *testing.T) {
	client := svc.NewEthClient()
	block  := query.GetBlockInfo(client, blockNum, false)
	assert.NotNil(t, block)
	transactions := query.GetTransactionInfo(block, 1)
	receipts,_ := query.GetReceiptsInfoByTx(client, *transactions[0])
	assert.NotNil(t, receipts)
}