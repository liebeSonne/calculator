package calc

import (
	"calculator/pkg/calc/command"
)

type Calculator interface {
	RunCommand(str string)
}

func NewCalculator(
	parser command.Parser,
) Calculator {
	return &calculator{
		parser: parser,
	}
}

type calculator struct {
	parser command.Parser
}

func (c *calculator) RunCommand(str string) {
	cmd := c.parser.CreateCommand(str)
	if cmd != nil {
		(*cmd).Execute()
	}
}
