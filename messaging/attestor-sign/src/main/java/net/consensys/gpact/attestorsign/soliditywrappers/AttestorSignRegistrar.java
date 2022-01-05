package net.consensys.gpact.attestorsign.soliditywrappers;

import io.reactivex.Flowable;
import io.reactivex.functions.Function;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.DynamicBytes;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.RemoteFunctionCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.BaseEventResponse;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 4.8.7-SNAPSHOT.
 */
@SuppressWarnings("rawtypes")
public class AttestorSignRegistrar extends Contract {
    public static final String BINARY = "608060405234801561001057600080fd5b50600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350611a7f806100616000396000f3fe608060405234801561001057600080fd5b50600436106100ff5760003560e01c80639cabe58b11610097578063d4c0d34d11610066578063d4c0d34d14610217578063ea13ec8b1461022a578063f2fde38b1461023d578063f5e232ea1461025057600080fd5b80639cabe58b146101b0578063a64ce199146101c3578063ad107bb4146101f1578063b9c362091461020457600080fd5b8063715018a6116100d3578063715018a6146101675780638bd6ac821461016f5780638d7678fd146101825780638da5cb5b1461019557600080fd5b8062ab24141461010457806315a09825146101195780633156d37c1461012c57806348bcbd2d1461013f575b600080fd5b6101176101123660046113dd565b610274565b005b610117610127366004611412565b6102f5565b61011761013a366004611412565b6103c8565b61015261014d366004611412565b6104b3565b60405190151581526020015b60405180910390f35b6101176104e3565b61015261017d36600461148a565b610557565b6101176101903660046114d6565b6105b7565b6000546040516001600160a01b03909116815260200161015e565b6101526101be36600461156b565b610675565b6101e36101d136600461166b565b60009081526001602052604090205490565b60405190815260200161015e565b6101526101ff36600461156b565b6108be565b610117610212366004611684565b6108ef565b6101176102253660046113dd565b61093e565b61015261023836600461156b565b61098f565b61011761024b3660046116a6565b610bc4565b6101e361025e36600461166b565b6000908152600160208190526040909120015490565b6000546001600160a01b031633146102a75760405162461bcd60e51b815260040161029e906116c8565b60405180910390fd5b6102b18383610cae565b600083815260016020819052604082208101546102cd91611713565b60008581526001602081905260409091200181905590506102ef848284610d50565b50505050565b6000546001600160a01b0316331461031f5760405162461bcd60e51b815260040161029e906116c8565b6000828152600160205260409020546103985760405162461bcd60e51b815260206004820152603560248201527f43616e206e6f7420616464207369676e657220666f7220626c6f636b636861696044820152741b881dda5d1a081e995c9bc81d1a1c995cda1bdb19605a1b606482015260840161029e565b6103a28282610cae565b60008281526001602081905260408220018054916103bf8361172b565b91905055505050565b6000546001600160a01b031633146103f25760405162461bcd60e51b815260040161029e906116c8565b6103fc8282610e11565b600082815260016020819052604082208101546104199190611746565b6000848152600160205260409020549091508110156104985760405162461bcd60e51b815260206004820152603560248201527f50726f706f736564206e6577206e756d626572206f66207369676e65727320696044820152741cc81b195cdcc81d1a185b881d1a1c995cda1bdb19605a1b606482015260840161029e565b60009283526001602081905260409093209092019190915550565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff165b92915050565b6000546001600160a01b0316331461050d5760405162461bcd60e51b815260040161029e906116c8565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008381526001602052604081205482908110156105ac5760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768207369676e65727360701b604482015260640161029e565b506001949350505050565b6000546001600160a01b031633146105e15760405162461bcd60e51b815260040161029e906116c8565b60005b8281101561062e5761061c858585848181106106025761060261175d565b905060200201602081019061061791906116a6565b610cae565b806106268161172b565b9150506105e4565b5060008481526001602081905260408220015461064c908490611713565b600086815260016020819052604090912001819055905061066e858284610d50565b5050505050565b6000898881146106be5760405162461bcd60e51b81526020600482015260146024820152730e6d2cea440d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b8087146107045760405162461bcd60e51b81526020600482015260146024820152730e6d2cea640d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b80851461074a5760405162461bcd60e51b81526020600482015260146024820152730e6d2ceac40d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b60005b818110156108aa5760008e8152600160205260408120600201908e8e848181106107795761077961175d565b905060200201602081019061078e91906116a6565b6001600160a01b0316815260208101919091526040016000205460ff166107c75760405162461bcd60e51b815260040161029e90611773565b6108518d8d838181106107dc576107dc61175d565b90506020020160208101906107f191906116a6565b86868e8e868181106108055761080561175d565b905060200201358d8d8781811061081e5761081e61175d565b905060200201358c8c888181106108375761083761175d565b905060200201602081019061084c91906117b8565b610eaf565b6108985760405162461bcd60e51b81526020600482015260186024820152775369676e617475726520646964206e6f742076657269667960401b604482015260640161029e565b806108a28161172b565b91505061074d565b5060019d9c50505050505050505050505050565b60006108cb8c8c8c610557565b506108df8c8c8c8c8c8c8c8c8c8c8c61098f565b9c9b505050505050505050505050565b6000546001600160a01b031633146109195760405162461bcd60e51b815260040161029e906116c8565b61093a82600160008581526020019081526020016000206001015483610d50565b5050565b6000546001600160a01b031633146109685760405162461bcd60e51b815260040161029e906116c8565b6109728383610e11565b600083815260016020819052604082208101546102cd9190611746565b6000898881146109d85760405162461bcd60e51b81526020600482015260146024820152730e6d2cea440d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b808714610a1e5760405162461bcd60e51b81526020600482015260146024820152730e6d2cea640d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b808514610a645760405162461bcd60e51b81526020600482015260146024820152730e6d2ceac40d8cadccee8d040dad2e6dac2e8c6d60631b604482015260640161029e565b60005b818110156108aa5760008e8152600160205260408120600201908e8e84818110610a9357610a9361175d565b9050602002016020810190610aa891906116a6565b6001600160a01b0316815260208101919091526040016000205460ff16610ae15760405162461bcd60e51b815260040161029e90611773565b610b6b8d8d83818110610af657610af661175d565b9050602002016020810190610b0b91906116a6565b86868e8e86818110610b1f57610b1f61175d565b905060200201358d8d87818110610b3857610b3861175d565b905060200201358c8c88818110610b5157610b5161175d565b9050602002016020810190610b6691906117b8565b61105e565b610bb25760405162461bcd60e51b81526020600482015260186024820152775369676e617475726520646964206e6f742076657269667960401b604482015260640161029e565b80610bbc8161172b565b915050610a67565b6000546001600160a01b03163314610bee5760405162461bcd60e51b815260040161029e906116c8565b6001600160a01b038116610c535760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161029e565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff1615610d1c5760405162461bcd60e51b81526020600482015260156024820152745369676e657220616c72656164792065786973747360581b604482015260640161029e565b60009182526001602081815260408085206001600160a01b039094168552600290930190529120805460ff19169091179055565b80821015610dae5760405162461bcd60e51b815260206004820152602560248201527f4e756d626572206f66207369676e657273206c657373207468616e20746872656044820152641cda1bdb1960da1b606482015260840161029e565b80610dfb5760405162461bcd60e51b815260206004820152601960248201527f5468726573686f6c642063616e206e6f74206265207a65726f00000000000000604482015260640161029e565b6000928352600160205260409092209190915550565b60008281526001602090815260408083206001600160a01b038516845260020190915290205460ff16610e7e5760405162461bcd60e51b815260206004820152601560248201527414da59db995c88191bd95cc81b9bdd08195e1a5cdd605a1b604482015260640161029e565b60009182526001602090815260408084206001600160a01b039093168452600290920190529020805460ff19169055565b60007ff4de966b0393276bf043f8f91ad6d569aff8ea50f2e5856fb2bb3cfd90ff01568686604051610ee29291906117db565b60405180910390a160008686604051610efc92919061180a565b604051809103902090507f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9610f338460ff1661111b565b604051610f40919061181a565b60405180910390a18260ff16601b14158015610f6057508260ff16601c14155b15610f6f576000915050611054565b7f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9610f9989611244565b604051610fa6919061181a565b60405180910390a16040805160008082526020820180845284905260ff861692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611002573d6000803e3d6000fd5b5050506020604051035190507f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae961103882611244565b604051611045919061181a565b60405180910390a16001925050505b9695505050505050565b600080868660405161107192919061180a565b604051809103902090508260ff16601b1415801561109357508260ff16601c14155b156110a2576000915050611054565b60408051600081526020810180835283905260ff851691810191909152606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156110f5573d6000803e3d6000fd5b505050602060405103516001600160a01b0316886001600160a01b031614915050611054565b60608161113f5750506040805180820190915260018152600360fc1b602082015290565b8160005b811561116957806111538161172b565b91506111629050600a83611885565b9150611143565b60008167ffffffffffffffff81111561118457611184611899565b6040519080825280601f01601f1916602001820160405280156111ae576020820181803683370190505b509050815b851561123b576111c4600182611746565b905060006111d3600a88611885565b6111de90600a6118af565b6111e89088611746565b6111f39060306118ce565b905060008160f81b9050808484815181106112105761121061175d565b60200101906001600160f81b031916908160001a905350611232600a89611885565b975050506111b3565b50949350505050565b60408051602880825260608281019093526000919060208201818036833701905050905060005b6014811015611384576000611281826013611746565b61128c9060086118af565b6112979060026119d7565b6112aa906001600160a01b038716611885565b60f81b9050600060108260f81c6112c191906119e3565b60f81b905060008160f81c60106112d89190611a05565b8360f81c6112e69190611a26565b60f81b90506112f48261138b565b856113008660026118af565b815181106113105761131061175d565b60200101906001600160f81b031916908160001a9053506113308161138b565b8561133c8660026118af565b611347906001611713565b815181106113575761135761175d565b60200101906001600160f81b031916908160001a905350505050808061137c9061172b565b91505061126b565b5092915050565b6000600a60f883901c10156113b2576113a960f883901c60306118ce565b60f81b92915050565b6113a960f883901c60576118ce565b919050565b80356001600160a01b03811681146113c157600080fd5b6000806000606084860312156113f257600080fd5b83359250611402602085016113c6565b9150604084013590509250925092565b6000806040838503121561142557600080fd5b82359150611435602084016113c6565b90509250929050565b60008083601f84011261145057600080fd5b50813567ffffffffffffffff81111561146857600080fd5b6020830191508360208260051b850101111561148357600080fd5b9250929050565b60008060006040848603121561149f57600080fd5b83359250602084013567ffffffffffffffff8111156114bd57600080fd5b6114c98682870161143e565b9497909650939450505050565b600080600080606085870312156114ec57600080fd5b84359350602085013567ffffffffffffffff81111561150a57600080fd5b6115168782880161143e565b9598909750949560400135949350505050565b60008083601f84011261153b57600080fd5b50813567ffffffffffffffff81111561155357600080fd5b60208301915083602082850101111561148357600080fd5b600080600080600080600080600080600060c08c8e03121561158c57600080fd5b8b359a5067ffffffffffffffff8060208e013511156115aa57600080fd5b6115ba8e60208f01358f0161143e565b909b50995060408d01358110156115d057600080fd5b6115e08e60408f01358f0161143e565b909950975060608d01358110156115f657600080fd5b6116068e60608f01358f0161143e565b909750955060808d013581101561161c57600080fd5b61162c8e60808f01358f0161143e565b909550935060a08d013581101561164257600080fd5b506116538d60a08e01358e01611529565b81935080925050509295989b509295989b9093969950565b60006020828403121561167d57600080fd5b5035919050565b6000806040838503121561169757600080fd5b50508035926020909101359150565b6000602082840312156116b857600080fd5b6116c1826113c6565b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052601160045260246000fd5b60008219821115611726576117266116fd565b500190565b600060001982141561173f5761173f6116fd565b5060010190565b600082821015611758576117586116fd565b500390565b634e487b7160e01b600052603260045260246000fd5b60208082526025908201527f5369676e6572206e6f74207369676e657220666f72207468697320626c6f636b60408201526431b430b4b760d91b606082015260800190565b6000602082840312156117ca57600080fd5b813560ff811681146116c157600080fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b8183823760009101908152919050565b600060208083528351808285015260005b818110156118475785810183015185820160400152820161182b565b81811115611859576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052601260045260246000fd5b6000826118945761189461186f565b500490565b634e487b7160e01b600052604160045260246000fd5b60008160001904831182151516156118c9576118c96116fd565b500290565b600060ff821660ff84168060ff038211156118eb576118eb6116fd565b019392505050565b600181815b8085111561192e578160001904821115611914576119146116fd565b8085161561192157918102915b93841c93908002906118f8565b509250929050565b600082611945575060016104dd565b81611952575060006104dd565b816001811461196857600281146119725761198e565b60019150506104dd565b60ff841115611983576119836116fd565b50506001821b6104dd565b5060208310610133831016604e8410600b84101617156119b1575081810a6104dd565b6119bb83836118f3565b80600019048211156119cf576119cf6116fd565b029392505050565b60006116c18383611936565b600060ff8316806119f6576119f661186f565b8060ff84160491505092915050565b600060ff821660ff84168160ff04811182151516156119cf576119cf6116fd565b600060ff821660ff841680821015611a4057611a406116fd565b9003939250505056fea26469706673582212202cb049250753a59f9aa9a34db29008465a386fbd12206eb15251c7df818bd24f64736f6c63430008090033";

