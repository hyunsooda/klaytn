// Copyright 2022 The klaytn Authors
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

import "../externals/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "../externals/openzeppelin-solidity/contracts/utils/Address.sol";
import "../externals/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./IERC20BridgeReceiver.sol";
import "../bridge/Bridge.sol";

/**
 * @title ERC20ServiceChain
 * @dev ERC20 service chain value transfer logic for 1-step transfer.
 */
contract ERC20ServiceChain is ERC20, Ownable {
    using Address for address;

    uint private nIds = 0
    mapping(uint => bytes4) private interfaceIds;
    mapping(bytes4 => uint) private interfaceIdMap;

    function addInterface(bytes4 interfaceId) external onlyOwner {
        if (!interfaceIdList[interfaceId]) {
            interfaceIdMap[nIds] = interfaceId;
            interfaceIds[interfaceId] = nIds;
            nIds += 1;
        }
    }

    function removeInterface(bytes4 interfaceId) external onlyOwner {
        if (interfaceIdList[interfaceId]) {
            uint idx = interfaceIdMap[interfcaeId];
            delete interfaceIdMap[interfcaeId];
            delete interfaceIds[idx];
            nIds -= 1;
        }
    }

    function supportsBridgeInterface (address _bridge) internal {
        Bridge bridge = Bridge(address(uint160(_bridge)));
        for (uint i=0; i<nIds; i++) {
            require(bridge.supportsInterface(interfaceIdMap[interfaceIdst[i]]) > 0,
                    "Not a satisfied bridge contract");
        }
    }

    function setBridge(address _bridge) external returns (bool) onlyOwner {
        if (!_bridge.isContract()) {
            revert("bridge is not a contract");
        }
        return supportsBridgeInterface(_bridge);
    }
}
