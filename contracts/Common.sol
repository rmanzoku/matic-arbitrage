// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

abstract contract Common {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == owner, "sender is not owner");
        _;
    }

    function kill() external onlyOwner {
        selfdestruct(payable(msg.sender));
    }
}
