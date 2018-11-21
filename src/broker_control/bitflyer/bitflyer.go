package bitflyer

import (
	bitflyer "github.com/ryomak/go-bitflyer"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

var client_ *bitflyer.Client

func Init(b conf.Broker) {
	client_ = bitflyer.NewClient(b.APIKey, b.SecretKey)
}

func GetBalance() ([]bitflyer.Balance, error) {
	return client_.GetBalance()
}
