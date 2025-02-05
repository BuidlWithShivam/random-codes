package main

import (
	"errors"
)

type LFUCache[K KeyConstraint, T any] struct {
	capacity    int
	cache       map[K]T
	cacheToHeap map[K]int
	order       *Heap[int, K]
}

func NewLFUCache[K KeyConstraint, T any](capacity int) *LFUCache[K, T] {
	return &LFUCache[K, T]{
		capacity:    capacity,
		cache:       make(map[K]T, capacity),
		cacheToHeap: make(map[K]int),
		order:       NewHeap[int, K](capacity, false),
	}
}

func (l *LFUCache[K, T]) Get(key K) (T, error) {
	value, ok := l.cache[key]
	if !ok {
		return getZero[T](), errors.New("key not found")
	}
	l.cacheToHeap[key] = l.order.IncreaseKey(l.cacheToHeap[key], l.order.heap[l.cacheToHeap[key]].key+1)
	return value, nil
}

func (l *LFUCache[K, T]) Put(key K, value T) {
	index := 0
	if len(l.cache) == l.capacity {
		pair := l.order.Peek()
		l.order.heap[0] = &Pair[int, K]{
			0,
			key,
		}
		delete(l.cacheToHeap, pair.val)
		delete(l.cache, pair.val)
	} else {
		index = l.order.Insert(0, key)
	}
	l.cache[key] = value
	l.cacheToHeap[key] = index
}

func (l *LFUCache[K, T]) Remove(key K) error {
	index := l.cacheToHeap[key]
	if index != len(l.order.heap)-1 {
		l.order.heap[index], l.order.heap[len(l.order.heap)-1] =
			swap[int, K](l.order.heap[index], l.order.heap[len(l.order.heap)-1])
		l.order.heap = l.order.heap[:len(l.order.heap)-1]
		newIndex := Heapify[int, K](l.order.heap, index, false)
		l.cacheToHeap[l.order.heap[newIndex].val] = newIndex
	} else {
		l.order.heap = l.order.heap[:len(l.order.heap)-1]
	}
	delete(l.cache, key)
	delete(l.cacheToHeap, key)
	return nil
}
