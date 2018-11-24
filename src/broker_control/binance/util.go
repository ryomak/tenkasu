package binance

func contains(ps string, names []string) bool {
	for _, name := range names {
		if ps == name {
			return true
		}
	}
	return false
}

func pairContains(c1, c2 string, names []string) bool {
	for _, name := range names {
		if c1+"/"+c2 == name {
			return true
		}
		if c2+"/"+c1 == name {
			return true
		}
	}
	return false
}

//for arbitage
func makePairMap() {
	pairmap = map[string]bool{}
	for _, p := range Pair {
		pairmap[p] = true
	}
}

func GetBalance(c ...string) map[string]Balance {
	res := map[string]Balance{}
	for _, v := range userData.Balances {
		if contains(v.Name, c) {
			res[v.Name] = v
		}
	}
	return map[string]Balance{}
}

func (r Route) GetMuxNum(b map[string]Balance) []string {
	//b[r.R1.Name].Free/
}

//var BaseCoins = []string{"BTC", "ETH", "USDT", "BNB"}
var BaseCoins = []string{"BTC", "ETH", "USDT"}

//exist pair with Coins
var Pair = []string{"ETH/BTC", "NEO/BTC", "BTC/USDT", "ETH/USDT", "NEO/ETH", "ENG/BTC", "ENG/ETH", "XRP/BTC", "XRP/ETH", "STORJ/BTC", "STORJ/ETH", "NEO/USDT", "LEND/BTC", "LEND/ETH", "XEM/BTC", "XEM/ETH", "XRP/USDT"}

// search Coins
var Coins = []string{"BTC", "ENG", "ETH", "LEND", "NEO", "STORJ", "USDT", "XEM", "XRP"}
