package goodinfo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/scraper"
)

func NewGoodInfoExtractor(category string, source interface{}) (e extractor.Extractor) {
	doc := createDocument(source)
	data := make(extractor.DataSelector)
	temp := make(map[string]*goquery.Selection)
	base := extractor.BaseExtractor{Data: data, Doc: doc, Temp: temp}
	switch category {
	case "info":
		e = &InfoExtractor{base}
	}
	return
}

func createDocument(source interface{}) (doc *goquery.Document) {
	switch source.(type) {
	case string:
		doc, _ = goquery.NewDocumentFromReader(strings.NewReader((source.(string))))
	case *scraper.Page:
		doc, _ = goquery.NewDocumentFromReader(strings.NewReader(source.(*scraper.Page).Body))
	}
	return doc
}
