package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

type EvaluateSuite struct{}

var _ = Suite(&EvaluateSuite{})

func TestEvaluate(t *testing.T) {
	TestingT(t)
}

func (s *EvaluateSuite) TestSimpleExpressions(c *C) {
	tests := []struct {
		expr string
		want float64
	}{
		{"2 3 +", 5.0},
		{"5 2 -", 3.0},
		{"4 3 *", 12.0},
		{"6 2 /", 3.0},
	}
	for _, tt := range tests {
		got, err := EvaluatePostifx(tt.expr)
		c.Assert(err, IsNil)
		c.Assert(got, Equals, tt.want)
	}
}

func (s *EvaluateSuite) TestComplexExpressions(c *C) {
	tests := []struct {
		expr string
		want float64
	}{
		{"2 3 4 * +", 14.0},
		{"5 1 2 + 4 * + 3 -", 14.0},
		{"3 4 2 * 1 5 - 2 3 ^ ^ / +", 3.0001220703125},
	}
	for _, tt := range tests {
		got, err := EvaluatePostifx(tt.expr)
		c.Assert(err, IsNil)
		c.Assert(got, Equals, tt.want)
	}
}

func (s *EvaluateSuite) TestInvalidExpression(c *C) {
	tests := []struct {
		expr string
	}{
		{""},
		{"2 + 3"},
		{"2 x 3 +"},
	}
	for _, tt := range tests {
		_, err := EvaluatePostifx(tt.expr)
		c.Assert(err, NotNil)
		c.Assert(err, FitsTypeOf, InvalidExpressionError{})
	}
}
