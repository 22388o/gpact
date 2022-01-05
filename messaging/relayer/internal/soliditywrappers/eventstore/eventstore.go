// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eventstore

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

// EventStoreMetaData contains all meta data concerning the EventStore contract.
var EventStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_functionCall\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_log\",\"type\":\"string\"}],\"name\":\"Dump\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_encodedEvent\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"decodeAndVerifyEvent\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sourceBlockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceCbcAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_eventData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"relayEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610fc3380380610fc383398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b610f05806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631cd8253a1461003b5780634c1ce90214610050575b600080fd5b61004e610049366004610b0d565b610063565b005b61004e61005e366004610ba5565b61088e565b60608060608060006100aa87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250925061098b915050565b90506100bd605563ffffffff8316610bfa565b6100c8906004610c19565b861461011b5760405162461bcd60e51b815260206004820152601a60248201527f5369676e617475726520696e636f7272656374206c656e67746800000000000060448201526064015b60405180910390fd5b8063ffffffff1667ffffffffffffffff81111561013a5761013a610c31565b604051908082528060200260200182016040528015610163578160200160208202803683370190505b5094508063ffffffff1667ffffffffffffffff81111561018557610185610c31565b6040519080825280602002602001820160405280156101ae578160200160208202803683370190505b5093508063ffffffff1667ffffffffffffffff8111156101d0576101d0610c31565b6040519080825280602002602001820160405280156101f9578160200160208202803683370190505b5092508063ffffffff1667ffffffffffffffff81111561021b5761021b610c31565b604051908082528060200260200182016040528015610244578160200160208202803683370190505b509150600460005b8263ffffffff168110156104325761029b89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109f1915050565b8782815181106102ad576102ad610c47565b6001600160a01b03909216602092830291909101909101526102d0601483610c19565b915061031389898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109f9915050565b86828151811061032557610325610c47565b60200260200101818152505060208261033e9190610c19565b915061038189898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109f9915050565b85828151811061039357610393610c47565b6020026020010181815250506020826103ac9190610c19565b91506103ef89898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610a5e915050565b84828151811061040157610401610c47565b60ff9092166020928302919091019091015261041e600183610c19565b91508061042a81610c5d565b91505061024c565b505060008b8b7f59f736dc5e15c4b12526487502645403b0a4316d82eba7e9ecdc2a050c10ad278c8c60405160200161046f959493929190610c78565b60408051601f198184030181529082905260005463ea13ec8b60e01b83529092506001600160a01b03169063ea13ec8b906104b8908f908a908a908a908a908990600401610d6d565b60206040518083038186803b1580156104d057600080fd5b505afa1580156104e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105089190610e16565b507f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9604051610556906020808252600b908201526a76657269667920646f6e6560a81b604082015260600190565b60405180910390a18051602082012060005b8363ffffffff168110156106b2576002600083815260200190815260200160002060030160008983815181106105a0576105a0610c47565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156105d1576106a0565b60016002600084815260200190815260200160002060030160008a84815181106105fd576105fd610c47565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055506002600083815260200190815260200160002060020188828151811061066457610664610c47565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b039092169190911790555b806106aa81610c5d565b915050610568565b5060008181526002602052604090206001015460ff16156106d95750505050505050610886565b60008060009054906101000a90046001600160a01b03166001600160a01b031663a64ce1998f6040518263ffffffff1660e01b815260040161071d91815260200190565b60206040518083038186803b15801561073557600080fd5b505afa158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076d9190610e3f565b60008381526002602081905260409091200154909150811161087d57600082815260026020526040908190206001908101805460ff19169091179055517f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9906107f7906020808252600d908201526c199d5b98dd1a5bdb8818d85b1b609a1b604082015260600190565b60405180910390a1600160009054906101000a90046001600160a01b03166001600160a01b031663408840528f8f8f8f8f8f6040518763ffffffff1660e01b815260040161084a96959493929190610e81565b600060405180830381600087803b15801561086457600080fd5b505af1158015610878573d6000803e3d6000fd5b505050505b50505050505050505b505050505050565b600084846040516108a0929190610ebf565b6040519081900381206000805463a64ce19960e01b8452600484018b9052919350916001600160a01b039091169063a64ce1999060240160206040518083038186803b1580156108ef57600080fd5b505afa158015610903573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109279190610e3f565b600083815260026020819052604090912001549091508111156109815760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768207369676e65727360701b6044820152606401610112565b5050505050505050565b6000610998826004610c19565b835110156109e85760405162461bcd60e51b815260206004820152601d60248201527f736c6963696e67206f7574206f662072616e6765202875696e743332290000006044820152606401610112565b50016004015190565b016014015190565b60008060005b6020811015610a5657610a13816008610bfa565b85610a1e8387610c19565b81518110610a2e57610a2e610c47565b01602001516001600160f81b031916901c919091179080610a4e81610c5d565b9150506109ff565b509392505050565b6000610a6b826001610c19565b83511015610abb5760405162461bcd60e51b815260206004820152601c60248201527f736c6963696e67206f7574206f662072616e6765202875696e743829000000006044820152606401610112565b50016001015190565b60008083601f840112610ad657600080fd5b50813567ffffffffffffffff811115610aee57600080fd5b602083019150836020828501011115610b0657600080fd5b9250929050565b60008060008060008060808789031215610b2657600080fd5b8635955060208701356001600160a01b0381168114610b4457600080fd5b9450604087013567ffffffffffffffff80821115610b6157600080fd5b610b6d8a838b01610ac4565b90965094506060890135915080821115610b8657600080fd5b50610b9389828a01610ac4565b979a9699509497509295939492505050565b60008060008060008060808789031215610bbe57600080fd5b8635955060208701359450604087013567ffffffffffffffff80821115610b6157600080fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610c1457610c14610be4565b500290565b60008219821115610c2c57610c2c610be4565b500190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000600019821415610c7157610c71610be4565b5060010190565b8581526bffffffffffffffffffffffff198560601b1660208201528360348201528183605483013760009101605401908152949350505050565b600081518084526020808501945080840160005b83811015610ce257815187529582019590820190600101610cc6565b509495945050505050565b600081518084526020808501945080840160005b83811015610ce257815160ff1687529582019590820190600101610d01565b6000815180845260005b81811015610d4657602081850181015186830182015201610d2a565b81811115610d58576000602083870101525b50601f01601f19169290920160200192915050565b600060c08201888352602060c08185015281895180845260e086019150828b01935060005b81811015610db75784516001600160a01b031683529383019391830191600101610d92565b50508481036040860152610dcb818a610cb2565b925050508281036060840152610de18187610cb2565b90508281036080840152610df58186610ced565b905082810360a0840152610e098185610d20565b9998505050505050505050565b600060208284031215610e2857600080fd5b81518015158114610e3857600080fd5b9392505050565b600060208284031215610e5157600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8681526001600160a01b0386166020820152608060408201819052600090610eac9083018688610e58565b8281036060840152610e09818587610e58565b818382376000910190815291905056fea264697066735822122016ead7e87ff39a26b896148315b0783ef9f38920a9099024fbc5d301089264d964736f6c63430008090033",
}

// EventStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use EventStoreMetaData.ABI instead.
var EventStoreABI = EventStoreMetaData.ABI

// EventStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventStoreMetaData.Bin instead.
var EventStoreBin = EventStoreMetaData.Bin

// DeployEventStore deploys a new Ethereum contract, binding an instance of EventStore to it.
func DeployEventStore(auth *bind.TransactOpts, backend bind.ContractBackend, _registrar common.Address, _functionCall common.Address) (common.Address, *types.Transaction, *EventStore, error) {
	parsed, err := EventStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventStoreBin), backend, _registrar, _functionCall)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EventStore{EventStoreCaller: EventStoreCaller{contract: contract}, EventStoreTransactor: EventStoreTransactor{contract: contract}, EventStoreFilterer: EventStoreFilterer{contract: contract}}, nil
}

// EventStore is an auto generated Go binding around an Ethereum contract.
type EventStore struct {
	EventStoreCaller     // Read-only binding to the contract
	EventStoreTransactor // Write-only binding to the contract
	EventStoreFilterer   // Log filterer for contract events
}

// EventStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventStoreSession struct {
	Contract     *EventStore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventStoreCallerSession struct {
	Contract *EventStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EventStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventStoreTransactorSession struct {
	Contract     *EventStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EventStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventStoreRaw struct {
	Contract *EventStore // Generic contract binding to access the raw methods on
}

// EventStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventStoreCallerRaw struct {
	Contract *EventStoreCaller // Generic read-only contract binding to access the raw methods on
}

// EventStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventStoreTransactorRaw struct {
	Contract *EventStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventStore creates a new instance of EventStore, bound to a specific deployed contract.
func NewEventStore(address common.Address, backend bind.ContractBackend) (*EventStore, error) {
	contract, err := bindEventStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventStore{EventStoreCaller: EventStoreCaller{contract: contract}, EventStoreTransactor: EventStoreTransactor{contract: contract}, EventStoreFilterer: EventStoreFilterer{contract: contract}}, nil
}

// NewEventStoreCaller creates a new read-only instance of EventStore, bound to a specific deployed contract.
func NewEventStoreCaller(address common.Address, caller bind.ContractCaller) (*EventStoreCaller, error) {
	contract, err := bindEventStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventStoreCaller{contract: contract}, nil
}

// NewEventStoreTransactor creates a new write-only instance of EventStore, bound to a specific deployed contract.
func NewEventStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*EventStoreTransactor, error) {
	contract, err := bindEventStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventStoreTransactor{contract: contract}, nil
}

