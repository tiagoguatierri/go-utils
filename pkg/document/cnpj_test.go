package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CNPJTestSuite struct {
	suite.Suite
	FormattedValidCNPJNumber   string
	FormattedInvalidCNPJNumber string
	PlainValidCNPJNumber       string
	PlainInvalidCNPJNumber     string
}

func (suite *CNPJTestSuite) SetupTest() {
	suite.FormattedValidCNPJNumber = "20.397.213/0001-06"
	suite.FormattedInvalidCNPJNumber = "12.345.678/9000-00"
	suite.PlainValidCNPJNumber = "20397213000106"
	suite.PlainInvalidCNPJNumber = "12345678900000"
}

func (suite *CNPJTestSuite) TestFormattedValidCNPJNumber() {
	cnpj := NewCNPJ(suite.FormattedValidCNPJNumber)
	assert.Equal(suite.T(), true, cnpj.IsValid())
}

func (suite *CNPJTestSuite) TestFormattedInvalidCNPJNumber() {
	cnpj := NewCNPJ(suite.FormattedInvalidCNPJNumber)
	assert.Equal(suite.T(), false, cnpj.IsValid())
}

func (suite *CNPJTestSuite) TestPlainValidCNPJNumber() {
	cnpj := NewCNPJ(suite.PlainValidCNPJNumber)
	assert.Equal(suite.T(), true, cnpj.IsValid())
}

func (suite *CNPJTestSuite) TestPlainInvalidCNPJNumber() {
	cnpj := NewCNPJ(suite.PlainInvalidCNPJNumber)
	assert.Equal(suite.T(), false, cnpj.IsValid())
}

func (suite *CNPJTestSuite) TestFormattedCNPJNumber() {
	cnpj := NewCNPJ(suite.PlainValidCNPJNumber)
	assert.Equal(suite.T(), suite.FormattedValidCNPJNumber, cnpj.Formatted())
}

func (suite *CNPJTestSuite) TestPlainCNPJNumber() {
	cnpj := NewCNPJ(suite.FormattedValidCNPJNumber)
	assert.Equal(suite.T(), suite.PlainValidCNPJNumber, cnpj.String())
}

func (suite *CNPJTestSuite) TestIncompleteCNPJNumber() {
	number := "123123123"
	cnpj := NewCNPJ(number)
	assert.Equal(suite.T(), false, cnpj.IsValid())
	assert.Equal(suite.T(), number, cnpj.Formatted())
}

func (suite *CNPJTestSuite) TestAllEqualCNPJNumber() {
	number := "99.999.999/9999-99"
	cnpj := NewCNPJ(number)
	assert.Equal(suite.T(), false, cnpj.IsValid())
	assert.Equal(suite.T(), "99999999999999", cnpj.Formatted())
}

func TestCNPJTestSuite(t *testing.T) {
	suite.Run(t, new(CNPJTestSuite))
}
