package query_testing

import (
	"testing"
	"github.com/Vinoctis/ethereum-practice/svc"  
	"github.com/Vinoctis/ethereum-practice/query" 
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	client := svc.NewEthClient()
	address := "0xb0Ee223CfeCABE681B66aDd0eA3288ca1CAf95f0"
	query.GetBalance(client, address)
}

func TestGetTokenBalance(t *testing.T) {
	client := svc.NewEthClient()
	address := "0xb0Ee223CfeCABE681B66aDd0eA3288ca1CAf95f0"
	token   := "0x90F7C3583dbB7dAb975541fD55D2988c4Cf23363"
	value   := query.GetTokenBalance(client, address, token)
	assert.NotNil(t, value)
}