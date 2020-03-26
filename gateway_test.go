package twstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	stock := Get("2330")
	assert.NotNil(t, stock.BasicInfo)
	assert.NotNil(t, stock.ShareMeeting)
	assert.NotNil(t, stock.BzPerformance)
	assert.NotNil(t, stock.Transactions)
	assert.NotNil(t, stock.Revenues)
	assert.NotNil(t, stock.Dividends)
}
