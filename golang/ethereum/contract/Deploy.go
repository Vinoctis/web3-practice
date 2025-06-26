// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DeployContractMetaData contains all meta data concerning the DeployContract contract.
var DeployContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"myNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNum\",\"type\":\"uint256\"}],\"name\":\"setNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600a5f553480156012575f5ffd5b50610126806100205f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c8063cd16ecbf146034578063d4da005a14604c575b5f5ffd5b604a60048036038101906046919060a6565b6066565b005b6052606f565b604051605d919060d9565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6088816078565b81146091575f5ffd5b50565b5f8135905060a0816081565b92915050565b5f6020828403121560b85760b76074565b5b5f60c3848285016094565b91505092915050565b60d3816078565b82525050565b5f60208201905060ea5f83018460cc565b9291505056fea26469706673582212204a6f5828a0cff0e25cd8243f2260e6c7dd3e76b9427dabc8b4df1bd055a3e59364736f6c634300081e0033",
}

// DeployContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployContractMetaData.ABI instead.
var DeployContractABI = DeployContractMetaData.ABI

// DeployContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DeployContractMetaData.Bin instead.
var DeployContractBin = DeployContractMetaData.Bin

// DeployDeployContract deploys a new Ethereum contract, binding an instance of DeployContract to it.
func DeployDeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DeployContract, error) {
	parsed, err := DeployContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DeployContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DeployContract{DeployContractCaller: DeployContractCaller{contract: contract}, DeployContractTransactor: DeployContractTransactor{contract: contract}, DeployContractFilterer: DeployContractFilterer{contract: contract}}, nil
}

// DeployContract is an auto generated Go binding around an Ethereum contract.
type DeployContract struct {
	DeployContractCaller     // Read-only binding to the contract
	DeployContractTransactor // Write-only binding to the contract
	DeployContractFilterer   // Log filterer for contract events
}

// DeployContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployContractSession struct {
	Contract     *DeployContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployContractCallerSession struct {
	Contract *DeployContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DeployContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployContractTransactorSession struct {
	Contract     *DeployContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DeployContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployContractRaw struct {
	Contract *DeployContract // Generic contract binding to access the raw methods on
}

// DeployContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployContractCallerRaw struct {
	Contract *DeployContractCaller // Generic read-only contract binding to access the raw methods on
}

// DeployContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployContractTransactorRaw struct {
	Contract *DeployContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployContract creates a new instance of DeployContract, bound to a specific deployed contract.
func NewDeployContract(address common.Address, backend bind.ContractBackend) (*DeployContract, error) {
	contract, err := bindDeployContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeployContract{DeployContractCaller: DeployContractCaller{contract: contract}, DeployContractTransactor: DeployContractTransactor{contract: contract}, DeployContractFilterer: DeployContractFilterer{contract: contract}}, nil
}

// NewDeployContractCaller creates a new read-only instance of DeployContract, bound to a specific deployed contract.
func NewDeployContractCaller(address common.Address, caller bind.ContractCaller) (*DeployContractCaller, error) {
	contract, err := bindDeployContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployContractCaller{contract: contract}, nil
}

// NewDeployContractTransactor creates a new write-only instance of DeployContract, bound to a specific deployed contract.
func NewDeployContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployContractTransactor, error) {
	contract, err := bindDeployContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployContractTransactor{contract: contract}, nil
}

// NewDeployContractFilterer creates a new log filterer instance of DeployContract, bound to a specific deployed contract.
func NewDeployContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployContractFilterer, error) {
	contract, err := bindDeployContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployContractFilterer{contract: contract}, nil
}

// bindDeployContract binds a generic wrapper to an already deployed contract.
func bindDeployContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeployContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployContract *DeployContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployContract.Contract.DeployContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployContract *DeployContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployContract.Contract.DeployContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployContract *DeployContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployContract.Contract.DeployContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployContract *DeployContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployContract *DeployContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployContract *DeployContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployContract.Contract.contract.Transact(opts, method, params...)
}

// MyNum is a free data retrieval call binding the contract method 0xd4da005a.
//
// Solidity: function myNum() view returns(uint256)
func (_DeployContract *DeployContractCaller) MyNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DeployContract.contract.Call(opts, &out, "myNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyNum is a free data retrieval call binding the contract method 0xd4da005a.
//
// Solidity: function myNum() view returns(uint256)
func (_DeployContract *DeployContractSession) MyNum() (*big.Int, error) {
	return _DeployContract.Contract.MyNum(&_DeployContract.CallOpts)
}

// MyNum is a free data retrieval call binding the contract method 0xd4da005a.
//
// Solidity: function myNum() view returns(uint256)
func (_DeployContract *DeployContractCallerSession) MyNum() (*big.Int, error) {
	return _DeployContract.Contract.MyNum(&_DeployContract.CallOpts)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _newNum) returns()
func (_DeployContract *DeployContractTransactor) SetNum(opts *bind.TransactOpts, _newNum *big.Int) (*types.Transaction, error) {
	return _DeployContract.contract.Transact(opts, "setNum", _newNum)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _newNum) returns()
func (_DeployContract *DeployContractSession) SetNum(_newNum *big.Int) (*types.Transaction, error) {
	return _DeployContract.Contract.SetNum(&_DeployContract.TransactOpts, _newNum)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _newNum) returns()
func (_DeployContract *DeployContractTransactorSession) SetNum(_newNum *big.Int) (*types.Transaction, error) {
	return _DeployContract.Contract.SetNum(&_DeployContract.TransactOpts, _newNum)
}
