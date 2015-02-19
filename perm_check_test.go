package codility

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PermCheckTestSuite struct {
	suite.Suite
}

func (s *PermCheckTestSuite) TestPermCheck() {
	s.Equal(1, PermCheck([]int{1}))
	s.Equal(1, PermCheck([]int{1, 3, 2}))
	s.Equal(0, PermCheck([]int{4, 5, 6}))
	s.Equal(0, PermCheck([]int{1, 3}))
	s.Equal(0, PermCheck([]int{2}))
}

func TestPermCheckTestSuite(t *testing.T) {
	suite.Run(t, new(PermCheckTestSuite))
}
