package tw_stock

import (
	"github.com/XuVic/tw_stock/data"
	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/scraper"
)

type Gateway interface {
	Get(string) *data.Stock
}

func gatewayFactory(category string) Gateway {
	switch category {
	case "goodinfo":
		e := extractor.NewExtractor()
		return &GoodinfoGateway{e, scraper.GoodinfoClient}
	default:
		e := extractor.NewExtractor()
		return &GoodinfoGateway{e, scraper.GoodinfoClient}
	}
}

func Query(stockID string, gateway Gateway) *data.Stock {
	return gateway.Get(stockID)
}

func Get(stockID string) *data.Stock {
	return Query(stockID, gatewayFactory("goodinfo"))
}
