package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	TradeBox := new(TradeBox)
	TradeBox.New()
	getAll := TradeBox.GetAll()
	fmt.Println(time.Now())

	for i := 0; i < len(getAll); i++ {
		fmt.Println(getAll["BTC "+strconv.Itoa(i)])
	}
	fmt.Println(time.Now())
}
