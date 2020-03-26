# TWSTOCK

TWSTOCK provide a simple API to allow consumer to get rich information of Taiwan stock. The work behind this package is to scrape the html data from [goodinfo](https://goodinfo.tw/StockInfo/index.asp) website.

## Table of Contents

* [Installation](#installation)
* [API](#api)
* [Stock Schema](#schema)

### Installation

Please note that TWSTOCK requires Go1.14+.

    $ go get github.com/XuVic/tw_stock

(optional) To run unit tests:
Before running unit tests, you have to change work directory to the directory where TWSTOCK was installed to.

    $ go test -v

(optional) To run benchmarks:
TWSTOCK only implement benchmark test on web scraping function.

    $ go test -bench=. -run=None ./scraper/

### API

To get stock information, consumer can just pass stockID(e.g 2881) to `Get` function.

Example:

    stock := twstock.Get("2330")

Function Signature:

    Get(stockID string) *twstock.Data.Stock

### Schema

Stock struct is composed of others six structs, each structs combine different type of information about stock.

    type Stock struct {
	    *BasicInfo
	    *ShareHolder
	    BzPerformance []*BzPerformance
	    Revenues      []*Revenue
	    Transactions  []*Transaction
	    Dividends     []*Dividend
	    Timestamp     time.Time
    }

[To see more detailed fields of each struct.](https://docs.google.com/spreadsheets/d/1tSOa7QKc_RB97iKGHBZmpqQGbr2vfDoW5YQCN_BBuYE/edit?usp=sharing)