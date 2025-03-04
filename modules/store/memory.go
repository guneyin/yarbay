package store

import "sync"

type memory struct {
	sync.Map
}

func (m *memory) Set(key string, value any) {
	m.Store(key, value)
}

func (m *memory) Get(key string) Value[any] {
	val, _ := m.Load(key)
	return Value[any]{val}
}

func (m *memory) GetOnce(key string) (*Value[any], error) {
	val, ok := m.LoadAndDelete(key)
	if !ok {
		return nil, ErrRecordNotFound
	}
	return &Value[any]{val}, nil
}

func (m *memory) Delete(key string) {
	m.Delete(key)
}
