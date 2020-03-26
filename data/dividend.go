package data

import (
	"time"

	. "github.com/XuVic/tw_stock/helper"
)

func NewDividend(source interface{}) []*Dividend {
	dataSlice := source.([]map[string]string)
	res := make([]*Dividend, len(dataSlice))
	for i, data := range dataSlice {
		dData := &Dividend{
			Cash: ToFloat(data["cash"], false, false), Stock: ToFloat(data["stock"], false, false),
			Yield: ToFloat(data["yield"], false, true), EarningShare: ToFloat(data["earningShare"], false, true),
			Timestamp: ToTime(data["time"]),
		}
		res[i] = dData
	}
	return res
}

type Dividend struct {
	Cash, Stock, Yield, EarningShare float64
	Timestamp                        time.Time
}
