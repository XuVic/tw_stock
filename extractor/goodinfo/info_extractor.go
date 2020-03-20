package goodinfo

import (
	"strings"

	"github.com/XuVic/tw_stock/extractor"

	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/tw_stock/helper"
	"github.com/XuVic/tw_stock/scraper"
)

func NewInfoExtractor(page *scraper.Page) *InfoExtractor {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(page.Body))
	temp := make(map[string]*goquery.Selection)
	data := make(extractor.DataSelector)
	return &InfoExtractor{Data: data, Page: page, Doc: doc, temp: temp}
}

func NewInfoExtractorFromStr(str string) *InfoExtractor {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(str))
	temp := make(map[string]*goquery.Selection)
	data := make(extractor.DataSelector)
	return &InfoExtractor{Data: data, Doc: doc, temp: temp}
}

type InfoExtractor struct {
	Data extractor.DataSelector
	Doc  *goquery.Document
	Page *scraper.Page
	temp map[string]*goquery.Selection
}

func (e *InfoExtractor) Extract() map[string]interface{} {
	e.checkDoc()
	e.getData()
	return e.Data
}

func (e *InfoExtractor) infoTable() *goquery.Selection {
	if e.temp["info_table"] != nil {
		return e.temp["info_table"]
	}

	selection := e.Doc.Find("table.solid_1_padding_4_6_tbl>tbody")
	e.temp["info_table"] = selection
	return selection
}

func (e *InfoExtractor) getData() {
	table := e.infoTable()
	var strCriteria = map[string]string{
		"stockID": "tr:nth-child(2)>td:nth-child(2)", "stockName": "tr:nth-child(2)>td:nth-child(4)",
		"industry": "tr:nth-child(3)>td:nth-child(2)", "listed": "tr:nth-child(3)>td:nth-child(4)",
		"comName": "tr:nth-child(4)>td:nth-child(2)", "chairman": "tr:nth-child(12)>td:nth-child(2)",
		"business": "tr:nth-child(23)>td:nth-child(2)",
	}
	var intCriteria = map[string]string{
		"cap": "tr:nth-child(8)>td:nth-child(2)", "marketCap": "tr:nth-child(9)>td:nth-child(2)",
		"shares": "tr:nth-child(10)>td:nth-child(2)", "preShares": "tr:nth-child(11)>td:nth-child(2)",
	}
	e.Data.SelectFrom(table, strCriteria, "toStr")
	e.Data.SelectFrom(table, intCriteria, "toInt")
}

func (e *InfoExtractor) checkDoc() {
	if e.Doc == nil {
		panic(helper.NotNull("Doc attribute"))
	}
}
