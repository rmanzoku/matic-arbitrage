import { ethers } from 'hardhat'
import { V1 } from "typechain/V1"
import { addrV1, addrWMATIC } from "./common"

const main = async () => {
    const contract = (await ethers.getContractAt("V1", addrV1)) as V1
    const tx = await contract.withdraw(addrWMATIC)
    console.log(tx)
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });