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
	file, err := os.Create("src/broker_control/binance/Pair.txt")
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

	file1, err := os.Create("src/broker_control/binance/Coin.txt")
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
	for _, p := range Pair {
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

var Pair = []string{
	"ETH/BTC", "LTC/BTC", "BNB/BTC", "NEO/BTC", "QTUM/ETH", "EOS/ETH", "SNT/ETH", "BNT/ETH", "GAS/BTC", "BNB/ETH", "BTC/USDT", "ETH/USDT", "HSR/BTC", "OAX/ETH", "DNT/ETH", "MCO/ETH", "ICN/ETH", "MCO/BTC", "WTC/BTC", "WTC/ETH", "LRC/BTC", "LRC/ETH", "QTUM/BTC", "OMG/BTC", "OMG/ETH", "ZRX/BTC", "ZRX/ETH", "STRAT/BTC", "STRAT/ETH", "SNGLS/BTC", "SNGLS/ETH", "BQX/BTC", "BQX/ETH", "KNC/BTC", "KNC/ETH", "FUN/BTC", "FUN/ETH", "SNM/BTC", "SNM/ETH", "NEO/ETH", "IOTA/BTC", "IOTA/ETH", "LINK/BTC", "LINK/ETH", "XVG/BTC", "XVG/ETH", "SALT/BTC", "SALT/ETH", "MDA/BTC", "MDA/ETH", "MTL/BTC", "MTL/ETH", "SUB/BTC", "SUB/ETH", "EOS/BTC", "SNT/BTC", "ETC/ETH", "ETC/BTC", "MTH/BTC", "MTH/ETH", "ENG/BTC", "ENG/ETH", "DNT/BTC", "ZEC/BTC", "ZEC/ETH", "BNT/BTC", "AST/BTC", "AST/ETH", "DASH/BTC", "DASH/ETH", "OAX/BTC", "ICN/BTC", "BTG/BTC", "BTG/ETH", "EVX/BTC", "EVX/ETH", "REQ/BTC", "REQ/ETH", "VIB/BTC", "VIB/ETH", "HSR/ETH", "TRX/BTC", "TRX/ETH", "POWR/BTC", "POWR/ETH", "ARK/BTC", "ARK/ETH", "XRP/BTC", "XRP/ETH", "MOD/BTC", "MOD/ETH", "ENJ/BTC", "ENJ/ETH", "STORJ/BTC", "STORJ/ETH", "BNB/USDT", "VEN/BNB", "POWR/BNB", "VEN/BTC", "VEN/ETH", "KMD/BTC", "KMD/ETH", "NULS/BNB", "RCN/BTC", "RCN/ETH", "RCN/BNB", "NULS/BTC", "NULS/ETH", "RDN/BTC", "RDN/ETH", "RDN/BNB", "XMR/BTC", "XMR/ETH", "DLT/BNB", "WTC/BNB", "DLT/BTC", "DLT/ETH", "AMB/BTC", "AMB/ETH", "AMB/BNB", "BAT/BTC", "BAT/ETH", "BAT/BNB", "BCPT/BTC", "BCPT/ETH", "BCPT/BNB", "ARN/BTC", "ARN/ETH", "GVT/BTC", "GVT/ETH", "CDT/BTC", "CDT/ETH", "GXS/BTC", "GXS/ETH", "NEO/USDT", "NEO/BNB", "POE/BTC", "POE/ETH", "QSP/BTC", "QSP/ETH", "QSP/BNB", "BTS/BTC", "BTS/ETH", "BTS/BNB", "XZC/BTC", "XZC/ETH", "XZC/BNB", "LSK/BTC", "LSK/ETH", "LSK/BNB", "TNT/BTC", "TNT/ETH", "FUEL/BTC", "FUEL/ETH", "MANA/BTC", "MANA/ETH", "BCD/BTC", "BCD/ETH", "DGD/BTC", "DGD/ETH", "IOTA/BNB", "ADX/BTC", "ADX/ETH", "ADX/BNB", "ADA/BTC", "ADA/ETH", "PPT/BTC", "PPT/ETH", "CMT/BTC", "CMT/ETH", "CMT/BNB", "XLM/BTC", "XLM/ETH", "XLM/BNB", "CND/BTC", "CND/ETH", "CND/BNB", "LEND/BTC", "LEND/ETH", "WABI/BTC", "WABI/ETH", "WABI/BNB", "LTC/ETH", "LTC/USDT", "LTC/BNB", "TNB/BTC", "TNB/ETH", "WAVES/BTC", "WAVES/ETH", "WAVES/BNB", "GTO/BTC", "GTO/ETH", "GTO/BNB", "ICX/BTC", "ICX/ETH", "ICX/BNB", "OST/BTC", "OST/ETH", "OST/BNB", "ELF/BTC", "ELF/ETH", "AION/BTC", "AION/ETH", "AION/BNB", "NEBL/BTC", "NEBL/ETH", "NEBL/BNB", "BRD/BTC", "BRD/ETH", "BRD/BNB", "MCO/BNB", "EDO/BTC", "EDO/ETH", "WINGS/BTC", "WINGS/ETH", "NAV/BTC", "NAV/ETH", "NAV/BNB", "LUN/BTC", "LUN/ETH", "TRIG/BTC", "TRIG/ETH", "TRIG/BNB", "APPC/BTC", "APPC/ETH", "APPC/BNB", "VIBE/BTC", "VIBE/ETH", "RLC/BTC", "RLC/ETH", "RLC/BNB", "INS/BTC", "INS/ETH", "PIVX/BTC", "PIVX/ETH", "PIVX/BNB", "IOST/BTC", "IOST/ETH", "CHAT/BTC", "CHAT/ETH", "STEEM/BTC", "STEEM/ETH", "STEEM/BNB", "VIA/BTC", "VIA/ETH", "VIA/BNB", "BLZ/BTC", "BLZ/ETH", "BLZ/BNB", "AE/BTC", "AE/ETH", "AE/BNB", "RPX/BTC", "RPX/ETH", "RPX/BNB", "NCASH/BTC", "NCASH/ETH", "NCASH/BNB", "POA/BTC", "POA/ETH", "POA/BNB", "ZIL/BTC", "ZIL/ETH", "ZIL/BNB", "ONT/BTC", "ONT/ETH", "ONT/BNB", "STORM/BTC", "STORM/ETH", "STORM/BNB", "QTUM/BNB", "QTUM/USDT", "XEM/BTC", "XEM/ETH", "XEM/BNB", "WAN/BTC", "WAN/ETH", "WAN/BNB", "WPR/BTC", "WPR/ETH", "QLC/BTC", "QLC/ETH", "SYS/BTC", "SYS/ETH", "SYS/BNB", "QLC/BNB", "GRS/BTC", "GRS/ETH", "ADA/USDT", "ADA/BNB", "CLOAK/BTC", "CLOAK/ETH", "GNT/BTC", "GNT/ETH", "GNT/BNB", "LOOM/BTC", "LOOM/ETH", "LOOM/BNB", "XRP/USDT", "BCN/BTC", "BCN/ETH", "BCN/BNB", "REP/BTC", "REP/ETH", "REP/BNB", "TUSD/BTC", "TUSD/ETH", "TUSD/BNB", "ZEN/BTC", "ZEN/ETH", "ZEN/BNB", "SKY/BTC", "SKY/ETH", "SKY/BNB", "EOS/USDT", "EOS/BNB", "CVC/BTC", "CVC/ETH", "CVC/BNB", "THETA/BTC", "THETA/ETH", "THETA/BNB", "XRP/BNB", "TUSD/USDT", "IOTA/USDT", "XLM/USDT", "IOTX/BTC", "IOTX/ETH", "QKC/BTC", "QKC/ETH", "AGI/BTC", "AGI/ETH", "AGI/BNB", "NXS/BTC", "NXS/ETH", "NXS/BNB", "ENJ/BNB", "DATA/BTC", "DATA/ETH", "ONT/USDT", "TRX/USDT", "ETC/USDT", "ETC/BNB", "ICX/USDT", "SC/BTC", "SC/ETH", "SC/BNB", "NPXS/BTC", "NPXS/ETH", "VEN/USDT", "KEY/BTC", "KEY/ETH", "NAS/BTC", "NAS/ETH", "NAS/BNB", "MFT/BTC", "MFT/ETH", "MFT/BNB", "DENT/BTC", "DENT/ETH"}

var Coins = []string{"ADA", "ADX", "AE", "AGI", "AION", "AMB", "APPC", "ARK", "ARN", "AST", "BAT", "BCD", "BCN", "BCPT", "BLZ", "BNB", "BNT", "BQX", "BRD", "BTC", "BTG", "BTS", "CDT", "CHAT", "CLOAK", "CMT", "CND", "CVC", "DASH", "DATA", "DENT", "DGD", "DLT", "DNT", "EDO", "ELF", "ENG", "ENJ", "EOS", "ETC", "ETH", "EVX", "FUEL", "FUN", "GAS", "GNT", "GRS", "GTO", "GVT", "GXS", "HSR", "ICN", "ICX", "INS", "IOST", "IOTA", "IOTX", "KEY", "KMD", "KNC", "LEND", "LINK", "LOOM", "LRC", "LSK", "LTC", "LUN", "MANA", "MCO", "MDA", "MFT", "MOD", "MTH", "MTL", "NAS", "NAV", "NCASH", "NEBL", "NEO", "NPXS", "NULS", "NXS", "OAX", "OMG", "ONT", "OST", "PIVX", "POA", "POE", "POWR", "PPT", "QKC", "QLC", "QSP", "QTUM", "RCN", "RDN", "REP", "REQ", "RLC", "RPX", "SALT", "SC", "SKY", "SNGLS", "SNM", "SNT", "STEEM", "STORJ", "STORM", "STRAT", "SUB", "SYS", "THETA", "TNB", "TNT", "TRIG", "TRX", "TUSD", "USDT", "VEN", "VIA", "VIB", "VIBE", "WABI", "WAN", "WAVES", "WINGS", "WPR", "WTC", "XEM", "XLM", "XMR", "XRP", "XVG", "XZC", "ZEC", "ZEN", "ZIL", "ZRX"}
