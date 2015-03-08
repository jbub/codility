package permcheck

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SolutionTestSuite struct {
	suite.Suite
}

func (s *SolutionTestSuite) TestSolution() {
	s.Equal(1, Solution([]int{1}))
	s.Equal(1, Solution([]int{1, 3, 2}))
	s.Equal(0, Solution([]int{4, 5, 6}))
	s.Equal(0, Solution([]int{1, 3}))
	s.Equal(0, Solution([]int{2}))
}

func TestSolutionTestSuite(t *testing.T) {
	suite.Run(t, new(SolutionTestSuite))
}
