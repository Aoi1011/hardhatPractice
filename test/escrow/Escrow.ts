import { expect } from "chai";
import { expectRevert } from "@openzeppelin/test-helpers";
import { ethers } from "hardhat";

const humanReadableUnixTimeStamp = timeStampInt => {
  return new Date(timeStampInt * 1000);
};

describe("Escrow Events and State", () => {
  let provider;
  let Escrow, escrow, seller, firstBuyer, secondBuyer;

  let closedEvent,
    confirmPurchaseEvent,
    sellerRefundBuyerEvent,
    confirmReceivedEvent,
    sellerRefundedEvent,
    restartedEvent,
    endEvent;

  beforeEach(async () => {
    provider = ethers.getDefaultProvider();

    Escrow = await ethers.getContractFactory("Escrow");
    escrow = await Escrow.deploy({ value: ethers.utils.parseEther("2.0") });

    [seller, firstBuyer, secondBuyer, _] = await ethers.getSigners();

    closedEvent = new Promise((resolve, reject) => {
      escrow.on("Closed", (when, event) => {
        event.removeListener();

        resolve({
          when,
        });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 6000);
    });

    confirmPurchaseEvent = new Promise((resolve, reject) => {
      escrow.on("ConfirmPurchase", (when, by, event) => {
        event.removeListener();

        resolve({
          when,
          by,
        });
      });

      setTimeout(() => {
        reject(new Error("timeout"));
      }, 60000);
    });
  });
});
