package data

import . "github.com/XuVic/tw_stock/helper"

type BasicInfo struct {
	StockID, StockName, Industry, ComName, Listed, Chairman, Business string
	Cap, MarketCap, Shares, PreShares                                 float64
}

func NewBasicInfo(source interface{}) *BasicInfo {
	data := source.(map[string]string)
	return &BasicInfo{
		StockID: data["stockID"], StockName: data["stockName"], Industry: data["industry"],
		ComName: data["conName"], Chairman: data["chairman"], Business: data["business"],
		Cap: ToFloat(data["cap"], true, false), MarketCap: ToFloat(data["marketCap"], true, true),
		Shares: ToFloat(data["shares"], false, false), PreShares: ToFloat(data["preShares"], false, false),
	}
}
