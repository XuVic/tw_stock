package scraper

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// StockClient type is used to fetch html page according to given stock_id
type StockClient struct {
	URLs map[string]*url.URL
}

// GoodinfoClient is a default stock client which scrape html page from goodinfo website.
var GoodinfoClient *StockClient = NewStockClient(GoodinfoURLs)

// NewStockClient is used to create a StockCleint instance.
func NewStockClient(strUrls map[string]string) *StockClient {
	newUrls := make(map[string]*url.URL, len(strUrls))

	for key, strU := range strUrls {
		u, _ := url.Parse(strU)
		newUrls[key] = u
	}
	return &StockClient{newUrls}
}

// Fetch method is used to fetch html pages.
func (c *StockClient) Fetch(stockID string) Pages {
	return c.fetchAsync(stockID)
}

func (c *StockClient) updateUrls(stockID string) {
	for _, u := range c.URLs {
		q := u.Query()
		q.Set("STOCK_ID", stockID)
		u.RawQuery = q.Encode()
	}
}

func (*StockClient) getPage(u *url.URL, category string) *Page {
	res, _ := http.Get(u.String())
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return &Page{URL: u, Category: category, Body: string(body)}
}

func (c *StockClient) getPageAsync(u *url.URL, category string, ch chan *Page) {
	page := c.getPage(u, category)
	ch <- page
}

func (c *StockClient) fetch(stockID string) map[string]*Page {
	c.updateUrls(stockID)
	pages := make(map[string]*Page)

	for key, u := range c.URLs {
		page := c.getPage(u, key)
		pages[page.Category] = page
	}
	return pages
}

func (c *StockClient) fetchAsync(stockID string) Pages {
	c.updateUrls(stockID)
	chs := make([]chan *Page, len(c.URLs))
	pages := make(Pages, len(c.URLs))

	i := 0
	for key, u := range c.URLs {
		chs[i] = make(chan *Page)
		go c.getPageAsync(u, key, chs[i])
		i++
	}

	for _, ch := range chs {
		page := <-ch
		pages.Set(page)
	}
	return pages
}
