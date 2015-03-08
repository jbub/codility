package frogjmp

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SolutionTestSuite struct {
	suite.Suite
}

func (s *SolutionTestSuite) TestSolution() {
	s.Equal(0, Solution(10, 10, 40))
	s.Equal(1, Solution(10, 30, 20))
	s.Equal(2, Solution(10, 30, 15))
	s.Equal(3, Solution(10, 85, 30))
	s.Equal(1, Solution(10, 11, 1))
	s.Equal(2, Solution(10, 12, 1))
	s.Equal(1, Solution(10, 12, 3))
}

func TestSolutionTestSuite(t *testing.T) {
	suite.Run(t, new(SolutionTestSuite))
}
