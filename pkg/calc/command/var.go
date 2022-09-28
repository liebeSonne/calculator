package command

import (
	"calculator/pkg/calc/expression"
)

func NewCommandVar(id string, storage expression.Storage) Command {
	return &commandVar{
		id:      id,
		storage: storage,
	}
}

type commandVar struct {
	id      string
	storage expression.Storage
}

func (c *commandVar) Execute() {
	if !c.storage.Has(c.id) {
		variable := expression.NewVariable(c.id, nil)
		c.storage.SetVar(c.id, &variable)
	}
}
