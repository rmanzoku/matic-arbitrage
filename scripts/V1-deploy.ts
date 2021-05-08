import { ethers } from 'hardhat'
import { V1 } from "typechain/V1"
import { addrV1 } from "./common"

const main = async () => {
    if (addrV1) {
        return
    }
    const Contract = await ethers.getContractFactory("V1");

    const contract = (await Contract.deploy()) as V1
    await contract.deployed();
    console.log('deployed txHash', contract.deployTransaction.hash);
    console.log('deployed address', contract.address);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });