package main

import (
	"fmt"
)

func main() {
	a := ""
	for _, p := range Pair {
		if contains(Coins, p) {
			a += ",\"" + p + "\""
		}
	}
	fmt.Println(a)
}
