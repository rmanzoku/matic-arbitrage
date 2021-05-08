import { HardhatUserConfig } from "hardhat/config";
import { HttpNetworkAccountsConfig, HardhatNetworkAccountsUserConfig } from "hardhat/types"

import "hardhat-typechain"
import "@nomiclabs/hardhat-waffle"
import "hardhat-gas-reporter";

import "solc"

const accounts = (): HttpNetworkAccountsConfig => {
  if (!process.env.PRIV_KEY) {
    return []
  }
  return [process.env.PRIV_KEY!]
}

const accounts2 = (): HardhatNetworkAccountsUserConfig => {
  return [{
    privateKey: process.env.PRIV_KEY!,
    balance: "10000000000000000"
  }]
}

const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",
  networks: {
    hardhat: {
      //  accounts: accounts2(),
      forking: {
        url: `https://rpc-mainnet.matic.network`,
        enabled: true
      }
    },
    matic: {
      url: "https://rpc-mainnet.matic.network",
      chainId: 137,
      gas: 8500000,
      gasPrice: 1000000001,
      accounts: accounts()
    },
  },
  solidity: {
    compilers: [
      {
        version: "0.8.4",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          }
        }
      },
      {
        version: "0.4.26",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200
          }
        }
      },
    ]
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  },
  mocha: {
    timeout: 200000
  },
  gasReporter: {
    enabled: true,
    currency: "ETH",
    gasPrice: 1
  }
};

export default config;
