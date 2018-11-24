package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	b "github.com/adshao/go-binance"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

var client_ *b.Client
var listenKey string
var userData struct {
	Type     string    `json:"e"`
	Balances []Balance `json:"B"`
}

type Balance struct {
	Name string  `json:"a"`
	Free float64 `json:"f"`
	Lock float64 `json:"l"`
}

func Init(broker conf.Broker) {
	client_ = b.NewClient(broker.APIKey, broker.SecretKey)
	InitListenKey()
}

func GetAccount() (*b.Account, error) {
	return client_.NewGetAccountService().Do(context.Background())
}

func GetPrice() ([]*b.SymbolPrice, error) {
	return client_.NewListPricesService().Do(context.Background())
}

func GetTicker(str string) (*b.BookTicker, error) {
	return client_.NewBookTickerService().Symbol(str).Do(context.Background())
}

func GetAllOrder(symbol string) ([]*b.Order, error) {
	return client_.NewListOrdersService().Symbol(symbol).Do(context.Background())
}

func GetOrder(symbol string, id int64) (*b.Order, error) {
	return client_.NewGetOrderService().Symbol(symbol).OrderID(id).Do(context.Background())
}

func BuyMarketOrder(symbol, amount string) (*b.CreateOrderResponse, error) {
	return client_.NewCreateOrderService().Symbol(symbol).
		Side(b.SideTypeBuy).Type(b.OrderTypeMarket).
		TimeInForce(b.TimeInForceGTC).Quantity(amount).Do(context.Background())
}

func SellMarketOrder(symbol, amount string) (*b.CreateOrderResponse, error) {
	return client_.NewCreateOrderService().Symbol(symbol).
		Side(b.SideTypeSell).Type(b.OrderTypeMarket).
		Quantity(amount).Do(context.Background())
}

func InitListenKey() {
	key, err := client_.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		panic(err)
	}
	listenKey = key
}

func SetAccount() {
	//websocket
	wsHandler := func(message []byte) {
		if err := json.Unmarshal(bytes, &userData); err != nil {
			log.Println(err)
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := b.WsUserDataServe(listenKey, wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
