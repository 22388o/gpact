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
package net.consensys.gpact.examples.sfc.erc20bridge;

import java.math.BigInteger;
import net.consensys.gpact.common.*;
import net.consensys.gpact.functioncall.CrossControlManagerGroup;
import net.consensys.gpact.helpers.CredentialsCreator;
import net.consensys.gpact.helpers.SfcExampleSystemManager;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.web3j.crypto.Credentials;

/**
 * Sample code showing how to use the Simple Function Call protocol ERC 20 Mass Conservation and
 * Minter Burner bridges.
 */
public class ERC20TokenBridgeExample {
  static final Logger LOG = LogManager.getLogger(ERC20TokenBridgeExample.class);

  public static final boolean BLOCKCHAIN_B_MASS_CONSERVATION = false;

  public static final int NUM_TIMES_EXECUTE = 2;

  public static void main(String[] args) throws Exception {
    main(args, BLOCKCHAIN_B_MASS_CONSERVATION);
  }

  public static void main(String[] args, boolean blockchainBIsMassConservation) throws Exception {
    StatsHolder.log("Example: SFC: Token Bridge");
    LOG.info("Started");
    LOG.info("Blockchain B is Mass Conservation: {}", blockchainBIsMassConservation);

    if (args.length != 1) {
      LOG.info("Usage: [properties file name]");
      return;
    }

    SfcExampleSystemManager exampleManager = new SfcExampleSystemManager(args[0]);
    exampleManager.sfcStandardExampleConfig(2);

    BlockchainConfig root = exampleManager.getRootBcInfo();
    BlockchainConfig bc2 = exampleManager.getBc2Info();
    CrossControlManagerGroup crossControlManagerGroup =
        exampleManager.getCrossControlManagerGroup();

    final int CHAIN_A_TOKEN_SUPPLY = 1000;
    final int CHAIN_B_TOKEN_SUPPLY = 1000;

    // Set-up classes to manage blockchains.
    Credentials erc20OwnerCreds = CredentialsCreator.createCredentials();
    MassConservationERC20Bridge chainA =
        new MassConservationERC20Bridge(
            "ChainA",
            BigInteger.valueOf(CHAIN_A_TOKEN_SUPPLY),
            erc20OwnerCreds,
            root.bcId,
            root.uri,
            root.gasPriceStrategy,
            root.period);
    AbstractERC20Bridge chainB;
    if (blockchainBIsMassConservation) {
      chainB =
          new MassConservationERC20Bridge(
              "ChainB",
              BigInteger.valueOf(CHAIN_B_TOKEN_SUPPLY),
              erc20OwnerCreds,
              bc2.bcId,
              bc2.uri,
              bc2.gasPriceStrategy,
              bc2.period);
    } else {
      chainB =
          new MinterBurnerERC20Bridge(
              "ChainB",
              BigInteger.valueOf(CHAIN_B_TOKEN_SUPPLY),
              erc20OwnerCreds,
              bc2.bcId,
              bc2.uri,
              bc2.gasPriceStrategy,
              bc2.period);
    }

    // Deploy application contracts.
    BlockchainId chainABcId = chainA.getBlockchainId();
    chainA.deployContracts(crossControlManagerGroup.getCbcAddress(chainABcId));
    BlockchainId chainBBcId = chainB.getBlockchainId();
    chainB.deployContracts(crossControlManagerGroup.getCbcAddress(chainBBcId));

    // Connect the ERC 20 Bridges
    chainA.addBlockchain(chainBBcId, chainB.getErc20BridgeContractAddress());
    chainB.addBlockchain(chainABcId, chainA.getErc20BridgeContractAddress());

    // Register the ERC20 contracts with each blockchain.
    chainA.addRemoteERC20(chainBBcId, chainB.getErc20ContractAddress());
    chainB.addRemoteERC20(chainABcId, chainA.getErc20ContractAddress());

    // Create some users and give them some tokens.
    Erc20User user1 =
        new Erc20User(
            "User1",
            root.bcId,
            chainA.getErc20ContractAddress(),
            chainA.getErc20BridgeContractAddress(),
            bc2.bcId,
            chainB.getErc20ContractAddress(),
            chainA.getErc20BridgeContractAddress());
    Erc20User user2 =
        new Erc20User(
            "User2",
            root.bcId,
            chainA.getErc20ContractAddress(),
            chainA.getErc20BridgeContractAddress(),
            bc2.bcId,
            chainB.getErc20ContractAddress(),
            chainA.getErc20BridgeContractAddress());
    Erc20User user3 =
        new Erc20User(
            "User3",
            root.bcId,
            chainA.getErc20ContractAddress(),
            chainA.getErc20BridgeContractAddress(),
            bc2.bcId,
            chainB.getErc20ContractAddress(),
            chainA.getErc20BridgeContractAddress());

    user1.createCbcManager(
        root,
        crossControlManagerGroup.getCbcAddress(chainABcId),
        crossControlManagerGroup.getMessageVerification(chainABcId),
        bc2,
        crossControlManagerGroup.getCbcAddress(chainBBcId),
        crossControlManagerGroup.getMessageVerification(chainBBcId));
    user2.createCbcManager(
        root,
        crossControlManagerGroup.getCbcAddress(chainABcId),
        crossControlManagerGroup.getMessageVerification(chainABcId),
        bc2,
        crossControlManagerGroup.getCbcAddress(chainBBcId),
        crossControlManagerGroup.getMessageVerification(chainBBcId));
    user3.createCbcManager(
        root,
        crossControlManagerGroup.getCbcAddress(chainABcId),
        crossControlManagerGroup.getMessageVerification(chainABcId),
        bc2,
        crossControlManagerGroup.getCbcAddress(chainBBcId),
        crossControlManagerGroup.getMessageVerification(chainBBcId));

    // Give some balance to the users
    chainA.giveTokens(user1, 500);
    chainA.giveTokens(user2, 200);
    chainA.giveTokens(user3, 300);

    if (blockchainBIsMassConservation) {
      chainB.giveTokensToERC20Bridge(1000);
    }

    Erc20User[] users = new Erc20User[] {user1, user2, user3};

    chainA.showErc20Balances(users);
    chainB.showErc20Balances(users);

    for (int numExecutions = 0; numExecutions < NUM_TIMES_EXECUTE; numExecutions++) {
      LOG.info("Execution: {}  *****************", numExecutions);
      StatsHolder.log("Execution: " + numExecutions + " **************************");

      user1.transfer(true, 20);
      user2.transfer(true, user2.getAddress(), 30);
      user3.transfer(true, user2.getAddress(), 10);

      user2.transfer(false, 39);
      user1.transfer(false, 18);

      chainA.showErc20Balances(users);
      chainB.showErc20Balances(users);
    }

    chainA.shutdown();
    chainB.shutdown();

    StatsHolder.log("End");
    StatsHolder.print();
  }
}
