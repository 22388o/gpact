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

// TestRegistrarMetaData contains all meta data concerning the TestRegistrar contract.
var TestRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"addSignerSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"addSignersSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"}],\"name\":\"getSigningThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mightBeSigner\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"}],\"name\":\"numSigners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"removeSignerSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_plainText\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_plainText\",\"type\":\"bytes\"}],\"name\":\"verifyAndCheckThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350611177806100616000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c80638da5cb5b11610097578063d4c0d34d11610066578063d4c0d34d14610223578063ea13ec8b14610236578063f2fde38b14610249578063f5e232ea1461025c57600080fd5b80638da5cb5b146101b4578063a64ce199146101cf578063ad107bb4146101fd578063b9c362091461021057600080fd5b806348bcbd2d116100d357806348bcbd2d14610134578063715018a6146101865780638bd6ac821461018e5780638d7678fd146101a157600080fd5b8062ab2414146100f957806315a098251461010e5780633156d37c14610121575b600080fd5b61010c610107366004610d78565b610280565b005b61010c61011c366004610dad565b610301565b61010c61012f366004610dad565b6103d4565b610171610142366004610dad565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff1692915050565b60405190151581526020015b60405180910390f35b61010c6104bf565b61017161019c366004610e25565b610533565b61010c6101af366004610e71565b610593565b6000546040516001600160a01b03909116815260200161017d565b6101ef6101dd366004610ec4565b60009081526001602052604090205490565b60405190815260200161017d565b61017161020b366004610f1f565b610651565b61010c61021e36600461101f565b610682565b61010c610231366004610d78565b6106d1565b610171610244366004610f1f565b610722565b61010c610257366004611041565b6109ae565b6101ef61026a366004610ec4565b6000908152600160208190526040909120015490565b6000546001600160a01b031633146102b35760405162461bcd60e51b81526004016102aa90611063565b60405180910390fd5b6102bd8383610a98565b600083815260016020819052604082208101546102d9916110ae565b60008581526001602081905260409091200181905590506102fb848284610b3a565b50505050565b6000546001600160a01b0316331461032b5760405162461bcd60e51b81526004016102aa90611063565b6000828152600160205260409020546103a45760405162461bcd60e51b815260206004820152603560248201527f43616e206e6f7420616464207369676e657220666f7220626c6f636b636861696044820152741b881dda5d1a081e995c9bc81d1a1c995cda1bdb19605a1b60648201526084016102aa565b6103ae8282610a98565b60008281526001602081905260408220018054916103cb836110c6565b91905055505050565b6000546001600160a01b031633146103fe5760405162461bcd60e51b81526004016102aa90611063565b6104088282610bfb565b6000828152600160208190526040822081015461042591906110e1565b6000848152600160205260409020549091508110156104a45760405162461bcd60e51b815260206004820152603560248201527f50726f706f736564206e6577206e756d626572206f66207369676e65727320696044820152741cc81b195cdcc81d1a185b881d1a1c995cda1bdb19605a1b60648201526084016102aa565b60009283526001602081905260409093209092019190915550565b6000546001600160a01b031633146104e95760405162461bcd60e51b81526004016102aa90611063565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008381526001602052604081205482908110156105885760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768207369676e65727360701b60448201526064016102aa565b506001949350505050565b6000546001600160a01b031633146105bd5760405162461bcd60e51b81526004016102aa90611063565b60005b8281101561060a576105f8858585848181106105de576105de6110f8565b90506020020160208101906105f39190611041565b610a98565b80610602816110c6565b9150506105c0565b506000848152600160208190526040822001546106289084906110ae565b600086815260016020819052604090912001819055905061064a858284610b3a565b5050505050565b600061065e8c8c8c610533565b506106728c8c8c8c8c8c8c8c8c8c8c610722565b9c9b505050505050505050505050565b6000546001600160a01b031633146106ac5760405162461bcd60e51b81526004016102aa90611063565b6106cd82600160008581526020019081526020016000206001015483610b3a565b5050565b6000546001600160a01b031633146106fb5760405162461bcd60e51b81526004016102aa90611063565b6107058383610bfb565b600083815260016020819052604082208101546102d991906110e1565b60008988811461076b5760405162461bcd60e51b81526020600482015260146024820152730e6d2cea440d8cadccee8d040dad2e6dac2e8c6d60631b60448201526064016102aa565b8087146107b15760405162461bcd60e51b81526020600482015260146024820152730e6d2cea640d8cadccee8d040dad2e6dac2e8c6d60631b60448201526064016102aa565b8085146107f75760405162461bcd60e51b81526020600482015260146024820152730e6d2ceac40d8cadccee8d040dad2e6dac2e8c6d60631b60448201526064016102aa565b60005b8181101561099a5760008e8152600160205260408120600201908e8e84818110610826576108266110f8565b905060200201602081019061083b9190611041565b6001600160a01b0316815260208101919091526040016000205460ff166108b25760405162461bcd60e51b815260206004820152602560248201527f5369676e6572206e6f74207369676e657220666f72207468697320626c6f636b60448201526431b430b4b760d91b60648201526084016102aa565b61093c8d8d838181106108c7576108c76110f8565b90506020020160208101906108dc9190611041565b86868e8e868181106108f0576108f06110f8565b905060200201358d8d87818110610909576109096110f8565b905060200201358c8c88818110610922576109226110f8565b9050602002016020810190610937919061110e565b610c99565b6109885760405162461bcd60e51b815260206004820152601860248201527f5369676e617475726520646964206e6f7420766572696679000000000000000060448201526064016102aa565b80610992816110c6565b9150506107fa565b5060019d9c50505050505050505050505050565b6000546001600160a01b031633146109d85760405162461bcd60e51b81526004016102aa90611063565b6001600160a01b038116610a3d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102aa565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff1615610b065760405162461bcd60e51b81526020600482015260156024820152745369676e657220616c72656164792065786973747360581b60448201526064016102aa565b60009182526001602081815260408085206001600160a01b039094168552600290930190529120805460ff19169091179055565b80821015610b985760405162461bcd60e51b815260206004820152602560248201527f4e756d626572206f66207369676e657273206c657373207468616e20746872656044820152641cda1bdb1960da1b60648201526084016102aa565b80610be55760405162461bcd60e51b815260206004820152601960248201527f5468726573686f6c642063616e206e6f74206265207a65726f0000000000000060448201526064016102aa565b6000928352600160205260409092209190915550565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff16610c685760405162461bcd60e51b815260206004820152601560248201527414da59db995c88191bd95cc81b9bdd08195e1a5cdd605a1b60448201526064016102aa565b60009182526001602090815260408084206001600160a01b039093168452600290920190529020805460ff19169055565b6000808686604051610cac929190611131565b604051809103902090508260ff16601b14158015610cce57508260ff16601c14155b15610cdd576000915050610d52565b60408051600081526020810180835283905260ff851691810191909152606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610d30573d6000803e3d6000fd5b505050602060405103516001600160a01b0316886001600160a01b0316149150505b9695505050505050565b80356001600160a01b0381168114610d7357600080fd5b919050565b600080600060608486031215610d8d57600080fd5b83359250610d9d60208501610d5c565b9150604084013590509250925092565b60008060408385031215610dc057600080fd5b82359150610dd060208401610d5c565b90509250929050565b60008083601f840112610deb57600080fd5b50813567ffffffffffffffff811115610e0357600080fd5b6020830191508360208260051b8501011115610e1e57600080fd5b9250929050565b600080600060408486031215610e3a57600080fd5b83359250602084013567ffffffffffffffff811115610e5857600080fd5b610e6486828701610dd9565b9497909650939450505050565b60008060008060608587031215610e8757600080fd5b84359350602085013567ffffffffffffffff811115610ea557600080fd5b610eb187828801610dd9565b9598909750949560400135949350505050565b600060208284031215610ed657600080fd5b5035919050565b60008083601f840112610eef57600080fd5b50813567ffffffffffffffff811115610f0757600080fd5b602083019150836020828501011115610e1e57600080fd5b600080600080600080600080600080600060c08c8e031215610f4057600080fd5b8b359a5067ffffffffffffffff8060208e01351115610f5e57600080fd5b610f6e8e60208f01358f01610dd9565b909b50995060408d0135811015610f8457600080fd5b610f948e60408f01358f01610dd9565b909950975060608d0135811015610faa57600080fd5b610fba8e60608f01358f01610dd9565b909750955060808d0135811015610fd057600080fd5b610fe08e60808f01358f01610dd9565b909550935060a08d0135811015610ff657600080fd5b506110078d60a08e01358e01610edd565b81935080925050509295989b509295989b9093969950565b6000806040838503121561103257600080fd5b50508035926020909101359150565b60006020828403121561105357600080fd5b61105c82610d5c565b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082198211156110c1576110c1611098565b500190565b60006000198214156110da576110da611098565b5060010190565b6000828210156110f3576110f3611098565b500390565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561112057600080fd5b813560ff8116811461105c57600080fd5b818382376000910190815291905056fea2646970667358221220dfcd5e6de4c84907a5d6c8394c9f0ce39eb7175d526075a7c4fc027032143ebc64736f6c63430008090033",
}

// TestRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use TestRegistrarMetaData.ABI instead.
var TestRegistrarABI = TestRegistrarMetaData.ABI

// TestRegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestRegistrarMetaData.Bin instead.
var TestRegistrarBin = TestRegistrarMetaData.Bin

// DeployTestRegistrar deploys a new Ethereum contract, binding an instance of TestRegistrar to it.
func DeployTestRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestRegistrar, error) {
	parsed, err := TestRegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestRegistrarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestRegistrar{TestRegistrarCaller: TestRegistrarCaller{contract: contract}, TestRegistrarTransactor: TestRegistrarTransactor{contract: contract}, TestRegistrarFilterer: TestRegistrarFilterer{contract: contract}}, nil
}

// TestRegistrar is an auto generated Go binding around an Ethereum contract.
type TestRegistrar struct {
	TestRegistrarCaller     // Read-only binding to the contract
	TestRegistrarTransactor // Write-only binding to the contract
	TestRegistrarFilterer   // Log filterer for contract events
}

// TestRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestRegistrarSession struct {
	Contract     *TestRegistrar    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestRegistrarCallerSession struct {
	Contract *TestRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestRegistrarTransactorSession struct {
	Contract     *TestRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRegistrarRaw struct {
	Contract *TestRegistrar // Generic contract binding to access the raw methods on
}

// TestRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestRegistrarCallerRaw struct {
	Contract *TestRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// TestRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestRegistrarTransactorRaw struct {
	Contract *TestRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestRegistrar creates a new instance of TestRegistrar, bound to a specific deployed contract.
func NewTestRegistrar(address common.Address, backend bind.ContractBackend) (*TestRegistrar, error) {
	contract, err := bindTestRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestRegistrar{TestRegistrarCaller: TestRegistrarCaller{contract: contract}, TestRegistrarTransactor: TestRegistrarTransactor{contract: contract}, TestRegistrarFilterer: TestRegistrarFilterer{contract: contract}}, nil
}

// NewTestRegistrarCaller creates a new read-only instance of TestRegistrar, bound to a specific deployed contract.
func NewTestRegistrarCaller(address common.Address, caller bind.ContractCaller) (*TestRegistrarCaller, error) {
	contract, err := bindTestRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestRegistrarCaller{contract: contract}, nil
}

// NewTestRegistrarTransactor creates a new write-only instance of TestRegistrar, bound to a specific deployed contract.
func NewTestRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*TestRegistrarTransactor, error) {
	contract, err := bindTestRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestRegistrarTransactor{contract: contract}, nil
}

// NewTestRegistrarFilterer creates a new log filterer instance of TestRegistrar, bound to a specific deployed contract.
func NewTestRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*TestRegistrarFilterer, error) {
	contract, err := bindTestRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestRegistrarFilterer{contract: contract}, nil
}

