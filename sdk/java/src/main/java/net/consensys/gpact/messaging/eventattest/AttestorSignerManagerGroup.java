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
package net.consensys.gpact.messaging.eventattest;

import java.math.BigInteger;
import java.util.*;
import net.consensys.gpact.common.BlockchainConfig;
import net.consensys.gpact.common.BlockchainId;
import net.consensys.gpact.messaging.MessagingManagerGroupInterface;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.web3j.crypto.Credentials;

/** Manage multiple blockchains, each holding a set of registrar and verification contracts */
public class AttestorSignerManagerGroup implements MessagingManagerGroupInterface {
  static final Logger LOG = LogManager.getLogger(AttestorSignerManagerGroup.class);

  Map<BlockchainId, AttestorSignerManager> blockchains = new HashMap<>();

  @Override
  public void addBlockchainAndDeployContracts(Credentials creds, BlockchainConfig bcInfo)
      throws Exception {
    BlockchainId blockchainId = bcInfo.bcId;
    if (this.blockchains.containsKey(blockchainId)) {
      return;
      // throw new Exception("Blockchain already added: " + blockchainId);
    }
    LOG.debug("Deploying Cross-Blockchain Control contracts for blockchain id {}", blockchainId);

    AttestorSignerManager manager =
        new AttestorSignerManager(
            creds, blockchainId, bcInfo.uri, bcInfo.gasPriceStrategy, bcInfo.period);
    manager.deployContracts();

    this.blockchains.put(blockchainId, manager);
  }

  @Override
  public void addBlockchainAndLoadContracts(
      Credentials creds, BlockchainConfig bcInfo, ArrayList<String> addresses) throws Exception {
    BlockchainId blockchainId = bcInfo.bcId;
    if (this.blockchains.containsKey(blockchainId)) {
      throw new Exception("Blockchain already added: " + blockchainId);
    }

    AttestorSignerManager manager =
        new AttestorSignerManager(
            creds, blockchainId, bcInfo.uri, bcInfo.gasPriceStrategy, bcInfo.period);
    manager.loadContracts(addresses);

    this.blockchains.put(blockchainId, manager);
  }

  @Override
  public void registerSignerOnAllBlockchains(String signersAddress) throws Exception {
    for (BlockchainId bcId1 : this.blockchains.keySet()) {
      registerSigner(signersAddress, bcId1);
    }
  }

  @Override
  public void registerSigner(String signersAddress, BlockchainId bcId1) throws Exception {
    // Add the signer (their address / public key) to each blockchain
    for (BlockchainId bcId2 : this.blockchains.keySet()) {
      AttestorSignerManager manager = this.blockchains.get(bcId2);
      manager.registerSigner(bcId1, signersAddress);
    }
  }

  @Override
  public void registerFirstSignerOnAllBlockchains(String signersAddress) throws Exception {
    for (BlockchainId bcId1 : this.blockchains.keySet()) {
      registerFirstSigner(signersAddress, bcId1);
    }
  }

  @Override
  public void registerFirstSigner(String signersAddress, BlockchainId bcId1) throws Exception {
    // Add the signer (their address / public key) to each blockchain
    for (BlockchainId bcId2 : this.blockchains.keySet()) {
      AttestorSignerManager manager = this.blockchains.get(bcId2);
      manager.registerSigner(bcId1, signersAddress, BigInteger.ONE);
    }
  }

  @Override
  public String getVerifierAddress(final BlockchainId bcId) throws Exception {
    if (!this.blockchains.containsKey(bcId)) {
      throw new Exception("Blockchain not found: " + bcId);
    }
    return this.blockchains.get(bcId).getVerifierAddress();
  }
}
