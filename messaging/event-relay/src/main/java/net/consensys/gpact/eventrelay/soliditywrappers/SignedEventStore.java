package net.consensys.gpact.eventrelay.soliditywrappers;

import io.reactivex.Flowable;
import io.reactivex.functions.Function;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.FunctionEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
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
public class SignedEventStore extends Contract {
    public static final String BINARY = "608060405234801561001057600080fd5b50604051610fc3380380610fc383398101604081905261002f9161007c565b600080546001600160a01b039384166001600160a01b031991821617909155600180549290931691161790556100af565b80516001600160a01b038116811461007757600080fd5b919050565b6000806040838503121561008f57600080fd5b61009883610060565b91506100a660208401610060565b90509250929050565b610f05806100be6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631cd8253a1461003b5780634c1ce90214610050575b600080fd5b61004e610049366004610b0d565b610063565b005b61004e61005e366004610ba5565b61088e565b60608060608060006100aa87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250925061098b915050565b90506100bd605563ffffffff8316610bfa565b6100c8906004610c19565b861461011b5760405162461bcd60e51b815260206004820152601a60248201527f5369676e617475726520696e636f7272656374206c656e67746800000000000060448201526064015b60405180910390fd5b8063ffffffff1667ffffffffffffffff81111561013a5761013a610c31565b604051908082528060200260200182016040528015610163578160200160208202803683370190505b5094508063ffffffff1667ffffffffffffffff81111561018557610185610c31565b6040519080825280602002602001820160405280156101ae578160200160208202803683370190505b5093508063ffffffff1667ffffffffffffffff8111156101d0576101d0610c31565b6040519080825280602002602001820160405280156101f9578160200160208202803683370190505b5092508063ffffffff1667ffffffffffffffff81111561021b5761021b610c31565b604051908082528060200260200182016040528015610244578160200160208202803683370190505b509150600460005b8263ffffffff168110156104325761029b89898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109f1915050565b8782815181106102ad576102ad610c47565b6001600160a01b03909216602092830291909101909101526102d0601483610c19565b915061031389898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109f9915050565b86828151811061032557610325610c47565b60200260200101818152505060208261033e9190610c19565b915061038189898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508692506109f9915050565b85828151811061039357610393610c47565b6020026020010181815250506020826103ac9190610c19565b91506103ef89898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250610a5e915050565b84828151811061040157610401610c47565b60ff9092166020928302919091019091015261041e600183610c19565b91508061042a81610c5d565b91505061024c565b505060008b8b7f59f736dc5e15c4b12526487502645403b0a4316d82eba7e9ecdc2a050c10ad278c8c60405160200161046f959493929190610c78565b60408051601f198184030181529082905260005463ea13ec8b60e01b83529092506001600160a01b03169063ea13ec8b906104b8908f908a908a908a908a908990600401610d6d565b60206040518083038186803b1580156104d057600080fd5b505afa1580156104e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105089190610e16565b507f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9604051610556906020808252600b908201526a76657269667920646f6e6560a81b604082015260600190565b60405180910390a18051602082012060005b8363ffffffff168110156106b2576002600083815260200190815260200160002060030160008983815181106105a0576105a0610c47565b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16156105d1576106a0565b60016002600084815260200190815260200160002060030160008a84815181106105fd576105fd610c47565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055506002600083815260200190815260200160002060020188828151811061066457610664610c47565b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b039092169190911790555b806106aa81610c5d565b915050610568565b5060008181526002602052604090206001015460ff16156106d95750505050505050610886565b60008060009054906101000a90046001600160a01b03166001600160a01b031663a64ce1998f6040518263ffffffff1660e01b815260040161071d91815260200190565b60206040518083038186803b15801561073557600080fd5b505afa158015610749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076d9190610e3f565b60008381526002602081905260409091200154909150811161087d57600082815260026020526040908190206001908101805460ff19169091179055517f922800f9d4416ff1a5be6eba1df0d7e09a89197bd2ad2473ca9b5614794f9ae9906107f7906020808252600d908201526c199d5b98dd1a5bdb8818d85b1b609a1b604082015260600190565b60405180910390a1600160009054906101000a90046001600160a01b03166001600160a01b031663408840528f8f8f8f8f8f6040518763ffffffff1660e01b815260040161084a96959493929190610e81565b600060405180830381600087803b15801561086457600080fd5b505af1158015610878573d6000803e3d6000fd5b505050505b50505050505050505b505050505050565b600084846040516108a0929190610ebf565b6040519081900381206000805463a64ce19960e01b8452600484018b9052919350916001600160a01b039091169063a64ce1999060240160206040518083038186803b1580156108ef57600080fd5b505afa158015610903573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109279190610e3f565b600083815260026020819052604090912001549091508111156109815760405162461bcd60e51b81526020600482015260126024820152714e6f7420656e6f756768207369676e65727360701b6044820152606401610112565b5050505050505050565b6000610998826004610c19565b835110156109e85760405162461bcd60e51b815260206004820152601d60248201527f736c6963696e67206f7574206f662072616e6765202875696e743332290000006044820152606401610112565b50016004015190565b016014015190565b60008060005b6020811015610a5657610a13816008610bfa565b85610a1e8387610c19565b81518110610a2e57610a2e610c47565b01602001516001600160f81b031916901c919091179080610a4e81610c5d565b9150506109ff565b509392505050565b6000610a6b826001610c19565b83511015610abb5760405162461bcd60e51b815260206004820152601c60248201527f736c6963696e67206f7574206f662072616e6765202875696e743829000000006044820152606401610112565b50016001015190565b60008083601f840112610ad657600080fd5b50813567ffffffffffffffff811115610aee57600080fd5b602083019150836020828501011115610b0657600080fd5b9250929050565b60008060008060008060808789031215610b2657600080fd5b8635955060208701356001600160a01b0381168114610b4457600080fd5b9450604087013567ffffffffffffffff80821115610b6157600080fd5b610b6d8a838b01610ac4565b90965094506060890135915080821115610b8657600080fd5b50610b9389828a01610ac4565b979a9699509497509295939492505050565b60008060008060008060808789031215610bbe57600080fd5b8635955060208701359450604087013567ffffffffffffffff80821115610b6157600080fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610c1457610c14610be4565b500290565b60008219821115610c2c57610c2c610be4565b500190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000600019821415610c7157610c71610be4565b5060010190565b8581526bffffffffffffffffffffffff198560601b1660208201528360348201528183605483013760009101605401908152949350505050565b600081518084526020808501945080840160005b83811015610ce257815187529582019590820190600101610cc6565b509495945050505050565b600081518084526020808501945080840160005b83811015610ce257815160ff1687529582019590820190600101610d01565b6000815180845260005b81811015610d4657602081850181015186830182015201610d2a565b81811115610d58576000602083870101525b50601f01601f19169290920160200192915050565b600060c08201888352602060c08185015281895180845260e086019150828b01935060005b81811015610db75784516001600160a01b031683529383019391830191600101610d92565b50508481036040860152610dcb818a610cb2565b925050508281036060840152610de18187610cb2565b90508281036080840152610df58186610ced565b905082810360a0840152610e098185610d20565b9998505050505050505050565b600060208284031215610e2857600080fd5b81518015158114610e3857600080fd5b9392505050565b600060208284031215610e5157600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8681526001600160a01b0386166020820152608060408201819052600090610eac9083018688610e58565b8281036060840152610e09818587610e58565b818382376000910190815291905056fea264697066735822122016ead7e87ff39a26b896148315b0783ef9f38920a9099024fbc5d301089264d964736f6c63430008090033";

