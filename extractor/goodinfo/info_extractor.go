package goodinfo

import (
	"github.com/XuVic/tw_stock/extractor"

	"github.com/PuerkitoBio/goquery"
)

func NewInfoExtractor(source interface{}) *InfoExtractor {
	doc := createDocument(source)
	data := make(extractor.DataSelector)
	temp := make(map[string]*goquery.Selection)
	base := extractor.BaseExtractor{Data: data, Doc: doc, Temp: temp}
	return &InfoExtractor{base}
}

type InfoExtractor struct {
	extractor.BaseExtractor
}

func (e *InfoExtractor) Extract() map[string]interface{} {
	e.CheckDoc()
	e.getData()
	return e.Data
}

func (e *InfoExtractor) infoTable() *goquery.Selection {
	if _, ok := e.Temp["info_table"]; ok {
		return e.Temp["info_table"]
	}

	selection := e.Doc.Find("table.solid_1_padding_4_6_tbl>tbody")
	e.Temp["info_table"] = selection
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
