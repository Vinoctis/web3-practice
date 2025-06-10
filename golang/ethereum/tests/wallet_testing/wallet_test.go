package wallet_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Vinoctis/ethereum-practice/wallet"
	// "github.com/Vinoctis/ethereum-practice/svc"
)

func TestCreateWallet(t *testing.T) {
	privateKey := wallet.CreateWallet()
	assert.NotNil(t, privateKey)
}