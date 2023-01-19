package calc

import (
	"io"

	"calculator/pkg/calc/command"
)

type Calculator interface {
	RunCommand(str string)
}

func NewCalculator(
	parser command.Parser,
	out io.Writer,
) Calculator {
	return &calculator{
		parser: parser,
		out:    out,
	}
}

type calculator struct {
	parser command.Parser
	out    io.Writer
}

func (c *calculator) RunCommand(str string) {
	cmd := c.parser.CreateCommand(str)
	if cmd != nil {
		(*cmd).Execute()
	} else {
		c.out.Write([]byte("[!] unknown command\n"))
	}
}
