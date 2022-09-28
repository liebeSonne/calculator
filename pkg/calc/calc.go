package calc

import (
	"calculator/pkg/calc/command"
	"calculator/pkg/calc/expression"
)

type Calculator interface {
	RunCommand(str string)
}

func NewCalculator(
	storage expression.Storage,
	parser command.Parser,
) Calculator {
	return &calculator{
		storage: storage,
		parser:  parser,
	}
}

type calculator struct {
	storage expression.Storage
	parser  command.Parser
}

func (c *calculator) RunCommand(str string) {
	cmd := c.parser.CreateCommand(str)
	if cmd != nil {
		(*cmd).Execute()
	}
}
