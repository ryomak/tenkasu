package binance

import (
	"context"

	b "github.com/adshao/go-binance"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

var client_ *b.Client

func Init(broker conf.Broker) {
	client_ = b.NewClient(broker.APIKey, broker.SecretKey)
}

func GetAccount() (*b.Account, error) {
	return client_.NewGetAccountService().Do(context.Background())
}
