package wallet

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Vinoctis/ethereum-practice/wallet"
	"github.com/Vinoctis/ethereum-practice/svc"
)


func TestDeployContract(t *testing.T) {
	ethClient := svc.NewEthClient()
	assert.NotNil(t, ethClient)
	privateKeyHex := "b0adaf63abfaf4e6171f6de80ae3f657d0858db8adc8343768763c2e856532e1"
	wallet.DeployContract(ethClient, privateKeyHex)
}