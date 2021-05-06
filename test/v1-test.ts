import "@nomiclabs/hardhat-waffle"
import { ethers } from 'hardhat'
import { expect, use } from 'chai'
import { BigNumberish, BytesLike, PayableOverrides, utils, ContractTransaction } from "ethers"
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/dist/src/signer-with-address"
import { V1 } from "typechain/V1"
const { reverts } = require('truffle-assertions')

let owner: SignerWithAddress
let addr1: SignerWithAddress
let addr2: SignerWithAddress

const oneEther = utils.parseEther("1")
const kEther = utils.parseEther("1000")
const proxyAddress = "0x0000000000000000000000000000000000000000";

ethers.getSigners().then((signers) => {
    owner = signers[0]
    addr1 = signers[1]
    addr2 = signers[2]
})

describe("testing for V1", async () => {
    let contract: V1

    beforeEach(async () => {
        const Contract = await ethers.getContractFactory("V1");
        contract = (await Contract.deploy()) as V1
        await contract.deployed()

    })
    describe("Basic Approvals", async () => {
        before(async () => {

        });

        beforeEach(async () => {
        });

        it("faild when not an owner", async () => {
        });

    })
})
