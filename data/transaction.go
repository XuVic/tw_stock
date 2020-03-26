package data

import (
	"time"

	. "github.com/XuVic/tw_stock/helper"
)

func NewTransaction(source interface{}) []*Transaction {
	dataSlice := source.([]map[string]string)
	res := make([]*Transaction, len(dataSlice))
	for i, data := range dataSlice {
		tData := &Transaction{
			OpenPrice: ToFloat(data["openPrice"], false, false), ClosePrice: ToFloat(data["closePrice"], false, false),
			HighPrice: ToFloat(data["highPrice"], false, false), LowPrice: ToFloat(data["lowPrice"], false, false),
			UpdownRate: ToFloat(data["updownRate"], false, true), Timestamp: ToTime(data["time"]),
		}
		res[i] = tData
	}
	return res
}

type Transaction struct {
	OpenPrice, ClosePrice, HighPrice, LowPrice, UpdownRate float64
	Timestamp                                              time.Time
}
