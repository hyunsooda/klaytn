// Copyright 2024 The klaytn Authors
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

// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "../../system_contracts/kaiabridge/EnumerableSetUpgradable.sol";

contract EnumSet is EnumerableSetUpgradable {
    using EnumerableSet for EnumerableSet.UintSet;

    EnumerableSet.UintSet s;

    function add(uint64 v) public {
        setAdd(s, v);
    }

    function remove(uint64 v) public {
        setRemove(s, v);
    }

    function contains(uint64 v) public view returns (bool) {
        return setContains(s, v);
    }

    function length() public view returns (uint256) {
        return setLength(s);
    }

    function getAll() public view returns (uint64[] memory) {
        return getAll(s);
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() { _disableInitializers(); }
}
