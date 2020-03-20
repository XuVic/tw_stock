package main

import (
	"fmt"
	"io/ioutil"

	"github.com/XuVic/tw_stock/scraper"
)

func initData() {
	pages := scraper.GoodinfoClient.Fetch("2881")
	count := 0
	for key, page := range pages {
		err := ioutil.WriteFile("./testdata/"+key+".html", []byte(page.Body), 0644)
		checkErr(err)
		count++
	}
	fmt.Printf("Successfully create %d test data in testdata folder.\n", count)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	initData()
}
