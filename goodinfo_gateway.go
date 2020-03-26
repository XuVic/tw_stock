package tw_stock

import (
	"github.com/XuVic/tw_stock/data"
	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/extractor/goodinfo"
	"github.com/XuVic/tw_stock/scraper"
)

type GoodinfoGateway struct {
	*extractor.Extractor
	*scraper.StockClient
}

func (g *GoodinfoGateway) Get(stock string) *data.Stock {
	g.setup(stock)
	g.ExtractAll()
	return data.NewStock(g.Data)
}

func (g *GoodinfoGateway) setup(stock string) {
	pages := g.Fetch(stock)
	criterias := make([]extractor.Criteria, len(pages))
	i := 0
	for k, page := range pages {
		switch k {
		case "basicinfo":
			criterias[i] = goodinfo.BasicInfo(page)
		case "shareholder":
			criterias[i] = goodinfo.ShareHolder(page)
		case "bzperformance":
			criterias[i] = goodinfo.BzPerformance(page)
		case "revenues":
			criterias[i] = goodinfo.Revenues(page)
		case "transactions":
			criterias[i] = goodinfo.Transactions(page)
		case "dividends":
			criterias[i] = goodinfo.Dividends(page)
		}
		i++
	}
	g.Setup(criterias)
}
