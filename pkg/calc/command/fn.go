package command

import (
	"calculator/pkg/calc/expression"
	"calculator/pkg/calc/operation"
)

func NewCommandFn(id, id2 string, storage expression.Storage) Command {
	return &commandFn{
		id:      id,
		id2:     id2,
		storage: storage,
	}
}

type commandFn struct {
	id      string
	id2     string
	storage expression.Storage
}

func (c commandFn) Execute() {
	if c.storage.Has(c.id) {
		return
	}

	let := c.storage.Get(c.id2)
	fn := expression.NewLetFunction(c.id, let)
	c.storage.SetFn(c.id, &fn)
}

func NewCommandFnOp(id, id2 string, operation operation.Operation, id3 string, storage expression.Storage) Command {
	return &commandFnOp{
		id:        id,
		id2:       id2,
		operation: operation,
		id3:       id3,
		storage:   storage,
	}
}

type commandFnOp struct {
	id        string
	id2       string
	operation operation.Operation
	id3       string
	storage   expression.Storage
}

func (c commandFnOp) Execute() {
	if c.storage.Has(c.id) {
		return
	}

	if !c.storage.Has(c.id2) {
		return
	}
	if !c.storage.Has(c.id3) {
		return
	}

	operand1 := c.storage.Get(c.id2)
	operand2 := c.storage.Get(c.id3)
	fn := expression.NewOperationFunction(c.id, operand1, c.operation, operand2)
	c.storage.SetFn(c.id, &fn)
}
