package main

import (
	"container/list"
	"sync"
)

type LRU[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}
type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	mu       sync.Mutex
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	c := &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
	c.ll.Init()
	return c
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, found := c.items[key]
	if !found {
		return value, false
	}
	c.ll.MoveToFront(e)
	return e.Value.(*entry[K, V]).value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.capacity <= 0 {
		return
	}
	if e, found := c.items[key]; found {
		e.Value.(*entry[K, V]).value = value
		return
	}
	if c.ll.Len() >= c.capacity {
		oldest := c.ll.Back()
		c.ll.Remove(oldest)
		delete(c.items, oldest.Value.(*entry[K, V]).key)
	}
	c.items[key] = c.ll.PushFront(&entry[K, V]{key: key, value: value})
}
