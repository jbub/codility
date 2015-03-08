package tapeequilibrium

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SolutionTestSuite struct {
	suite.Suite
}

func (s *SolutionTestSuite) TestSolution() {
	s.Equal(1, Solution([]int{3, 1, 2, 4, 3}))
	s.Equal(30, Solution([]int{15, -15}))
	s.Equal(3, Solution([]int{-10, -5, -3, -4, -5}))
	s.Equal(0, Solution([]int{15}))
	s.Equal(1, Solution([]int{15, 14}))
}

func TestSolutionTestSuite(t *testing.T) {
	suite.Run(t, new(SolutionTestSuite))
}
