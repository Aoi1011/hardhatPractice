// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

contract Escrow {
    uint256 public price;
    address payable public seller;
    address payable public buyer;

    address[] public previousBuyers;

    enum State {
        Sale,
        Locked,
        Release,
        Closed,
        Complete
    }

    State public state;

    modifier condition(bool _condition) {
        require(_condition);
        _;
    }

    modifier onlyBuyer() {
        require(msg.sender == buyer, "Only buyer can call this.");
        _;
    }

    modifier onlySeller() {
        require(msg.sender == seller, "Only seller can call this.");
        _;
    }

    modifier notSeller() {
        require(msg.sender != seller, "Seller shouldn't call this.");
        _;
    }

    modifier inState(State _state) {
        require(state == _state, "Invalid state.");
        _;
    }

    event Closed(uint256 when);

    event ConfirmPurchase(uint256 when, address by);

    event ConfirmReceived(uint256 when, address by);

    event SellerRefundBuyer(uint256 when);

    event SellerRefunded(uint256 when);

    event Restarted(uint256 when);

    event End(uint256 when);
}
