package cache

import (
	"calculator/pkg/calc/expression"
)

func NewCacheForStorage(storage expression.Storage, cache CacheStorage) expression.Storage {
	return &cacheForStorage{
		storage: storage,
		cache:   cache,
	}
}

type cacheForStorage struct {
	storage expression.Storage
	cache   CacheStorage
}

func (c cacheForStorage) SetFn(id string, fn *expression.Function) {
	cFn := NewCacheFunction(fn, c.cache)
	c.storage.SetFn(id, &cFn)
}

func (c cacheForStorage) SetVar(id string, v *expression.Variable) {
	cv := NewCacheVariable(v, c.cache)
	c.storage.SetVar(id, &cv)
}

func (c cacheForStorage) Get(id string) *expression.Computable {
	return c.storage.Get(id)
}

func (c cacheForStorage) Has(id string) bool {
	return c.storage.Has(id)
}

func (c cacheForStorage) IsFn(id string) bool {
	return c.storage.IsFn(id)
}

func (c cacheForStorage) GetVars() []*expression.Variable {
	return c.storage.GetVars()
}

func (c cacheForStorage) GetFns() []*expression.Function {
	return c.storage.GetFns()
}

func (c cacheForStorage) GetVar(id string) *expression.Variable {
	return c.storage.GetVar(id)
}
