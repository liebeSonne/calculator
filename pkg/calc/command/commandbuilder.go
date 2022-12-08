package command

import (
	"calculator/pkg/calc/expression"
	"calculator/pkg/calc/operation"
)

type CmcBuilder interface {
	CreateFnIdCommand(id, id2 string) *Command
	CreateFnOpCommand(id, id2 string, op operation.Operation, id3 string) *Command
}

func NewCmdBuilder(storage expression.Storage) CmcBuilder {
	return &cmdBuilder{
		storage: storage,
	}
}

type cmdBuilder struct {
	storage expression.Storage
}

func (p *cmdBuilder) CreateFnIdCommand(id, id2 string) *Command {
	cmd := NewCommandFn(id, id2, p.storage)
	return &cmd
}

func (p *cmdBuilder) CreateFnOpCommand(id, id2 string, op operation.Operation, id3 string) *Command {
	cmd := NewCommandFnOp(id, id2, op, id3, p.storage)
	return &cmd
}
