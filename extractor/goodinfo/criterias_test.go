package goodinfo

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/XuVic/tw_stock/extractor"
	"github.com/stretchr/testify/assert"
)

func testData(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()

	htmlBody, _ := ioutil.ReadAll(f)
	return string(htmlBody)
}

// func TestSelectionFind(t *testing.T) {
// 	// selection := Revenue(testData("../../testdata/Revenue.html")).Selection
// 	// // t.Log(selection.Text())
// 	// selection.Each(func(i int, s *goquery.Selection) {
// 	// 	t.Log(s.Find("td:nth-child(1)").Text())
// 	// })
// }

func TestCriteria(t *testing.T) {
	testcases := []struct {
		Name     string
		Criteria extractor.Criteria
		Stream   bool
	}{{
		"stockinfo",
		InfoCriteria(testData("../../testdata/Info.html")),
		false,
	}, {
		"shareholderinfo",
		ShareHolderInfo(testData("../../testdata/Overview.html")),
		false,
	}, {
		"bzperformance",
		BzPerformance(testData("../../testdata/Performance.html")),
		true,
	}, {
		"revenue",
		Revenue(testData("../../testdata/Revenue.html")),
		true,
	}}
	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			assert.NotNil(t, testcase.Criteria.Rules)
			e := extractor.NewExtractor()
			if testcase.Stream {
				assert.Empty(t, e.Stream)
				e.Extract(testcase.Criteria)
				assert.NotEmpty(t, e.Stream)
				t.Log(e.Stream)
			} else {
				assert.Empty(t, e.Data)
				e.Extract(testcase.Criteria)
				assert.NotEmpty(t, e.Data)
			}
		})
	}
}
