package binance

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var routes []Route
var pairmap map[string]bool

type Route struct {
	R1     *Order
	R2     *Order
	R3     *Order
	Result float64
}

type Order struct {
	Name  string
	Type  string //next action
	Price float64
}

func init() {
	readRoute("Route.txt")
	makePairMap()
}

func (r Route) Start() {
  firstAmount:= "0"
	r1 := r.R1.Name
	r2 := r.R2.Name
	r3 := r.R3.Name
	GetBalance([]string{r1, r2, r3})
	if r.R1.Type == "buy" {
		BuyMarketOrder(r2+r1, firstAmount)
	} else {
		SellMarketOrder(r1+r2, firstAmount)
	}
	time.Sleep(1)
	if r.R2.Type == "buy" {
		BuyMarketOrder(r3+r2, firstAmount)
	} else {
		SellMarketOrder(r2+r3, firstAmount)
	}
	time.Sleep(1)
	if r.R3.Type == "buy" {
		BuyMarketOrder(r1+r3, firstAmount)
	} else {
		SellMarketOrder(r3+r1, firstAmount)
	}
}

func (o *Order) setRecentData(str string, price float64) {
	o.Type = str
	o.Price = price
}

func WinRoute(plus float64) []Route {
	res := make([]Route, 0)
	for _, r := range routes {
		result := r.calc()
		if result > (1 + plus) {
			r.Result = result
			res = append(res, r)
		}
	}
	return res
}

func (route *Route) calc() float64 {
	var result float64
	t0 := route.R1.Name
	t1 := route.R2.Name
	t2 := route.R3.Name
	if pairmap[t1+"/"+t0] {
		t, err := GetTicker(t1 + t0)
		if err != nil {
			fmt.Println("bid:", t0, ":", t1, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.BidPrice, 64)
		route.R1.setRecentData("buy", tmp)
		result = tmp
	} else {
		route.R1.Type = "sell"
		t, err := GetTicker(t0 + t1)
		if err != nil {
			fmt.Println("ask:", t1, ":", t0, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.AskPrice, 64)
		route.R1.setRecentData("sell", tmp)
		result = 1 / tmp
	}
	if pairmap[t2+"/"+t1] {
		route.R2.Type = "buy"
		t, err := GetTicker(t2 + t1)
		if err != nil {
			fmt.Println("bid:", t0, ":", t1, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.BidPrice, 64)
		route.R2.setRecentData("buy", tmp)
		result = result * tmp
	} else {
		route.R2.Type = "sell"
		t, err := GetTicker(t1 + t2)
		if err != nil {
			fmt.Println("ask:", t2, ":", t1, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.AskPrice, 64)
		route.R2.setRecentData("sell", tmp)
		result = result / tmp
	}
	if pairmap[t0+"/"+t2] {
		route.R3.Type = "buy"
		t, err := GetTicker(t0 + t2)
		if err != nil {
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.BidPrice, 64)
		result = result * tmp
		route.R3.setRecentData("buy", tmp)
	} else {
		route.R3.Type = "sell"
		t, err := GetTicker(t2 + t0)
		if err != nil {
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.AskPrice, 64)
		result = result / tmp
		route.R3.setRecentData("sell", tmp)
	}
	return result
}

func readRoute(path string) {
	fp, err := os.Open("Route.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		} else if err != nil {
			panic(err)
		}
		rs := strings.Split(string(line), ",")
		routes = append(routes, Route{&Order{Name: rs[0]}, &Order{Name: rs[1]}, &Order{Name: rs[2]}, 0.0})
	}
}
