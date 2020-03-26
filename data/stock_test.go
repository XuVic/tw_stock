package data

import (
	"testing"

	"github.com/XuVic/twstock/extractor"
	"github.com/XuVic/twstock/extractor/goodinfo"
	. "github.com/XuVic/twstock/helper"
	"github.com/stretchr/testify/assert"
)

func TestNewStock(t *testing.T) {
	basicinfo := goodinfo.BasicInfo(TestData("../testdata/basicinfo.html"))
	shareholder := goodinfo.ShareHolder(TestData("../testdata/shareholder.html"))
	bzperformance := goodinfo.BzPerformance(TestData("../testdata/bzperformance.html"))
	revenues := goodinfo.Revenues(TestData("../testdata/revenues.html"))
	transactions := goodinfo.Transactions(TestData("../testdata/transactions.html"))
	dividends := goodinfo.Dividends(TestData("../testdata/dividends.html"))

	e := extractor.NewExtractor()
	e.Setup([]extractor.Criteria{basicinfo, shareholder, bzperformance, revenues, transactions, dividends})
	e.ExtractAll()

	stock := NewStock(e.Data)
	assert.NotNil(t, stock.StockID)
}
