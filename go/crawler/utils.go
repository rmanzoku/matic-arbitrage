package crawler

import (
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
)

func EtherToWei(ether float64) (*big.Int, error) {
	etherDecimal := decimal.NewFromFloat(ether)
	base := decimal.NewFromInt(params.Ether)

	retDecimal := etherDecimal.Mul(base).Floor()
	ret, ok := new(big.Int).SetString(retDecimal.String(), 10)
	if !ok {
		return nil, errors.New("Invalit number " + retDecimal.String())
	}
	return ret, nil
}

func GweiToWei(gwei float64) (*big.Int, error) {
	gweiDecimal := decimal.NewFromFloat(gwei)
	base := decimal.NewFromInt(params.GWei)

	retDecimal := gweiDecimal.Mul(base).Floor()
	ret, ok := new(big.Int).SetString(retDecimal.String(), 10)
	if !ok {
		return nil, errors.New("Invalit number " + retDecimal.String())
	}
	return ret, nil
}

// ToGwei ...
func ToGwei(wei *big.Int) float64 {
	weiDecimal, _ := decimal.NewFromString(wei.String())
	base := decimal.NewFromInt(params.GWei)
	ret, _ := weiDecimal.Div(base).Float64()
	return ret
}

// ToEther ...
func ToEther(wei *big.Int) *big.Int {
	return new(big.Int).Quo(wei, big.NewInt(params.Ether))
}

func EncodeToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func DecodeHex(s string) ([]byte, error) {
	if s[0:2] != "0x" {
		return nil, errors.New("hex must start with 0x")
	}
	return hex.DecodeString(s[2:])
}
