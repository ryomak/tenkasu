package broker

import (
	"github.com/ryomak/tenkasu/src/broker_control/binance"
	"github.com/ryomak/tenkasu/src/broker_control/bitflyer"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

func Init() {
	brokers := conf.GetBrokers()
	binance.Init(brokers["binance"])
	bitflyer.Init(brokers["bitflyer"])
}
