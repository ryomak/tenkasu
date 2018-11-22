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

var BaseCoins = []string{"BTC", "ETH", "USDT", "BNB"}

var Pair = []string{"ETH/BTC", "LTC/BTC", "BNB/BTC", "NEO/BTC", "EOS/ETH", "BNB/ETH", "BTC/USDT", "ETH/USDT", "DNT/ETH", "NEO/ETH", "IOTA/BTC", "IOTA/ETH", "EOS/BTC", "ETC/ETH", "ETC/BTC", "DNT/BTC", "REQ/BTC", "REQ/ETH", "XRP/BTC", "XRP/ETH", "ENJ/BTC", "ENJ/ETH", "STORJ/BTC", "STORJ/ETH", "BNB/USDT", "NEO/USDT", "NEO/BNB", "IOTA/BNB", "LEND/BTC", "LEND/ETH", "LTC/ETH", "LTC/USDT", "LTC/BNB", "RLC/BTC", "RLC/ETH", "RLC/BNB", "XEM/BTC", "XEM/ETH", "XEM/BNB", "CLOAK/BTC", "CLOAK/ETH", "LOOM/BTC", "LOOM/ETH", "LOOM/BNB", "XRP/USDT", "TUSD/BTC", "TUSD/ETH", "TUSD/BNB", "EOS/USDT", "EOS/BNB", "XRP/BNB", "TUSD/USDT", "IOTA/USDT", "IOTX/BTC", "IOTX/ETH", "ENJ/BNB", "ETC/USDT", "ETC/BNB", "KEY/BTC", "KEY/ETH", "DENT/BTC", "DENT/ETH"}

var Coins = []string{"BNB", "BTC", "CLOAK", "DENT", "DNT", "ENJ", "EOS", "ETC", "ETH", "IOTA", "IOTX", "KEY", "LEND", "LOOM", "LTC", "NEO", "REQ", "RLC", "STORJ", "TUSD", "USDT", "XEM", "XRP"}
