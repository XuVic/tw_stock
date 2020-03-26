package data

import (
	"time"

	. "github.com/XuVic/tw_stock/helper"
)

type Dividend struct {
	Cash, Stock, Yield, EarningShare float64
	Timestamp                        time.Time
}

type Revenue struct {
	TurnoverMon, GrowthRateY, GrowthRateM, AccuTurnover float64
	Timestamp                                           time.Time
}

type BzPerformance struct {
	FinRate, AvgPrice, UpdownRate, ROE, ROA, EPS, EPSGrowth, Turnover, Netincome, ProfitRate float64
	Timestamp                                                                                time.Time
}

type BasicInfo struct {
	StockID, StockName, Industry, ComName, Listed, Chairman, Business string
	Cap, MarketCap, Shares, PreShares                                 float64
}

type ShareHolder struct {
	DirShares, ForeShares, GovShares, NatShares, NatJShares, NatFShares float64
	BoardMeeting, ShareMeeting, DividendDate, DividTraDate              time.Time
}

type Transaction struct {
	OpenPrice, ClosePrice, HighPrice, LowPrice, UpdownRate float64
	Timestamp                                              time.Time
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

func NewRevenues(source interface{}) []*Revenue {
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

func NewDividends(source interface{}) []*Dividend {
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

func NewBzPerformance(source interface{}) []*BzPerformance {
	dataSlice := source.([]map[string]string)
	res := make([]*BzPerformance, len(dataSlice))
	for i, data := range dataSlice {
		bzData := &BzPerformance{
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

func NewTransactions(source interface{}) []*Transaction {
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
