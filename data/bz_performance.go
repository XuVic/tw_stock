package data

import (
	"time"

	. "github.com/XuVic/tw_stock/helper"
)

func NewBzPerformance(source interface{}) []*BzPerformace {
	dataSlice := source.([]map[string]string)
	res := make([]*BzPerformace, len(dataSlice))
	for i, data := range dataSlice {
		bzData := &BzPerformace{
			FinRate: ToFloat(data["finRate"], false, false), AvgPrice: ToFloat(data["avgPrice"], false, false),
			UpdownRate: ToFloat(data["updownRate"], false, true), ROE: ToFloat(data["ROE"], false, true),
			ROA: ToFloat(data["ROA"], false, true), EPS: ToFloat(data["EPS"], false, false),
			EPSGrowth: ToFloat(data["EPSGrowth"], false, false), Turnover: ToFloat(data["turnover"], true, false),
			Netincome: ToFloat(data["netincome"], true, false), ProfitRate: ToFloat(data["profitRate"], false, true),
			Timestamp: ToTime(data["time"]),
		}
		res[i] = bzData
	}
	return res
}

type BzPerformace struct {
	FinRate, AvgPrice, UpdownRate, ROE, ROA, EPS, EPSGrowth, Turnover, Netincome, ProfitRate float64
	Timestamp                                                                                time.Time
}
