package main

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
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/consensys/gpact/messaging/relayer/internal/crypto"
	"github.com/consensys/gpact/messaging/relayer/internal/mqserver"
	"github.com/consensys/gpact/messaging/relayer/internal/soliditywrappers/application"
	"github.com/consensys/gpact/messaging/relayer/internal/soliditywrappers/erc20"
	"github.com/consensys/gpact/messaging/relayer/internal/soliditywrappers/eventstore"
	"github.com/consensys/gpact/messaging/relayer/internal/soliditywrappers/registrar"
	"github.com/consensys/gpact/messaging/relayer/internal/soliditywrappers/sfc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
)

// TODO, Need a separate package to put all core components.
var s *mqserver.MQServer

func createUser() ([]byte, *bind.TransactOpts) {
	key, err := crypto.Secp256k1GenerateKey()
	if err != nil {
		panic(err)
	}
	// Create the same key
	x, y := secp256k1.S256().ScalarBaseMult(key)
	prv := &ecdsa.PrivateKey{}
	prv.D = big.NewInt(0).SetBytes(key)
	prv.PublicKey = ecdsa.PublicKey{
		X:     x,
		Y:     y,
		Curve: secp256k1.S256(),
	}
	auth := bind.NewKeyedTransactor(prv)
	auth.Nonce = big.NewInt(int64(0))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(0)
	return key, auth
}

