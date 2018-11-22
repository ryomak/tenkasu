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

func PriceSymbol(names ...string) ([]b.SymbolPrice, error) {
	prices, err := client_.NewListPricesService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	res := make([]b.SymbolPrice, len(names))
	for _, p := range prices {
		if contains(p, names) {
			res = append(res, *p)
		}
	}
	return res, nil
}
