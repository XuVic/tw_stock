package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUrls(t *testing.T) {
	assert.Empty(t, GoodinfoCleint.URLs[0].RawQuery)
	GoodinfoCleint.updateUrls("2881")
	assert.NotEmpty(t, GoodinfoCleint.URLs[0].RawQuery)
	assert.Equal(t, "STOCK_ID=2881", GoodinfoCleint.URLs[0].RawQuery)
	GoodinfoCleint.updateUrls("1234")
	assert.Equal(t, "STOCK_ID=1234", GoodinfoCleint.URLs[0].RawQuery)
}

func TestGetPage(t *testing.T) {
	GoodinfoCleint.updateUrls("2881")
	url := GoodinfoCleint.URLs[0]
	page := GoodinfoCleint.getPage(url)
	assert.NotEmpty(t, page.Body)
	assert.Equal(t, url, page.URL)
}

func TestFetchAsync(t *testing.T) {
	pages := GoodinfoCleint.fetchAsync("2881")
	assert.NotEmpty(t, pages)
}

func BenchmarkFetchAsync(*testing.B) {
	GoodinfoCleint.fetchAsync("2881")
}

func BenchmarkFetch(*testing.B) {
	GoodinfoCleint.fetch("2881")
}
