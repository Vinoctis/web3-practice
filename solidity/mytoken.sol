// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyToken is ERC20, Ownable {
    uint256 public constant RATE = 100000000;//100000000 MyToken per eth
    uint256 public constant MIN_ETH = 0.001 ether;

    constructor(address initialOwner) ERC20("RCCDemoToken", "RDT") Ownable(msg.sender) {}

    function mint() public payable {
        require(msg.value >= MIN_ETH, "not enough ETH sent");
        uint256 tokenToMint = msg.value * RATE;
        _mint(msg.sender, tokenToMint);
    }

    function withDrawEth() public onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "no ETH to withDraw");
        payable(owner()).transfer(balance);
    }

    receive() external payable {
        mint();
    }


}