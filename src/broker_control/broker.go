package broker

import (
	"fmt"

	"github.com/ryomak/tenkasu/src/broker_control/binance"
)

func GetAccount() {
	/*a, err := binance.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	b, err := bitflyer.GetBalance()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	*/
}

func GetRoute() {
	for _, v := range binance.SearchRoute() {
		fmt.Printf("%+v\n", v)
	}
}
