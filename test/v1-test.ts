import "@nomiclabs/hardhat-waffle"
import { ethers } from 'hardhat'
import { expect, use } from 'chai'
import { BigNumberish, BytesLike, PayableOverrides, utils, ContractTransaction } from "ethers"
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/dist/src/signer-with-address"
import { V1 } from "typechain/V1"
import { WMATIC } from "typechain/WMATIC"
const { reverts } = require('truffle-assertions')

let owner: SignerWithAddress
let addr1: SignerWithAddress
let addr2: SignerWithAddress

const oneEther = utils.parseEther("1")
const kEther = utils.parseEther("1000")
const proxyAddress = "0x0000000000000000000000000000000000000000";

const addrWMATIC = "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"
const addrWETH = "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"

const forth = [
    addrWMATIC,
    addrWETH
]
const back = [
    addrWETH,
    addrWMATIC,
]
const val = oneEther.mul(10)

const addrQuickSwap = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff"
const addrElk = "0xf38a7A7Ac2D745E2204c13F824c00139DF831FFf"
const addrSushiSwap = "0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"
// ISwapper private quick =
//     ISwapper(0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff);
// ISwapper private elk = ISwapper(0xf38a7A7Ac2D745E2204c13F824c00139DF831FFf);
// ISwapper private sushi =
//     ISwapper(0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506);

ethers.getSigners().then((signers) => {
    owner = signers[0]
    addr1 = signers[1]
    addr2 = signers[2]
})

describe("testing for V1", async () => {
    let contract: V1
    let wmaticContract: WMATIC

    beforeEach(async () => {
        const Contract = await ethers.getContractFactory("V1");
        contract = (await Contract.deploy()) as V1
        await contract.deployed()

        wmaticContract = (await ethers.getContractAt("WMATIC", addrWMATIC, owner)) as WMATIC
        // const wmatic = await contract.WETH()
        // wmatic.map((a) => {
        //     if (a != addrWMATIC) {
        //         console.error(a, addrWMATIC)
        //         return
        //     }
        // })

        await wmaticContract.deposit({ value: kEther })
        await wmaticContract.transfer(contract.address, kEther)
    })
    describe("Basic Approvals", async () => {
        before(async () => {

        });

        beforeEach(async () => {
        });

        it("dry", async () => {
            const ret = await contract.dry(addrElk, addrQuickSwap, val, forth, back)
            console.log(ret.toString())
        });

        it("withdraw", async () => {
            const before = await owner.getBalance()
            await contract.withdraw(addrWMATIC)
            const after = await owner.getBalance()
            console.log(after.toString(), before.toString())
            expect(after).to.equal(before.add(kEther))
        });

        it("elk and quick", async () => {
            const a = await contract.swap(addrElk, addrQuickSwap, val, forth, back)
            //console.log(a)
        });
        it("quick and elk", async () => {
            const a = await contract.swap(addrQuickSwap, addrElk, val, forth, back)
            //console.log(a)
        });

        it("quick and sushi", async () => {
            const a = await contract.swap(addrQuickSwap, addrSushiSwap, val, forth, back)
            //console.log(a)
        });

        it("sushi and quick", async () => {
            const a = await contract.swap(addrSushiSwap, addrQuickSwap, val, forth, back)
            //console.log(a)
        });

        it("elk and sushi", async () => {
            const a = await contract.swap(addrElk, addrSushiSwap, val, forth, back)
            //console.log(a)
        });

        it("sushi and elk", async () => {
            const a = await contract.swap(addrSushiSwap, addrElk, val, forth, back)
            //console.log(a)
        });

    })
})
