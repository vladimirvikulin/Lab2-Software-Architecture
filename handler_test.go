package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

func (s *EvaluateSuite) TestComputeHandler(c *C) {
	input := "2 2 +"
	output := &bytes.Buffer{}

	handler := ComputeHandler{
		Input:  strings.NewReader(input),
		Output: output,
	}

	err := handler.Compute()

	c.Assert(err, Equals, nil)
	c.Assert(output.String(), Equals, "4.0")
}
