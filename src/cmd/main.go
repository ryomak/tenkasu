package main

import (
	"flag"

	"github.com/ryomak/tenkasu/src/broker_control"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

var (
	configPath = flag.String("conf", "config.toml", "config toml")
)

func init() {
	flag.Parse()
	conf.LoadConfig(*configPath)
	broker.InitBroker()
}

func main() {
	broker.GetAccount()
}
