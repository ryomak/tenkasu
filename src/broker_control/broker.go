package broker

import (
	"github.com/ryomak/tenkasu/src/broker_control/binance"
  "fmt"
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
	go binance.SetAccount()
	se := new(binance.Route)
	for _, v := range binance.WinRoute(0.000003) {
    fmt.Println(v.R1.Name+"=>"+v.R2.Name+"=>"+v.R3.Name+":",v.Result)
		if v.Result > se.Result && v.R1.Name == "ETH" {
			se = &v
		}
	}
  //se.Start()
}
