package command

import "calculator/pkg/calc/expression"

func NewCommandLetVal(id string, value *float64, storage expression.Storage) Command {
	return &commandLetVal{
		id:      id,
		value:   value,
		storage: storage,
	}
}

type commandLetVal struct {
	id      string
	value   *float64
	storage expression.Storage
}

func (c *commandLetVal) Execute() {
	if c.storage.Has(c.id) {
		if c.storage.IsFn(c.id) {
			return
		}
		let := c.storage.GetVar(c.id)
		(*let).SetValue(c.value)
	} else {
		variable := expression.NewVariable(c.id, c.value)
		c.storage.SetVar(c.id, &variable)
	}
}

func NewCommandLetLet(id, letId string, storage expression.Storage) Command {
	return &commandLetLet{
		id:      id,
		letId:   letId,
		storage: storage,
	}
}

type commandLetLet struct {
	id      string
	letId   string
	storage expression.Storage
}

func (c *commandLetLet) Execute() {
	if c.storage.IsFn(c.id) {
		return
	}
	let := c.storage.Get(c.letId)
	var value *float64
	if let != nil {
		value = (*let).Value()
	}
	if c.storage.Has(c.id) {
		let := c.storage.GetVar(c.id)
		(*let).SetValue(value)
	} else {
		variable := expression.NewVariable(c.id, value)
		c.storage.SetVar(c.id, &variable)
	}
}
