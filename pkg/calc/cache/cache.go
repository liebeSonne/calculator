package cache

import (
	"reflect"
)

type ValueStorage interface {
	Get(id string) *float64
	Set(id string, value *float64)
	Has(id string) bool
	Del(id string)
}

type valueStorage struct {
	storage map[string]*float64
}

func (s *valueStorage) Get(id string) *float64 {
	if value, ok := s.storage[id]; ok {
		return value
	}
	return nil
}

func (s *valueStorage) Set(id string, value *float64) {
	s.storage[id] = value
}

func (s *valueStorage) Has(id string) bool {
	if _, ok := s.storage[id]; ok {
		return true
	}
	return false
}

func (s *valueStorage) Del(id string) {
	delete(s.storage, id)
}

type CacheStorage interface {
	OnLet(id string)
	OnValue(id string, value *float64)
	SetDependencies(id string, ids []string)
	ValueStorage
}

func NewCacheStorage() CacheStorage {
	return &cacheStorage{
		//cache:        NewValueStorage(),
		dependencies: make(map[string][]string),
		valueStorage: valueStorage{
			storage: make(map[string]*float64),
		},
	}
}

type cacheStorage struct {
	valueStorage
	dependencies map[string][]string
}

// Очищаем кеш зависимых элементов при изменении значения
func (s *cacheStorage) OnLet(id string) {
	for _, did := range s.getDependencies(id) {
		if s.Has(did) {
			// чистим кеш у зависимостей зависимости
			s.OnLet(did)
		}
		s.Del(did)
	}
}

// сохраняем кеш при вычислении значения
func (s *cacheStorage) OnValue(id string, value *float64) {
	s.Set(id, value)
}

// Сохраняем зависимости
func (s *cacheStorage) SetDependencies(id string, ids []string) {
	for _, did := range ids {
		if _, ok := s.dependencies[did]; ok {
			if is, _ := in_array(id, s.dependencies[did]); !is {
				s.dependencies[did] = append(s.dependencies[did], id)
			}
		} else {
			s.dependencies[did] = []string{id}
		}
	}
}

func (s *cacheStorage) getDependencies(id string) []string {
	if ids, ok := s.dependencies[id]; ok {
		return ids
	}
	return make([]string, 0)
}

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
