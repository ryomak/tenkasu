package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ryomak/tenkasu/src/broker_control/binance"
	"github.com/ryomak/tenkasu/src/lib/conf"
)

var brokers map[string]conf.Broker

func init() {
	conf.LoadConfig("config.toml")
	brokers = conf.GetBrokers()
	binance.Init(brokers["binance"])
}

func main() {
	writeSymbol()
}

func writeSymbol() {
	ps := makeSymbol()
	file, err := os.Create("Pair.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rs := make([]string, len(ps)*2)
	for i, p := range ps {
		rs[2*i] = p.C1
		rs[2*i+1] = p.C2
		str := p.C1 + p.C2 + "," + p.C1 + "," + p.C2 + "," + strconv.FormatFloat(p.Ask, 'f', 4, 64) + "," + strconv.FormatFloat(p.Bid, 'f', 4, 64) + "\n"
		file.Write(([]byte)(str))
	}

	file1, err := os.Create("Coin.txt")
	if err != nil {
		panic(err)
	}
	defer file1.Close()
	m := make(map[string]struct{}) // 空のstructを使う
	for _, ele := range rs {
		m[ele] = struct{}{}
	}
	uniq := []string{}
	for i := range m {
		uniq = append(uniq, i)
	}
	for _, v := range uniq {
		file1.Write(([]byte)(v + ","))
	}
}

func makeSymbol() []binance.PS {
	ps := make([]binance.PS, 0)
	for _, p := range binance.Pair {
		tmp := strings.Split(p, "/")
		t, err := binance.GetTicker(tmp[0] + tmp[1])
		if err != nil {
			fmt.Println(p)
			fmt.Println(err)
			continue
		}
		bid, _ := strconv.ParseFloat(t.BidPrice, 64)
		ask, _ := strconv.ParseFloat(t.AskPrice, 64)
		ps = append(ps, binance.PS{tmp[0], tmp[1], ask, bid})
	}
	return ps
}
