package command

import (
	"io"
	"regexp"
	"strconv"

	"calculator/pkg/calc/expression"
	"calculator/pkg/calc/operation"
)

const (
	idParam        = "id"
	id2Param       = "id2"
	id3Param       = "id3"
	valParam       = "value"
	operationParam = "operation"

	operationPlus     = "+"
	operationMinus    = "-"
	operationMultiply = "*"
	operationDivision = "/"
)

var (
	idReg        = `[a-zA-Z][a-zA-Z0-9_]*`
	valReg       = `[0-9]+(\.[0-9]*)?`
	operationReg = `\+|\-|\*|\/`

	cmdVarRegexp       = regexp.MustCompile(`^var (?P<` + idParam + `>` + idReg + `)$`)
	cmdLetValRegexp    = regexp.MustCompile(`^let (?P<` + idParam + `>` + idReg + `)=(?P<` + valParam + `>` + valReg + `)$`)
	cmdLetLetRegexp    = regexp.MustCompile(`^let (?P<` + idParam + `>` + idReg + `)=(?P<` + id2Param + `>` + idReg + `)$`)
	cmdFnIdRegexp      = regexp.MustCompile(`^fn (?P<` + idParam + `>` + idReg + `)=(?P<` + id2Param + `>` + idReg + `)$`)
	cmdFnOpRegexp      = regexp.MustCompile(`^fn (?P<` + idParam + `>` + idReg + `)=(?P<` + id2Param + `>` + idReg + `)(?P<` + operationParam + `>` + operationReg + `)(?P<` + id3Param + `>` + idReg + `)$`)
	cmdPrintRegexp     = regexp.MustCompile(`^print (?P<` + idParam + `>` + idReg + `)$`)
	cmdPrintVarsRegexp = regexp.MustCompile(`^printvars$`)
	cmdPrintFnsRegexp  = regexp.MustCompile(`^printfns$`)
)

type Parser interface {
	CreateCommand(str string) *Command
}

func NewParser(storage expression.Storage, out io.Writer) Parser {
	return &parser{
		storage: storage,
		out:     out,
	}
}

type parser struct {
	storage expression.Storage
	out     io.Writer
}

func (p *parser) CreateCommand(str string) *Command {
	var cmd *Command

	cmd = p.createVarCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createLetValCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createLetLetCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createFnIdCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createFnOpCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createPrintCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createPrintVarsCommand(str)
	if cmd != nil {
		return cmd
	}

	cmd = p.createPrintFnsCommand(str)
	if cmd != nil {
		return cmd
	}

	return nil
}

func (p *parser) createVarCommand(str string) *Command {
	re := cmdVarRegexp
	if re.MatchString(str) {
		data := p.getSubmatchData(re, str)
		id := p.getDataId(data)
		if id == nil {
			return nil
		}
		cmd := NewCommandVar(*id, p.storage)
		return &cmd
	}
	return nil
}

func (p *parser) createLetValCommand(str string) *Command {
	re := cmdLetValRegexp
	if re.MatchString(str) {
		data := p.getSubmatchData(re, str)
		id := p.getDataId(data)
		value := p.getDataValue(data)
		if id == nil || value == nil {
			return nil
		}
		cmd := NewCommandLetVal(*id, value, p.storage)
		return &cmd
	}
	return nil
}

func (p *parser) createLetLetCommand(str string) *Command {
	re := cmdLetLetRegexp
	if re.MatchString(str) {
		data := p.getSubmatchData(re, str)
		id := p.getDataId(data)
		id2 := p.getDataId2(data)
		if id == nil || id2 == nil {
			return nil
		}
		cmd := NewCommandLetLet(*id, *id2, p.storage)
		return &cmd
	}
	return nil
}

func (p *parser) createFnIdCommand(str string) *Command {
	re := cmdFnIdRegexp
	if re.MatchString(str) {
		data := p.getSubmatchData(re, str)
		id := p.getDataId(data)
		id2 := p.getDataId2(data)
		if id == nil || id2 == nil {
			return nil
		}
		cmd := NewCommandFn(*id, *id2, p.storage)
		return &cmd
	}
	return nil
}

func (p *parser) createFnOpCommand(str string) *Command {
	re := cmdFnOpRegexp
	if re.MatchString(str) {
		data := p.getSubmatchData(re, str)
		id := p.getDataId(data)
		op := p.getDataOperation(data)
		id2 := p.getDataId2(data)
		id3 := p.getDataId3(data)
		if id == nil || id2 == nil || id3 == nil || op == nil {
			return nil
		}
		operation := p.createOperator(*op)
		if operation == nil {
			return nil
		}

		cmd := NewCommandFnOp(*id, *id2, *operation, *id3, p.storage)
		return &cmd
	}
	return nil
}

func (p *parser) createPrintCommand(str string) *Command {
	re := cmdPrintRegexp
	if re.MatchString(str) {
		data := p.getSubmatchData(re, str)
		id := p.getDataId(data)
		if id == nil {
			return nil
		}
		cmd := NewCommandPrint(*id, p.storage, p.out)
		return &cmd
	}
	return nil
}

func (p *parser) createPrintVarsCommand(str string) *Command {
	re := cmdPrintVarsRegexp
	if re.MatchString(str) {
		cmd := NewCommandPrintVars(p.storage, p.out)
		return &cmd
	}
	return nil
}

func (p *parser) createPrintFnsCommand(str string) *Command {
	re := cmdPrintFnsRegexp
	if re.MatchString(str) {
		cmd := NewCommandPrintFns(p.storage, p.out)
		return &cmd
	}
	return nil
}

func (p *parser) createOperator(str string) *operation.Operation {
	switch str {
	case operationPlus:
		operation := operation.NewPlusOperation()
		return &operation
	case operationMinus:
		operation := operation.NewMinusOperation()
		return &operation
	case operationMultiply:
		operation := operation.NewMultiplyOperation()
		return &operation
	case operationDivision:
		operation := operation.NewDivisionOperation()
		return &operation
	}
	return nil
}

func (p *parser) getSubmatchData(re *regexp.Regexp, message string) map[string]string {
	values := re.FindStringSubmatch(message)
	keys := re.SubexpNames()

	data := make(map[string]string)

	for i := range keys {
		if i == 0 {
			continue
		}
		if keys[i] != "" && values[i] != "" {
			data[keys[i]] = values[i]
		}
		data[keys[i]] = values[i]
	}

	return data
}

func (p *parser) getDataId(data map[string]string) *string {
	if id, ok := data[idParam]; ok {
		return &id
	}
	return nil
}

func (p *parser) getDataId2(data map[string]string) *string {
	if id, ok := data[id2Param]; ok {
		return &id
	}
	return nil
}

func (p *parser) getDataId3(data map[string]string) *string {
	if id, ok := data[id3Param]; ok {
		return &id
	}
	return nil
}

func (p *parser) getDataValue(data map[string]string) *float64 {
	if dataValue, ok := data[valParam]; ok {
		value, err := strconv.ParseFloat(dataValue, 64)
		if err != nil {
			return nil
		}
		return &value
	}
	return nil
}

func (p *parser) getDataOperation(data map[string]string) *string {
	if operation, ok := data[operationParam]; ok {
		return &operation
	}
	return nil
}
