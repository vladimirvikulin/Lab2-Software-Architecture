package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
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
      got, err := Evaluate(tt.expr)
	  c.Assert(err, IsNil)
      c.Assert(got, Equals, tt.want)
   }
}