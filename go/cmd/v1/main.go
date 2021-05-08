package main

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	v1 "github.com/rmanzoku/matic-arbitrage/go/contracts/v1"
	"github.com/rmanzoku/matic-arbitrage/go/crawler"
)

var (
	txFee, _ = crawler.GweiToWei(300700)
	value, _ = crawler.EtherToWei(1)

	baseToken = crawler.AddrWMATIC
	swappers  = []common.Address{crawler.AddrQuickSwap, crawler.AddrElk, crawler.AddrSushiSwap}

	swapTokens = []common.Address{
		crawler.AddrWETH,
		common.HexToAddress("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"), // wbtc
		common.HexToAddress("0xe1c8f3d529bea8e3fa1fac5b416335a2f998ee1c"), // elk
		common.HexToAddress("0x831753DD7087CaC61aB5644b308642cc1c33Dc13"), // quick
		common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174"), // usdc
		common.HexToAddress("0x8f3cf7ad23cd3cadbd9735aff958023239c6a063"), // dai
		common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f"), // usdt
		common.HexToAddress("0x6ae7dfc73e0dde2aa99ac063dcf7e8a63265108c"), // jpyc
		// common.HexToAddress("0x05089c9ebffa4f0aca269e32056b1b36b37ed71b"), // krill
		common.HexToAddress("0x289cf2b63c5edeeeab89663639674d9233e8668e"), // Fish
		common.HexToAddress("0x652a7b75c229850714d4a11e856052aac3e9b065"), // WOLF
		common.HexToAddress("0xd6df932a45c0f255f85145f286ea0b292b21c90b"), // aave

	}
)

func handler(ctx context.Context, c *crawler.Crawler) (err error) {
	contract, err := v1.NewV1(common.HexToAddress("0xc86ba4527797f46569bEb67fb7fE2F0B2E6d1fB8"), c.EthClient)
	if err != nil {
		return err
	}

	valueWithFee := big.NewInt(0).Add(value, txFee)

	wg := sync.WaitGroup{}

	for _, swapper1 := range swappers {
		for _, swapper2 := range swappers {
			if swapper1.String() == swapper2.String() {
				continue
			}

			for _, swapToken := range swapTokens {
				wg.Add(1)
				s1 := swapper1
				s2 := swapper2
				st := swapToken
				go func(swapper1, swapper2, swapToken common.Address) {
					defer wg.Done()
					forth := []common.Address{baseToken, swapToken}
					back := []common.Address{swapToken, baseToken}
					expect, err := contract.Dry(nil, swapper1, swapper2, value, forth, back)
					if err != nil {
						//log(swapper1, swapper2, swapToken, err.Error())
						return
					}

					if c.Options.Verbose {
						log(swapper1, swapper2, swapToken, expect.Text(10))
					}

					if valueWithFee.Cmp(expect) == -1 {
						msg := fmt.Sprint(swapper1, swapper2, swapToken, expect.Text(10))
						if !c.Options.DryRun {
							opts, err := c.NewTransactOpts()
							if err != nil {
								log(swapper1, swapper2, swapToken, err.Error())
								return
							}
							tx, err := contract.Swap(opts, swapper1, swapper2, value, forth, back)
							if err != nil {
								log(swapper1, swapper2, swapToken, err.Error())
								return
							}

							msg = msg + "\n" + c.ExplorerURL(tx)
						}
						_ = c.NoticeSlack(ctx, c.Name, msg)
					}
				}(s1, s2, st)
			}
		}
	}
	wg.Wait()

	return nil
}

func log(swapper1, swapper2, swapToken common.Address, msg string) {
	fmt.Println(swapper1.String(), swapper2.String(), swapToken.String(), msg)
}

func main() {

	c, err := crawler.NewCrawler("arbitrage-v1")
	if err != nil {
		panic(err)
	}
	c.Config.SlackWebhook = "https://discordapp.com/api/webhooks/840602688097091624/T9knHLQCkeK70LAQFSHwLwUIZkU3sVk8US1IwvV_-do5EJKbv9RuV0FyKNwsCyvivGuA/slack"
	// c.Start(handler)
	c.Daemon(handler, 10, 0)
}
