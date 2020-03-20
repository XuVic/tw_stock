package scraper

import (
	"net/url"
)

// Page struct is used to organize html page content.
type Page struct {
	URL      *url.URL
	Body     string
	Category string
}

type Pages map[string]*Page

func (pages Pages) Get(key string) *Page {
	page, ok := pages[key]
	if ok {
		return page
	}
	return nil
}

func (pages Pages) Set(page *Page) {
	pages[page.Category] = page
}
