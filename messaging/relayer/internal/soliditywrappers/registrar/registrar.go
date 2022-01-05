// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrar

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

// RegistrarMetaData contains all meta data concerning the Registrar contract.
var RegistrarMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_Log\",\"type\":\"string\"}],\"name\":\"Dump\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Dump2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"addSignerSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"addSignersSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"}],\"name\":\"getSigningThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mightBeSigner\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"}],\"name\":\"numSigners\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"removeSignerSetThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bcId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newSigningThreshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_plainText\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_plainText\",\"type\":\"bytes\"}],\"name\":\"verify2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockchainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigR\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sigS\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_sigV\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_plainText\",\"type\":\"bytes\"}],\"name\":\"verifyAndCheckThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350611a7f806100616000396000f3fe608060405234801561001057600080fd5b50600436106100ff5760003560e01c80639cabe58b11610097578063d4c0d34d11610066578063d4c0d34d14610217578063ea13ec8b1461022a578063f2fde38b1461023d578063f5e232ea1461025057600080fd5b80639cabe58b146101b0578063a64ce199146101c3578063ad107bb4146101f1578063b9c362091461020457600080fd5b8063715018a6116100d3578063715018a6146101675780638bd6ac821461016f5780638d7678fd146101825780638da5cb5b1461019557600080fd5b8062ab24141461010457806315a09825146101195780633156d37c1461012c57806348bcbd2d1461013f575b600080fd5b6101176101123660046113dd565b610274565b005b610117610127366004611412565b6102f5565b61011761013a366004611412565b6103c8565b61015261014d366004611412565b6104b3565b60405190151581526020015b60405180910390f35b6101176104e3565b61015261017d36600461148a565b610557565b6101176101903660046114d6565b6105b7565b6000546040516001600160a01b03909116815260200161015e565b6101526101be36600461156b565b610675565b6101e36101d136600461166b565b60009081526001602052604090205490565b60405190815260200161015e565b6101526101ff36600461156b565b6108be565b610117610212366004611684565b6108ef565b6101176102253660046113dd565b61093e565b61015261023836600461156b565b61098f565b61011761024b3660046116a6565b610bc4565b6101e361025e36600461166b565b6000908152600160208190526040909120015490565b6000546001600160a01b031633146102a75760405162461bcd60e51b815260040161029e906116c8565b60405180910390fd5b6102b18383610cae565b600083815260016020819052604082208101546102cd91611713565b60008581526001602081905260409091200181905590506102ef848284610d50565b50505050565b6000546001600160a01b0316331461031f5760405162461bcd60e51b815260040161029e906116c8565b6000828152600160205260409020546103985760405162461bcd60e51b815260206004820152603560248201527f43616e206e6f7420616464207369676e657220666f7220626c6f636b636861696044820152741b881dda5d1a081e995c9bc81d1a1c995cda1bdb19605a1b606482015260840161029e565b6103a28282610cae565b60008281526001602081905260408220018054916103bf8361172b565b91905055505050565b6000546001600160a01b031633146103f25760405162461bcd60e51b815260040161029e906116c8565b6103fc8282610e11565b600082815260016020819052604082208101546104199190611746565b6000848152600160205260409020549091508110156104985760405162461bcd60e51b815260206004820152603560248201527f50726f706f736564206e6577206e756d626572206f66207369676e65727320696044820152741cc81b195cdcc81d1a185b881d1a1c995cda1bdb19605a1b606482015260840161029e565b60009283526001602081905260409093209092019190915550565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff165b92915050565b6000546001600160a01b0316331461050d5760405162461bcd60e51b815260040161029e906116c8565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008381526001602052604081205482908110156105ac5760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768207369676e65727360701b604482015260640161029e565b506001949350505050565b6000546001600160a01b031633146105e15760405162461bcd60e51b815260040161029e906116c8565b60005b8281101561062e5761061c858585848181106106025761060261175d565b905060200201602081019061061791906116a6565b610cae565b806106268161172b565b9150506105e4565b5060008481526001602081905260408220015461064c908490611713565b600086815260016020819052604090912001819055905061066e858284610d50565b5050505050565b6000898881146106be5760405162461bcd60e51b81526020600482015260146024820152730e6d2cea440d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b8087146107045760405162461bcd60e51b81526020600482015260146024820152730e6d2cea640d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b80851461074a5760405162461bcd60e51b81526020600482015260146024820152730e6d2ceac40d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b60005b818110156108aa5760008e8152600160205260408120600201908e8e848181106107795761077961175d565b905060200201602081019061078e91906116a6565b6001600160a01b0316815260208101919091526040016000205460ff166107c75760405162461bcd60e51b815260040161029e90611773565b6108518d8d838181106107dc576107dc61175d565b90506020020160208101906107f191906116a6565b86868e8e868181106108055761080561175d565b905060200201358d8d8781811061081e5761081e61175d565b905060200201358c8c888181106108375761083761175d565b905060200201602081019061084c91906117b8565b610eaf565b6108985760405162461bcd60e51b81526020600482015260186024820152775369676e617475726520646964206e6f742076657269667960401b604482015260640161029e565b806108a28161172b565b91505061074d565b5060019d9c50505050505050505050505050565b60006108cb8c8c8c610557565b506108df8c8c8c8c8c8c8c8c8c8c8c61098f565b9c9b505050505050505050505050565b6000546001600160a01b031633146109195760405162461bcd60e51b815260040161029e906116c8565b61093a82600160008581526020019081526020016000206001015483610d50565b5050565b6000546001600160a01b031633146109685760405162461bcd60e51b815260040161029e906116c8565b6109728383610e11565b600083815260016020819052604082208101546102cd9190611746565b6000898881146109d85760405162461bcd60e51b81526020600482015260146024820152730e6d2cea440d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b808714610a1e5760405162461bcd60e51b81526020600482015260146024820152730e6d2cea640d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b808514610a645760405162461bcd60e51b81526020600482015260146024820152730e6d2ceac40d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b60005b818110156108aa5760008e8152600160205260408120600201908e8e84818110610a9357610a9361175d565b9050602002016020810190610aa891906116a6565b6001600160a01b0316815260208101919091526040016000205460ff16610ae15760405162461bcd60e51b815260040161029e90611773565b610b6b8d8d83818110610af657610af661175d565b9050602002016020810190610b0b91906116a6565b86868e8e86818110610b1f57610b1f61175d565b905060200201358d8d87818110610b3857610b3861175d565b905060200201358c8c88818110610b5157610b5161175d565b9050602002016020810190610b6691906117b8565b61105e565b610bb25760405162461bcd60e51b81526020600482015260186024820152775369676e617475726520646964206e6f742076657269667960401b604482015260640161029e565b80610bbc8161172b565b915050610a67565b6000546001600160a01b03163314610bee5760405162461bcd60e51b815260040161029e906116c8565b6001600160a01b038116610c535760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161029e565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff1615610d1c5760405162461bcd60e51b81526020600482015260156024820152745369676e657220616c72656164792065786973747360581b604482015260640161029e565b60009182526001602081815260408085206001600160a01b039094168552600290930190529120805460ff19169091179055565b80821015610dae5760405162461bcd60e51b815260206004820152602560248201527f4e756d626572206f66207369676e657273206c657373207468616e20746872656044820152641cda1bdb1960da1b606482015260840161029e565b80610dfb5760405162461bcd60e51b815260206004820152601960248201527f5468726573686f6c642063616e206e6f74206265207a65726f00000000000000604482015260640161029e565b6000928352600160205260409092209190915550565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff16610e7e5760405162461bcd60e51b815260206004820152601560248201527414da59db995c88191bd95cc81b9bdd08195e1a5cdd605a1b604482015260640161029e565b60009182526001602090815260408084206001600160a01b039093168452600290920190529020805460ff19169055565b60007ff4de966b0393276bf043f8f91ad6d569aff8ea50f2e5856fb2bb3cfd90ff01568686604051610ee29291906117db565b60405180910390a160008686604051610efc92919061180a565b604051809103902090507f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9610f338460ff1661111b565b604051610f40919061181a565b60405180910390a18260ff16601b14158015610f6057508260ff16601c14155b15610f6f576000915050611054565b7f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9610f9989611244565b604051610fa6919061181a565b60405180910390a16040805160008082526020820180845284905260ff861692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611002573d6000803e3d6000fd5b5050506020604051035190507f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae961103882611244565b604051611045919061181a565b60405180910390a16001925050505b9695505050505050565b600080868660405161107192919061180a565b604051809103902090508260ff16601b1415801561109357508260ff16601c14155b156110a2576000915050611054565b60408051600081526020810180835283905260ff851691810191909152606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156110f5573d6000803e3d6000fd5b505050602060405103516001600160a01b0316886001600160a01b031614915050611054565b60608161113f5750506040805180820190915260018152600360fc1b602082015290565b8160005b811561116957806111538161172b565b91506111629050600a83611885565b9150611143565b60008167ffffffffffffffff81111561118457611184611899565b6040519080825280601f01601f1916602001820160405280156111ae576020820181803683370190505b509050815b851561123b576111c4600182611746565b905060006111d3600a88611885565b6111de90600a6118af565b6111e89088611746565b6111f39060306118ce565b905060008160f81b9050808484815181106112105761121061175d565b60200101906001600160f81b031916908160001a905350611232600a89611885565b975050506111b3565b50949350505050565b60408051602880825260608281019093526000919060208201818036833701905050905060005b6014811015611384576000611281826013611746565b61128c9060086118af565b6112979060026119d7565b6112aa906001600160a01b038716611885565b60f81b9050600060108260f81c6112c191906119e3565b60f81b905060008160f81c60106112d89190611a05565b8360f81c6112e69190611a26565b60f81b90506112f48261138b565b856113008660026118af565b815181106113105761131061175d565b60200101906001600160f81b031916908160001a9053506113308161138b565b8561133c8660026118af565b611347906001611713565b815181106113575761135761175d565b60200101906001600160f81b031916908160001a905350505050808061137c9061172b565b91505061126b565b5092915050565b6000600a60f883901c10156113b2576113a960f883901c60306118ce565b60f81b92915050565b6113a960f883901c60576118ce565b919050565b80356001600160a01b03811681146113c157600080fd5b6000806000606084860312156113f257600080fd5b83359250611402602085016113c6565b9150604084013590509250925092565b6000806040838503121561142557600080fd5b82359150611435602084016113c6565b90509250929050565b60008083601f84011261145057600080fd5b50813567ffffffffffffffff81111561146857600080fd5b6020830191508360208260051b850101111561148357600080fd5b9250929050565b60008060006040848603121561149f57600080fd5b83359250602084013567ffffffffffffffff8111156114bd57600080fd5b6114c98682870161143e565b9497909650939450505050565b600080600080606085870312156114ec57600080fd5b84359350602085013567ffffffffffffffff81111561150a57600080fd5b6115168782880161143e565b9598909750949560400135949350505050565b60008083601f84011261153b57600080fd5b50813567ffffffffffffffff81111561155357600080fd5b60208301915083602082850101111561148357600080fd5b600080600080600080600080600080600060c08c8e03121561158c57600080fd5b8b359a5067ffffffffffffffff8060208e013511156115aa57600080fd5b6115ba8e60208f01358f0161143e565b909b50995060408d01358110156115d057600080fd5b6115e08e60408f01358f0161143e565b909950975060608d01358110156115f657600080fd5b6116068e60608f01358f0161143e565b909750955060808d013581101561161c57600080fd5b61162c8e60808f01358f0161143e565b909550935060a08d013581101561164257600080fd5b506116538d60a08e01358e01611529565b81935080925050509295989b509295989b9093969950565b60006020828403121561167d57600080fd5b5035919050565b6000806040838503121561169757600080fd5b50508035926020909101359150565b6000602082840312156116b857600080fd5b6116c1826113c6565b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b60008219821115611726576117266116fd565b500190565b600060001982141561173f5761173f6116fd565b5060010190565b600082821015611758576117586116fd565b500390565b634e487b7160e01b600052603260045260246000fd5b60208082526025908201527f5369676e6572206e6f74207369676e657220666f72207468697320626c6f636b60408201526431b430b4b760d91b606082015260800190565b6000602082840312156117ca57600080fd5b813560ff811681146116c157600080fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b8183823760009101908152919050565b600060208083528351808285015260005b818110156118475785810183015185820160400152820161182b565b81811115611859576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052601260045260246000fd5b6000826118945761189461186f565b500490565b634e487b7160e01b600052604160045260246000fd5b60008160001904831182151516156118c9576118c96116fd565b500290565b600060ff821660ff84168060ff038211156118eb576118eb6116fd565b019392505050565b600181815b8085111561192e578160001904821115611914576119146116fd565b8085161561192157918102915b93841c93908002906118f8565b509250929050565b600082611945575060016104dd565b81611952575060006104dd565b816001811461196857600281146119725761198e565b60019150506104dd565b60ff841115611983576119836116fd565b50506001821b6104dd565b5060208310610133831016604e8410600b84101617156119b1575081810a6104dd565b6119bb83836118f3565b80600019048211156119cf576119cf6116fd565b029392505050565b60006116c18383611936565b600060ff8316806119f6576119f661186f565b8060ff84160491505092915050565b600060ff821660ff84168160ff04811182151516156119cf576119cf6116fd565b600060ff821660ff841680821015611a4057611a406116fd565b9003939250505056fea26469706673582212202cb049250753a59f9aa9a34db29008465a386fbd12206eb15251c7df818bd24f64736f6c63430008090033",
}

// RegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrarMetaData.ABI instead.
var RegistrarABI = RegistrarMetaData.ABI

// RegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistrarMetaData.Bin instead.
var RegistrarBin = RegistrarMetaData.Bin

// DeployRegistrar deploys a new Ethereum contract, binding an instance of Registrar to it.
func DeployRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registrar, error) {
	parsed, err := RegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistrarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}, RegistrarFilterer: RegistrarFilterer{contract: contract}}, nil
}

// Registrar is an auto generated Go binding around an Ethereum contract.
type Registrar struct {
	RegistrarCaller     // Read-only binding to the contract
	RegistrarTransactor // Write-only binding to the contract
	RegistrarFilterer   // Log filterer for contract events
}

// RegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarSession struct {
	Contract     *Registrar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarCallerSession struct {
	Contract *RegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarTransactorSession struct {
	Contract     *RegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarRaw struct {
	Contract *Registrar // Generic contract binding to access the raw methods on
}

// RegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarCallerRaw struct {
	Contract *RegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarTransactorRaw struct {
	Contract *RegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrar creates a new instance of Registrar, bound to a specific deployed contract.
func NewRegistrar(address common.Address, backend bind.ContractBackend) (*Registrar, error) {
	contract, err := bindRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}, RegistrarFilterer: RegistrarFilterer{contract: contract}}, nil
}

// NewRegistrarCaller creates a new read-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarCaller(address common.Address, caller bind.ContractCaller) (*RegistrarCaller, error) {
	contract, err := bindRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarCaller{contract: contract}, nil
}

