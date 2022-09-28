package command

import (
	"calculator/pkg/calc/expression"
	"fmt"
	"io"
	"sort"
	"strings"
)

func NewCommandPrintFns(storage expression.Storage, out io.Writer) Command {
	return &commandPrintFns{
		storage: storage,
		out:     out,
	}
}

type commandPrintFns struct {
	storage expression.Storage
	out     io.Writer
}

func (c commandPrintFns) Execute() {
	items := c.storage.GetFns()

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
