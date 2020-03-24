package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/tw_stock/helper"
	"github.com/XuVic/tw_stock/scraper"
)

// Extractor interface specify a method set to extract data from html page.
type Extractor interface {
	Extract() map[string]interface{}
}

type BaseExtractor struct {
	Data DataSelector
	Doc  *goquery.Document
	Page *scraper.Page
	Temp map[string]*goquery.Selection
}

func (e *BaseExtractor) CheckDoc() {
	if e.Doc == nil {
		panic(helper.NotNull("Doc attribute"))
	}
}
