package document

import (
	"strconv"
	"strings"
)

type Validator interface {
	Formatted() string
	IsValid() bool
	String() string
}

// Clean can be used for cleaning formatted documents
func clean(s string) string {
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	return s
}

func checkDigits(s string) bool {
	d := s[0]
	r := true
	for i := 1; i < len(s); i++ {
		if s[i] != d {
			r = false
			break
		}
	}
	return !r
}

func sumDigit(s string, table []int) int {

	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}
