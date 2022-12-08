package cache

import (
	"calculator/pkg/calc/expression"
)

func NewCacheFunction(fn *expression.Function, cache CacheStorage) expression.Function {
	return &cacheFunction{
		cache: cache,
		fn:    fn,
	}
}

type cacheFunction struct {
	cache CacheStorage
	fn    *expression.Function
}

func (c cacheFunction) Value() *float64 {
	if c.cache.Has((*c.fn).ID()) {
		return c.cache.Get((*c.fn).ID())
	}

	value := (*c.fn).Value()
	c.cache.OnValue((*c.fn).ID(), value)
	return value
}

func (c cacheFunction) ID() string {
	return (*c.fn).ID()
}

func NewCacheVariable(variable *expression.Variable, cache CacheStorage) expression.Variable {
	return &cacheVariable{
		cache:    cache,
		variable: variable,
	}
}

type cacheVariable struct {
	cache    CacheStorage
	variable *expression.Variable
}

func (c cacheVariable) Value() *float64 {
	return (*c.variable).Value()
}

func (c cacheVariable) ID() string {
	return (*c.variable).ID()
}

func (c cacheVariable) SetValue(value *float64) {
	(*c.variable).SetValue(value)
	c.cache.OnLet((*c.variable).ID())
}
