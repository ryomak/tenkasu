package main

import (
	"os"

	"github.com/ryomak/tenkasu/src/broker_control/binance"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

func init() {
	var brokers map[string]conf.Broker
	conf.LoadConfig("config.toml")
	brokers = conf.GetBrokers()
	binance.Init(brokers["binance"])
}

func main() {
	writeRoute()
}

func writeRoute() {
	routes := SearchRoute()
	file, err := os.Create("Route.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, r := range routes {
		file.Write(([]byte)(r.R1 + "," + r.R2 + "," + r.R3 + "\n"))
	}
}

func SearchRoute() []Route {
	routes := make([]Route, 0)
	for _, bc := range BaseCoins {
		t0 := bc
		for _, t1 := range Coins {
			if pairContains(t0, t1, Pair) {
				for _, t2 := range Coins {
					if pairContains(t1, t2, Pair) {
						if pairContains(t2, t0, Pair) {
							routes = append(routes, Route{t0, t1, t2, 0.0})
						}
					}
				}
			}
		}
	}
	return routes
}
