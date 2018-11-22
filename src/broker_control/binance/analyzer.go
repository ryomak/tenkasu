package binance

type PS struct {
	C1  string
	C2  string
	Ask float64
	Bid float64
}

type Route struct {
	Up  float64
	PSs []PS
}

/*
func Analyze(ps []PS) {
	first := ps[0]
	for _, p := range ps[1:] {

	}
}
*/
