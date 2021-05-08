package crawler

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v6"
	slack "github.com/catatsuy/notify_slack/slack"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// const (
// 	nilAddressString = "0x0000000000000000000000000000000000000000"
// )

type config struct {
	EthRPC          string  `env:"ETH_RPC" envDefault:"https://rpc-mainnet.matic.network"`
	PrivateKey      string  `env:"PRIV_KEY"`
	FixGasPriceGwei float64 `env:"FIX_GWEI" envDefault:"1.000001"`
	SlackWebhook    string  `env:"SLACK_WEBHOOK" envDefault:"https://discordapp.com/api/webhooks/742379201914601504/CrCFkASWoMvpzM7N-U-1rL6Q9NrepJXXOKRk8gt_h4_PbvTZJ2OthoH9y-xpRt1EEQ_2/slack"`
}

type Crawler struct {
	Name      string
	Config    *config
	Options   *Options
	Args      []string
	EthClient *ethclient.Client
	ChainId   *big.Int
}

type Options struct {
	DryRun bool `json:"dry_run"`
}

var DefaultOptions = Options{
	DryRun: false,
}

func UnmarshalOptions(data []byte) (Options, error) {
	r := DefaultOptions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Options) Marshal() ([]byte, error) {
	return json.Marshal(r)

}

func NewOptions() (o *Options) {
	ret := &Options{}
	flag.BoolVar(&ret.DryRun, "dry-run", DefaultOptions.DryRun, "Dry run")
	flag.Parse()
	return ret
}

func NewCrawler(name string) (*Crawler, error) {
	var err error
	ret := new(Crawler)
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	ret.Config = cfg

	ret.Name = name

	if cfg.EthRPC != "" {
		rpcClient, err := rpc.DialHTTP(cfg.EthRPC)
		if err != nil {
			return nil, err
		}
		//rpcClient.SetHeader("User-Agent", name)
		ret.EthClient = ethclient.NewClient(rpcClient)

		ret.ChainId, err = ret.EthClient.ChainID(context.TODO())
		if err != nil {
			return nil, err
		}
	}

	if ret.Options == nil {
		ret.Options = NewOptions()
	}

	return ret, err
}

type HandlerFunc = func(context.Context, *Crawler) error

func (c *Crawler) Start(handler HandlerFunc) {

	h := Notification(Recover(handler))
	if os.Getenv("AWS_EXECUTION_ENV") != "" {
		lambda.Start((func(ctx context.Context) error {
			return h(ctx, c)
		}))
	} else {
		if err := h(context.Background(), c); err != nil {
			log.Fatalf("%+v", err)
		}
	}
}

func (c *Crawler) Daemon(handler HandlerFunc, interval int64, timeout int64) {
	h := Notification(Recover(handler))

	since := time.Now().Unix()
	loop := func(ctx context.Context, c *Crawler) error {
		for {
			now := time.Now().Unix()
			if now > since+timeout {
				return nil
			}

			err := h(ctx, c)
			if err != nil {
				return err
			}
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}
	if os.Getenv("AWS_EXECUTION_ENV") != "" {
		log.Fatal(errors.New("Daemon not support on lambda"))
	} else {
		if err := loop(context.Background(), c); err != nil {
			log.Fatal(err)
		}
	}
}

func Recover(h HandlerFunc) HandlerFunc {
	return func(ctx context.Context, c *Crawler) (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = errors.New(fmt.Sprint("Panic:", e))
			}
		}()
		err = h(ctx, c)
		return err
	}
}

func Notification(h HandlerFunc) HandlerFunc {
	return func(ctx context.Context, c *Crawler) (err error) {
		err = h(ctx, c)
		if err != nil {
			msg := fmt.Sprintf("ERROR\n%s\n", err.Error())
			err2 := c.NoticeSlack(ctx, "Error from "+c.Name, msg)
			if err2 != nil {
				err = fmt.Errorf(err2.Error()+": %w", err)
			}
		}
		return err
	}
}

