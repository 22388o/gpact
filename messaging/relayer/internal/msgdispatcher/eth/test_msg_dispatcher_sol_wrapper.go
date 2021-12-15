// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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
)

// TestmsgdispatcherMetaData contains all meta data concerning the Testmsgdispatcher contract.
var TestmsgdispatcherMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"Dump\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"aFunc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"val\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060eb8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80633c6bb43614603757806368ec78f3146051575b600080fd5b603f60005481565b60405190815260200160405180910390f35b6060605c366004609d565b6062565b005b60008190556040518181527f9ebbca46e60e984aa522942ae547e34cddc84dec1c85971a1e80b27a23231b8a9060200160405180910390a150565b60006020828403121560ae57600080fd5b503591905056fea26469706673582212207c3e08b8110865c5a0f57ca6f8649dbd48a0249009eddb9aab572145cd45a79c64736f6c63430008090033",
}

// TestmsgdispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use TestmsgdispatcherMetaData.ABI instead.
var TestmsgdispatcherABI = TestmsgdispatcherMetaData.ABI

// TestmsgdispatcherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestmsgdispatcherMetaData.Bin instead.
var TestmsgdispatcherBin = TestmsgdispatcherMetaData.Bin

// DeployTestmsgdispatcher deploys a new Ethereum contract, binding an instance of Testmsgdispatcher to it.
func DeployTestmsgdispatcher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Testmsgdispatcher, error) {
	parsed, err := TestmsgdispatcherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestmsgdispatcherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testmsgdispatcher{TestmsgdispatcherCaller: TestmsgdispatcherCaller{contract: contract}, TestmsgdispatcherTransactor: TestmsgdispatcherTransactor{contract: contract}, TestmsgdispatcherFilterer: TestmsgdispatcherFilterer{contract: contract}}, nil
}

// Testmsgdispatcher is an auto generated Go binding around an Ethereum contract.
type Testmsgdispatcher struct {
	TestmsgdispatcherCaller     // Read-only binding to the contract
	TestmsgdispatcherTransactor // Write-only binding to the contract
	TestmsgdispatcherFilterer   // Log filterer for contract events
}

// TestmsgdispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestmsgdispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestmsgdispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestmsgdispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestmsgdispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestmsgdispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestmsgdispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestmsgdispatcherSession struct {
	Contract     *Testmsgdispatcher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestmsgdispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestmsgdispatcherCallerSession struct {
	Contract *TestmsgdispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TestmsgdispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestmsgdispatcherTransactorSession struct {
	Contract     *TestmsgdispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TestmsgdispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestmsgdispatcherRaw struct {
	Contract *Testmsgdispatcher // Generic contract binding to access the raw methods on
}

// TestmsgdispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestmsgdispatcherCallerRaw struct {
	Contract *TestmsgdispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// TestmsgdispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestmsgdispatcherTransactorRaw struct {
	Contract *TestmsgdispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestmsgdispatcher creates a new instance of Testmsgdispatcher, bound to a specific deployed contract.
func NewTestmsgdispatcher(address common.Address, backend bind.ContractBackend) (*Testmsgdispatcher, error) {
	contract, err := bindTestmsgdispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testmsgdispatcher{TestmsgdispatcherCaller: TestmsgdispatcherCaller{contract: contract}, TestmsgdispatcherTransactor: TestmsgdispatcherTransactor{contract: contract}, TestmsgdispatcherFilterer: TestmsgdispatcherFilterer{contract: contract}}, nil
}

// NewTestmsgdispatcherCaller creates a new read-only instance of Testmsgdispatcher, bound to a specific deployed contract.
func NewTestmsgdispatcherCaller(address common.Address, caller bind.ContractCaller) (*TestmsgdispatcherCaller, error) {
	contract, err := bindTestmsgdispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestmsgdispatcherCaller{contract: contract}, nil
}

// NewTestmsgdispatcherTransactor creates a new write-only instance of Testmsgdispatcher, bound to a specific deployed contract.
func NewTestmsgdispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*TestmsgdispatcherTransactor, error) {
	contract, err := bindTestmsgdispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestmsgdispatcherTransactor{contract: contract}, nil
}

// NewTestmsgdispatcherFilterer creates a new log filterer instance of Testmsgdispatcher, bound to a specific deployed contract.
func NewTestmsgdispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*TestmsgdispatcherFilterer, error) {
	contract, err := bindTestmsgdispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestmsgdispatcherFilterer{contract: contract}, nil
}

// bindTestmsgdispatcher binds a generic wrapper to an already deployed contract.
func bindTestmsgdispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestmsgdispatcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testmsgdispatcher *TestmsgdispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testmsgdispatcher.Contract.TestmsgdispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testmsgdispatcher *TestmsgdispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testmsgdispatcher.Contract.TestmsgdispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testmsgdispatcher *TestmsgdispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testmsgdispatcher.Contract.TestmsgdispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testmsgdispatcher *TestmsgdispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testmsgdispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testmsgdispatcher *TestmsgdispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testmsgdispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testmsgdispatcher *TestmsgdispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testmsgdispatcher.Contract.contract.Transact(opts, method, params...)
}

