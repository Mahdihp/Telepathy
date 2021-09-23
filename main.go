package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	TradeBox := new(TradeBox)
	TradeBox.Strart()
	getAll := TradeBox.GetAllSymbol()
	fmt.Println(time.Now(), "--------------------")

	for i := 0; i < len(getAll); i++ {
		fmt.Println(getAll["BTC "+strconv.Itoa(i)])
	}
	fmt.Println(time.Now(), "--------------------")
}
