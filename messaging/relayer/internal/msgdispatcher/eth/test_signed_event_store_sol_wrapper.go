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

// TestsignedeventstoreMetaData contains all meta data concerning the Testsignedeventstore contract.
var TestsignedeventstoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_functionCall\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_encodedEvent\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"decodeAndVerifyEvent\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_cbcAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_encodedEvent\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"relayEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610efc380380610efc83398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b610e3e806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631cd8253a1461003b5780634c1ce90214610050575b600080fd5b61004e610049366004610abe565b610063565b005b61004e61005e366004610b56565b6107e6565b60608060608060006100aa87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250925061093c915050565b90506100bd605563ffffffff8316610bab565b6100c8906004610bca565b861461011b5760405162461bcd60e51b815260206004820152601a60248201527f5369676e617475726520696e636f7272656374206c656e67746800000000000060448201526064015b60405180910390fd5b8063ffffffff1667ffffffffffffffff81111561013a5761013a610be2565b604051908082528060200260200182016040528015610163578160200160208202803683370190505b5094508063ffffffff1667ffffffffffffffff81111561018557610185610be2565b6040519080825280602002602001820160405280156101ae578160200160208202803683370190505b5093508063ffffffff1667ffffffffffffffff8111156101d0576101d0610be2565b6040519080825280602002602001820160405280156101f9578160200160208202803683370190505b5092508063ffffffff1667ffffffffffffffff81111561021b5761021b610be2565b604051908082528060200260200182016040528015610244578160200160208202803683370190505b509150600460005b8263ffffffff168110156104325761029b89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109a2915050565b8782815181106102ad576102ad610bf8565b6001600160a01b03909216602092830291909101909101526102d0601483610bca565b915061031389898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109aa915050565b86828151811061032557610325610bf8565b60200260200101818152505060208261033e9190610bca565b915061038189898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109aa915050565b85828151811061039357610393610bf8565b6020026020010181815250506020826103ac9190610bca565b91506103ef89898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610a0f915050565b84828151811061040157610401610bf8565b60ff9092166020928302919091019091015261041e600183610bca565b91508061042a81610c0e565b91505061024c565b505060008054906101000a90046001600160a01b03166001600160a01b031663ea13ec8b8c878787878f8f6040518863ffffffff1660e01b815260040161047f9796959493929190610cc0565b60206040518083038186803b15801561049757600080fd5b505afa1580156104ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cf9190610d6b565b50600089896040516104e2929190610d94565b6040518091039020905060005b8263ffffffff1681101561067b5760026000838152602001908152602001600020600301600088838151811061052757610527610bf8565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff161561059b5760405162461bcd60e51b815260206004820181905260248201527f52656c617965722068617320616c7265616479207369676e6564206576656e746044820152606401610112565b60016002600084815260200190815260200160002060030160008984815181106105c7576105c7610bf8565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055506002600083815260200190815260200160002060020187828151811061062e5761062e610bf8565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b039092169190911790558061067381610c0e565b9150506104ef565b5060008181526002602052604090206001015460ff16156106a1575050505050506107de565b6000805460405163a64ce19960e01b8152600481018f90526001600160a01b039091169063a64ce1999060240160206040518083038186803b1580156106e657600080fd5b505afa1580156106fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071e9190610da4565b6000838152600260208190526040909120015490915081116107d657600160009054906101000a90046001600160a01b03166001600160a01b031663408840528e8e8e8e8e8e6040518763ffffffff1660e01b815260040161078596959493929190610dbd565b600060405180830381600087803b15801561079f57600080fd5b505af11580156107b3573d6000803e3d6000fd5b50505060008381526002602052604090206001908101805460ff19169091179055505b505050505050505b505050505050565b600084846040516107f8929190610d94565b604080519182900390912060008181526002602052919091206001015490915060ff161561085b5760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481858dd1a5bdb995960821b6044820152606401610112565b6000805460405163a64ce19960e01b8152600481018a90526001600160a01b039091169063a64ce1999060240160206040518083038186803b1580156108a057600080fd5b505afa1580156108b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d89190610da4565b600083815260026020819052604090912001549091508111156109325760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768207369676e65727360701b6044820152606401610112565b5050505050505050565b6000610949826004610bca565b835110156109995760405162461bcd60e51b815260206004820152601d60248201527f736c6963696e67206f7574206f662072616e6765202875696e743332290000006044820152606401610112565b50016004015190565b016014015190565b60008060005b6020811015610a07576109c4816008610bab565b856109cf8387610bca565b815181106109df576109df610bf8565b01602001516001600160f81b031916901c9190911790806109ff81610c0e565b9150506109b0565b509392505050565b6000610a1c826001610bca565b83511015610a6c5760405162461bcd60e51b815260206004820152601c60248201527f736c6963696e67206f7574206f662072616e6765202875696e743829000000006044820152606401610112565b50016001015190565b60008083601f840112610a8757600080fd5b50813567ffffffffffffffff811115610a9f57600080fd5b602083019150836020828501011115610ab757600080fd5b9250929050565b60008060008060008060808789031215610ad757600080fd5b8635955060208701356001600160a01b0381168114610af557600080fd5b9450604087013567ffffffffffffffff80821115610b1257600080fd5b610b1e8a838b01610a75565b90965094506060890135915080821115610b3757600080fd5b50610b4489828a01610a75565b979a9699509497509295939492505050565b60008060008060008060808789031215610b6f57600080fd5b8635955060208701359450604087013567ffffffffffffffff80821115610b1257600080fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610bc557610bc5610b95565b500290565b60008219821115610bdd57610bdd610b95565b500190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000600019821415610c2257610c22610b95565b5060010190565b600081518084526020808501945080840160005b83811015610c5957815187529582019590820190600101610c3d565b509495945050505050565b600081518084526020808501945080840160005b83811015610c5957815160ff1687529582019590820190600101610c78565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600060c08201898352602060c081850152818a5180845260e086019150828c01935060005b81811015610d0a5784516001600160a01b031683529383019391830191600101610ce5565b50508481036040860152610d1e818b610c29565b925050508281036060840152610d348188610c29565b90508281036080840152610d488187610c64565b905082810360a0840152610d5d818587610c97565b9a9950505050505050505050565b600060208284031215610d7d57600080fd5b81518015158114610d8d57600080fd5b9392505050565b8183823760009101908152919050565b600060208284031215610db657600080fd5b5051919050565b8681526001600160a01b0386166020820152608060408201819052600090610de89083018688610c97565b8281036060840152610dfb818587610c97565b999850505050505050505056fea2646970667358221220b6a8094a02105bb5b9dbf1e4763132d740aa8e41f95a847fd4aea3eb42a2301a64736f6c63430008090033",
}

