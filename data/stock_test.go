package data

import (
	"testing"

	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/extractor/goodinfo"
	. "github.com/XuVic/tw_stock/helper"
)

func TestNewStock(t *testing.T) {
	basicinfo := goodinfo.BasicInfo(TestData("../testdata/Info.html"))
	shareholder := goodinfo.ShareHolderInfo(TestData("../testdata/Overview.html"))
	bzperformance := goodinfo.BzPerformance(TestData("../testdata/Performance.html"))
	revenue := goodinfo.Revenue(TestData("../testdata/Revenue.html"))
	transaction := goodinfo.Transaction(TestData("../testdata/Transaction.html"))
	dividend := goodinfo.DividendPolicy(TestData("../testdata/Dividend.html"))

	e := extractor.NewExtractor()
	e.ExtractorAll([]extractor.Criteria{basicinfo, shareholder, bzperformance, revenue, transaction, dividend})

	stock := NewStock(e.Data)
	t.Log(stock.StockID)
}
