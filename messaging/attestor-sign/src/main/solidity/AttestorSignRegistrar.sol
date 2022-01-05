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

import "../../../../../common/common/src/main/solidity/EcdsaSignatureVerification.sol";
import "../../../../../common/openzeppelin/src/main/solidity/access/Ownable.sol";

contract AttestorSignRegistrar is EcdsaSignatureVerification, Ownable {
    struct BlockchainRecord {
        uint256 signingThreshold;
        uint256 numSigners;
        mapping(address => bool) signers;
    }
    mapping(uint256 => BlockchainRecord) private blockchains;

    event Dump(string _Log);
    event Dump2(bytes data);

    function addSigner(uint256 _bcId, address _signer) external onlyOwner {
        require(
            blockchains[_bcId].signingThreshold != 0,
            "Can not add signer for blockchain with zero threshold"
        );
        internalAddSigner(_bcId, _signer);
        blockchains[_bcId].numSigners++;
    }

    function addSignerSetThreshold(
        uint256 _bcId,
        address _signer,
        uint256 _newSigningThreshold
    ) external onlyOwner {
        internalAddSigner(_bcId, _signer);
        uint256 newNumSigners = blockchains[_bcId].numSigners + 1;
        blockchains[_bcId].numSigners = newNumSigners;
        internalSetThreshold(_bcId, newNumSigners, _newSigningThreshold);
    }

    function addSignersSetThreshold(
        uint256 _bcId,
        address[] calldata _signers,
        uint256 _newSigningThreshold
    ) external onlyOwner {
        for (uint256 i = 0; i < _signers.length; i++) {
            internalAddSigner(_bcId, _signers[i]);
        }
        uint256 newNumSigners = blockchains[_bcId].numSigners + _signers.length;
        blockchains[_bcId].numSigners = newNumSigners;
        internalSetThreshold(_bcId, newNumSigners, _newSigningThreshold);
    }

    function removeSigner(uint256 _bcId, address _signer) external onlyOwner {
        internalRemoveSigner(_bcId, _signer);
        uint256 newNumSigners = blockchains[_bcId].numSigners - 1;
        require(
            newNumSigners >= blockchains[_bcId].signingThreshold,
            "Proposed new number of signers is less than threshold"
        );
        blockchains[_bcId].numSigners = newNumSigners;
    }

    function removeSignerSetThreshold(
        uint256 _bcId,
        address _signer,
        uint256 _newSigningThreshold
    ) external onlyOwner {
        internalRemoveSigner(_bcId, _signer);
        uint256 newNumSigners = blockchains[_bcId].numSigners - 1;
        blockchains[_bcId].numSigners = newNumSigners;
        internalSetThreshold(_bcId, newNumSigners, _newSigningThreshold);
    }

    function setThreshold(uint256 _bcId, uint256 _newSigningThreshold)
        external
        onlyOwner
    {
        internalSetThreshold(
            _bcId,
            blockchains[_bcId].numSigners,
            _newSigningThreshold
        );
    }

    function verifyAndCheckThreshold(
        uint256 _blockchainId,
        address[] calldata _signers,
        bytes32[] calldata _sigR,
        bytes32[] calldata _sigS,
        uint8[] calldata _sigV,
        bytes calldata _plainText
    ) external view returns (bool) {
        checkThreshold(_blockchainId, _signers);
        return verify(_blockchainId, _signers, _sigR, _sigS, _sigV, _plainText);
    }

    function verify(
        uint256 _blockchainId,
        address[] calldata _signers,
        bytes32[] calldata _sigR,
        bytes32[] calldata _sigS,
        uint8[] calldata _sigV,
        bytes calldata _plainText
    ) public view returns (bool) {
        uint256 signersLength = _signers.length;
        require(signersLength == _sigR.length, "sigR length mismatch");
        require(signersLength == _sigS.length, "sigS length mismatch");
        require(signersLength == _sigV.length, "sigV length mismatch");

        for (uint256 i = 0; i < signersLength; i++) {
            // Check that signer is a signer for this blockchain
            require(
                blockchains[_blockchainId].signers[_signers[i]],
                "Signer not signer for this blockchain"
            );
            // Verify the signature
            require(
                verifySigComponents(
                    _signers[i],
                    _plainText,
                    _sigR[i],
                    _sigS[i],
                    _sigV[i]
                ),
                "Signature did not verify"
            );
        }
        return true;
    }

     function uint2str(uint _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len;
        while (_i != 0) {
            k = k-1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }

    function char(bytes1 b) internal pure returns (bytes1 c) {
        if (uint8(b) < 10) return bytes1(uint8(b) + 0x30);
        else return bytes1(uint8(b) + 0x57);
    }

    function toAsciiString(address x) internal pure returns (string memory) {
        bytes memory s = new bytes(40);
        for (uint i = 0; i < 20; i++) {
            bytes1 b = bytes1(uint8(uint(uint160(x)) / (2**(8*(19 - i)))));
            bytes1 hi = bytes1(uint8(b) / 16);
            bytes1 lo = bytes1(uint8(b) - 16 * uint8(hi));
            s[2*i] = char(hi);
            s[2*i+1] = char(lo);            
        }
    return string(s);
}

    function verifySigComponents2(
        address _signer,
        bytes calldata _message,
        bytes32 _sigR,
        bytes32 _sigS,
        uint8 _sigV
    ) internal returns (bool) {

        emit Dump2(_message);

        bytes32 digest = keccak256(_message);

        emit Dump(uint2str(_sigV));
        if (_sigV != 27 && _sigV != 28) {
            return false;
        } else {
            // The signature is verified if the address recovered from the signature matches
            // the signer address (which maps to the public key).
            emit Dump(toAsciiString(_signer));
            address signer = ecrecover(digest, _sigV, _sigR, _sigS);
            emit Dump(toAsciiString(signer));
            return true;
        }
    }

    function verify2(
        uint256 _blockchainId,
        address[] calldata _signers,
        bytes32[] calldata _sigR,
        bytes32[] calldata _sigS,
        uint8[] calldata _sigV,
        bytes calldata _plainText
    ) public returns (bool) {
        uint256 signersLength = _signers.length;
        require(signersLength == _sigR.length, "sigR length mismatch");
        require(signersLength == _sigS.length, "sigS length mismatch");
        require(signersLength == _sigV.length, "sigV length mismatch");

        for (uint256 i = 0; i < signersLength; i++) {
            // Check that signer is a signer for this blockchain
            require(
                blockchains[_blockchainId].signers[_signers[i]],
                "Signer not signer for this blockchain"
            );
            // Verify the signature
            require(
                verifySigComponents2(
                    _signers[i],
                    _plainText,
                    _sigR[i],
                    _sigS[i],
                    _sigV[i]
                ),
                "Signature did not verify"
            );
        }
        return true;
    }

    function checkThreshold(uint256 _blockchainId, address[] calldata _signers)
        public
        view
        returns (bool)
    {
        uint256 signersLength = _signers.length;
        require(
            signersLength >= blockchains[_blockchainId].signingThreshold,
            "Not enough signers"
        );
        return true;
    }

    function getSigningThreshold(uint256 _blockchainId)
        external
        view
        returns (uint256)
    {
        return blockchains[_blockchainId].signingThreshold;
    }

    function numSigners(uint256 _blockchainId) external view returns (uint256) {
        return blockchains[_blockchainId].numSigners;
    }

    function isSigner(uint256 _blockchainId, address _mightBeSigner)
        external
        view
        returns (bool)
    {
        return blockchains[_blockchainId].signers[_mightBeSigner];
    }

    /************************************* PRIVATE FUNCTIONS BELOW HERE *************************************/
    /************************************* PRIVATE FUNCTIONS BELOW HERE *************************************/
    /************************************* PRIVATE FUNCTIONS BELOW HERE *************************************/

    function internalAddSigner(uint256 _bcId, address _signer) private {
        require(
            blockchains[_bcId].signers[_signer] == false,
            "Signer already exists"
        );
        blockchains[_bcId].signers[_signer] = true;
    }

    function internalRemoveSigner(uint256 _bcId, address _signer) private {
        require(blockchains[_bcId].signers[_signer], "Signer does not exist");
        blockchains[_bcId].signers[_signer] = false;
    }

    function internalSetThreshold(
        uint256 _bcId,
        uint256 _currentNumSigners,
        uint256 _newThreshold
    ) private {
        require(
            _currentNumSigners >= _newThreshold,
            "Number of signers less than threshold"
        );
        require(_newThreshold != 0, "Threshold can not be zero");
        blockchains[_bcId].signingThreshold = _newThreshold;
    }
}
