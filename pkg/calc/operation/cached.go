package operation

func NewCacheOperationStorage() CacheStorage {
	return CacheStorage{
		storage: make(map[cacheKey]*float64),
	}
}

type CacheStorage struct {
	storage map[cacheKey]*float64
}

func (s *CacheStorage) HasValue(key cacheKey) bool {
	if _, ok := s.storage[key]; ok {
		return true
	}
	return false
}

func (s *CacheStorage) GetValue(key cacheKey) *float64 {
	if result, ok := s.storage[key]; ok {
		return result
	}
	return nil
}

func (s *CacheStorage) SetValue(key cacheKey, value *float64) {
	s.storage[key] = value
}

func NewCachedOperation(operation Operation, name string, storage *CacheStorage) Operation {
	return &cacheOperation{
		storage:   storage,
		name:      name,
		operation: operation,
	}
}

type cacheKey struct {
	name string
	v1   *float64
	v2   *float64
}

type cacheOperation struct {
	operation Operation
	name      string
	storage   *CacheStorage
}

func (p *cacheOperation) Calculate(operand1 Operand, operand2 Operand) *float64 {
	key := cacheKey{p.name, operand1.Value(), operand2.Value()}

	if p.storage.HasValue(key) {
		return p.storage.GetValue(key)
	}

	result := p.operation.Calculate(operand1, operand2)
	p.storage.SetValue(key, result)
	return result
}