// bindTestRegistrar binds a generic wrapper to an already deployed contract.
func bindTestRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRegistrar *TestRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestRegistrar.Contract.TestRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRegistrar *TestRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRegistrar.Contract.TestRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRegistrar *TestRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRegistrar.Contract.TestRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRegistrar *TestRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRegistrar *TestRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRegistrar *TestRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRegistrar.Contract.contract.Transact(opts, method, params...)
}

// CheckThreshold is a free data retrieval call binding the contract method 0x8bd6ac82.
//
// Solidity: function checkThreshold(uint256 _blockchainId, address[] _signers) view returns(bool)
func (_TestRegistrar *TestRegistrarCaller) CheckThreshold(opts *bind.CallOpts, _blockchainId *big.Int, _signers []common.Address) (bool, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "checkThreshold", _blockchainId, _signers)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0x8bd6ac82.
//
// Solidity: function checkThreshold(uint256 _blockchainId, address[] _signers) view returns(bool)
func (_TestRegistrar *TestRegistrarSession) CheckThreshold(_blockchainId *big.Int, _signers []common.Address) (bool, error) {
	return _TestRegistrar.Contract.CheckThreshold(&_TestRegistrar.CallOpts, _blockchainId, _signers)
}

// CheckThreshold is a free data retrieval call binding the contract method 0x8bd6ac82.
//
// Solidity: function checkThreshold(uint256 _blockchainId, address[] _signers) view returns(bool)
func (_TestRegistrar *TestRegistrarCallerSession) CheckThreshold(_blockchainId *big.Int, _signers []common.Address) (bool, error) {
	return _TestRegistrar.Contract.CheckThreshold(&_TestRegistrar.CallOpts, _blockchainId, _signers)
}