// TestsignedeventstoreABI is the input ABI used to generate the binding from.
// Deprecated: Use TestsignedeventstoreMetaData.ABI instead.
var TestsignedeventstoreABI = TestsignedeventstoreMetaData.ABI

// TestsignedeventstoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestsignedeventstoreMetaData.Bin instead.
var TestsignedeventstoreBin = TestsignedeventstoreMetaData.Bin

// DeployTestsignedeventstore deploys a new Ethereum contract, binding an instance of Testsignedeventstore to it.
func DeployTestsignedeventstore(auth *bind.TransactOpts, backend bind.ContractBackend, _registrar common.Address, _functionCall common.Address) (common.Address, *types.Transaction, *Testsignedeventstore, error) {
	parsed, err := TestsignedeventstoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestsignedeventstoreBin), backend, _registrar, _functionCall)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testsignedeventstore{TestsignedeventstoreCaller: TestsignedeventstoreCaller{contract: contract}, TestsignedeventstoreTransactor: TestsignedeventstoreTransactor{contract: contract}, TestsignedeventstoreFilterer: TestsignedeventstoreFilterer{contract: contract}}, nil
}

// Testsignedeventstore is an auto generated Go binding around an Ethereum contract.
type Testsignedeventstore struct {
	TestsignedeventstoreCaller     // Read-only binding to the contract
	TestsignedeventstoreTransactor // Write-only binding to the contract
	TestsignedeventstoreFilterer   // Log filterer for contract events
}

// TestsignedeventstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestsignedeventstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestsignedeventstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestsignedeventstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestsignedeventstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestsignedeventstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestsignedeventstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestsignedeventstoreSession struct {
	Contract     *Testsignedeventstore // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestsignedeventstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestsignedeventstoreCallerSession struct {
	Contract *TestsignedeventstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TestsignedeventstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestsignedeventstoreTransactorSession struct {
	Contract     *TestsignedeventstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TestsignedeventstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestsignedeventstoreRaw struct {
	Contract *Testsignedeventstore // Generic contract binding to access the raw methods on
}

// TestsignedeventstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestsignedeventstoreCallerRaw struct {
	Contract *TestsignedeventstoreCaller // Generic read-only contract binding to access the raw methods on
}

// TestsignedeventstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestsignedeventstoreTransactorRaw struct {
	Contract *TestsignedeventstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestsignedeventstore creates a new instance of Testsignedeventstore, bound to a specific deployed contract.
func NewTestsignedeventstore(address common.Address, backend bind.ContractBackend) (*Testsignedeventstore, error) {
	contract, err := bindTestsignedeventstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testsignedeventstore{TestsignedeventstoreCaller: TestsignedeventstoreCaller{contract: contract}, TestsignedeventstoreTransactor: TestsignedeventstoreTransactor{contract: contract}, TestsignedeventstoreFilterer: TestsignedeventstoreFilterer{contract: contract}}, nil
}

// NewTestsignedeventstoreCaller creates a new read-only instance of Testsignedeventstore, bound to a specific deployed contract.
func NewTestsignedeventstoreCaller(address common.Address, caller bind.ContractCaller) (*TestsignedeventstoreCaller, error) {
	contract, err := bindTestsignedeventstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestsignedeventstoreCaller{contract: contract}, nil
}

// NewTestsignedeventstoreTransactor creates a new write-only instance of Testsignedeventstore, bound to a specific deployed contract.
func NewTestsignedeventstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*TestsignedeventstoreTransactor, error) {
	contract, err := bindTestsignedeventstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestsignedeventstoreTransactor{contract: contract}, nil
}

