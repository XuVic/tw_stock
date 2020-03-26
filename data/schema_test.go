package data

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/XuVic/twstock/extractor"
	"github.com/XuVic/twstock/extractor/goodinfo"
	. "github.com/XuVic/twstock/helper"
)

func TestNewData(t *testing.T) {
	basicinfo := goodinfo.BasicInfo(MockPage(TestData("../testdata/basicinfo.html")))
	shareholder := goodinfo.ShareHolder(MockPage(TestData("../testdata/shareholder.html")))
	bzperformance := goodinfo.BzPerformance(MockPage(TestData("../testdata/bzperformance.html")))
	revenues := goodinfo.Revenues(TestData("../testdata/revenues.html"))
	transactions := goodinfo.Transactions(TestData("../testdata/transactions.html"))
	dividends := goodinfo.Dividends(TestData("../testdata/dividends.html"))

	e := extractor.NewExtractor()
	e.Setup([]extractor.Criteria{basicinfo, shareholder, bzperformance, revenues, transactions, dividends})
	e.ExtractAll()

	basicinfoData := NewBasicInfo(e.Data[basicinfo.Name])
	assert.NotNil(t, basicinfoData.Business)

	shareholderData := NewShareHolder(e.Data[shareholder.Name])
	assert.NotNil(t, shareholderData.DirShares)

	bzData := NewBzPerformance(e.Data[bzperformance.Name])
	assert.NotNil(t, bzData[0].AvgPrice)

	revenueData := NewRevenues(e.Data[revenues.Name])
	assert.NotNil(t, revenueData[0].TurnoverMon)

	transactionData := NewTransactions(e.Data[transactions.Name])
	assert.NotNil(t, transactionData[0].OpenPrice)

	dividendData := NewDividends(e.Data[dividends.Name])
	assert.NotNil(t, dividendData[1].Stock)
	t.Log(dividendData[1])
}
