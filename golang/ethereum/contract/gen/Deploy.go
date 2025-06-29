// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

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

// TestDeployContractMetaData contains all meta data concerning the TestDeployContract contract.
var TestDeployContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"myNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newNum\",\"type\":\"uint256\"}],\"name\":\"setNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600a5f553480156012575f5ffd5b50610126806100205f395ff3fe6080604052348015600e575f5ffd5b50600436106030575f3560e01c8063cd16ecbf146034578063d4da005a14604c575b5f5ffd5b604a60048036038101906046919060a6565b6066565b005b6052606f565b604051605d919060d9565b60405180910390f35b805f8190555050565b5f5481565b5f5ffd5b5f819050919050565b6088816078565b81146091575f5ffd5b50565b5f8135905060a0816081565b92915050565b5f6020828403121560b85760b76074565b5b5f60c3848285016094565b91505092915050565b60d3816078565b82525050565b5f60208201905060ea5f83018460cc565b9291505056fea26469706673582212204a6f5828a0cff0e25cd8243f2260e6c7dd3e76b9427dabc8b4df1bd055a3e59364736f6c634300081e0033",
}

// TestDeployContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestDeployContractMetaData.ABI instead.
var TestDeployContractABI = TestDeployContractMetaData.ABI

// TestDeployContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestDeployContractMetaData.Bin instead.
var TestDeployContractBin = TestDeployContractMetaData.Bin

// DeployTestDeployContract deploys a new Ethereum contract, binding an instance of TestDeployContract to it.
func DeployTestDeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestDeployContract, error) {
	parsed, err := TestDeployContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestDeployContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestDeployContract{TestDeployContractCaller: TestDeployContractCaller{contract: contract}, TestDeployContractTransactor: TestDeployContractTransactor{contract: contract}, TestDeployContractFilterer: TestDeployContractFilterer{contract: contract}}, nil
}

// TestDeployContract is an auto generated Go binding around an Ethereum contract.
type TestDeployContract struct {
	TestDeployContractCaller     // Read-only binding to the contract
	TestDeployContractTransactor // Write-only binding to the contract
	TestDeployContractFilterer   // Log filterer for contract events
}

// TestDeployContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestDeployContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDeployContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestDeployContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDeployContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestDeployContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestDeployContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestDeployContractSession struct {
	Contract     *TestDeployContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestDeployContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestDeployContractCallerSession struct {
	Contract *TestDeployContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TestDeployContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestDeployContractTransactorSession struct {
	Contract     *TestDeployContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TestDeployContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestDeployContractRaw struct {
	Contract *TestDeployContract // Generic contract binding to access the raw methods on
}

// TestDeployContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestDeployContractCallerRaw struct {
	Contract *TestDeployContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestDeployContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestDeployContractTransactorRaw struct {
	Contract *TestDeployContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestDeployContract creates a new instance of TestDeployContract, bound to a specific deployed contract.
func NewTestDeployContract(address common.Address, backend bind.ContractBackend) (*TestDeployContract, error) {
	contract, err := bindTestDeployContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestDeployContract{TestDeployContractCaller: TestDeployContractCaller{contract: contract}, TestDeployContractTransactor: TestDeployContractTransactor{contract: contract}, TestDeployContractFilterer: TestDeployContractFilterer{contract: contract}}, nil
}

// NewTestDeployContractCaller creates a new read-only instance of TestDeployContract, bound to a specific deployed contract.
func NewTestDeployContractCaller(address common.Address, caller bind.ContractCaller) (*TestDeployContractCaller, error) {
	contract, err := bindTestDeployContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestDeployContractCaller{contract: contract}, nil
}

// NewTestDeployContractTransactor creates a new write-only instance of TestDeployContract, bound to a specific deployed contract.
func NewTestDeployContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestDeployContractTransactor, error) {
	contract, err := bindTestDeployContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestDeployContractTransactor{contract: contract}, nil
}

// NewTestDeployContractFilterer creates a new log filterer instance of TestDeployContract, bound to a specific deployed contract.
func NewTestDeployContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestDeployContractFilterer, error) {
	contract, err := bindTestDeployContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestDeployContractFilterer{contract: contract}, nil
}

// bindTestDeployContract binds a generic wrapper to an already deployed contract.
func bindTestDeployContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestDeployContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDeployContract *TestDeployContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDeployContract.Contract.TestDeployContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDeployContract *TestDeployContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDeployContract.Contract.TestDeployContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDeployContract *TestDeployContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDeployContract.Contract.TestDeployContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestDeployContract *TestDeployContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestDeployContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestDeployContract *TestDeployContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestDeployContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestDeployContract *TestDeployContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestDeployContract.Contract.contract.Transact(opts, method, params...)
}

// MyNum is a free data retrieval call binding the contract method 0xd4da005a.
//
// Solidity: function myNum() view returns(uint256)
func (_TestDeployContract *TestDeployContractCaller) MyNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestDeployContract.contract.Call(opts, &out, "myNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyNum is a free data retrieval call binding the contract method 0xd4da005a.
//
// Solidity: function myNum() view returns(uint256)
func (_TestDeployContract *TestDeployContractSession) MyNum() (*big.Int, error) {
	return _TestDeployContract.Contract.MyNum(&_TestDeployContract.CallOpts)
}

// MyNum is a free data retrieval call binding the contract method 0xd4da005a.
//
// Solidity: function myNum() view returns(uint256)
func (_TestDeployContract *TestDeployContractCallerSession) MyNum() (*big.Int, error) {
	return _TestDeployContract.Contract.MyNum(&_TestDeployContract.CallOpts)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _newNum) returns()
func (_TestDeployContract *TestDeployContractTransactor) SetNum(opts *bind.TransactOpts, _newNum *big.Int) (*types.Transaction, error) {
	return _TestDeployContract.contract.Transact(opts, "setNum", _newNum)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _newNum) returns()
func (_TestDeployContract *TestDeployContractSession) SetNum(_newNum *big.Int) (*types.Transaction, error) {
	return _TestDeployContract.Contract.SetNum(&_TestDeployContract.TransactOpts, _newNum)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _newNum) returns()
func (_TestDeployContract *TestDeployContractTransactorSession) SetNum(_newNum *big.Int) (*types.Transaction, error) {
	return _TestDeployContract.Contract.SetNum(&_TestDeployContract.TransactOpts, _newNum)
}
