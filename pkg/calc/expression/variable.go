package expression

type Variable interface {
	Computable
	ID() string
	SetValue(value *float64)
}

func NewVariable(id string, value *float64) Variable {
	return &variable{
		id:    id,
		value: value,
	}
}

type variable struct {
	id    string
	value *float64
}

func (v *variable) ID() string {
	return v.id
}

func (v *variable) Value() *float64 {
	return v.value
}

func (v *variable) SetValue(value *float64) {
	v.value = value
}