func (c *Crawler) NoticeSlack(ctx context.Context, name, msg string) (err error) {
	if c.Config.SlackWebhook == "" {
		fmt.Println(name, msg)
		return nil
	}

	cli, err := slack.NewClient(c.Config.SlackWebhook, nil)
	if err != nil {
		return
	}
	text := "```"
	text += msg
	text += "```"

	param := &slack.PostTextParam{
		Username: name,
		Text:     text,
	}
	return cli.PostText(ctx, param)
}

func (c *Crawler) WaitTransaction(ctx context.Context, tx *types.Transaction, conf uint64, timeoutSec int64) (err error) {
	startAt := time.Now().Unix()
	hash := tx.Hash()
	for {
		now := time.Now().Unix()
		if startAt+timeoutSec < now {
			return errors.New("wait transaction timeout: " + c.ExplorerURL(tx))
		}
		time.Sleep(3 * time.Second)

		_, isPending, err := c.EthClient.TransactionByHash(ctx, hash)
		if err != nil {
			return err
		}
		if isPending {
			continue
		}

		receipt, err := c.EthClient.TransactionReceipt(ctx, hash)
		if err != nil {
			return err
		}

		block, err := c.EthClient.BlockByNumber(ctx, nil)
		if err != nil {
			return err
		}
		if receipt.BlockNumber.Uint64()+conf > block.NumberU64() {
			continue
		}

		//EIP-658 0 indicating failure and 1 indicating success.
		if receipt.Status == 0 {
			return errors.New("transaction failure: " + c.ExplorerURL(tx))
		} else {
			return nil
		}
	}
}

func (c Crawler) NewTransactOpts() (*bind.TransactOpts, error) {
	key, err := crypto.HexToECDSA(c.Config.PrivateKey)
	if err != nil {
		return nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(key, c.ChainId)
	if err != nil {
		return nil, err
	}

	return c.transactOptsConf(opts)
}

// func (c Crawler) NewKMSTransactOpts(keyid string) (*bind.TransactOpts, error) {
// 	svc, err := kmsutil.NewKMSClient()
// 	if err != nil {
// 		return nil, err
// 	}

// 	opts, err := awseoa.NewKMSTransactor(svc, keyid, c.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return c.transactOptsConf(opts)
// }

// func (c Crawler) NewKeyStoreTransactOpts() (*bind.TransactOpts, error) {
// 	key, err := ioutil.ReadFile(c.Config.Keystore)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if c.Config.Password == "" {
// 		return nil, errors.New("Set keystore password by env PASSWORD")
// 	}
// 	opts, err := bind.NewTransactorWithChainID(bytes.NewReader(key), c.Config.Password, c.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return c.transactOptsConf(opts)
// }

func (c Crawler) transactOptsConf(opts *bind.TransactOpts) (*bind.TransactOpts, error) {
	opts.Context = context.TODO()
	var err error

	opts.GasPrice, err = GweiToWei(c.Config.FixGasPriceGwei)
	if err != nil {
		return nil, err
	}

	// n, err := c.EthClient.NonceAt(context.TODO(), opts.From, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// opts.Nonce = big.NewInt(0).SetUint64(n)

	return opts, nil
}

func (c Crawler) GasPrice(coefficient float64, upperGwei float64) (*big.Int, error) {

	suggestGasPrice, err := c.EthClient.SuggestGasPrice(context.TODO())
	if err != nil {
		return nil, err
	}
	sugest := ToGwei(suggestGasPrice)
	gasPrice := sugest * coefficient

	if gasPrice > upperGwei {
		return nil, fmt.Errorf("Gas price too expensive %f * %f > %f",
			sugest, coefficient, upperGwei,
		)
	}

	return GweiToWei(gasPrice)
}

func (c *Crawler) ExplorerURL(tx *types.Transaction) string {
	switch c.ChainId.Int64() {
	case 1:
		return "https://etherscan.io/tx/" + tx.Hash().String()
	case 4:
		return "https://rinkeby.etherscan.io/tx/" + tx.Hash().String()
	case 137:
		return "https://explorer-mainnet.maticvigil.com/tx/" + tx.Hash().String()
	case 80001:
		return "https://mumbai-explorer.matic.today/tx/" + tx.Hash().String()
	default:
		return fmt.Sprintf("chainId:%d\ttx:%s", tx.ChainId().Int64(), tx.Hash().String())
	}
}
