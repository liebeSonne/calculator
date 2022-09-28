package command

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"calculator/pkg/calc/expression"
)

func NewCommandPrintVars(storage expression.Storage, out io.Writer) Command {
	return &commandPrintVars{
		storage: storage,
		out:     out,
	}
}

type commandPrintVars struct {
	storage expression.Storage
	out     io.Writer
}

func (c commandPrintVars) Execute() {
	items := c.storage.GetVars()

	sort.SliceStable(items, func(i, j int) bool {
		iID := (*items[i]).ID()
		jID := (*items[j]).ID()
		return strings.Compare(iID, jID) < 0
	})

	for _, item := range items {
		value := (*item).Value()
		if value == nil {
			fmt.Fprintf(c.out, "%s:nan\n", (*item).ID())
		} else {
			fmt.Fprintf(c.out, "%s:%.2f\n", (*item).ID(), *value)
		}
	}
}
