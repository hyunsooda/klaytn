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

import "../externals/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./IKIP17BridgeReceiver.sol";
import "../kip17/KIP17.sol";


/**
 * @title KIP17ServiceChain
 * @dev KIP17 service chain value transfer logic for 1-step transfer.
 */
contract KIP17ServiceChain is KIP17, Ownable {
    address public bridge;

    constructor(address _bridge) internal {
        setBridge(_bridge);
    }

    function setBridge(address _bridge) public onlyOwner {
        if (!_bridge.isContract()) {
            revert("bridge is not a contract");
        }

        bridge = _bridge;
    }

    function requestValueTransfer(uint256 _uid, address _to, bytes calldata _extraData) external {
        transferFrom(msg.sender, bridge, _uid);

        IKIP17BridgeReceiver(bridge).onKIP17Received(msg.sender, _uid, _to, _extraData);
    }
}
