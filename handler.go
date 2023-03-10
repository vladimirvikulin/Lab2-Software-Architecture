package lab2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	inputBytes, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	input := string(inputBytes)

	result, err := Evaluate(input)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(ch.Output, "%.1f", result)
	if err != nil {
		return err
	}

	return nil
}
