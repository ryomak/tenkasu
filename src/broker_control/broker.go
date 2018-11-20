package broker

import (
	"fmt"

	"github.com/ryomak/tenkasu/src/broker_control/binance"
)

func GetAccount() {
	a, err := binance.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
