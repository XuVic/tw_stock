package data

import (
	"time"

	. "github.com/XuVic/tw_stock/helper"
)

func NewRevenue(source interface{}) []*Revenue {
	dataSlice := source.([]map[string]string)
	res := make([]*Revenue, len(dataSlice))
	for i, data := range dataSlice {
		reveData := &Revenue{
			TurnoverMon: ToFloat(data["turnoverMon"], true, false), GrowthRateM: ToFloat(data["growthRateM"], false, true),
			GrowthRateY: ToFloat(data["growthRateY"], false, true), AccuTurnover: ToFloat(data["accuTurnover"], true, false),
			Timestamp: ToTime(data["time"]),
		}
		res[i] = reveData
	}
	return res
}

type Revenue struct {
	TurnoverMon, GrowthRateY, GrowthRateM, AccuTurnover float64
	Timestamp                                           time.Time
}
