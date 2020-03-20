package helper

import (
	"net/url"

	"github.com/XuVic/tw_stock/scraper"
)

func MockPage(body string) *scraper.Page {
	fakeURL, _ := url.Parse("https://www.fake.com")
	return &scraper.Page{Body: body, URL: fakeURL, Category: "fake"}
}
