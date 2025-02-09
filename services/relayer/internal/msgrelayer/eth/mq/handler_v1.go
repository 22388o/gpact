package mq

/*
 * Copyright 2021 ConsenSys Software Inc
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

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/consensys/gpact/messaging/relayer/internal/logging"
	"github.com/consensys/gpact/messaging/relayer/internal/messages"
	v1 "github.com/consensys/gpact/messaging/relayer/internal/messages/v1"
	"github.com/consensys/gpact/messaging/relayer/internal/msgrelayer/eth/node"
	"github.com/consensys/gpact/messaging/relayer/internal/msgrelayer/eth/signer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// handleV1 handles message v1 with type v1.
func handleV1(req messages.Message) {
	instance := node.GetSingleInstance()

	msg, ok := req.(*v1.Message)
	if !ok {
		logging.Error("fail to decode message")
		return
	}
	logging.Info("Process message with ID: %v", msg.ID)

	data, err := hex.DecodeString(msg.Payload)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	raw := types.Log{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	srcID, err := strconv.Atoi(msg.Source.NetworkID)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	srcAddr := common.HexToAddress(msg.Source.ContractAddress)
	destID, err := strconv.Atoi(msg.Destination.NetworkID)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	destAddr := common.HexToAddress(msg.Destination.ContractAddress)

	toSign := make([]byte, 32)
	toSign[31] = byte(srcID)
	toSign = append(toSign, srcAddr.Bytes()...)
	toSign = append(toSign, []byte{89, 247, 54, 220, 94, 21, 196, 177, 37, 38, 72, 117, 2, 100, 84, 3, 176, 164, 49, 109, 130, 235, 167, 233, 236, 220, 42, 5, 12, 16, 173, 39}...)
	toSign = append(toSign, raw.Data...)
	logging.Info("Generated data to be signed: %v", hex.EncodeToString(toSign))

	_, addr, err := instance.Signer.GetAddr(big.NewInt(int64(destID)), destAddr)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	sigType, signature, err := instance.Signer.Sign(big.NewInt(int64(destID)), destAddr, toSign)
	if err != nil {
		logging.Error(err.Error())
		return
	}
	logging.Info("Signature generated with type %v: %v", signer.TypeToString(sigType), hex.EncodeToString(signature))

	signature[len(signature)-1] += 27
	// TODO, Add relayer communication to order signatures from different relayers.
	signature = append([]byte{0, 0, 0, 1}, append(addr.Bytes(), signature...)...)

	// Add proof
	msg.Proofs = append(msg.Proofs, v1.Proof{
		ProofType: signer.TypeToString(sigType),
		Created:   time.Now().UnixMilli(),
		Proof:     hex.EncodeToString(signature),
	})

	// Pass message to MQ.
	go instance.MQ.Request(msg.Version, msg.MsgType, msg)
}

// initV1 inits the handler.
func initV1() {
	_, ok := Handlers[v1.Version]
	if !ok {
		Handlers[v1.Version] = make(map[string]func(msg messages.Message))
	}
	Handlers[v1.Version][v1.MessageType] = handleV1
}
