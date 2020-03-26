package goodinfo

import (
	"testing"

	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/helper"
	"github.com/stretchr/testify/assert"
)

func TestCriteria(t *testing.T) {
	testcases := []struct {
		Name     string
		Criteria extractor.Criteria
	}{{
		"stockinfo",
		BasicInfo(helper.MockPage(helper.TestData("../../testdata/basicinfo.html"))),
	}, {
		"shareholderinfo",
		ShareHolder(helper.TestData("../../testdata/shareholder.html")),
	}, {
		"bzperformance",
		BzPerformance(helper.TestData("../../testdata/bzperformance.html")),
	}, {
		"revenue",
		Revenues(helper.TestData("../../testdata/revenues.html")),
	}, {
		"transaction",
		Transactions(helper.TestData("../../testdata/transactions.html")),
	}, {
		"dividendpolicy",
		Dividends(helper.TestData("../../testdata/dividends.html")),
	}}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			assert.NotNil(t, testcase.Criteria.Rules)
			e := extractor.NewExtractor()
			assert.Empty(t, e.Data)
			e.ExtractFrom(testcase.Criteria)
			assert.NotEmpty(t, e.Data)
		})
	}
}
