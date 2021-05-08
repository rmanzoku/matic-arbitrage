// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.4;

import {ISwapper} from "./interface/ISwapper.sol";
import {Common} from "./Common.sol";
import {IERC20} from "./interface/IERC20.sol";

import "hardhat/console.sol";

contract V1 is Common {
    uint256 unlimit = 2**256 - 1;

    receive() external payable {}

    function dry(
        address swapper1,
        address swapper2,
        uint256 val,
        address[] calldata forth,
        address[] calldata back
    ) external view returns (uint256) {
        uint256[] memory outgoings =
            ISwapper(swapper1).getAmountsOut(val, forth);

        uint256[] memory incomings =
            ISwapper(swapper2).getAmountsOut(
                outgoings[outgoings.length - 1],
                back
            );

        console.log(
            outgoings[0],
            outgoings[outgoings.length - 1],
            incomings[0],
            incomings[incomings.length - 1]
        );
        return incomings[incomings.length - 1];
    }

    function swap(
        address swapper1,
        address swapper2,
        uint256 val,
        address[] calldata forth,
        address[] calldata back
    ) external {
        for (uint256 i = 0; i < forth.length; i++) {
            IERC20 token = IERC20(forth[i]);
            if (token.allowance(address(this), swapper1) < 1) {
                token.approve(swapper1, unlimit);
            }
        }
        for (uint256 i = 0; i < back.length; i++) {
            IERC20 token = IERC20(forth[i]);
            if (token.allowance(address(this), swapper2) < 1) {
                token.approve(swapper2, unlimit);
            }
        }

        uint256[] memory outgoings =
            ISwapper(swapper1).swapExactTokensForTokens(
                val,
                0,
                forth,
                address(this),
                unlimit
            );

        uint256[] memory incomings =
            ISwapper(swapper2).swapExactTokensForTokens(
                outgoings[outgoings.length - 1],
                0,
                back,
                address(this),
                unlimit
            );

        console.log(
            outgoings[0],
            outgoings[outgoings.length - 1],
            incomings[0],
            incomings[incomings.length - 1]
        );
        require(incomings[incomings.length - 1] > outgoings[0], "loss");
    }

    function withdraw(address _token) external onlyOwner {
        IERC20 token = IERC20(_token);
        token.transfer(msg.sender, token.balanceOf(address(this)));
    }

    // function swapElkQuick(address[] calldata path, address[] calldata reverse)
    //     external
    //     payable
    // {
    //     console.log(msg.sender.balance);
    //     console.log(msg.sender);

    //     uint256 val = msg.value;
    //     uint256[] memory amounts =
    //         elk.swapExactTokensForTokens{value: val}(
    //             0,
    //             path,
    //             address(this),
    //             unlimit
    //         );

    //     for (uint256 i = 0; i < path.length; i++) {
    //         IERC20 token = IERC20(path[i]);
    //         address spender = address(elk);

    //         if (token.allowance(address(this), spender) < amounts[i]) {
    //             token.approve(spender, unlimit);
    //         }
    //     }

    //     uint256[] memory amounts2 =
    //         elk.swapExactTokensForMATIC(
    //             amounts[1],
    //             0,
    //             reverse,
    //             address(this),
    //             unlimit
    //         );

    //     console.log(amounts[0], amounts[1], amounts2[0], amounts2[1]);
    // }

    // function WETH() external view returns (address[3] memory) {
    //     address[3] memory ret =
    //         [quickswap.WETH(), elk.WMATIC(), sushiswap.WETH()];
    //     return ret;
    // }
}
