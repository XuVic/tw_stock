package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getTime(s string) time.Time {
	t, _ := time.Parse("20060102", s)
	return t
}

func TestToTime(t *testing.T) {
	testcases := []struct {
		S      string
		expect time.Time
	}{
		{"1995/4/16", getTime("19950416")},
		{"12/12", getTime("20201212")},
		{"2020/12", getTime("20201201")},
		{"2020/12/12", getTime("20201212")},
		{"2020", getTime("20200101")},
		{"19Q3", getTime("20190101")},
	}

	for _, testcase := range testcases {
		res := ToTime(testcase.S)
		assert.Equal(t, testcase.expect, res)
	}
}

func TestToFloat(t *testing.T) {
	testcases := []struct {
		S      string
		expect float64
	}{
		{"1150 億元", float64(115000000000)},
		{"10233603995股", float64(10233603995)},
		{"21.2", float64(21.2)},
		{"21.6%", float64(0.216)},
		{"1,056億", float64(105600000000)},
		{"-13.5", float64(-13.5)},
		{"+5.6", float64(5.6)},
		{"-13.5%", float64(-0.135)},
		{"-", float64(0)},
		{"", float64(0)},
	}

	for _, testcase := range testcases {
		res := ToFloat(testcase.S, false, false)
		assert.Equal(t, testcase.expect, res)
	}
}

func TestTrimString(t *testing.T) {
	testcases := []struct {
		S      string
		expect string
	}{
		{"abc (asdf)", "abc"},
		{"	2001/12/19   (成立 18年)", "2001/12/19"},
		{"	10,233,603,995 股 (含私募0股)", "10233603995股"},
		{"1150 億元", "1150億元"},
		{"21.6%", "21.6%"},
	}

	for _, testcase := range testcases {
		res := TrimString(testcase.S)
		assert.Equal(t, testcase.expect, res)
	}
}
