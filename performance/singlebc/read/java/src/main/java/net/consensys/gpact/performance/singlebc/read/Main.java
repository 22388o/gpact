/*
 * Copyright 2019 ConsenSys Software Inc
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
package net.consensys.gpact.performance.singlebc.read;

import java.math.BigInteger;
import net.consensys.gpact.common.BlockchainConfig;
import net.consensys.gpact.common.StatsHolder;
import net.consensys.gpact.helpers.CredentialsCreator;
import net.consensys.gpact.helpers.PropertiesLoader;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.core.methods.response.TransactionReceipt;

public class Main {
  static final Logger LOG = LogManager.getLogger(Main.class);

  // Running multiple times will reveal any gas difference due to rerunning.
  static int NUM_TIMES_EXECUTE = 4;

  public static void main(String[] args) throws Exception {
    StatsHolder.log("Example: Single Blockchain: Read");
    LOG.info("Started");

    if (args.length != 1) {
      LOG.info("Usage: [properties file name]");
      return;
    }

    PropertiesLoader propsLoader = new PropertiesLoader(args[0]);
    Credentials creds = CredentialsCreator.createCredentials();
    BlockchainConfig root = propsLoader.getBlockchainInfo("ROOT");

    Bc1ContractA bc1ContractABlockchain =
        new Bc1ContractA(creds, root.bcId, root.uri, root.gasPriceStrategy, root.period);
    Bc2ContractB bc2ContractBBlockchain =
        new Bc2ContractB(creds, root.bcId, root.uri, root.gasPriceStrategy, root.period);

    BigInteger val = BigInteger.valueOf(7);

    // Set-up client side and deploy contracts on the blockchains.
    bc2ContractBBlockchain.deployContracts(val);
    String contractBContractAddress = bc2ContractBBlockchain.contractB.getContractAddress();

    bc1ContractABlockchain.deployContracts(contractBContractAddress);
    String contractAContractAddress = bc1ContractABlockchain.contractA.getContractAddress();

    for (int numExecutions = 0; numExecutions < NUM_TIMES_EXECUTE; numExecutions++) {
      LOG.info("Execution: {}  *****************", numExecutions);
      StatsHolder.log("Execution: " + numExecutions + " **************************");

      TransactionReceipt txR = bc1ContractABlockchain.doRead();
      bc1ContractABlockchain.showEvents(txR);
      bc1ContractABlockchain.showValueRead();

      bc2ContractBBlockchain.separatedGet();
      bc1ContractABlockchain.separatedGet();
    }

    bc1ContractABlockchain.shutdown();
    bc2ContractBBlockchain.shutdown();

    StatsHolder.log("End");
    StatsHolder.print();
  }
}
