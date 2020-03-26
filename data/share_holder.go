package data

import (
	"time"

	. "github.com/XuVic/tw_stock/helper"
)

type ShareHolder struct {
	DirShares, ForeShares, GovShares, NatShares, NatJShares, NatFShares float64
	BoardMeeting, ShareMeeting, DividendDate, DividTraDate              time.Time
}

func NewShareHolder(source interface{}) *ShareHolder {
	data := source.(map[string]string)
	return &ShareHolder{
		DirShares: ToFloat(data["dirShares"], false, true), ForeShares: ToFloat(data["foreShare"], false, false),
		GovShares: ToFloat(data["govShares"], false, true), NatShares: ToFloat(data["natShares"], false, false),
		NatJShares: ToFloat(data["natJShares"], false, true), NatFShares: ToFloat(data["natFShares"], false, true),
		BoardMeeting: ToTime(data["boardMeeting"]), ShareMeeting: ToTime(data["shareMeeting"]),
		DividendDate: ToTime(data["dividendDate"]), DividTraDate: ToTime(data["dividTraDate"]),
	}
}
