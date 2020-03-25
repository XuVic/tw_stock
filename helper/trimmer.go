package helper

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const empty = ""
const digitsRex = `[^0-9.]`

var units = map[string]int{
	"億": 100000000,
}

func StringTrim(s string) string {
	newS := strings.TrimSpace(s)
	newS = strings.ReplaceAll(newS, "\u00a0", empty)
	re := regexp.MustCompile(`(?s)\((.*)\)|%`)
	newS = re.ReplaceAllString(newS, empty)
	return newS
}

func ToInt(s string) int {
	s = StringTrim(s)
	unit := 1

	for k, v := range units {
		if strings.Contains(s, k) {
			unit = v
			break
		}
	}

	re := regexp.MustCompile(digitsRex)
	strNumber := re.ReplaceAllString(s, empty)
	number, _ := strconv.ParseFloat(strNumber, 64)

	return int(number) * unit
}

func ToFloat(s string) float64 {
	s = StringTrim(s)

	re := regexp.MustCompile(digitsRex)
	strNumber := re.ReplaceAllString(s, empty)
	number, _ := strconv.ParseFloat(strNumber, 32)
	return math.Round(number*100) / 100
}
