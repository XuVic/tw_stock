package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/tw_stock/helper"
)

type DataCollector map[string]interface{}

func (d DataCollector) Collect(key string, selection *goquery.Selection, rules map[string]string) {
	dict := d.getText(selection, rules)
	d[key] = dict
}

func (d DataCollector) CollectStream(key string, selection *goquery.Selection, rules map[string]string) {
	var slice []map[string]string

	selection.Each(func(i int, s *goquery.Selection) {
		temp := d.getText(s, rules)
		slice = append(slice, temp)
	})
	d[key] = slice
}

func (d DataCollector) getText(selection *goquery.Selection, rules map[string]string) map[string]string {
	res := make(map[string]string)

	for k, rule := range rules {
		res[k] = helper.TrimString(selection.Find(rule).Text())
	}

	return res
}
