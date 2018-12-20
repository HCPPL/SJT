var SJTBridge = artifacts.require("./SJTBridge.sol");

module.exports = function(deployer) {
  deployer.deploy(SJTBridge);
};
