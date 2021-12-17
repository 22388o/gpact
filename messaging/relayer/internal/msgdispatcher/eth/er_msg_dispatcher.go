/*
Package eventrelay - Message dispatcher for Ethereum Clients for Event Relay messaging.
*/
package eth

import (
	//"context"
	// "fmt"
	"math/big"

	log "github.com/consensys/gpact/messaging/relayer/internal/logging"
	"github.com/consensys/gpact/messaging/relayer/internal/soliditywrappers/eventstore"

	gethbind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	// gethtypes "github.com/ethereum/go-ethereum/core/types"
	// //gethcrypto "github.com/ethereum/go-ethereum/crypto"
	// gethclient "github.com/ethereum/go-ethereum/ethclient"
	// gethrpc "github.com/ethereum/go-ethereum/rpc"
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

type EventRelayMsgDispatcher struct {
	m          *MsgDispatcher
	eventStore *eventstore.Eventstore
}

// NewMsgDispatcher creates a new message dispatcher instance.
func NewMEventRelayMsgDispatcher(m *MsgDispatcher) *EventRelayMsgDispatcher {
	var e = EventRelayMsgDispatcher{}
	e.m = m
	return &e
}

func (e *EventRelayMsgDispatcher) Connect() {
	// TODO do something with the error!
	eventStoreContract, _ := eventstore.NewEventstore(*(e.m.EventStoreAddress), e.m.Connection)
	e.eventStore = eventStoreContract

}

func (e *EventRelayMsgDispatcher) RelaySignedEvent(
	sourceBcId big.Int,
	cbcAddress gethcommon.Address,
	eventData []byte,
	encodedSignatureOrProof []byte) error {

	privKey, nonce := e.m.KeyManager.GetNextKeyAndIncNonce()

	// 	tx := gethtypes.NewTx(&gethtypes.DynamicFeeTx{
	// 		ChainID:   m.chainID,
	// 		Nonce:     nonce,
	// 		GasFeeCap: feeCap,
	// 		GasTipCap: tip,
	// 		Gas:       gas,
	// 		To:        m.eventStoreAddress,
	// 		Value:     value,
	// 		Data:      txData,
	// 	})

	// 	signedTx, err := gethtypes.SignTx(tx, gethtypes.NewEIP155Signer(m.chainID), privKey)

	auth := gethbind.NewKeyedTransactor(privKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = big.NewInt(0)

	// Submit event information to contract.
	txReceipt, err := e.eventStore.RelayEvent(auth, &sourceBcId, cbcAddress, eventData, encodedSignatureOrProof)
	log.Info("Err: %s", err.Error())
	log.Info("tx receipt: %s", txReceipt.Hash().Hex())

	return err

}