// GetSigningThreshold is a free data retrieval call binding the contract method 0xa64ce199.
//
// Solidity: function getSigningThreshold(uint256 _blockchainId) view returns(uint256)
func (_TestRegistrar *TestRegistrarCaller) GetSigningThreshold(opts *bind.CallOpts, _blockchainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "getSigningThreshold", _blockchainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSigningThreshold is a free data retrieval call binding the contract method 0xa64ce199.
//
// Solidity: function getSigningThreshold(uint256 _blockchainId) view returns(uint256)
func (_TestRegistrar *TestRegistrarSession) GetSigningThreshold(_blockchainId *big.Int) (*big.Int, error) {
	return _TestRegistrar.Contract.GetSigningThreshold(&_TestRegistrar.CallOpts, _blockchainId)
}

// GetSigningThreshold is a free data retrieval call binding the contract method 0xa64ce199.
//
// Solidity: function getSigningThreshold(uint256 _blockchainId) view returns(uint256)
func (_TestRegistrar *TestRegistrarCallerSession) GetSigningThreshold(_blockchainId *big.Int) (*big.Int, error) {
	return _TestRegistrar.Contract.GetSigningThreshold(&_TestRegistrar.CallOpts, _blockchainId)
}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 _blockchainId, address _mightBeSigner) view returns(bool)
func (_TestRegistrar *TestRegistrarCaller) IsSigner(opts *bind.CallOpts, _blockchainId *big.Int, _mightBeSigner common.Address) (bool, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "isSigner", _blockchainId, _mightBeSigner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 _blockchainId, address _mightBeSigner) view returns(bool)
func (_TestRegistrar *TestRegistrarSession) IsSigner(_blockchainId *big.Int, _mightBeSigner common.Address) (bool, error) {
	return _TestRegistrar.Contract.IsSigner(&_TestRegistrar.CallOpts, _blockchainId, _mightBeSigner)
}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 _blockchainId, address _mightBeSigner) view returns(bool)
func (_TestRegistrar *TestRegistrarCallerSession) IsSigner(_blockchainId *big.Int, _mightBeSigner common.Address) (bool, error) {
	return _TestRegistrar.Contract.IsSigner(&_TestRegistrar.CallOpts, _blockchainId, _mightBeSigner)
}