func main() {
	chainA, err := ethclient.Dial("ws://127.0.0.1:8311")
	if err != nil {
		panic(err)
	}
	defer chainA.Close()

	chainB, err := ethclient.Dial("ws://127.0.0.1:8321")
	if err != nil {
		panic(err)
	}
	defer chainB.Close()

	// Create a super admin user
	_, admin := createUser()
	// Create a relayer signer
	signer, relayer := createUser()
	// User1 in chainA
	_, user1 := createUser()
	// User2 in chainB
	_, user2 := createUser()

	var tx *types.Transaction
	// Deploy erc20 contracts on both chains
	fmt.Println("Deploy ERC20 contract on chainA...")
	erc20AddrA, tx, ierc20A, err := erc20.DeployERC20PresetFixedSupply(admin, chainA, "chainA", "ca", big.NewInt(10000), admin.From)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Hash())
	waitForReceipt(chainA, tx)
	fmt.Println(erc20AddrA)

	fmt.Println("Deploy ERC20 contract on chainB...")
	erc20AddrB, tx, ierc20B, err := erc20.DeployERC20PresetFixedSupply(admin, chainB, "chainB", "ca", big.NewInt(10000), admin.From)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.Hash())
	waitForReceipt(chainB, tx)
	fmt.Println(erc20AddrB)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Deploy sfc contracts on both chains
	fmt.Println("Deploy SFC contract on chainA...")
	sfcAddrA, tx, isfcA, err := sfc.DeploySfc(admin, chainA, big.NewInt(31), big.NewInt(100000))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	fmt.Println(sfcAddrA)

	fmt.Println("Deploy SFC contract on chainB...")
	sfcAddrB, tx, isfcB, err := sfc.DeploySfc(admin, chainB, big.NewInt(32), big.NewInt(100000))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println(sfcAddrB)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Deploy registrar on both chains
	fmt.Println("Deploy registrar on chainA...")
	regAddrA, tx, iregA, err := registrar.DeployRegistrar(admin, chainA)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	fmt.Println(regAddrA)

	fmt.Println("Deploy registrar on chainB...")
	regAddrB, tx, iregB, err := registrar.DeployRegistrar(admin, chainB)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println(regAddrB)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Deploy verifier on both chains
	fmt.Println("Deploy verifier on chainA...")
	verifierAddrA, tx, _, err := eventstore.DeployEventStore(admin, chainA, regAddrA, sfcAddrA)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	fmt.Println(verifierAddrA)

	fmt.Println("Deploy verifier on chainB...")
	verifierAddrB, tx, iverifierB, err := eventstore.DeployEventStore(admin, chainB, regAddrB, sfcAddrB)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println(verifierAddrB)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Add relayer
	fmt.Println("Add signer on chainA...")
	tx, err = iregA.AddSignerSetThreshold(admin, big.NewInt(32), relayer.From, big.NewInt(1))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	fmt.Println("Done")

	fmt.Println("Add signer on chainB...")
	tx, err = iregB.AddSignerSetThreshold(admin, big.NewInt(31), relayer.From, big.NewInt(1))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println("Done")

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Add verifier
	tx, err = isfcA.AddVerifier(admin, big.NewInt(32), verifierAddrA)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)

	tx, err = isfcB.AddVerifier(admin, big.NewInt(31), verifierAddrB)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	tx, err = isfcA.AddRemoteCrosschainControl(admin, big.NewInt(32), sfcAddrB)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)

	tx, err = isfcB.AddRemoteCrosschainControl(admin, big.NewInt(31), sfcAddrA)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Deploy bridge
	fmt.Println("Deploy bridge on chainA...")
	bridgeAddrA, tx, ibridgeA, err := application.DeploySfcERC20Bridge(admin, chainA, sfcAddrA)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	fmt.Println(bridgeAddrA)

	fmt.Println("Deploy bridge on chainB...")
	bridgeAddrB, tx, ibridgeB, err := application.DeploySfcERC20Bridge(admin, chainB, sfcAddrB)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println(bridgeAddrB)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Setup bridge
	fmt.Println("Change blockchain mapping...")
	tx, err = ibridgeA.ChangeBlockchainMapping(admin, big.NewInt(32), bridgeAddrB)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	tx, err = ibridgeB.ChangeBlockchainMapping(admin, big.NewInt(31), bridgeAddrA)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println("Done")

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	fmt.Println("Add contract first mapping...")
	tx, err = ibridgeA.AddContractFirstMapping(admin, erc20AddrA, big.NewInt(32), erc20AddrB, true)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	tx, err = ibridgeB.AddContractFirstMapping(admin, erc20AddrB, big.NewInt(31), erc20AddrA, true)
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)
	fmt.Println("Done")

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Transfer some tokens to user1.
	fmt.Println("Transfer token to user1...")
	tx, err = ierc20A.Transfer(admin, user1.From, big.NewInt(50))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)
	fmt.Println("Done")

	tx, err = ierc20B.Transfer(admin, bridgeAddrB, big.NewInt(1000))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainB, tx)

	admin.Nonce.Add(admin.Nonce, big.NewInt(1))

	// Check balance before
	balA, _ := ierc20A.BalanceOf(nil, user1.From)
	balB, _ := ierc20B.BalanceOf(nil, user2.From)
	fmt.Println("Before: A - ", balA, " B - ", balB)

	// User1 do crosschain transfer to user2.
	fmt.Println("Crosschain transfer...")
	tx, err = ierc20A.Approve(user1, bridgeAddrA, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	waitForReceipt(chainA, tx)

	user1.Nonce.Add(user1.Nonce, big.NewInt(1))

	tx, err = ibridgeA.TransferToOtherBlockchain(user1, big.NewInt(32), erc20AddrA, user2.From, big.NewInt(10))
	if err != nil {
		panic(err)
	}

	sink := make(chan *sfc.SfcCrossCall)
	_, err = isfcA.WatchCrossCall(&bind.WatchOpts{Start: nil, Context: nil}, sink)
	if err != nil {
		panic(err)
	}
	event := <-sink
	waitForReceipt(chainA, tx)
	fmt.Println("Done...")
	user1.Nonce.Add(user1.Nonce, big.NewInt(1))

	// fmt.Println("Event is ", event)

	pc := make([]byte, 32)
	pc[31] = 31
	pc = append(pc, sfcAddrA.Bytes()...)
	pc = append(pc, []byte{89, 247, 54, 220, 94, 21, 196, 177, 37, 38, 72, 117, 2, 100, 84, 3, 176, 164, 49, 109, 130, 235, 167, 233, 236, 220, 42, 5, 12, 16, 173, 39}...)
	pc = append(pc, event.Raw.Data...)

	signature, err := crypto.Secp256k1Sign(signer, pc)
	if err != nil {
		panic(err)
	}
	signature[len(signature)-1] += 27
	fmt.Println(signature)

	tx, err = iverifierB.RelayEvent(relayer, big.NewInt(31), sfcAddrA, event.Raw.Data, append([]byte{0, 0, 0, 1}, append(relayer.From.Bytes(), signature...)...))
	if err != nil {
		panic(err)
	}
	sink2 := make(chan *eventstore.EventStoreDump)
	_, err = iverifierB.WatchDump(&bind.WatchOpts{Start: nil, Context: nil}, sink2)
	if err != nil {
		panic(err)
	}

	sink3 := make(chan *sfc.SfcDump)
	_, err = isfcB.WatchDump(&bind.WatchOpts{Start: nil, Context: nil}, sink3)
	if err != nil {
		panic(err)
	}

	end := make(chan bool)
	go func() {
		// fmt.Println("Start watching...")
		for {
			select {
			case log := <-sink2:
				fmt.Println("DEBUG1: ", log.Log)
			case log := <-sink3:
				fmt.Println("DEBUG3: ", log.Log)
			case <-end:
				fmt.Println("exit")
				return
			}
		}
	}()
	waitForReceipt(chainB, tx)
	fmt.Println("Done")

	balA, _ = ierc20A.BalanceOf(nil, user1.From)
	balB, _ = ierc20B.BalanceOf(nil, user2.From)
	fmt.Println("After: A - ", balA, " B - ", balB)

	end <- true
	time.Sleep(1 * time.Second)
}

func waitForReceipt(conn *ethclient.Client, tx *types.Transaction) error {
	for {
		rept, err := conn.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil {
			if rept.Status != types.ReceiptStatusSuccessful {
				fmt.Println("Transaction failed... ", rept)
				return nil
			} else {
			}
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
