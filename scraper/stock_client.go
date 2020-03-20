package scraper

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

var GoodinfoURLs = []string{"https://goodinfo.tw/StockInfo/StockDetail.asp",
	"https://goodinfo.tw/StockInfo/StockBzPerformance.asp",
	"https://goodinfo.tw/StockInfo/ShowSaleMonChart.asp",
}

// StockClient type is used to fetch html page according to given stock_id
type StockClient struct {
	URLs []*url.URL
}

var GoodinfoCleint *StockClient = NewStockClient(GoodinfoURLs)

// NewStockClient is used to create a StockCleint instance.
func NewStockClient(strUrls []string) *StockClient {
	newUrls := make([]*url.URL, len(strUrls))

	for i, strU := range strUrls {
		u, _ := url.Parse(strU)
		newUrls[i] = u
	}
	return &StockClient{newUrls}
}

// Fetch method is used to fetch html pages.
func (c *StockClient) Fetch(stockID string) map[string]*Page {
	return c.fetchAsync(stockID)
}

func (c *StockClient) updateUrls(stockID string) {
	for _, u := range c.URLs {
		q := u.Query()
		q.Set("STOCK_ID", stockID)
		u.RawQuery = q.Encode()
	}
}

func (*StockClient) getPage(u *url.URL) *Page {
	res, _ := http.Get(u.String())
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return &Page{URL: u, Type: "html", Body: string(body)}
}

func (c *StockClient) getPageAsync(u *url.URL, ch chan *Page) {
	page := c.getPage(u)
	ch <- page
}

func (c *StockClient) fetch(stockID string) map[string]*Page {
	c.updateUrls(stockID)
	pages := make(map[string]*Page)

	for _, u := range c.URLs {
		page := c.getPage(u)
		pages[u.String()] = page
	}
	return pages
}

func (c *StockClient) fetchAsync(stockID string) map[string]*Page {
	c.updateUrls(stockID)
	chs := make([]chan *Page, len(c.URLs))
	pages := make(map[string]*Page)

	for i, u := range c.URLs {
		chs[i] = make(chan *Page)
		go c.getPageAsync(u, chs[i])
	}

	for _, ch := range chs {
		page := <-ch
		pages[page.URL.String()] = page
	}
	return pages
}
