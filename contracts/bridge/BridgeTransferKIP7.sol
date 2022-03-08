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

import "../kip7/IKIP7.sol";
import "../kip7/KIP7Mintable.sol";
import "../kip7/KIP7Burnable.sol";

import "../sc_kip7/IKIP7BridgeReceiver.sol";
import "./BridgeTransfer.sol";


contract BridgeTransferKIP7 is BridgeTokens, IKIP7BridgeReceiver, BridgeTransfer {
    // handleKIP7Transfer sends the token by the request.
    function handleKIP7Transfer(
        bytes32 _requestTxHash,
        address _from,
        address _to,
        address _tokenAddress,
        uint256 _value,
        uint64 _requestedNonce,
        uint64 _requestedBlockNumber,
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
            TokenType.KIP7,
            _from,
            _to,
            _tokenAddress,
            _value,
            _requestedNonce,
            lowerHandleNonce,
            _extraData
        );

        if (modeMintBurn) {
            KIP7Mintable(_tokenAddress).mint(_to, _value);
        } else {
            IKIP7(_tokenAddress).transfer(_to, _value);
        }
    }

    // _requestKIP7Transfer requests transfer KIP7 to _to on relative chain.
    function _requestKIP7Transfer(
        address _tokenAddress,
        address _from,
        address _to,
        uint256 _value,
        uint256 _feeLimit,
        bytes memory _extraData
    )
        internal
        onlyRegisteredToken(_tokenAddress)
        onlyUnlockedToken(_tokenAddress)
    {
        require(isRunning, "stopped bridge");
        require(_value > 0, "zero msg.value @@@");

        uint256 fee = _payKIP7FeeAndRefundChange(_from, _tokenAddress, _feeLimit);

        if (modeMintBurn) {
            KIP7Burnable(_tokenAddress).burn(_value);
        }

        emit RequestValueTransfer(
            TokenType.KIP7,
            _from,
            _to,
            _tokenAddress,
            _value,
            requestNonce,
            fee,
            _extraData
        );
        requestNonce++;
    }

    // onKIP7Received function of KIP7 token for 1-step deposits to the Bridge.
    function onKIP7Received(
        address _from,
        address _to,
        uint256 _value,
        uint256 _feeLimit,
        bytes memory _extraData
    )
        public
    {
        _requestKIP7Transfer(msg.sender, _from, _to, _value, _feeLimit, _extraData);
    }

    // requestKIP7Transfer requests transfer KIP7 to _to on relative chain.
    function requestKIP7Transfer(
        address _tokenAddress,
        address _to,
        uint256 _value,
        uint256 _feeLimit,
        bytes memory _extraData
    )
        public
    {
        IKIP7(_tokenAddress).transferFrom(msg.sender, address(this), _value.add(_feeLimit));
        _requestKIP7Transfer(_tokenAddress, msg.sender, _to, _value, _feeLimit, _extraData);
    }


    // setKIP7Fee sets the fee of the token transfer.
    function setKIP7Fee(address _token, uint256 _fee, uint64 _requestNonce)
        external
        onlyOperators
    {
        if (!_voteConfiguration(_requestNonce)) {
            return;
        }
        _setKIP7Fee(_token, _fee);
    }
}