    public static final String FUNC_ADDSIGNER = "addSigner";

    public static final String FUNC_ADDSIGNERSETTHRESHOLD = "addSignerSetThreshold";

    public static final String FUNC_ADDSIGNERSSETTHRESHOLD = "addSignersSetThreshold";

    public static final String FUNC_CHECKTHRESHOLD = "checkThreshold";

    public static final String FUNC_GETSIGNINGTHRESHOLD = "getSigningThreshold";

    public static final String FUNC_ISSIGNER = "isSigner";

    public static final String FUNC_NUMSIGNERS = "numSigners";

    public static final String FUNC_OWNER = "owner";

    public static final String FUNC_REMOVESIGNER = "removeSigner";

    public static final String FUNC_REMOVESIGNERSETTHRESHOLD = "removeSignerSetThreshold";

    public static final String FUNC_RENOUNCEOWNERSHIP = "renounceOwnership";

    public static final String FUNC_SETTHRESHOLD = "setThreshold";

    public static final String FUNC_TRANSFEROWNERSHIP = "transferOwnership";

    public static final String FUNC_VERIFY = "verify";

    public static final String FUNC_VERIFY2 = "verify2";

    public static final String FUNC_VERIFYANDCHECKTHRESHOLD = "verifyAndCheckThreshold";