// NewRegistrarTransactor creates a new write-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarTransactor, error) {
	contract, err := bindRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarTransactor{contract: contract}, nil
}

// NewRegistrarFilterer creates a new log filterer instance of Registrar, bound to a specific deployed contract.
func NewRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrarFilterer, error) {
	contract, err := bindRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrarFilterer{contract: contract}, nil
}

// bindRegistrar binds a generic wrapper to an already deployed contract.
func bindRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.RegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transact(opts, method, params...)
}

// CheckThreshold is a free data retrieval call binding the contract method 0x8bd6ac82.
//
// Solidity: function checkThreshold(uint256 _blockchainId, address[] _signers) view returns(bool)
func (_Registrar *RegistrarCaller) CheckThreshold(opts *bind.CallOpts, _blockchainId *big.Int, _signers []common.Address) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "checkThreshold", _blockchainId, _signers)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0x8bd6ac82.
//
// Solidity: function checkThreshold(uint256 _blockchainId, address[] _signers) view returns(bool)
func (_Registrar *RegistrarSession) CheckThreshold(_blockchainId *big.Int, _signers []common.Address) (bool, error) {
	return _Registrar.Contract.CheckThreshold(&_Registrar.CallOpts, _blockchainId, _signers)
}

// CheckThreshold is a free data retrieval call binding the contract method 0x8bd6ac82.
//
// Solidity: function checkThreshold(uint256 _blockchainId, address[] _signers) view returns(bool)
func (_Registrar *RegistrarCallerSession) CheckThreshold(_blockchainId *big.Int, _signers []common.Address) (bool, error) {
	return _Registrar.Contract.CheckThreshold(&_Registrar.CallOpts, _blockchainId, _signers)
}

