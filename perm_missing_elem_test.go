package codility

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PermMissingElemTestSuite struct {
	suite.Suite
}

func (s *PermMissingElemTestSuite) TestFrogJmp() {
	s.Equal(4, PermMissingElem([]int{2, 3, 1, 5}))
	s.Equal(3, PermMissingElem([]int{1, 4, 7, 6, 5, 2}))
	s.Equal(2, PermMissingElem([]int{1}))
	s.Equal(1, PermMissingElem([]int{}))
}

func TestPermMissingElemTestSuite(t *testing.T) {
	suite.Run(t, new(PermMissingElemTestSuite))
}