    public static final Event DUMP_EVENT = new Event("Dump", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
    ;

    public static final Event DUMP2_EVENT = new Event("Dump2", 
            Arrays.<TypeReference<?>>asList(new TypeReference<DynamicBytes>() {}));
    ;

    public static final Event OWNERSHIPTRANSFERRED_EVENT = new Event("OwnershipTransferred", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}));
    ;

    @Deprecated
    protected AttestorSignRegistrar(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected AttestorSignRegistrar(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected AttestorSignRegistrar(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected AttestorSignRegistrar(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public List<DumpEventResponse> getDumpEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(DUMP_EVENT, transactionReceipt);
        ArrayList<DumpEventResponse> responses = new ArrayList<DumpEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            DumpEventResponse typedResponse = new DumpEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._Log = (String) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<DumpEventResponse> dumpEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new Function<Log, DumpEventResponse>() {
            @Override
            public DumpEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(DUMP_EVENT, log);
                DumpEventResponse typedResponse = new DumpEventResponse();
                typedResponse.log = log;
                typedResponse._Log = (String) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<DumpEventResponse> dumpEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(DUMP_EVENT));
        return dumpEventFlowable(filter);
    }

    public List<Dump2EventResponse> getDump2Events(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(DUMP2_EVENT, transactionReceipt);
        ArrayList<Dump2EventResponse> responses = new ArrayList<Dump2EventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            Dump2EventResponse typedResponse = new Dump2EventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.data = (byte[]) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<Dump2EventResponse> dump2EventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new Function<Log, Dump2EventResponse>() {
            @Override
            public Dump2EventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(DUMP2_EVENT, log);
                Dump2EventResponse typedResponse = new Dump2EventResponse();
                typedResponse.log = log;
                typedResponse.data = (byte[]) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<Dump2EventResponse> dump2EventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(DUMP2_EVENT));
        return dump2EventFlowable(filter);
    }

    public List<OwnershipTransferredEventResponse> getOwnershipTransferredEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(OWNERSHIPTRANSFERRED_EVENT, transactionReceipt);
        ArrayList<OwnershipTransferredEventResponse> responses = new ArrayList<OwnershipTransferredEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            OwnershipTransferredEventResponse typedResponse = new OwnershipTransferredEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse.previousOwner = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse.newOwner = (String) eventValues.getIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<OwnershipTransferredEventResponse> ownershipTransferredEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new Function<Log, OwnershipTransferredEventResponse>() {
            @Override
            public OwnershipTransferredEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(OWNERSHIPTRANSFERRED_EVENT, log);
                OwnershipTransferredEventResponse typedResponse = new OwnershipTransferredEventResponse();
                typedResponse.log = log;
                typedResponse.previousOwner = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse.newOwner = (String) eventValues.getIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<OwnershipTransferredEventResponse> ownershipTransferredEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(OWNERSHIPTRANSFERRED_EVENT));
        return ownershipTransferredEventFlowable(filter);
    }

    public RemoteFunctionCall<TransactionReceipt> addSigner(BigInteger _bcId, String _signer) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_ADDSIGNER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_bcId), 
                new org.web3j.abi.datatypes.Address(160, _signer)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> addSignerSetThreshold(BigInteger _bcId, String _signer, BigInteger _newSigningThreshold) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_ADDSIGNERSETTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_bcId), 
                new org.web3j.abi.datatypes.Address(160, _signer), 
                new org.web3j.abi.datatypes.generated.Uint256(_newSigningThreshold)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> addSignersSetThreshold(BigInteger _bcId, List<String> _signers, BigInteger _newSigningThreshold) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_ADDSIGNERSSETTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_bcId), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.Address>(
                        org.web3j.abi.datatypes.Address.class,
                        org.web3j.abi.Utils.typeMap(_signers, org.web3j.abi.datatypes.Address.class)), 
                new org.web3j.abi.datatypes.generated.Uint256(_newSigningThreshold)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> checkThreshold(BigInteger _blockchainId, List<String> _signers) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_CHECKTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.Address>(
                        org.web3j.abi.datatypes.Address.class,
                        org.web3j.abi.Utils.typeMap(_signers, org.web3j.abi.datatypes.Address.class))), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<BigInteger> getSigningThreshold(BigInteger _blockchainId) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_GETSIGNINGTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<Boolean> isSigner(BigInteger _blockchainId, String _mightBeSigner) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_ISSIGNER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId), 
                new org.web3j.abi.datatypes.Address(160, _mightBeSigner)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<BigInteger> numSigners(BigInteger _blockchainId) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_NUMSIGNERS, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteFunctionCall<String> owner() {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_OWNER, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Address>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteFunctionCall<TransactionReceipt> removeSigner(BigInteger _bcId, String _signer) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_REMOVESIGNER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_bcId), 
                new org.web3j.abi.datatypes.Address(160, _signer)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> removeSignerSetThreshold(BigInteger _bcId, String _signer, BigInteger _newSigningThreshold) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_REMOVESIGNERSETTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_bcId), 
                new org.web3j.abi.datatypes.Address(160, _signer), 
                new org.web3j.abi.datatypes.generated.Uint256(_newSigningThreshold)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> renounceOwnership() {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_RENOUNCEOWNERSHIP, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> setThreshold(BigInteger _bcId, BigInteger _newSigningThreshold) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_SETTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_bcId), 
                new org.web3j.abi.datatypes.generated.Uint256(_newSigningThreshold)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<TransactionReceipt> transferOwnership(String newOwner) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_TRANSFEROWNERSHIP, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, newOwner)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> verify(BigInteger _blockchainId, List<String> _signers, List<byte[]> _sigR, List<byte[]> _sigS, List<BigInteger> _sigV, byte[] _plainText) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_VERIFY, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.Address>(
                        org.web3j.abi.datatypes.Address.class,
                        org.web3j.abi.Utils.typeMap(_signers, org.web3j.abi.datatypes.Address.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Bytes32>(
                        org.web3j.abi.datatypes.generated.Bytes32.class,
                        org.web3j.abi.Utils.typeMap(_sigR, org.web3j.abi.datatypes.generated.Bytes32.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Bytes32>(
                        org.web3j.abi.datatypes.generated.Bytes32.class,
                        org.web3j.abi.Utils.typeMap(_sigS, org.web3j.abi.datatypes.generated.Bytes32.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Uint8>(
                        org.web3j.abi.datatypes.generated.Uint8.class,
                        org.web3j.abi.Utils.typeMap(_sigV, org.web3j.abi.datatypes.generated.Uint8.class)), 
                new org.web3j.abi.datatypes.DynamicBytes(_plainText)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteFunctionCall<TransactionReceipt> verify2(BigInteger _blockchainId, List<String> _signers, List<byte[]> _sigR, List<byte[]> _sigS, List<BigInteger> _sigV, byte[] _plainText) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_VERIFY2, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.Address>(
                        org.web3j.abi.datatypes.Address.class,
                        org.web3j.abi.Utils.typeMap(_signers, org.web3j.abi.datatypes.Address.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Bytes32>(
                        org.web3j.abi.datatypes.generated.Bytes32.class,
                        org.web3j.abi.Utils.typeMap(_sigR, org.web3j.abi.datatypes.generated.Bytes32.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Bytes32>(
                        org.web3j.abi.datatypes.generated.Bytes32.class,
                        org.web3j.abi.Utils.typeMap(_sigS, org.web3j.abi.datatypes.generated.Bytes32.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Uint8>(
                        org.web3j.abi.datatypes.generated.Uint8.class,
                        org.web3j.abi.Utils.typeMap(_sigV, org.web3j.abi.datatypes.generated.Uint8.class)), 
                new org.web3j.abi.datatypes.DynamicBytes(_plainText)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteFunctionCall<Boolean> verifyAndCheckThreshold(BigInteger _blockchainId, List<String> _signers, List<byte[]> _sigR, List<byte[]> _sigS, List<BigInteger> _sigV, byte[] _plainText) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(FUNC_VERIFYANDCHECKTHRESHOLD, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_blockchainId), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.Address>(
                        org.web3j.abi.datatypes.Address.class,
                        org.web3j.abi.Utils.typeMap(_signers, org.web3j.abi.datatypes.Address.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Bytes32>(
                        org.web3j.abi.datatypes.generated.Bytes32.class,
                        org.web3j.abi.Utils.typeMap(_sigR, org.web3j.abi.datatypes.generated.Bytes32.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Bytes32>(
                        org.web3j.abi.datatypes.generated.Bytes32.class,
                        org.web3j.abi.Utils.typeMap(_sigS, org.web3j.abi.datatypes.generated.Bytes32.class)), 
                new org.web3j.abi.datatypes.DynamicArray<org.web3j.abi.datatypes.generated.Uint8>(
                        org.web3j.abi.datatypes.generated.Uint8.class,
                        org.web3j.abi.Utils.typeMap(_sigV, org.web3j.abi.datatypes.generated.Uint8.class)), 
                new org.web3j.abi.datatypes.DynamicBytes(_plainText)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    @Deprecated
    public static AttestorSignRegistrar load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new AttestorSignRegistrar(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static AttestorSignRegistrar load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new AttestorSignRegistrar(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static AttestorSignRegistrar load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new AttestorSignRegistrar(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static AttestorSignRegistrar load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new AttestorSignRegistrar(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<AttestorSignRegistrar> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(AttestorSignRegistrar.class, web3j, credentials, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<AttestorSignRegistrar> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(AttestorSignRegistrar.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    public static RemoteCall<AttestorSignRegistrar> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return deployRemoteCall(AttestorSignRegistrar.class, web3j, transactionManager, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<AttestorSignRegistrar> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(AttestorSignRegistrar.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    public static class DumpEventResponse extends BaseEventResponse {
        public String _Log;
    }

    public static class Dump2EventResponse extends BaseEventResponse {
        public byte[] data;
    }

    public static class OwnershipTransferredEventResponse extends BaseEventResponse {
        public String previousOwner;

        public String newOwner;
    }
}