    public static final String FUNC_DECODEANDVERIFYEVENT = "decodeAndVerifyEvent";

    public static final String FUNC_RELAYEVENT = "relayEvent";

    public static final Event DUMP_EVENT = new Event("Dump", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
    ;

    @Deprecated
    protected SignedEventStore(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected SignedEventStore(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected SignedEventStore(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected SignedEventStore(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public List<DumpEventResponse> getDumpEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(DUMP_EVENT, transactionReceipt);
        ArrayList<DumpEventResponse> responses = new ArrayList<DumpEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            DumpEventResponse typedResponse = new DumpEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._log = (String) eventValues.getNonIndexedValues().get(0).getValue();
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
                typedResponse._log = (String) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<DumpEventResponse> dumpEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(DUMP_EVENT));
        return dumpEventFlowable(filter);
    }

    public RemoteFunctionCall<TransactionReceipt> relayEvent(BigInteger _sourceBlockchainId, String _sourceCbcAddress, byte[] _eventData, byte[] _signature) {
        final org.web3j.abi.datatypes.Function function = new org.web3j.abi.datatypes.Function(
                FUNC_RELAYEVENT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(_sourceBlockchainId), 
                new org.web3j.abi.datatypes.Address(160, _sourceCbcAddress), 
                new org.web3j.abi.datatypes.DynamicBytes(_eventData), 
                new org.web3j.abi.datatypes.DynamicBytes(_signature)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    @Deprecated
    public static SignedEventStore load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new SignedEventStore(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static SignedEventStore load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new SignedEventStore(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static SignedEventStore load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new SignedEventStore(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static SignedEventStore load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new SignedEventStore(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<SignedEventStore> deploy(Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider, String _registrar, String _functionCall) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, _registrar), 
                new org.web3j.abi.datatypes.Address(160, _functionCall)));
        return deployRemoteCall(SignedEventStore.class, web3j, credentials, contractGasProvider, BINARY, encodedConstructor);
    }

    public static RemoteCall<SignedEventStore> deploy(Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider, String _registrar, String _functionCall) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, _registrar), 
                new org.web3j.abi.datatypes.Address(160, _functionCall)));
        return deployRemoteCall(SignedEventStore.class, web3j, transactionManager, contractGasProvider, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<SignedEventStore> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit, String _registrar, String _functionCall) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, _registrar), 
                new org.web3j.abi.datatypes.Address(160, _functionCall)));
        return deployRemoteCall(SignedEventStore.class, web3j, credentials, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    @Deprecated
    public static RemoteCall<SignedEventStore> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit, String _registrar, String _functionCall) {
        String encodedConstructor = FunctionEncoder.encodeConstructor(Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(160, _registrar), 
                new org.web3j.abi.datatypes.Address(160, _functionCall)));
        return deployRemoteCall(SignedEventStore.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, encodedConstructor);
    }

    public static class DumpEventResponse extends BaseEventResponse {
        public String _log;
    }
}
