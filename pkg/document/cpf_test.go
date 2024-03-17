package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CPFTestSuite struct {
	suite.Suite
	FormattedValidCPFNumber   string
	FormattedInvalidCPFNumber string
	PlainValidCPFNumber       string
	PlainInvalidCPFNumber     string
}

func (suite *CPFTestSuite) SetupTest() {
	suite.FormattedValidCPFNumber = "226.391.320-32"
	suite.FormattedInvalidCPFNumber = "123.456.789-00"
	suite.PlainValidCPFNumber = "22639132032"
	suite.PlainInvalidCPFNumber = "12345678900"
}

func (suite *CPFTestSuite) TestFormattedValidCPFNumber() {
	cpf := NewCPF(suite.FormattedValidCPFNumber)
	assert.Equal(suite.T(), true, cpf.IsValid())
}

func (suite *CPFTestSuite) TestFormattedInvalidCPFNumber() {
	cpf := NewCPF(suite.FormattedInvalidCPFNumber)
	assert.Equal(suite.T(), false, cpf.IsValid())
}

func (suite *CPFTestSuite) TestPlainValidCPFNumber() {
	cpf := NewCPF(suite.PlainValidCPFNumber)
	assert.Equal(suite.T(), true, cpf.IsValid())
}

func (suite *CPFTestSuite) TestPlainInvalidCPFNumber() {
	cpf := NewCPF(suite.PlainInvalidCPFNumber)
	assert.Equal(suite.T(), false, cpf.IsValid())
}

func (suite *CPFTestSuite) TestFormattedCPFNumber() {
	cpf := NewCPF(suite.PlainValidCPFNumber)
	assert.Equal(suite.T(), suite.FormattedValidCPFNumber, cpf.Formatted())
}

func (suite *CPFTestSuite) TestPlainCPFNumber() {
	cpf := NewCPF(suite.FormattedValidCPFNumber)
	assert.Equal(suite.T(), suite.PlainValidCPFNumber, cpf.String())
}

func TestCPFTestSuite(t *testing.T) {
	suite.Run(t, new(CPFTestSuite))
}
