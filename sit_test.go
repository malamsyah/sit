package sit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AssertionTestSuite struct {
	suite.Suite
	sit *SIT
}

func (s *AssertionTestSuite) SetupSuite() {
	s.sit = NewSIT(s.Suite.T(), "http://localhost:8080")
}

func (s *AssertionTestSuite) SetupTest() {
	err := s.sit.Reset()
	assert.NoError(s.T(), err)
}

func TestAssertionTest(t *testing.T) {
	suite.Run(t, new(AssertionTestSuite))
}
