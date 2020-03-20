package extractor

// Extractor interface specify a method set to extract data from html page.
type Extractor interface {
	Extract() map[string]interface{}
}

