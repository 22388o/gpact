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
pragma solidity >=0.8;

import "./CallPathCallExecutionTree.sol";

contract CallExecutionTreeTest is CallPathCallExecutionTree {
    function extractTargetFromCallGraph1(
        bytes memory _callGraph,
        uint256[] memory _callPath
    )
        external
        pure
        returns (
            uint256 targetBlockchainId,
            address targetContract,
            bytes memory functionCall
        )
    {
        return extractTargetFromCallGraph(_callGraph, _callPath, true);
    }
}
