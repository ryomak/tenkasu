package broker

import (
	"github.com/ryomak/tenkasu/src/broker_control/binance"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

func InitBroker() {
	brokers := conf.GetBrokers()
	binance.Init(brokers["binance"])
}
