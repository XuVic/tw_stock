package extractor

import (
	"github.com/XuVic/tw_stock/scraper"
)


// Extractor interface specify a method set to extract data from html page. 
type Extractor interface {
	Extract(page scraper.Page) map[string]interface{}
}
