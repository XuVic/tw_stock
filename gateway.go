package twstock

import (
	"github.com/XuVic/twstock/data"
	"github.com/XuVic/twstock/extractor"
	"github.com/XuVic/twstock/scraper"
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
