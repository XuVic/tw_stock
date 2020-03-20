package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/tw_stock/helper"
)

type DataSelector map[string]interface{}

func (d DataSelector) SelectFrom(selection *goquery.Selection, criterions map[string]string, category string) {
	switch category {
	case "toStr":
		for k, c := range criterions {
			d[k] = helper.StringTrim(selection.Find(c).Text())
		}
	case "toInt":
		for k, c := range criterions {
			d[k] = helper.ToInt(selection.Find(c).Text())
		}
	}
}
