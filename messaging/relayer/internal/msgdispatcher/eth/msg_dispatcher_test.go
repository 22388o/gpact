/*
Package eth - Message dispatcher for Ethereum Clients.
*/
package eth

import (
	//	"crypto"
	"testing"

	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	//	log "github.com/consensys/gpact/messaging/relayer/internal/logging"

	//	"github.com/stretchr/testify/assert"

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
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := gethcrypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := gethcrypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := gethbind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployMsgdispatchertest(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance

	// Submit a transaction that calls a function.

	// Check that the transaction has executed by checking the contract state.

}
