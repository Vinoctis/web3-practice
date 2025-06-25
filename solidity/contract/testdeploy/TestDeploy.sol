// SPDX-License-Identifier: MIT
pragma solidity 0.8.30;

contract TestDeploy {
    uint public myNum = 10;
    function setNum(uint _newNum) external {
        myNum = _newNum;
    }
}

