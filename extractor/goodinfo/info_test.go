package goodinfo

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/XuVic/tw_stock/extractor"
	"github.com/XuVic/tw_stock/helper"
	"github.com/stretchr/testify/assert"
)

func testData(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()

	htmlBody, _ := ioutil.ReadAll(f)
	return string(htmlBody)
}

func TestInfoTable(t *testing.T) {
	infoExtractor := NewInfoExtractorFromStr(testData("../../testdata/Info.html"))
	assert.NotNil(t, infoExtractor.Doc)

	table := infoExtractor.infoTable()
	assert.NotNil(t, table.Text())
	title := table.Find("tr:first-child").Text()
	assert.Equal(t, "公司基本資料", helper.StringTrim(title))
}

func TestBasicInfo(t *testing.T) {
	infoExtractor := NewInfoExtractorFromStr(testData("../../testdata/Info.html"))
	infoExtractor.getData()
	assert.NotNil(t, infoExtractor.Data)
	assert.NotEqual(t, 0, len(infoExtractor.Data))
}

func TestExtract(t *testing.T) {
	mockPage := helper.MockPage(testData("../../testdata/Info.html"))
	var e extractor.Extractor = NewInfoExtractor(mockPage)
	data := e.Extract()
	assert.NotNil(t, data)
}
