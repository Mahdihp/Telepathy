package main

import (
	"strconv"
)

type TradeBox struct {
	SymbolBoxs map[string]SymbolBox
}

func (this *TradeBox) Strart() {
	mapp := make(map[string]SymbolBox)

	for i := 0; i < 20; i++ {
		var sym = SymbolBox{SymbolName: "BTC " + strconv.Itoa(i), Price: 100000.0}
		mapp["BTC "+strconv.Itoa(i)] = sym
	}
	this.SymbolBoxs = mapp
}
func (this *TradeBox) GetAll() map[string]SymbolBox {
	return this.SymbolBoxs
}
func (this *TradeBox) StopTrade() {

}

func (this *TradeBox) GetSymbolBox(SymbolName string) SymbolBox {
	return this.SymbolBoxs[SymbolName]
}