// GetSigningThreshold is a free data retrieval call binding the contract method 0xa64ce199.
//
// Solidity: function getSigningThreshold(uint256 _blockchainId) view returns(uint256)
func (_Registrar *RegistrarCaller) GetSigningThreshold(opts *bind.CallOpts, _blockchainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "getSigningThreshold", _blockchainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSigningThreshold is a free data retrieval call binding the contract method 0xa64ce199.
//
// Solidity: function getSigningThreshold(uint256 _blockchainId) view returns(uint256)
func (_Registrar *RegistrarSession) GetSigningThreshold(_blockchainId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.GetSigningThreshold(&_Registrar.CallOpts, _blockchainId)
}

// GetSigningThreshold is a free data retrieval call binding the contract method 0xa64ce199.
//
// Solidity: function getSigningThreshold(uint256 _blockchainId) view returns(uint256)
func (_Registrar *RegistrarCallerSession) GetSigningThreshold(_blockchainId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.GetSigningThreshold(&_Registrar.CallOpts, _blockchainId)
}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 _blockchainId, address _mightBeSigner) view returns(bool)
func (_Registrar *RegistrarCaller) IsSigner(opts *bind.CallOpts, _blockchainId *big.Int, _mightBeSigner common.Address) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "isSigner", _blockchainId, _mightBeSigner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 _blockchainId, address _mightBeSigner) view returns(bool)
func (_Registrar *RegistrarSession) IsSigner(_blockchainId *big.Int, _mightBeSigner common.Address) (bool, error) {
	return _Registrar.Contract.IsSigner(&_Registrar.CallOpts, _blockchainId, _mightBeSigner)
}

