/*
Package eth - Message dispatcher for Ethereum Clients.
*/
package eth

import (
	//	"crypto"
	"testing"

	"github.com/stretchr/testify/assert"

	"context"
	"crypto/ecdsa"
	"math/big"

	rcrypto "github.com/consensys/gpact/messaging/relayer/internal/crypto"
	log "github.com/consensys/gpact/messaging/relayer/internal/logging"

	gethbind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethclient "github.com/ethereum/go-ethereum/ethclient"
	//gethrpc "github.com/ethereum/go-ethereum/rpc"
)

/*
 * Copyright 2021 ConsenSys Software Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

func TestFreeConsortiumTransaciton(t *testing.T) {

	// Deploy contract
	client, err := gethclient.Dial("http://127.0.0.1:8310")
	assert.Nil(t, err)

	blockchainId := big.NewInt(31)

	// Use a different account each time to ensure this test execution doesn't clash with
	// any other execution.
	privKeyBytes, err := rcrypto.Secp256k1GenerateKey()
	assert.Nil(t, err)
	privateKey, err := gethcrypto.ToECDSA(privKeyBytes)
	assert.Nil(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	assert.True(t, ok, "error casting public key to ECDSA")

	fromAddress := gethcrypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	assert.Nil(t, err)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	assert.Nil(t, err)

	auth := gethbind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice

	testRegistrarAddress, testRegistrarTxReceipt, registrarInstance, err := DeployTestRegistrar(auth, client)
	assert.Nil(t, err)
	log.Info("Registrar address: %s", testRegistrarAddress.Hex())
	log.Info("Registrar deploy tx receipt: %s", testRegistrarTxReceipt.Hash().Hex())

	nonce++
	auth.Nonce = big.NewInt(int64(nonce))

	sfcTimeHorizon := big.NewInt(1000) // Time horizon in seconds.
	testSfcAddress, testSfcTxReceipt, _ /*testSfcInstance */, err := DeployTestsfc(auth, client, blockchainId, sfcTimeHorizon)
	assert.Nil(t, err)
	log.Info("Sfc address: %s", testSfcAddress.Hex())
	log.Info("Sfc deploy tx receipt: %s", testSfcTxReceipt.Hash().Hex())

	nonce++
	auth.Nonce = big.NewInt(int64(nonce))

	eventStoreAddress, eventStoreTxReceipt, _ /* eventStoreInstance */, err := DeployTestsignedeventstore(auth, client, testRegistrarAddress, testSfcAddress)
	assert.Nil(t, err)
	log.Info("Signed Event Store address: %s", eventStoreAddress.Hex())
	log.Info("Signed Event Store deploy tx receipt: %s", eventStoreTxReceipt.Hash().Hex())

	nonce++
	auth.Nonce = big.NewInt(int64(nonce))

	testMsgDispatcherAddress, testMsgDispatcherTxReceipt, _ /*testMsgDispatcherInstance*/, err := DeployTestmsgdispatcher(auth, client)
	assert.Nil(t, err)
	log.Info("Test contract address: %s", testMsgDispatcherAddress.Hex())
	log.Info("Test contract deploy tx receipt: %s", testMsgDispatcherTxReceipt.Hash().Hex())

	nonce++
	auth.Nonce = big.NewInt(int64(nonce))

	// Configure registrar, SFC, and Signed Event Store contracts.
	sourceBcId := big.NewInt(2)
	// TODO have separate key pair for event signing.
	signerAddress := fromAddress // Reuse the credentials used for deploying the contracts as the one and only event signer.
	signingThreshold := big.NewInt(1)
	registrarInstance.AddSignerSetThreshold(auth, sourceBcId, signerAddress, signingThreshold)

	// Configure the message dispatcher.
	var config = MsgDispatcherConfig{}
	config.endpoint = "http://127.0.0.1:8310"
	config.http = false
	config.apiAuthKey = ""
	config.eventStoreAddress = &eventStoreAddress
	config.blockchainType = ZeroCostConsortium

	msgDispatcher, err := NewMsgDispatcher(&config)
	assert.Nil(t, err)
	// NOTE: Connect before adding keys.
	err = msgDispatcher.Connect()
	assert.Nil(t, err)

	keyManager := msgDispatcher.GetKeyManager()
	keyManager.AddKey(privKeyBytes)

	// Submit a transaction that calls a function.
	// TODO ABI Encode packed with selector:
	// crossCallHandler(
	//     uint256 _sourceBcId,
	//     address _cbcAddress,
	//     bytes calldata _eventData,
	//     bytes calldata _signature
	// where: _eventData is:
	//        bytes32 txId;
	// uint256 timestamp;
	// address caller;
	// uint256 destBcId;
	// address destContract;
	// bytes memory functionCall;
	// (txId, timestamp, caller, destBcId, destContract, functionCall) = abi
	// 	.decode(
	// 		_eventData,
	// 		(bytes32, uint256, address, uint256, address, bytes)
	// and where: _signature is:
	//   uint32 len = BytesUtil.bytesToUint32(_signature, 0);
	// {
	// 	require(
	// 		_signature.length == LEN_OF_LEN + len * LEN_OF_SIG,
	// 		"Signature incorrect length"
	// 	);

	// 	signers = new address[](len);
	// 	sigRs = new bytes32[](len);
	// 	sigSs = new bytes32[](len);
	// 	sigVs = new uint8[](len);

	// 	uint256 offset = LEN_OF_LEN;
	// 	for (uint256 i = 0; i < len; i++) {
	// 		signers[i] = BytesUtil.bytesToAddress2(_signature, offset);
	// 		offset += 20;
	// 		sigRs[i] = BytesUtil.bytesToBytes32(_signature, offset);
	// 		offset += 32;
	// 		sigSs[i] = BytesUtil.bytesToBytes32(_signature, offset);
	// 		offset += 32;
	// 		sigVs[i] = BytesUtil.bytesToUint8(_signature, offset);
	// 		offset += 1;

	// TODO the code below havs problems reported as gas issues, but are probably nonce related.
	// encodedTransaction := make([]byte, 32)
	// hash, err := msgDispatcher.SubmitTransaction(encodedTransaction)
	// assert.Nil(t, err)
	// log.Info("Submitted transaction hash: %s", hash.Hex())

	// Check that the transaction has executed by checking the contract state.

}
