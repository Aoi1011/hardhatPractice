// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

contract Token {
    string public name = "My Hardhat Token";
    string public symbol = "MHT";

    uint256 public totalSupply = 1000000;

    address public owner;

    mapping(address => uint256) public balances;

    constructor() {
        balances[msg.sender] = totalSupply;
        owner = msg.sender;
    }

    function transfer(address _to, uint256 _amount) external {
        require(balances[msg.sender] >= _amount, "Not enough tokens");

        balances[msg.sender] -= _amount;
        balances[_to] += _amount;
    }

    function balanceOf(address _account) external view returns (uint256) {
        return balances[_account];
    }
}