// NumSigners is a free data retrieval call binding the contract method 0xf5e232ea.
//
// Solidity: function numSigners(uint256 _blockchainId) view returns(uint256)
func (_TestRegistrar *TestRegistrarCaller) NumSigners(opts *bind.CallOpts, _blockchainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "numSigners", _blockchainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumSigners is a free data retrieval call binding the contract method 0xf5e232ea.
//
// Solidity: function numSigners(uint256 _blockchainId) view returns(uint256)
func (_TestRegistrar *TestRegistrarSession) NumSigners(_blockchainId *big.Int) (*big.Int, error) {
	return _TestRegistrar.Contract.NumSigners(&_TestRegistrar.CallOpts, _blockchainId)
}

// NumSigners is a free data retrieval call binding the contract method 0xf5e232ea.
//
// Solidity: function numSigners(uint256 _blockchainId) view returns(uint256)
func (_TestRegistrar *TestRegistrarCallerSession) NumSigners(_blockchainId *big.Int) (*big.Int, error) {
	return _TestRegistrar.Contract.NumSigners(&_TestRegistrar.CallOpts, _blockchainId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestRegistrar *TestRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestRegistrar *TestRegistrarSession) Owner() (common.Address, error) {
	return _TestRegistrar.Contract.Owner(&_TestRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestRegistrar *TestRegistrarCallerSession) Owner() (common.Address, error) {
	return _TestRegistrar.Contract.Owner(&_TestRegistrar.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xea13ec8b.
//
// Solidity: function verify(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_TestRegistrar *TestRegistrarCaller) Verify(opts *bind.CallOpts, _blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "verify", _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xea13ec8b.
//
// Solidity: function verify(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_TestRegistrar *TestRegistrarSession) Verify(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _TestRegistrar.Contract.Verify(&_TestRegistrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// Verify is a free data retrieval call binding the contract method 0xea13ec8b.
//
// Solidity: function verify(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_TestRegistrar *TestRegistrarCallerSession) Verify(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _TestRegistrar.Contract.Verify(&_TestRegistrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// VerifyAndCheckThreshold is a free data retrieval call binding the contract method 0xad107bb4.
//
// Solidity: function verifyAndCheckThreshold(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_TestRegistrar *TestRegistrarCaller) VerifyAndCheckThreshold(opts *bind.CallOpts, _blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	var out []interface{}
	err := _TestRegistrar.contract.Call(opts, &out, "verifyAndCheckThreshold", _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAndCheckThreshold is a free data retrieval call binding the contract method 0xad107bb4.
//
// Solidity: function verifyAndCheckThreshold(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_TestRegistrar *TestRegistrarSession) VerifyAndCheckThreshold(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _TestRegistrar.Contract.VerifyAndCheckThreshold(&_TestRegistrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// VerifyAndCheckThreshold is a free data retrieval call binding the contract method 0xad107bb4.
//
// Solidity: function verifyAndCheckThreshold(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_TestRegistrar *TestRegistrarCallerSession) VerifyAndCheckThreshold(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _TestRegistrar.Contract.VerifyAndCheckThreshold(&_TestRegistrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// AddSigner is a paid mutator transaction binding the contract method 0x15a09825.
//
// Solidity: function addSigner(uint256 _bcId, address _signer) returns()
func (_TestRegistrar *TestRegistrarTransactor) AddSigner(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "addSigner", _bcId, _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0x15a09825.
//
// Solidity: function addSigner(uint256 _bcId, address _signer) returns()
func (_TestRegistrar *TestRegistrarSession) AddSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _TestRegistrar.Contract.AddSigner(&_TestRegistrar.TransactOpts, _bcId, _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0x15a09825.
//
// Solidity: function addSigner(uint256 _bcId, address _signer) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) AddSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _TestRegistrar.Contract.AddSigner(&_TestRegistrar.TransactOpts, _bcId, _signer)
}

// AddSignerSetThreshold is a paid mutator transaction binding the contract method 0x00ab2414.
//
// Solidity: function addSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactor) AddSignerSetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "addSignerSetThreshold", _bcId, _signer, _newSigningThreshold)
}

// AddSignerSetThreshold is a paid mutator transaction binding the contract method 0x00ab2414.
//
// Solidity: function addSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarSession) AddSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.AddSignerSetThreshold(&_TestRegistrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// AddSignerSetThreshold is a paid mutator transaction binding the contract method 0x00ab2414.
//
// Solidity: function addSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) AddSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.AddSignerSetThreshold(&_TestRegistrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// AddSignersSetThreshold is a paid mutator transaction binding the contract method 0x8d7678fd.
//
// Solidity: function addSignersSetThreshold(uint256 _bcId, address[] _signers, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactor) AddSignersSetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _signers []common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "addSignersSetThreshold", _bcId, _signers, _newSigningThreshold)
}

// AddSignersSetThreshold is a paid mutator transaction binding the contract method 0x8d7678fd.
//
// Solidity: function addSignersSetThreshold(uint256 _bcId, address[] _signers, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarSession) AddSignersSetThreshold(_bcId *big.Int, _signers []common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.AddSignersSetThreshold(&_TestRegistrar.TransactOpts, _bcId, _signers, _newSigningThreshold)
}

// AddSignersSetThreshold is a paid mutator transaction binding the contract method 0x8d7678fd.
//
// Solidity: function addSignersSetThreshold(uint256 _bcId, address[] _signers, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) AddSignersSetThreshold(_bcId *big.Int, _signers []common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.AddSignersSetThreshold(&_TestRegistrar.TransactOpts, _bcId, _signers, _newSigningThreshold)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3156d37c.
//
// Solidity: function removeSigner(uint256 _bcId, address _signer) returns()
func (_TestRegistrar *TestRegistrarTransactor) RemoveSigner(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "removeSigner", _bcId, _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3156d37c.
//
// Solidity: function removeSigner(uint256 _bcId, address _signer) returns()
func (_TestRegistrar *TestRegistrarSession) RemoveSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _TestRegistrar.Contract.RemoveSigner(&_TestRegistrar.TransactOpts, _bcId, _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3156d37c.
//
// Solidity: function removeSigner(uint256 _bcId, address _signer) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) RemoveSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _TestRegistrar.Contract.RemoveSigner(&_TestRegistrar.TransactOpts, _bcId, _signer)
}

// RemoveSignerSetThreshold is a paid mutator transaction binding the contract method 0xd4c0d34d.
//
// Solidity: function removeSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactor) RemoveSignerSetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "removeSignerSetThreshold", _bcId, _signer, _newSigningThreshold)
}

// RemoveSignerSetThreshold is a paid mutator transaction binding the contract method 0xd4c0d34d.
//
// Solidity: function removeSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarSession) RemoveSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.RemoveSignerSetThreshold(&_TestRegistrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// RemoveSignerSetThreshold is a paid mutator transaction binding the contract method 0xd4c0d34d.
//
// Solidity: function removeSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) RemoveSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.RemoveSignerSetThreshold(&_TestRegistrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestRegistrar *TestRegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestRegistrar *TestRegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestRegistrar.Contract.RenounceOwnership(&_TestRegistrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestRegistrar *TestRegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestRegistrar.Contract.RenounceOwnership(&_TestRegistrar.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _bcId, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactor) SetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "setThreshold", _bcId, _newSigningThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _bcId, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarSession) SetThreshold(_bcId *big.Int, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.SetThreshold(&_TestRegistrar.TransactOpts, _bcId, _newSigningThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _bcId, uint256 _newSigningThreshold) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) SetThreshold(_bcId *big.Int, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _TestRegistrar.Contract.SetThreshold(&_TestRegistrar.TransactOpts, _bcId, _newSigningThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestRegistrar *TestRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestRegistrar *TestRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestRegistrar.Contract.TransferOwnership(&_TestRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestRegistrar *TestRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestRegistrar.Contract.TransferOwnership(&_TestRegistrar.TransactOpts, newOwner)
}

// TestRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestRegistrar contract.
type TestRegistrarOwnershipTransferredIterator struct {
	Event *TestRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestRegistrarOwnershipTransferred)
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
		it.Event = new(TestRegistrarOwnershipTransferred)
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
func (it *TestRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the TestRegistrar contract.
type TestRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestRegistrar *TestRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestRegistrarOwnershipTransferredIterator{contract: _TestRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestRegistrar *TestRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestRegistrarOwnershipTransferred)
				if err := _TestRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestRegistrar *TestRegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*TestRegistrarOwnershipTransferred, error) {
	event := new(TestRegistrarOwnershipTransferred)
	if err := _TestRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
