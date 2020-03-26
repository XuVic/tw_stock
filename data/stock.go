package data

import "time"

type Stock struct {
	*BasicInfo
	*ShareHolder
	BzPerformance []*BzPerformance
	Revenues      []*Revenue
	Transactions  []*Transaction
	Dividends     []*Dividend
	Timestamp     time.Time
}

func NewStock(data map[string]interface{}) *Stock {
	var basicinfo *BasicInfo
	var shareholder *ShareHolder
	var bzperformance []*BzPerformance
	var revenues []*Revenue
	var transactions []*Transaction
	var dividends []*Dividend

	for k, source := range data {
		switch k {
		case "basicinfo":
			basicinfo = NewBasicInfo(source)
		case "shareholder":
			shareholder = NewShareHolder(source)
		case "bzperformance":
			bzperformance = NewBzPerformance(source)
		case "revenues":
			revenues = NewRevenues(source)
		case "transactions":
			transactions = NewTransactions(source)
		case "dividends":
			dividends = NewDividends(source)
		}
	}
	return &Stock{basicinfo, shareholder, bzperformance, revenues, transactions, dividends, time.Now()}
}
