package scraper

import (
	"net/url"
)

// Page struct is used to organize html page content.
type Page struct {
	URL  *url.URL
	Body string
	Type string
}

// type Pages map[string]*Page
