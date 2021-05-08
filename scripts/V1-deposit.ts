import { ethers } from 'hardhat'
import { V1 } from "typechain/V1"
import { WMATIC } from "typechain/WMATIC"
import { addrV1, addrWMATIC, oneEther } from "./common"


const main = async () => {
    const wmatic = (await ethers.getContractAt("WMATIC", addrWMATIC)) as WMATIC

    const value = oneEther.mul(5)
    const tx = await wmatic.deposit({ value })
    console.log("convert:", tx.hash)

    const tx2 = await wmatic.transfer(addrV1, value)
    console.log("deposit:", tx2.hash)
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });