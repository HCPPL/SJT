var SJTCoin = artifacts.require("./SJTCoin.sol");

module.exports = function(deployer) {
  deployer.deploy(SJTCoin);
};
