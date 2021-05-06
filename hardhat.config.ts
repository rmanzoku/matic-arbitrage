import { HardhatUserConfig } from "hardhat/config";
import { HttpNetworkAccountsConfig } from "hardhat/types"

import "hardhat-typechain"
import "@nomiclabs/hardhat-waffle"
import "hardhat-gas-reporter";

import "solc"

const accounts = (): HttpNetworkAccountsConfig => {
  if (!process.env.PRIV_KEY) {
    return "remote"
  }
  return [process.env.PRIV_KEY!]
}

const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
      forking: {
        url: `https://mainnet.infura.io/v3/${process.env.INFURA_KEY}`,
        enabled: false
      }
    },
    matic: {
      url: "https://rpc-mainnet.matic.network",
      chainId: 137,
      gas: 8500000,
      gasPrice: 1000000001,
    },
  },
  solidity: {
    version: "0.8.4",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  },
  mocha: {
    timeout: 20000
  },
  gasReporter: {
    enabled: true,
    currency: "ETH",
    gasPrice: 1
  }
};

export default config;
