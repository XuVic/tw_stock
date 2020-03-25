package extractor

import "github.com/PuerkitoBio/goquery"

type Criteria struct {
	Name      string
	Selection *goquery.Selection
	Rules     map[string]string
	Stream    bool
}
