package main

type SymbolBox struct {
	CreateDateTime     int64   `json:"createDateTime"`
	LastUpdateDateTime int64   `json:"lastUpdateDateTime"`
	SymbolName         string  `json:"symbolName"`
	Price              float64 `json:"price"`
}
