package trmutex

import (
	"fmt"
	"sync"
)

type Mutex struct {
	id      string
	factory *Factory
}

func (m *Mutex) Lock() {
	m.factory.acquire(m.id).Lock()
}

func (m *Mutex) Unlock() {
	m.factory.release(m.id).Unlock()
}

type mutexItem struct {
	ref   int
	mutex sync.Mutex
}

type Factory struct {
	items map[string]*mutexItem
	mutex sync.Mutex
}

func NewFactory() *Factory {
	return &Factory{
		items: make(map[string]*mutexItem),
	}
}

func (f *Factory) Require(id string) Mutex {
	return Mutex{
		id:      id,
		factory: f,
	}
}

func (f *Factory) acquire(id string) *sync.Mutex {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	item, ok := f.items[id]
	if !ok {
		item = &mutexItem{ref: 1}
		f.items[id] = item
	} else {
		item.ref++
	}
	return &item.mutex
}

func (f *Factory) release(id string) *sync.Mutex {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	item, ok := f.items[id]
	if !ok {
		panic(fmt.Errorf("item(%d) not found", id))
	}
	if item.ref <= 0 {
		panic(fmt.Errorf("invalid ref: %d", item.ref))
	}
	item.ref--
	if item.ref == 0 {
		delete(f.items, id)
	}
	return &item.mutex
}
