package contract_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Vinoctis/ethereum-practice/contract"
	"github.com/Vinoctis/ethereum-practice/contract/gen"
	"github.com/Vinoctis/ethereum-practice/svc"
)

//合约的部署
func TestDeployContract(t *testing.T) {
	ethClient := svc.NewEthClient()
	assert.NotNil(t, ethClient)
	privateKeyHex := "b0adaf63abfaf4e6171f6de80ae3f657d0858db8adc8343768763c2e856532e1"
	contract.DeployContract(ethClient, privateKeyHex)
}

//合约的加载
func TestLoadContract(t *testing.T) {
	client := svc.NewEthClient()
	contractAddr := "0x9f5609BdEf5EF65f34F9f8ED43cEE6192345cB86"
	instance := contract.LoadContract(client, contractAddr)
	assert.NotNil(t, instance)
	//检查类型是否正确
	assert.IsType(t, (*gen.TestDeployContract)(nil), instance)
}

//合约的调用
func TestCallContract(t *testing.T) {
	privateKey := "b0adaf63abfaf4e6171f6de80ae3f657d0858db8adc8343768763c2e856532e1"
	client := svc.NewEthClient()
	contractAddr := "0x9f5609BdEf5EF65f34F9f8ED43cEE6192345cB86"
	instance := contract.LoadContract(client, contractAddr)
	assert.NotNil(t, instance)
	assert.IsType(t, (*gen.TestDeployContract)(nil), instance)
	contract.CallContract(client, instance, privateKey)
}

func TestCallContractWithABI(t *testing.T) {
	privateKey := "b0adaf63abfaf4e6171f6de80ae3f657d0858db8adc8343768763c2e856532e1"
	client := svc.NewEthClient()
	contractAddr := "0x9f5609BdEf5EF65f34F9f8ED43cEE6192345cB86"
	abi := "/Users/vinoctis/Work/www/web3-practice/solidity/contract/testdeploy/testDeploy.abi"
	contract.CallContractWithABI(client, contractAddr, abi, privateKey)
}