// NewEventStoreFilterer creates a new log filterer instance of EventStore, bound to a specific deployed contract.
func NewEventStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*EventStoreFilterer, error) {
	contract, err := bindEventStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventStoreFilterer{contract: contract}, nil
}

// bindEventStore binds a generic wrapper to an already deployed contract.
func bindEventStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventStore *EventStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventStore.Contract.EventStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventStore *EventStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventStore.Contract.EventStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventStore *EventStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventStore.Contract.EventStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventStore *EventStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventStore *EventStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventStore *EventStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventStore.Contract.contract.Transact(opts, method, params...)
}

// DecodeAndVerifyEvent is a free data retrieval call binding the contract method 0x4c1ce902.
//
// Solidity: function decodeAndVerifyEvent(uint256 _blockchainId, bytes32 , bytes _encodedEvent, bytes ) view returns()
func (_EventStore *EventStoreCaller) DecodeAndVerifyEvent(opts *bind.CallOpts, _blockchainId *big.Int, arg1 [32]byte, _encodedEvent []byte, arg3 []byte) error {
	var out []interface{}
	err := _EventStore.contract.Call(opts, &out, "decodeAndVerifyEvent", _blockchainId, arg1, _encodedEvent, arg3)

	if err != nil {
		return err
	}

	return err

}

// DecodeAndVerifyEvent is a free data retrieval call binding the contract method 0x4c1ce902.
//
// Solidity: function decodeAndVerifyEvent(uint256 _blockchainId, bytes32 , bytes _encodedEvent, bytes ) view returns()
func (_EventStore *EventStoreSession) DecodeAndVerifyEvent(_blockchainId *big.Int, arg1 [32]byte, _encodedEvent []byte, arg3 []byte) error {
	return _EventStore.Contract.DecodeAndVerifyEvent(&_EventStore.CallOpts, _blockchainId, arg1, _encodedEvent, arg3)
}

