package codility

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TapeEquilibriumTestSuite struct {
	suite.Suite
}

func (s *TapeEquilibriumTestSuite) TestTapeEquilibrium() {
	s.Equal(1, TapeEquilibrium([]int{3, 1, 2, 4, 3}))
	s.Equal(30, TapeEquilibrium([]int{15, -15}))
	s.Equal(3, TapeEquilibrium([]int{-10, -5, -3, -4, -5}))
	s.Equal(0, TapeEquilibrium([]int{15}))
	s.Equal(1, TapeEquilibrium([]int{15, 14}))
}

func TestTapeEquilibriumTestSuite(t *testing.T) {
	suite.Run(t, new(TapeEquilibriumTestSuite))
}
