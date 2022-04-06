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

import "../externals/openzeppelin-solidity/contracts/token/ERC721/ERC721.sol";

import "../externals/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./IERC721BridgeReceiver.sol";
import "../bridge/Bridge.sol";

/**
 * @title ERC721ServiceChain
 * @dev ERC721 service chain value transfer logic for 1-step transfer.
 */
contract ERC721ServiceChain is ERC721, Ownable {
    address public bridge;

    // Pre-calculated interface ID
    bytes4 private constant _BRIDGE_TRANSFER_INTERFACE_ID = 0x00000001;
    bytes4 private constant _BRIDGE_TRANSFER_KALY_INTERFACE_ID = 0x00000002;
    bytes4 private constant _BRIDGE_TRANSFER_ERC20_INTERFACE_ID = 0x00000003;
    bytes4 private constant _BRIDGE_TRANSFER_ERC721_INTERFACE_ID = 0x00000004;

    constructor(address _bridge) internal {
        setBridge(_bridge); 
    }
    
    function supportsBridgeInterface (address _bridge) internal view returns (bool) {
        Bridge bridge = Bridge(address(uint160(_bridge)));
        require( 
            bridge.supportsInterface(_BRIDGE_TRANSFER_INTERFACE_ID) &&
            bridge.supportsInterface(_BRIDGE_TRANSFER_ERC721_INTERFACE_ID),
            "Not a satisfied bridge contract");
    }

    function setBridge(address _bridge) public onlyOwner {
        if (!_bridge.isContract()) {
            revert("bridge is not a contract");
        }
        supportsBridgeInterface(_bridge);
        bridge = _bridge;
    }

    function requestValueTransfer(uint256 _uid, address _to, bytes calldata _extraData) external {
        transferFrom(msg.sender, bridge, _uid);

        IERC721BridgeReceiver(bridge).onERC721Received(msg.sender, _uid, _to, _extraData);
    }
}
