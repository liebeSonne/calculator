package expression

type Storage interface {
	SetFn(id string, fn *Function)
	SetVar(id string, v *Variable)
	Get(id string) *Computable
	Has(id string) bool
	IsFn(id string) bool
	GetVars() []*Variable
	GetFns() []*Function
	GetVar(id string) *Variable
}

func NewStorage() Storage {
	return &storage{
		functions: make(map[string]*Function),
		variables: make(map[string]*Variable),
	}
}

type storage struct {
	functions map[string]*Function
	variables map[string]*Variable
}

func (s *storage) GetVars() []*Variable {
	var result []*Variable
	for _, item := range s.variables {
		result = append(result, item)
	}
	return result
}

func (s *storage) GetFns() []*Function {
	var result []*Function
	for _, item := range s.functions {
		result = append(result, item)
	}
	return result
}

func (s *storage) SetFn(id string, fn *Function) {
	s.functions[id] = fn
}

func (s *storage) SetVar(id string, v *Variable) {
	s.variables[id] = v
}

func (s *storage) Get(id string) *Computable {
	if item, ok := s.variables[id]; ok {
		var c Computable
		c = *item
		return &c
	}
	if item, ok := s.functions[id]; ok {
		var c Computable
		c = *item
		return &c
	}
	return nil
}

func (s *storage) Has(id string) bool {
	if _, ok := s.variables[id]; ok {
		return true
	}
	if _, ok := s.functions[id]; ok {
		return true
	}
	return false
}

func (s *storage) IsFn(id string) bool {
	if _, ok := s.functions[id]; ok {
		return true
	}
	return false
}

func (s *storage) GetVar(id string) *Variable {
	if item, ok := s.variables[id]; ok {
		return item
	}
	return nil
}
