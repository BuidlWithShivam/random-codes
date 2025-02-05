package datastructures

import (
	"cmp"
	"fmt"
)

type Pair[K cmp.Ordered, T any] struct {
	key K
	val T
}

func (p *Pair[K, T]) Compare(pair *Pair[K, T], reverse bool) bool {
	if reverse {
		return cmp.Compare(p.key, pair.key) > 0
	}
	return cmp.Compare(p.key, pair.key) < 0
}

type Heap[K cmp.Ordered, T any] struct {
	heap     []*Pair[K, T]
	capacity int
	maxHeap  bool
}

func NewHeap[K cmp.Ordered, T any](capacity int, maxHeap bool) *Heap[K, T] {
	return &Heap[K, T]{
		heap:     make([]*Pair[K, T], 0, capacity),
		capacity: capacity,
		maxHeap:  maxHeap,
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func swap[K cmp.Ordered, T any](first, second *Pair[K, T]) (*Pair[K, T], *Pair[K, T]) {
	return second, first
}

func (h *Heap[K, T]) Insert(key K, val T) int {
	pair := &Pair[K, T]{key, val}
	if len(h.heap) == h.capacity {
		if pair.Compare(h.Peek(), h.maxHeap) {
			return -1
		} else {
			h.heap[0] = pair
			currentIndex := 0
			for {
				leftIndex := left(currentIndex)
				rightIndex := right(currentIndex)
				var scope int // smallest if min heap and largest if max heap
				scope = currentIndex
				if leftIndex < len(h.heap) && h.heap[leftIndex].Compare(h.heap[scope], h.maxHeap) {
					scope = leftIndex
				}

				if rightIndex < len(h.heap) && h.heap[rightIndex].Compare(h.heap[scope], h.maxHeap) {
					scope = rightIndex
				}

				if scope == currentIndex {
					return currentIndex
				}
				h.heap[currentIndex], h.heap[scope] = swap(h.heap[currentIndex], h.heap[scope])
				currentIndex = scope
			}
		}
	} else {
		h.heap = append(h.heap, pair)
		currentIndex := len(h.heap) - 1
		for currentIndex > 0 && pair.Compare(h.heap[parent(currentIndex)], h.maxHeap) {
			h.heap[currentIndex], h.heap[parent(currentIndex)] = swap(h.heap[currentIndex], h.heap[parent(currentIndex)])
			currentIndex = parent(currentIndex)
		}
		return currentIndex
	}
}

func (h *Heap[K, T]) Peek() *Pair[K, T] {
	if len(h.heap) == 0 {
		return nil
	}
	return h.heap[0]
}

func (h *Heap[K, T]) Pop() *Pair[K, T] {
	if len(h.heap) == 0 {
		return nil
	}
	pair := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	currentIndex := 0
	for {
		leftIndex := left(currentIndex)
		rightIndex := right(currentIndex)
		var scope int // smallest if min heap, largest if max heap
		scope = currentIndex
		if leftIndex < len(h.heap) && h.heap[leftIndex].Compare(h.heap[scope], h.maxHeap) {
			scope = leftIndex
		}

		if rightIndex < len(h.heap) && h.heap[rightIndex].Compare(h.heap[scope], h.maxHeap) {
			scope = rightIndex
		}

		if scope == currentIndex {
			break
		}
		h.heap[currentIndex], h.heap[scope] = swap(h.heap[currentIndex], h.heap[scope])
		currentIndex = scope
	}
	return pair
}

func (h *Heap[K, T]) IncreaseKey(index int, key K) int {
	h.heap[index].key = key
	if !h.maxHeap {
		return Heapify[K, T](h.heap, index, h.maxHeap)
	} else {
		currentIndex := index
		for currentIndex > 0 && h.heap[currentIndex].Compare(h.heap[parent(currentIndex)], h.maxHeap) {
			h.heap[currentIndex], h.heap[parent(currentIndex)] = swap(h.heap[currentIndex], h.heap[parent(currentIndex)])
			currentIndex = parent(currentIndex)
		}
		return currentIndex
	}
}

func Heapify[K cmp.Ordered, T any](heap []*Pair[K, T], index int, maxHeap bool) int {
	leftIndex := left(index)
	rightIndex := right(index)
	var scope int // smallest if min heap, largest if max heap
	scope = index

	if leftIndex < len(heap) && heap[leftIndex].Compare(heap[scope], maxHeap) {
		scope = leftIndex
	}

	if rightIndex < len(heap) && heap[rightIndex].Compare(heap[scope], maxHeap) {
		scope = rightIndex
	}

	if scope == index {
		return index
	}

	heap[index], heap[scope] = swap(heap[index], heap[scope])
	return Heapify[K, T](heap, scope, maxHeap)
}

func BuildHeap[K cmp.Ordered, T any](heap []*Pair[K, T], minHeap bool) []*Pair[K, T] {
	length := len(heap)
	startIndex := length/2 - 1
	for i := startIndex; i >= 0; i-- {
		Heapify(heap, i, minHeap)
	}
	return heap
}

func (h *Heap[K, T]) PrintHeap() {
	for _, pair := range h.heap {
		fmt.Print(*pair)
	}
	fmt.Println()
}
