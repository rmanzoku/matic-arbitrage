package main

import (
	"context"

	"github.com/rmanzoku/matic-arbitrage/go/crawler"
)

func handler(ctx context.Context, c *crawler.Crawler) (err error) {
	c.NoticeSlack(ctx, c.Name, "hello")
	return nil
}

func main() {

	c, err := crawler.NewCrawler("arbitrage-v1")
	if err != nil {
		panic(err)
	}
	c.Start(handler)
}
