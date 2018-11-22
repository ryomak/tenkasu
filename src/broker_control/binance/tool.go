package binance

import (
	"context"

	binance "github.com/adshao/go-binance"
)

func GetPrice() ([]*binance.SymbolPrice, error) {
	return client_.NewListPricesService().Do(context.Background())
}

func GetTicker(str string) (*binance.BookTicker, error) {
	return client_.NewBookTickerService().Symbol(str).Do(context.Background())
}
