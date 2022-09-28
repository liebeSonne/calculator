package command

import (
	"fmt"
	"io"

	"calculator/pkg/calc/expression"
)

func NewCommandPrint(id string, storage expression.Storage, out io.Writer) Command {
	return &commandPrint{
		id:      id,
		storage: storage,
		out:     out,
	}
}

type commandPrint struct {
	id      string
	storage expression.Storage
	out     io.Writer
}

func (c *commandPrint) Execute() {
	if !c.storage.Has(c.id) {
		return
	}

	let := c.storage.Get(c.id)

	var value *float64
	value = (*let).Value()

	if value == nil {
		fmt.Fprintln(c.out, "nan")
		return
	}

	fmt.Fprintf(c.out, "%.2f\n", *value)
}
