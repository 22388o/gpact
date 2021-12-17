/*
Package eth - Message dispatcher for Ethereum Clients.
*/
package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	log "github.com/consensys/gpact/messaging/relayer/internal/logging"

	//gethcommon "github.com/ethereum/go-ethereum/common"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
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

// KeyManager holds one or more Ethereum keys, or manages access to the keys.
// TODO: For the moment, just use one key pair.
type KeyManager struct {
	keys          []EthKey
	current       int
	msgDispatcher *MsgDispatcher
}

type EthKey struct {
	privKey *ecdsa.PrivateKey
	nonce   uint64
}

// NewKeyManager creates a new key manager.
func NewKeyManager(m *MsgDispatcher) *KeyManager {
	var k = KeyManager{}
	k.msgDispatcher = m
	return &k
}

// AddKey adds a key that can be used by the key manager.
func (k *KeyManager) AddKey(privKeyBytes []byte) (int, error) {
	privKey, err := gethcrypto.ToECDSA(privKeyBytes)
	if err != nil {
		return 0, err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = fmt.Errorf("error casting public key to ECDSA: %v", err.Error())
		log.Fatal("error casting public key to ECDSA")
		return 0, err
	}

	addr := gethcrypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := k.msgDispatcher.connection.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	key := EthKey{privKey, nonce}
	k.keys = append(k.keys, key)

	return len(k.keys), nil

}

// GetNextKeyAndIncNonce returns the next private key to use and increments
// nonce value for the key, so it will be ready for next time.
func (k *KeyManager) GetNextKeyAndIncNonce() (*ecdsa.PrivateKey, uint64) {
	// TODO this operation needs locking around it.

	k.current = (k.current + 1) % len(k.keys)
	key := &k.keys[k.current]
	oldNonce := key.nonce
	key.nonce++

	return key.privKey, oldNonce
}
