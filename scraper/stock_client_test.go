package scraper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUrls(t *testing.T) {
	assert.Empty(t, GoodinfoClient.URLs["Overview"].RawQuery)
	GoodinfoClient.updateUrls("2881")
	assert.NotEmpty(t, GoodinfoClient.URLs["Overview"].RawQuery)
	assert.Equal(t, "STOCK_ID=2881", GoodinfoClient.URLs["Overview"].RawQuery)
	GoodinfoClient.updateUrls("1234")
	assert.Equal(t, "STOCK_ID=1234", GoodinfoClient.URLs["Overview"].RawQuery)
}

func TestGetPage(t *testing.T) {
	GoodinfoClient.updateUrls("2881")
	url := GoodinfoClient.URLs["Overview"]
	page := GoodinfoClient.getPage(url, "Overview")
	assert.NotEmpty(t, page.Body)
	assert.Equal(t, url, page.URL)
}

func TestFetchAsync(t *testing.T) {
	pages := GoodinfoClient.fetchAsync("2881")
	assert.NotEmpty(t, pages)
	assert.NotEmpty(t, pages.Get("Overview"))
}

func BenchmarkFetchAsync(*testing.B) {
	GoodinfoClient.fetchAsync("2881")
}

func BenchmarkFetch(*testing.B) {
	GoodinfoClient.fetch("2881")
}
