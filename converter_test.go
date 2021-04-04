package doku2md

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConverterTestSuite struct {
	suite.Suite
	converter *Converter
}

func TestConverterTestSuite(t *testing.T) {
	suite.Run(t, new(ConverterTestSuite))
}

func (s *ConverterTestSuite) SetupTest() {
	s.converter = &Converter{}
}

func (s *ConverterTestSuite) TestDokuToMd_Header() {
	assert.Equal(s.T(), "# Header 1", s.converter.DokuToMd("====== Header 1 ======"))
	assert.Equal(s.T(), "## Header 2", s.converter.DokuToMd("===== Header 2 ====="))
	assert.Equal(s.T(), "### Header 3", s.converter.DokuToMd("==== Header 3 ===="))
	assert.Equal(s.T(), "#### Header 4", s.converter.DokuToMd("=== Header 4 ==="))
	assert.Equal(s.T(), "##### Header 5", s.converter.DokuToMd("== Header 5 =="))
	assert.Equal(s.T(), "###### Header 6", s.converter.DokuToMd("= Header 6 ="))
}

func (s *ConverterTestSuite) TestDokuToMd_Italic() {
	assert.Equal(s.T(), "*italic*", s.converter.DokuToMd("//italic//"))
	assert.Equal(s.T(), "http://www.google.com https://www.google.com", s.converter.DokuToMd("http://www.google.com https://www.google.com"))
}
