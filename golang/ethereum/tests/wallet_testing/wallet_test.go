package wallet

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Vinoctis/ethereum-practice/wallet"
)

func TestCrateWllet(t *testing.T) {
	privateKey := wallet.CreateWallet()
	assert.NotNil(t, privateKey)
}