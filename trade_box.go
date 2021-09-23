package main

import (
	"strconv"
	"time"
)

type TradeBox struct {
	SymbolBoxs map[string]SymbolBox
}

func (this *TradeBox) Strart() {
	this.SymbolBoxs = make(map[string]SymbolBox)

	for i := 0; i < 20; i++ {
		var symbol = SymbolBox{SymbolName: "BTC " + strconv.Itoa(i),
			Price: 100000.0, CreateDateTime: time.Now().Unix(), LastUpdateDateTime: time.Now().Unix()}
		this.SymbolBoxs["BTC "+strconv.Itoa(i)] = symbol
	}
}
func (this *TradeBox) GetAllSymbol() map[string]SymbolBox {
	return this.SymbolBoxs
}
func (this *TradeBox) StopTrade() {

}

func (this *TradeBox) GetSymbolBox(SymbolName string) SymbolBox {
	return this.SymbolBoxs[SymbolName]
}
