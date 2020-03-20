package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/XuVic/tw_stock/scraper"
)

func initData() {
	pages := scraper.GoodinfoCleint.Fetch("2881")
	count := 1
	for _, page := range pages {
		err := ioutil.WriteFile("./testdata/data"+strconv.Itoa(count)+".html", []byte(page.Body), 0644)
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
