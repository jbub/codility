package codility

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type FrogJmpTestSuite struct {
	suite.Suite
}

func (s *FrogJmpTestSuite) TestFrogJmp() {
	s.Equal(0, FrogJmp(10, 10, 40))
	s.Equal(1, FrogJmp(10, 30, 20))
	s.Equal(2, FrogJmp(10, 30, 15))
	s.Equal(3, FrogJmp(10, 85, 30))
	s.Equal(1, FrogJmp(10, 11, 1))
	s.Equal(2, FrogJmp(10, 12, 1))
	s.Equal(1, FrogJmp(10, 12, 3))
}

func TestFrogJmpTestSuite(t *testing.T) {
	suite.Run(t, new(FrogJmpTestSuite))
}