// IsSigner is a free data retrieval call binding the contract method 0x48bcbd2d.
//
// Solidity: function isSigner(uint256 _blockchainId, address _mightBeSigner) view returns(bool)
func (_Registrar *RegistrarCallerSession) IsSigner(_blockchainId *big.Int, _mightBeSigner common.Address) (bool, error) {
	return _Registrar.Contract.IsSigner(&_Registrar.CallOpts, _blockchainId, _mightBeSigner)
}

// NumSigners is a free data retrieval call binding the contract method 0xf5e232ea.
//
// Solidity: function numSigners(uint256 _blockchainId) view returns(uint256)
func (_Registrar *RegistrarCaller) NumSigners(opts *bind.CallOpts, _blockchainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "numSigners", _blockchainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumSigners is a free data retrieval call binding the contract method 0xf5e232ea.
//
// Solidity: function numSigners(uint256 _blockchainId) view returns(uint256)
func (_Registrar *RegistrarSession) NumSigners(_blockchainId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.NumSigners(&_Registrar.CallOpts, _blockchainId)
}

// NumSigners is a free data retrieval call binding the contract method 0xf5e232ea.
//
// Solidity: function numSigners(uint256 _blockchainId) view returns(uint256)
func (_Registrar *RegistrarCallerSession) NumSigners(_blockchainId *big.Int) (*big.Int, error) {
	return _Registrar.Contract.NumSigners(&_Registrar.CallOpts, _blockchainId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrar *RegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrar *RegistrarSession) Owner() (common.Address, error) {
	return _Registrar.Contract.Owner(&_Registrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registrar *RegistrarCallerSession) Owner() (common.Address, error) {
	return _Registrar.Contract.Owner(&_Registrar.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xea13ec8b.
//
// Solidity: function verify(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_Registrar *RegistrarCaller) Verify(opts *bind.CallOpts, _blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "verify", _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xea13ec8b.
//
// Solidity: function verify(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_Registrar *RegistrarSession) Verify(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _Registrar.Contract.Verify(&_Registrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// Verify is a free data retrieval call binding the contract method 0xea13ec8b.
//
// Solidity: function verify(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_Registrar *RegistrarCallerSession) Verify(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _Registrar.Contract.Verify(&_Registrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// VerifyAndCheckThreshold is a free data retrieval call binding the contract method 0xad107bb4.
//
// Solidity: function verifyAndCheckThreshold(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_Registrar *RegistrarCaller) VerifyAndCheckThreshold(opts *bind.CallOpts, _blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	var out []interface{}
	err := _Registrar.contract.Call(opts, &out, "verifyAndCheckThreshold", _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyAndCheckThreshold is a free data retrieval call binding the contract method 0xad107bb4.
//
// Solidity: function verifyAndCheckThreshold(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_Registrar *RegistrarSession) VerifyAndCheckThreshold(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _Registrar.Contract.VerifyAndCheckThreshold(&_Registrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// VerifyAndCheckThreshold is a free data retrieval call binding the contract method 0xad107bb4.
//
// Solidity: function verifyAndCheckThreshold(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) view returns(bool)
func (_Registrar *RegistrarCallerSession) VerifyAndCheckThreshold(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (bool, error) {
	return _Registrar.Contract.VerifyAndCheckThreshold(&_Registrar.CallOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// AddSigner is a paid mutator transaction binding the contract method 0x15a09825.
//
// Solidity: function addSigner(uint256 _bcId, address _signer) returns()
func (_Registrar *RegistrarTransactor) AddSigner(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "addSigner", _bcId, _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0x15a09825.
//
// Solidity: function addSigner(uint256 _bcId, address _signer) returns()
func (_Registrar *RegistrarSession) AddSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.AddSigner(&_Registrar.TransactOpts, _bcId, _signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0x15a09825.
//
// Solidity: function addSigner(uint256 _bcId, address _signer) returns()
func (_Registrar *RegistrarTransactorSession) AddSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.AddSigner(&_Registrar.TransactOpts, _bcId, _signer)
}

// AddSignerSetThreshold is a paid mutator transaction binding the contract method 0x00ab2414.
//
// Solidity: function addSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactor) AddSignerSetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "addSignerSetThreshold", _bcId, _signer, _newSigningThreshold)
}

// AddSignerSetThreshold is a paid mutator transaction binding the contract method 0x00ab2414.
//
// Solidity: function addSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarSession) AddSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.AddSignerSetThreshold(&_Registrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// AddSignerSetThreshold is a paid mutator transaction binding the contract method 0x00ab2414.
//
// Solidity: function addSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactorSession) AddSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.AddSignerSetThreshold(&_Registrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// AddSignersSetThreshold is a paid mutator transaction binding the contract method 0x8d7678fd.
//
// Solidity: function addSignersSetThreshold(uint256 _bcId, address[] _signers, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactor) AddSignersSetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _signers []common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "addSignersSetThreshold", _bcId, _signers, _newSigningThreshold)
}

// AddSignersSetThreshold is a paid mutator transaction binding the contract method 0x8d7678fd.
//
// Solidity: function addSignersSetThreshold(uint256 _bcId, address[] _signers, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarSession) AddSignersSetThreshold(_bcId *big.Int, _signers []common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.AddSignersSetThreshold(&_Registrar.TransactOpts, _bcId, _signers, _newSigningThreshold)
}

// AddSignersSetThreshold is a paid mutator transaction binding the contract method 0x8d7678fd.
//
// Solidity: function addSignersSetThreshold(uint256 _bcId, address[] _signers, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactorSession) AddSignersSetThreshold(_bcId *big.Int, _signers []common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.AddSignersSetThreshold(&_Registrar.TransactOpts, _bcId, _signers, _newSigningThreshold)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3156d37c.
//
// Solidity: function removeSigner(uint256 _bcId, address _signer) returns()
func (_Registrar *RegistrarTransactor) RemoveSigner(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "removeSigner", _bcId, _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3156d37c.
//
// Solidity: function removeSigner(uint256 _bcId, address _signer) returns()
func (_Registrar *RegistrarSession) RemoveSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.RemoveSigner(&_Registrar.TransactOpts, _bcId, _signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x3156d37c.
//
// Solidity: function removeSigner(uint256 _bcId, address _signer) returns()
func (_Registrar *RegistrarTransactorSession) RemoveSigner(_bcId *big.Int, _signer common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.RemoveSigner(&_Registrar.TransactOpts, _bcId, _signer)
}

// RemoveSignerSetThreshold is a paid mutator transaction binding the contract method 0xd4c0d34d.
//
// Solidity: function removeSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactor) RemoveSignerSetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "removeSignerSetThreshold", _bcId, _signer, _newSigningThreshold)
}

// RemoveSignerSetThreshold is a paid mutator transaction binding the contract method 0xd4c0d34d.
//
// Solidity: function removeSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarSession) RemoveSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.RemoveSignerSetThreshold(&_Registrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// RemoveSignerSetThreshold is a paid mutator transaction binding the contract method 0xd4c0d34d.
//
// Solidity: function removeSignerSetThreshold(uint256 _bcId, address _signer, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactorSession) RemoveSignerSetThreshold(_bcId *big.Int, _signer common.Address, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.RemoveSignerSetThreshold(&_Registrar.TransactOpts, _bcId, _signer, _newSigningThreshold)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrar *RegistrarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrar *RegistrarSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registrar.Contract.RenounceOwnership(&_Registrar.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registrar *RegistrarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registrar.Contract.RenounceOwnership(&_Registrar.TransactOpts)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _bcId, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactor) SetThreshold(opts *bind.TransactOpts, _bcId *big.Int, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "setThreshold", _bcId, _newSigningThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _bcId, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarSession) SetThreshold(_bcId *big.Int, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.SetThreshold(&_Registrar.TransactOpts, _bcId, _newSigningThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _bcId, uint256 _newSigningThreshold) returns()
func (_Registrar *RegistrarTransactorSession) SetThreshold(_bcId *big.Int, _newSigningThreshold *big.Int) (*types.Transaction, error) {
	return _Registrar.Contract.SetThreshold(&_Registrar.TransactOpts, _bcId, _newSigningThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrar *RegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrar *RegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.TransferOwnership(&_Registrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registrar *RegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.TransferOwnership(&_Registrar.TransactOpts, newOwner)
}

// Verify2 is a paid mutator transaction binding the contract method 0x9cabe58b.
//
// Solidity: function verify2(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) returns(bool)
func (_Registrar *RegistrarTransactor) Verify2(opts *bind.TransactOpts, _blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "verify2", _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// Verify2 is a paid mutator transaction binding the contract method 0x9cabe58b.
//
// Solidity: function verify2(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) returns(bool)
func (_Registrar *RegistrarSession) Verify2(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (*types.Transaction, error) {
	return _Registrar.Contract.Verify2(&_Registrar.TransactOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// Verify2 is a paid mutator transaction binding the contract method 0x9cabe58b.
//
// Solidity: function verify2(uint256 _blockchainId, address[] _signers, bytes32[] _sigR, bytes32[] _sigS, uint8[] _sigV, bytes _plainText) returns(bool)
func (_Registrar *RegistrarTransactorSession) Verify2(_blockchainId *big.Int, _signers []common.Address, _sigR [][32]byte, _sigS [][32]byte, _sigV []uint8, _plainText []byte) (*types.Transaction, error) {
	return _Registrar.Contract.Verify2(&_Registrar.TransactOpts, _blockchainId, _signers, _sigR, _sigS, _sigV, _plainText)
}

// RegistrarDumpIterator is returned from FilterDump and is used to iterate over the raw logs and unpacked data for Dump events raised by the Registrar contract.
type RegistrarDumpIterator struct {
	Event *RegistrarDump // Event containing the contract specifics and raw log

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
func (it *RegistrarDumpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarDump)
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
		it.Event = new(RegistrarDump)
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
func (it *RegistrarDumpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarDumpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarDump represents a Dump event raised by the Registrar contract.
type RegistrarDump struct {
	Log string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDump is a free log retrieval operation binding the contract event 0x922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9.
//
// Solidity: event Dump(string _Log)
func (_Registrar *RegistrarFilterer) FilterDump(opts *bind.FilterOpts) (*RegistrarDumpIterator, error) {

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "Dump")
	if err != nil {
		return nil, err
	}
	return &RegistrarDumpIterator{contract: _Registrar.contract, event: "Dump", logs: logs, sub: sub}, nil
}

// WatchDump is a free log subscription operation binding the contract event 0x922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9.
//
// Solidity: event Dump(string _Log)
func (_Registrar *RegistrarFilterer) WatchDump(opts *bind.WatchOpts, sink chan<- *RegistrarDump) (event.Subscription, error) {

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "Dump")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarDump)
				if err := _Registrar.contract.UnpackLog(event, "Dump", log); err != nil {
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
// Solidity: event Dump(string _Log)
func (_Registrar *RegistrarFilterer) ParseDump(log types.Log) (*RegistrarDump, error) {
	event := new(RegistrarDump)
	if err := _Registrar.contract.UnpackLog(event, "Dump", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarDump2Iterator is returned from FilterDump2 and is used to iterate over the raw logs and unpacked data for Dump2 events raised by the Registrar contract.
type RegistrarDump2Iterator struct {
	Event *RegistrarDump2 // Event containing the contract specifics and raw log

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
func (it *RegistrarDump2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarDump2)
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
		it.Event = new(RegistrarDump2)
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
func (it *RegistrarDump2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarDump2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarDump2 represents a Dump2 event raised by the Registrar contract.
type RegistrarDump2 struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDump2 is a free log retrieval operation binding the contract event 0xf4de966b0393276bf043f8f91ad6d569aff8ea50f2e5856fb2bb3cfd90ff0156.
//
// Solidity: event Dump2(bytes data)
func (_Registrar *RegistrarFilterer) FilterDump2(opts *bind.FilterOpts) (*RegistrarDump2Iterator, error) {

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "Dump2")
	if err != nil {
		return nil, err
	}
	return &RegistrarDump2Iterator{contract: _Registrar.contract, event: "Dump2", logs: logs, sub: sub}, nil
}

// WatchDump2 is a free log subscription operation binding the contract event 0xf4de966b0393276bf043f8f91ad6d569aff8ea50f2e5856fb2bb3cfd90ff0156.
//
// Solidity: event Dump2(bytes data)
func (_Registrar *RegistrarFilterer) WatchDump2(opts *bind.WatchOpts, sink chan<- *RegistrarDump2) (event.Subscription, error) {

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "Dump2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarDump2)
				if err := _Registrar.contract.UnpackLog(event, "Dump2", log); err != nil {
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

// ParseDump2 is a log parse operation binding the contract event 0xf4de966b0393276bf043f8f91ad6d569aff8ea50f2e5856fb2bb3cfd90ff0156.
//
// Solidity: event Dump2(bytes data)
func (_Registrar *RegistrarFilterer) ParseDump2(log types.Log) (*RegistrarDump2, error) {
	event := new(RegistrarDump2)
	if err := _Registrar.contract.UnpackLog(event, "Dump2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registrar contract.
type RegistrarOwnershipTransferredIterator struct {
	Event *RegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrarOwnershipTransferred)
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
		it.Event = new(RegistrarOwnershipTransferred)
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
func (it *RegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the Registrar contract.
type RegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registrar *RegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrarOwnershipTransferredIterator{contract: _Registrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registrar *RegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrarOwnershipTransferred)
				if err := _Registrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registrar *RegistrarFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrarOwnershipTransferred, error) {
	event := new(RegistrarOwnershipTransferred)
	if err := _Registrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