// Val is a free data retrieval call binding the contract method 0x3c6bb436.
//
// Solidity: function val() view returns(uint256)
func (_Testmsgdispatcher *TestmsgdispatcherCaller) Val(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Testmsgdispatcher.contract.Call(opts, &out, "val")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Val is a free data retrieval call binding the contract method 0x3c6bb436.
//
// Solidity: function val() view returns(uint256)
func (_Testmsgdispatcher *TestmsgdispatcherSession) Val() (*big.Int, error) {
	return _Testmsgdispatcher.Contract.Val(&_Testmsgdispatcher.CallOpts)
}

// Val is a free data retrieval call binding the contract method 0x3c6bb436.
//
// Solidity: function val() view returns(uint256)
func (_Testmsgdispatcher *TestmsgdispatcherCallerSession) Val() (*big.Int, error) {
	return _Testmsgdispatcher.Contract.Val(&_Testmsgdispatcher.CallOpts)
}

// AFunc is a paid mutator transaction binding the contract method 0x68ec78f3.
//
// Solidity: function aFunc(uint256 _val) returns()
func (_Testmsgdispatcher *TestmsgdispatcherTransactor) AFunc(opts *bind.TransactOpts, _val *big.Int) (*types.Transaction, error) {
	return _Testmsgdispatcher.contract.Transact(opts, "aFunc", _val)
}

// AFunc is a paid mutator transaction binding the contract method 0x68ec78f3.
//
// Solidity: function aFunc(uint256 _val) returns()
func (_Testmsgdispatcher *TestmsgdispatcherSession) AFunc(_val *big.Int) (*types.Transaction, error) {
	return _Testmsgdispatcher.Contract.AFunc(&_Testmsgdispatcher.TransactOpts, _val)
}

// AFunc is a paid mutator transaction binding the contract method 0x68ec78f3.
//
// Solidity: function aFunc(uint256 _val) returns()
func (_Testmsgdispatcher *TestmsgdispatcherTransactorSession) AFunc(_val *big.Int) (*types.Transaction, error) {
	return _Testmsgdispatcher.Contract.AFunc(&_Testmsgdispatcher.TransactOpts, _val)
}

// TestmsgdispatcherDumpIterator is returned from FilterDump and is used to iterate over the raw logs and unpacked data for Dump events raised by the Testmsgdispatcher contract.
type TestmsgdispatcherDumpIterator struct {
	Event *TestmsgdispatcherDump // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestmsgdispatcherDumpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestmsgdispatcherDump)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestmsgdispatcherDump)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestmsgdispatcherDumpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestmsgdispatcherDumpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestmsgdispatcherDump represents a Dump event raised by the Testmsgdispatcher contract.
type TestmsgdispatcherDump struct {
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDump is a free log retrieval operation binding the contract event 0x9ebbca46e60e984aa522942ae547e34cddc84dec1c85971a1e80b27a23231b8a.
//
// Solidity: event Dump(uint256 _val)
func (_Testmsgdispatcher *TestmsgdispatcherFilterer) FilterDump(opts *bind.FilterOpts) (*TestmsgdispatcherDumpIterator, error) {

	logs, sub, err := _Testmsgdispatcher.contract.FilterLogs(opts, "Dump")
	if err != nil {
		return nil, err
	}
	return &TestmsgdispatcherDumpIterator{contract: _Testmsgdispatcher.contract, event: "Dump", logs: logs, sub: sub}, nil
}

// WatchDump is a free log subscription operation binding the contract event 0x9ebbca46e60e984aa522942ae547e34cddc84dec1c85971a1e80b27a23231b8a.
//
// Solidity: event Dump(uint256 _val)
func (_Testmsgdispatcher *TestmsgdispatcherFilterer) WatchDump(opts *bind.WatchOpts, sink chan<- *TestmsgdispatcherDump) (event.Subscription, error) {

	logs, sub, err := _Testmsgdispatcher.contract.WatchLogs(opts, "Dump")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestmsgdispatcherDump)
				if err := _Testmsgdispatcher.contract.UnpackLog(event, "Dump", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDump is a log parse operation binding the contract event 0x9ebbca46e60e984aa522942ae547e34cddc84dec1c85971a1e80b27a23231b8a.
//
// Solidity: event Dump(uint256 _val)
func (_Testmsgdispatcher *TestmsgdispatcherFilterer) ParseDump(log types.Log) (*TestmsgdispatcherDump, error) {
	event := new(TestmsgdispatcherDump)
	if err := _Testmsgdispatcher.contract.UnpackLog(event, "Dump", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