// NewTestsignedeventstoreFilterer creates a new log filterer instance of Testsignedeventstore, bound to a specific deployed contract.
func NewTestsignedeventstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*TestsignedeventstoreFilterer, error) {
	contract, err := bindTestsignedeventstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestsignedeventstoreFilterer{contract: contract}, nil
}

// bindTestsignedeventstore binds a generic wrapper to an already deployed contract.
func bindTestsignedeventstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestsignedeventstoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testsignedeventstore *TestsignedeventstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testsignedeventstore.Contract.TestsignedeventstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testsignedeventstore *TestsignedeventstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testsignedeventstore.Contract.TestsignedeventstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testsignedeventstore *TestsignedeventstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testsignedeventstore.Contract.TestsignedeventstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testsignedeventstore *TestsignedeventstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testsignedeventstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testsignedeventstore *TestsignedeventstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testsignedeventstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testsignedeventstore *TestsignedeventstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testsignedeventstore.Contract.contract.Transact(opts, method, params...)
}

// DecodeAndVerifyEvent is a free data retrieval call binding the contract method 0x4c1ce902.
//
// Solidity: function decodeAndVerifyEvent(uint256 _blockchainId, bytes32 , bytes _encodedEvent, bytes ) view returns()
func (_Testsignedeventstore *TestsignedeventstoreCaller) DecodeAndVerifyEvent(opts *bind.CallOpts, _blockchainId *big.Int, arg1 [32]byte, _encodedEvent []byte, arg3 []byte) error {
	var out []interface{}
	err := _Testsignedeventstore.contract.Call(opts, &out, "decodeAndVerifyEvent", _blockchainId, arg1, _encodedEvent, arg3)

	if err != nil {
		return err
	}

	return err

}

// DecodeAndVerifyEvent is a free data retrieval call binding the contract method 0x4c1ce902.
//
// Solidity: function decodeAndVerifyEvent(uint256 _blockchainId, bytes32 , bytes _encodedEvent, bytes ) view returns()
func (_Testsignedeventstore *TestsignedeventstoreSession) DecodeAndVerifyEvent(_blockchainId *big.Int, arg1 [32]byte, _encodedEvent []byte, arg3 []byte) error {
	return _Testsignedeventstore.Contract.DecodeAndVerifyEvent(&_Testsignedeventstore.CallOpts, _blockchainId, arg1, _encodedEvent, arg3)
}

// DecodeAndVerifyEvent is a free data retrieval call binding the contract method 0x4c1ce902.
//
// Solidity: function decodeAndVerifyEvent(uint256 _blockchainId, bytes32 , bytes _encodedEvent, bytes ) view returns()
func (_Testsignedeventstore *TestsignedeventstoreCallerSession) DecodeAndVerifyEvent(_blockchainId *big.Int, arg1 [32]byte, _encodedEvent []byte, arg3 []byte) error {
	return _Testsignedeventstore.Contract.DecodeAndVerifyEvent(&_Testsignedeventstore.CallOpts, _blockchainId, arg1, _encodedEvent, arg3)
}

// RelayEvent is a paid mutator transaction binding the contract method 0x1cd8253a.
//
// Solidity: function relayEvent(uint256 _blockchainId, address _cbcAddress, bytes _encodedEvent, bytes _signature) returns()
func (_Testsignedeventstore *TestsignedeventstoreTransactor) RelayEvent(opts *bind.TransactOpts, _blockchainId *big.Int, _cbcAddress common.Address, _encodedEvent []byte, _signature []byte) (*types.Transaction, error) {
	return _Testsignedeventstore.contract.Transact(opts, "relayEvent", _blockchainId, _cbcAddress, _encodedEvent, _signature)
}

// RelayEvent is a paid mutator transaction binding the contract method 0x1cd8253a.
//
// Solidity: function relayEvent(uint256 _blockchainId, address _cbcAddress, bytes _encodedEvent, bytes _signature) returns()
func (_Testsignedeventstore *TestsignedeventstoreSession) RelayEvent(_blockchainId *big.Int, _cbcAddress common.Address, _encodedEvent []byte, _signature []byte) (*types.Transaction, error) {
	return _Testsignedeventstore.Contract.RelayEvent(&_Testsignedeventstore.TransactOpts, _blockchainId, _cbcAddress, _encodedEvent, _signature)
}

// RelayEvent is a paid mutator transaction binding the contract method 0x1cd8253a.
//
// Solidity: function relayEvent(uint256 _blockchainId, address _cbcAddress, bytes _encodedEvent, bytes _signature) returns()
func (_Testsignedeventstore *TestsignedeventstoreTransactorSession) RelayEvent(_blockchainId *big.Int, _cbcAddress common.Address, _encodedEvent []byte, _signature []byte) (*types.Transaction, error) {
	return _Testsignedeventstore.Contract.RelayEvent(&_Testsignedeventstore.TransactOpts, _blockchainId, _cbcAddress, _encodedEvent, _signature)
}
