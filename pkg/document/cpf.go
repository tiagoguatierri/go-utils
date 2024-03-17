package document

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

const (
	// CPF FormatPattern is the pattern for string replacement
	// with Regex
	CPFFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
)

// CPF type definition
type CPF string

// NewCPF is a helper function to convert and clean a new CPF
// from a string
func NewCPF(cpf string) CPF {
	return CPF(clean(cpf))
}

// Formatted returns a formatted CPF document
// 000.000.000-00
func (c *CPF) Formatted() string {
	if !c.IsValid() {
		return c.String()
	}

	expr, _ := regexp.Compile(CPFFormatPattern)
	return expr.ReplaceAllString(c.String(), "$1.$2.$3-$4")
}

// String returns plain CPF value
func (c *CPF) String() string {
	return string(*c)
}

// IsValid call private fn to validate CPF
func (c *CPF) IsValid() bool {
	return validateCPF(string(*c))
}

// You should use without punctuation
func validateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	// Check if all digits are same
	if !checkDigits(cpf) {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}
