package cmap

import (
	"reflect"
	"sync"
)

type ConcurrentMap interface {
	Get(key interface{}) interface{}
	Put(key, value interface{})
	Remove(key interface{})
	Clear()
	Len() int
	Contains(key interface{}) bool
	Keys() []interface{}
	Values() []interface{}
	ToMap() map[interface{}]interface{}
	KeyType() reflect.Type
	ValueType() reflect.Type
}

type ConcurrentMemMap struct {
	data      map[interface{}]interface{}
	rwmu      sync.RWMutex
	keyType   reflect.Type
	valueType reflect.Type
}

func NewConcurrentMap(keyType, valueType reflect.Type) ConcurrentMap {
	return &ConcurrentMemMap{
		keyType:   keyType,
		valueType: valueType,
		data:      make(map[interface{}]interface{}),
	}
}

func (m *ConcurrentMemMap) Get(key interface{}) interface{} {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	return m.data[key]
}

func (m *ConcurrentMemMap) Put(key, value interface{}) {
	if reflect.TypeOf(key) != m.keyType {
		return
	}
	if reflect.TypeOf(value) != m.valueType {
		return
	}

	m.rwmu.Lock()
	defer m.rwmu.Unlock()
	m.data[key] = value
}

func (m *ConcurrentMemMap) Remove(key interface{}) {
	m.rwmu.Lock()
	defer m.rwmu.Unlock()
	delete(m.data, key)
}

func (m *ConcurrentMemMap) Clear() {
	m.rwmu.Lock()
	defer m.rwmu.Unlock()
	m.data = make(map[interface{}]interface{})
}

func (m *ConcurrentMemMap) Len() int {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	return len(m.data)
}

func (m *ConcurrentMemMap) Contains(key interface{}) bool {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *ConcurrentMemMap) Keys() []interface{} {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	res := make([]interface{}, len(m.data))
	idx := 0
	for k, _ := range m.data {
		res[idx] = k
		idx++
	}
	return res
}

func (m *ConcurrentMemMap) Values() []interface{} {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	res := make([]interface{}, len(m.data))
	idx := 0
	for _, v := range m.data {
		res[idx] = v
		idx++
	}
	return res
}

func (m *ConcurrentMemMap) ToMap() map[interface{}]interface{} {
	m.rwmu.RLock()
	defer m.rwmu.RUnlock()
	res := make(map[interface{}]interface{})
	for k, v := range m.data {
		res[k] = v
	}
	return res
}

func (m *ConcurrentMemMap) KeyType() reflect.Type {
	return m.keyType
}
func (m *ConcurrentMemMap) ValueType() reflect.Type {
	return m.valueType
}
