package helper

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const empty = ""
const digitsRex = `[^0-9.]`
const redundancy = `\(.+\)|\s|,`
const date = `\d{1,4}\/\d{1,2}\/?\d{1,2}`
const iso8601 = "20060102"
const billionUnit = 100000000

func ToTime(s string) time.Time {
	days := strings.Split(s, "/")
	if len(days) == 1 {
		days = takeYear(s)
	}
	if len(days) == 2 {
		days = append(days, "01")
		if len(days[0]) == 2 {
			m := days[0]
			d := days[1]
			days[0] = strconv.Itoa(time.Now().Year())
			days[1] = m
			days[2] = d
		}
	}

	for i, day := range days {
		if len(day) == 1 {
			days[i] = "0" + day
		}
	}

	date := strings.Join(days, empty)
	res, _ := time.Parse(iso8601, date)
	return res
}

func ToFloat(s string, billion bool, percent bool) float64 {
	if strings.Contains(s, "%") {
		percent = true
	}
	if strings.Contains(s, "億") {
		billion = true
	}

	re := regexp.MustCompile(`[^\d^.^-]`)
	s = re.ReplaceAllString(s, empty)
	res, _ := strconv.ParseFloat(s, 64)
	if percent {
		res = res / 100
	}
	if billion {
		res = res * billionUnit
	}

	return math.Round(res*10000) / 10000
}

func TrimString(s string) string {
	newS := strings.TrimSpace(s)
	newS = strings.ReplaceAll(newS, "\u00a0", empty)
	re := regexp.MustCompile(redundancy)
	newS = re.ReplaceAllString(newS, empty)
	return newS
}

func takeYear(s string) []string {
	y := s
	if strings.Contains(s, "Q") {
		year := strconv.Itoa(time.Now().Year())
		y = fmt.Sprintf("%s%s", year[:2], s[:2])
	}
	return []string{y, "01", "01"}
}
