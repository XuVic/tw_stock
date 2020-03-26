package goodinfo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/twstock/extractor"
	"github.com/XuVic/twstock/scraper"
)

func createDocument(source interface{}) (doc *goquery.Document) {
	switch source.(type) {
	case string:
		doc, _ = goquery.NewDocumentFromReader(strings.NewReader((source.(string))))
	case *scraper.Page:
		doc, _ = goquery.NewDocumentFromReader(strings.NewReader(source.(*scraper.Page).Body))
	}
	return doc
}

func BasicInfo(source interface{}) extractor.Criteria {
	selection := createDocument(source).Find("table.solid_1_padding_4_6_tbl>tbody")
	rules := map[string]string{
		"stockID": "tr:nth-child(2)>td:nth-child(2)", "stockName": "tr:nth-child(2)>td:nth-child(4)",
		"industry": "tr:nth-child(3)>td:nth-child(2)", "listed": "tr:nth-child(3)>td:nth-child(4)",
		"comName": "tr:nth-child(4)>td:nth-child(2)", "chairman": "tr:nth-child(12)>td:nth-child(2)",
		"business": "tr:nth-child(23)>td:nth-child(2)", "cap": "tr:nth-child(8)>td:nth-child(2)",
		"marketCap": "tr:nth-child(9)>td:nth-child(2)", "shares": "tr:nth-child(10)>td:nth-child(2)",
		"preShares": "tr:nth-child(11)>td:nth-child(2)",
	}
	return extractor.Criteria{Name: "basicinfo", Selection: selection, Rules: rules}
}

func ShareHolder(source interface{}) extractor.Criteria {
	selection := createDocument(source).Selection
	rules := map[string]string{
		"dirShares":    "table.solid_1_padding_4_0_tbl:contains(全體董監) tr:nth-child(4) td:nth-child(3)",
		"foreShares":   "table.solid_1_padding_4_6_tbl:contains(持有股數排名) tr:nth-child(2) td:nth-child(5)",
		"govShares":    "table.solid_1_padding_4_6_tbl:contains(持有股數排名) tr:nth-child(5) td:nth-child(5)",
		"natShares":    "table.solid_1_padding_4_6_tbl:contains(持有股數排名) tr:nth-child(4) td:nth-child(5)",
		"natJShares":   "table.solid_1_padding_4_6_tbl:contains(持有股數排名) tr:nth-child(3) td:nth-child(5)",
		"natFShares":   "table.solid_1_padding_4_6_tbl:contains(持有股數排名) tr:nth-child(6) td:nth-child(5)",
		"boardMeeting": "table.solid_1_padding_4_4_tbl:contains(重 要 行 事 曆) tr:nth-child(3) td:nth-child(2)",
		"shareMeeting": "table.solid_1_padding_4_4_tbl:contains(重 要 行 事 曆) tr:nth-child(4) td:nth-child(2)",
		"dividendDate": "table.solid_1_padding_4_4_tbl:contains(重 要 行 事 曆) tr:nth-child(10) td:nth-child(2)",
		"dividTraDate": "table.solid_1_padding_4_4_tbl:contains(重 要 行 事 曆) tr:nth-child(8) td:nth-child(2)",
	}
	return extractor.Criteria{Name: "shareholder", Selection: selection, Rules: rules}
}

func BzPerformance(source interface{}) extractor.Criteria {
	selection := createDocument(source).Find("div#txtFinDetailData div#divFinDetail>table>tbody tr")
	rules := map[string]string{
		"time": "td:nth-child(1)", "finRate": "td:nth-child(3)", "avgPrice": "td:nth-child(5)",
		"updownRate": "td:nth-child(7)", "ROE": "td:nth-child(17)", "ROA": "td:nth-child(18)",
		"EPS": "td:nth-child(19)", "EPSGrowth": "td:nth-child(20)", "turnover": "td:nth-child(8)",
		"netincome": "td:nth-child(12)", "profitRate": "td:nth-child(16)",
	}
	return extractor.Criteria{Name: "bzperformance", Selection: selection, Rules: rules, Stream: true}
}

func Revenues(source interface{}) extractor.Criteria {
	selection := createDocument(source).Find("div#divDetail>table tbody tr")
	rules := map[string]string{
		"time": "td:nth-child(1)", "turnoverMon": "td:nth-child(13)", "growthRateY": "td:nth-child(15)",
		"growthRateM": "td:nth-child(14)", "accuTurnover": "td:nth-child(16)",
	}
	return extractor.Criteria{Name: "revenues", Selection: selection, Rules: rules, Stream: true}
}

func Transactions(source interface{}) extractor.Criteria {
	selection := createDocument(source).Find("div#divK_ChartDetail div#divPriceDetail>table tbody tr")
	rules := map[string]string{
		"time": "td:nth-child(1)", "openPrice": "td:nth-child(2)", "closePrice": "td:nth-child(4)",
		"highPrice": "td:nth-child(3)", "lowPrice": "td:nth-child(4)", "updownRate": "td:nth-child(7)",
		"corBuySell": "td:nth-child(16)",
	}
	return extractor.Criteria{Name: "transactions", Selection: selection, Rules: rules, Stream: true}
}

func Dividends(source interface{}) extractor.Criteria {
	selection := createDocument(source).Find("div#divDetail>table tbody tr").Slice(0, 10)
	rules := map[string]string{
		"time": "td:nth-child(1)", "cash": "td:nth-child(4)", "stock": "td:nth-child(7)",
		"yield": "td:nth-child(19)", "earningShare": "td:nth-child(24)",
	}
	return extractor.Criteria{Name: "dividends", Selection: selection, Rules: rules, Stream: true}
}
