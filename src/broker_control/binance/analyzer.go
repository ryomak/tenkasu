package binance

import (
	"fmt"
	"strconv"
)

type PS struct {
	C1  string
	C2  string
	Ask float64
	Bid float64
}

type Route struct {
	R1     string
	R2     string
	R3     string
	Result float64
}

/*
func Analyze(ps []PS) {
	first := ps[0]
	for _, p := range ps[1:] {

	}
}
*/
func calc(t0, t1, t2 string) float64 {
	var result float64
	if contains(t0+"/"+t1, Pair) {
		t, err := GetTicker(t0 + t1)
		if err != nil {
			fmt.Println("bid:", t0, ":", t1, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.BidPrice, 64)
		result = tmp
	} else {
		t, err := GetTicker(t1 + t0)
		if err != nil {
			fmt.Println("ask:", t1, ":", t0, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.AskPrice, 64)
		result = 1 / tmp
	}
	if contains(t1+"/"+t2, Pair) {
		t, err := GetTicker(t1 + t2)
		if err != nil {
			fmt.Println("bid:", t0, ":", t1, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.BidPrice, 64)
		result = result * tmp
	} else {
		t, err := GetTicker(t2 + t1)
		if err != nil {
			fmt.Println("ask:", t2, ":", t1, err)
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.AskPrice, 64)
		result = result / tmp
	}
	if contains(t2+"/"+t0, Pair) {
		t, err := GetTicker(t2 + t0)
		if err != nil {
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.BidPrice, 64)
		result = result * tmp
	} else {
		t, err := GetTicker(t0 + t2)
		if err != nil {
			return -1000.0
		}
		tmp, _ := strconv.ParseFloat(t.AskPrice, 64)
		result = result / tmp
	}
	return result
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
							routes = append(routes, Route{t0, t1, t2, calc(t0, t1, t2)})
						}
					}
				}
			}
		}
	}
	return routes
}
