/*
Package eth - Message dispatcher for Ethereum Clients.
*/
package eth

import (
	"context"
	"fmt"
	"math/big"

	log "github.com/consensys/gpact/messaging/relayer/internal/logging"

	//gethbind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	//gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethclient "github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"
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

// MsgDispatcher holds the context for submitting transactions
// to an Ethereum Client.
type MsgDispatcher struct {
	endpoint                  string      // URL without protocol specifier of Ethereum client.
	http                      bool        // HTTP or WS
	apiAuthKey                string      // Authentication key to access the Ethereum API.
	keyManager                *KeyManager // Holds all keys for this dispatcher.
	crosschainControlContract *gethcommon.Address
	blockchainType            ChainType

	chainID    *big.Int
	connection *gethclient.Client
}

// MsgDispatcherConfig holds variables needed to configure the message
// dispatcher component.
type MsgDispatcherConfig struct {
	endpoint                  string // URL without protocol specifier of Ethereum client.
	http                      bool   // HTTP or WS
	apiAuthKey                string // Authentication key to access the Ethereum API.
	crosschainControlContract *gethcommon.Address
	blockchainType            ChainType
}

// NewMsgDispatcher creates a new message dispatcher instance.
func NewMsgDispatcher(c *MsgDispatcherConfig) (*MsgDispatcher, error) {
	var m = MsgDispatcher{}
	m.endpoint = c.endpoint
	m.http = c.http
	m.apiAuthKey = c.apiAuthKey
	m.crosschainControlContract = c.crosschainControlContract
	m.blockchainType = c.blockchainType

	m.keyManager = NewKeyManager(&m)

	log.Info("Message Dispatcher (Eth) for %v", c.endpoint)

	return &m, nil
}

func (m *MsgDispatcher) GetKeyManager() (keyManager *KeyManager) {
	return m.keyManager
}

// Connect attempts to use the configuration to connect to the end point.
func (m *MsgDispatcher) Connect() error {
	log.Info("Connecting to: %v", m.endpoint)
	var rpcClient *gethrpc.Client
	var err error
	// Start http or ws client
	if m.http {
		rpcClient, err = gethrpc.DialHTTP(m.endpoint)
	} else {
		rpcClient, err = gethrpc.DialContext(context.Background(), m.endpoint)
	}
	if err != nil {
		return err
	}
	m.connection = gethclient.NewClient(rpcClient)

	// Derive information from the blockchain.
	m.chainID, err = m.connection.NetworkID(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (m *MsgDispatcher) SubmitTransaction(txData []byte) (gethcommon.Hash, error) {
	switch m.blockchainType {
	case ZeroCostConsortium:
		return m.submitFreeConsortiumTransaction(txData)
	case PublicEip1559:
		return m.submitEIP1559Transaction(txData)
	default:
		log.Panic("Unknown blockhcain type: %s", m.blockchainType.String())
		empty := make([]byte, gethcommon.HashLength)
		return gethcommon.BytesToHash(empty), fmt.Errorf("Unknown blockhcain type: %s", m.blockchainType.String())
	}
}

// submitEIP1559Transaction submits a transaction to an EIP1559 blockchain. Return the trransaction hash.
func (m *MsgDispatcher) submitEIP1559Transaction(txData []byte) (gethcommon.Hash, error) {

	feeCap := big.NewInt(10000000)
	tip := big.NewInt(10000000)

	gas := uint64(100000)
	value := big.NewInt(0)
	//txData := make([]byte, 0)

	privKey, nonce := m.keyManager.GetNextKeyAndIncNonce()

	tx := gethtypes.NewTx(&gethtypes.DynamicFeeTx{
		ChainID:   m.chainID,
		Nonce:     nonce,
		GasFeeCap: feeCap,
		GasTipCap: tip,
		Gas:       gas,
		To:        m.crosschainControlContract,
		Value:     value,
		Data:      txData,
	})

	signedTx, err := gethtypes.SignTx(tx, gethtypes.NewEIP155Signer(m.chainID), privKey)
	if err != nil {
		empty := make([]byte, gethcommon.HashLength)
		return gethcommon.BytesToHash(empty), err
	}

	err = m.connection.SendTransaction(context.Background(), signedTx)

	return signedTx.Hash(), err
}

// submitFreeConsortiumTransaction submits a transaction to a pre-EIP1559 blockchain, and assumes a
// gas cost of zero. Return the trransaction hash.
func (m *MsgDispatcher) submitFreeConsortiumTransaction(txData []byte) (gethcommon.Hash, error) {

	gasPrice := big.NewInt(0)
	gasLimit := uint64(0)
	value := big.NewInt(0)
	//txData := make([]byte, 0)

	privKey, nonce := m.keyManager.GetNextKeyAndIncNonce()

	tx := gethtypes.NewTx(&gethtypes.LegacyTx{
		Nonce:    nonce,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		To:       m.crosschainControlContract,
		Data:     txData,
	})

	signedTx, err := gethtypes.SignTx(tx, gethtypes.NewEIP155Signer(m.chainID), privKey)
	if err != nil {
		empty := make([]byte, gethcommon.HashLength)
		return gethcommon.BytesToHash(empty), err
	}

	err = m.connection.SendTransaction(context.Background(), signedTx)

	return signedTx.Hash(), err
}
