package permmissingelem

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SolutionTestSuite struct {
	suite.Suite
}

func (s *SolutionTestSuite) TestSolution() {
	s.Equal(4, Solution([]int{2, 3, 1, 5}))
	s.Equal(3, Solution([]int{1, 4, 7, 6, 5, 2}))
	s.Equal(2, Solution([]int{1}))
	s.Equal(1, Solution([]int{}))
}

func TestSolutionTestSuite(t *testing.T) {
	suite.Run(t, new(SolutionTestSuite))
}
