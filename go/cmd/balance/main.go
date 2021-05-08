package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rmanzoku/matic-arbitrage/go/contracts/wmatic"
	"github.com/rmanzoku/matic-arbitrage/go/crawler"
)

type Address struct {
	Name    string
	Address string
	WMATIC  bool
}

func (a Address) ToAddress() common.Address {
	return common.HexToAddress(a.Address)
}

var (
	addresses = []Address{
		{"Transactor", "0xCFf7E558C192e135dbEE190254D9eFa978Bee41B", false},
		{"V1", "0xc86ba4527797f46569bEb67fb7fE2F0B2E6d1fB8", true},
	}
)

func handler(ctx context.Context, c *crawler.Crawler) (err error) {
	contract, err := wmatic.NewWmatic(crawler.AddrWMATIC, c.EthClient)
	if err != nil {
		return err
	}

	msg := ""

	sum := big.NewInt(0)
	for _, a := range addresses {
		balance := big.NewInt(0)
		tag := ""

		if a.WMATIC {
			balance, err = contract.BalanceOf(nil, a.ToAddress())
			if err != nil {
				return err
			}
			tag = "(W)"
		} else {
			balance, err = c.EthClient.BalanceAt(ctx, a.ToAddress(), nil)
			if err != nil {
				return err
			}
		}

		sum = big.NewInt(0).Add(sum, balance)
		msg += fmt.Sprintln(a.Name+tag, crawler.ToEther(balance))
	}
	msg += fmt.Sprintln("SUM", crawler.ToEther(sum))

	return c.NoticeSlack(ctx, c.Name, msg)
}

func main() {
	c, err := crawler.NewCrawler("balance")
	if err != nil {
		panic(err)
	}
	c.Config.SlackWebhook = "https://discordapp.com/api/webhooks/840613416930246667/o1nHfMy-JpAAqaaAm3XIk9TxKZzNgsFR99xCeEsEQy-J6A9EysDQPwwvEDHu26URJENK/slack"
	c.Start(handler)
}
