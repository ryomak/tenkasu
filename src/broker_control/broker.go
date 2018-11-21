package broker

import (
	"fmt"

	"github.com/ryomak/tenkasu/src/broker_control/binance"
	"github.com/ryomak/tenkasu/src/broker_control/bitflyer"
)

func GetAccount() {
	a, err := binance.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	b, err := bitflyer.GetBalance()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
