package extractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/XuVic/tw_stock/helper"
)

func NewDataCollector() *DataCollector {
	data := make(map[string]map[string]string)
	stream := make(map[string]map[string]map[string]string)
	return &DataCollector{Data: data, Stream: stream}
}

type DataCollector struct {
	Data   map[string]map[string]string
	Stream map[string]map[string]map[string]string
}

func (d DataCollector) Collect(key string, selection *goquery.Selection, rules map[string]string) {
	dict := d.getText(selection, rules)
	d.Data[key] = dict
}

func (d DataCollector) CollectStream(key string, selection *goquery.Selection, rules map[string]string) {
	dict := make(map[string]map[string]string)

	selection.Each(func(i int, s *goquery.Selection) {
		temp := d.getText(s, rules)
		dict[temp["time"]] = temp
	})
	d.Stream[key] = dict
}

func (d DataCollector) getText(selection *goquery.Selection, rules map[string]string) map[string]string {
	res := make(map[string]string)

	for k, rule := range rules {
		res[k] = helper.StringTrim(selection.Find(rule).Text())
	}

	return res
}
