package helper

import (
	"io/ioutil"
	"net/url"
	"os"

	"github.com/XuVic/twstock/scraper"
)

func MockPage(body string) *scraper.Page {
	fakeURL, _ := url.Parse("https://www.fake.com")
	return &scraper.Page{Body: body, URL: fakeURL, Category: "fake"}
}

func TestData(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()

	htmlBody, _ := ioutil.ReadAll(f)
	return string(htmlBody)
}
