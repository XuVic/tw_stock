package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/extractor/goodinfo"
	. "github.com/XuVic/tw_stock/helper"
)

func TestNewData(t *testing.T) {
	basicinfo := goodinfo.BasicInfo(MockPage(TestData("../testdata/Info.html")))
	shareholder := goodinfo.ShareHolderInfo(MockPage(TestData("../testdata/Overview.html")))
	bzperformance := goodinfo.BzPerformance(MockPage(TestData("../testdata/Performance.html")))
	revenue := goodinfo.Revenue(TestData("../testdata/Revenue.html"))
	transaction := goodinfo.Transaction(TestData("../testdata/Transaction.html"))
	dividend := goodinfo.DividendPolicy(TestData("../testdata/Dividend.html"))

	e := extractor.NewExtractor()
	e.ExtractorAll([]extractor.Criteria{basicinfo, shareholder, bzperformance, revenue, transaction, dividend})

	basicinfoData := NewBasicInfo(e.Data[basicinfo.Name])
	assert.NotNil(t, basicinfoData.Business)

	shareholderData := NewShareHolder(e.Data[shareholder.Name])
	assert.NotNil(t, shareholderData.DirShares)

	bzData := NewBzPerformance(e.Data[bzperformance.Name])
	assert.NotNil(t, bzData[0].AvgPrice)

	revenueData := NewRevenue(e.Data[revenue.Name])
	assert.NotNil(t, revenueData[0].TurnoverMon)

	transactionData := NewTransaction(e.Data[transaction.Name])
	assert.NotNil(t, transactionData[0].OpenPrice)

	dividendData := NewDividend(e.Data[dividend.Name])
	assert.NotNil(t, dividendData[1].Stock)
	t.Log(dividendData[1])
}
