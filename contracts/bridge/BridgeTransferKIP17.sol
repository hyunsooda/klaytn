// Copyright 2019 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

pragma solidity 0.5.6;

import "../kip17/IKIP17.sol";
import "../kip17/KIP17MetadataMintable.sol";
import "../kip17/KIP17Burnable.sol";

import "../sc_kip17/IKIP17BridgeReceiver.sol";
import "./BridgeTransfer.sol";


contract BridgeTransferKIP17 is BridgeTokens, IKIP17BridgeReceiver, BridgeTransfer {
    // handleKIP17Transfer sends the KIP17 by the request.
    function handleKIP17Transfer(
        bytes32 _requestTxHash,
        address _from,
        address _to,
        address _tokenAddress,
        uint256 _tokenId,
        uint64 _requestedNonce,
        uint64 _requestedBlockNumber,
        string memory _tokenURI,
        bytes memory _extraData
    )
        public
        onlyOperators
    {
        _lowerHandleNonceCheck(_requestedNonce);

        if (!_voteValueTransfer(_requestedNonce)) {
            return;
        }

        _setHandledRequestTxHash(_requestTxHash);

        handleNoncesToBlockNums[_requestedNonce] = _requestedBlockNumber;
        _updateHandleNonce(_requestedNonce);

        emit HandleValueTransfer(
            _requestTxHash,
            TokenType.KIP17,
            _from,
            _to,
            _tokenAddress,
            _tokenId,
            _requestedNonce,
            lowerHandleNonce,
            _extraData
        );

        if (modeMintBurn) {
            require(KIP17MetadataMintable(_tokenAddress).mintWithTokenURI(_to, _tokenId, _tokenURI), "mint failed");
        } else {
            IKIP17(_tokenAddress).safeTransferFrom(address(this), _to, _tokenId);
        }
    }

    // _requestKIP17Transfer requests transfer KIP17 to _to on relative chain.
    function _requestKIP17Transfer(
        address _tokenAddress,
        address _from,
        address _to,
        uint256 _tokenId,
        bytes memory _extraData
    )
        internal
        onlyRegisteredToken(_tokenAddress)
        onlyUnlockedToken(_tokenAddress)
    {
        require(isRunning, "stopped bridge");

        if (modeMintBurn) {
            KIP17Burnable(_tokenAddress).burn(_tokenId);
        }

        emit RequestValueTransfer(
            TokenType.KIP17,
            _from,
            _to,
            _tokenAddress,
            _tokenId,
            requestNonce,
            0,
            _extraData
        );
        requestNonce++;
    }

    // onKIP17Received function of KIP17 token for 1-step deposits to the Bridge
    function onKIP17Received(
        address _from,
        uint256 _tokenId,
        address _to,
        bytes memory _extraData
    )
        public
    {
        _requestKIP17Transfer(msg.sender, _from, _to, _tokenId, _extraData);
    }

    // requestKIP17Transfer requests transfer KIP17 to _to on relative chain.
    function requestKIP17Transfer(
        address _tokenAddress,
        address _to,
        uint256 _tokenId,
        bytes memory _extraData
    )
        public
    {
        IKIP17(_tokenAddress).transferFrom(msg.sender, address(this), _tokenId);
        _requestKIP17Transfer(_tokenAddress, msg.sender, _to, _tokenId, _extraData);
    }
}
