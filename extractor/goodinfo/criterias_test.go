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
		BasicInfo(helper.MockPage(helper.TestData("../../testdata/Info.html"))),
	}, {
		"shareholderinfo",
		ShareHolderInfo(helper.TestData("../../testdata/Overview.html")),
	}, {
		"bzperformance",
		BzPerformance(helper.TestData("../../testdata/Performance.html")),
	}, {
		"revenue",
		Revenue(helper.TestData("../../testdata/Revenue.html")),
	}, {
		"transaction",
		Transaction(helper.TestData("../../testdata/Transaction.html")),
	}, {
		"dividendpolicy",
		DividendPolicy(helper.TestData("../../testdata/Dividend.html")),
	}}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			assert.NotNil(t, testcase.Criteria.Rules)
			e := extractor.NewExtractor()
			assert.Empty(t, e.Data)
			e.Extract(testcase.Criteria)
			assert.NotEmpty(t, e.Data)
		})
	}
}