// DecodeAndVerifyEvent is a free data retrieval call binding the contract method 0x4c1ce902.
//
// Solidity: function decodeAndVerifyEvent(uint256 _blockchainId, bytes32 , bytes _encodedEvent, bytes ) view returns()
func (_EventStore *EventStoreCallerSession) DecodeAndVerifyEvent(_blockchainId *big.Int, arg1 [32]byte, _encodedEvent []byte, arg3 []byte) error {
	return _EventStore.Contract.DecodeAndVerifyEvent(&_EventStore.CallOpts, _blockchainId, arg1, _encodedEvent, arg3)
}

// RelayEvent is a paid mutator transaction binding the contract method 0x1cd8253a.
//
// Solidity: function relayEvent(uint256 _sourceBlockchainId, address _sourceCbcAddress, bytes _eventData, bytes _signature) returns()
func (_EventStore *EventStoreTransactor) RelayEvent(opts *bind.TransactOpts, _sourceBlockchainId *big.Int, _sourceCbcAddress common.Address, _eventData []byte, _signature []byte) (*types.Transaction, error) {
	return _EventStore.contract.Transact(opts, "relayEvent", _sourceBlockchainId, _sourceCbcAddress, _eventData, _signature)
}

// RelayEvent is a paid mutator transaction binding the contract method 0x1cd8253a.
//
// Solidity: function relayEvent(uint256 _sourceBlockchainId, address _sourceCbcAddress, bytes _eventData, bytes _signature) returns()
func (_EventStore *EventStoreSession) RelayEvent(_sourceBlockchainId *big.Int, _sourceCbcAddress common.Address, _eventData []byte, _signature []byte) (*types.Transaction, error) {
	return _EventStore.Contract.RelayEvent(&_EventStore.TransactOpts, _sourceBlockchainId, _sourceCbcAddress, _eventData, _signature)
}

// RelayEvent is a paid mutator transaction binding the contract method 0x1cd8253a.
//
// Solidity: function relayEvent(uint256 _sourceBlockchainId, address _sourceCbcAddress, bytes _eventData, bytes _signature) returns()
func (_EventStore *EventStoreTransactorSession) RelayEvent(_sourceBlockchainId *big.Int, _sourceCbcAddress common.Address, _eventData []byte, _signature []byte) (*types.Transaction, error) {
	return _EventStore.Contract.RelayEvent(&_EventStore.TransactOpts, _sourceBlockchainId, _sourceCbcAddress, _eventData, _signature)
}

// EventStoreDumpIterator is returned from FilterDump and is used to iterate over the raw logs and unpacked data for Dump events raised by the EventStore contract.
type EventStoreDumpIterator struct {
	Event *EventStoreDump // Event containing the contract specifics and raw log

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
func (it *EventStoreDumpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventStoreDump)
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
		it.Event = new(EventStoreDump)
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
func (it *EventStoreDumpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventStoreDumpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventStoreDump represents a Dump event raised by the EventStore contract.
type EventStoreDump struct {
	Log string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDump is a free log retrieval operation binding the contract event 0x922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9.
//
// Solidity: event Dump(string _log)
func (_EventStore *EventStoreFilterer) FilterDump(opts *bind.FilterOpts) (*EventStoreDumpIterator, error) {

	logs, sub, err := _EventStore.contract.FilterLogs(opts, "Dump")
	if err != nil {
		return nil, err
	}
	return &EventStoreDumpIterator{contract: _EventStore.contract, event: "Dump", logs: logs, sub: sub}, nil
}

// WatchDump is a free log subscription operation binding the contract event 0x922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9.
//
// Solidity: event Dump(string _log)
func (_EventStore *EventStoreFilterer) WatchDump(opts *bind.WatchOpts, sink chan<- *EventStoreDump) (event.Subscription, error) {

	logs, sub, err := _EventStore.contract.WatchLogs(opts, "Dump")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventStoreDump)
				if err := _EventStore.contract.UnpackLog(event, "Dump", log); err != nil {
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

// ParseDump is a log parse operation binding the contract event 0x922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9.
//
// Solidity: event Dump(string _log)
func (_EventStore *EventStoreFilterer) ParseDump(log types.Log) (*EventStoreDump, error) {
	event := new(EventStoreDump)
	if err := _EventStore.contract.UnpackLog(event, "Dump", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
