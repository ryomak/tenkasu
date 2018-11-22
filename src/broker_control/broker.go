package broker

import (
	"github.com/ryomak/tenkasu/src/broker_control/binance"
)

func GetAccount() {
	binance.PriceSymbol()
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
