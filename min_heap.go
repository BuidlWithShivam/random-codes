package main

import (
	"cmp"
	"fmt"
)

// add O(log n)
// remove O(log n)
// update O(log n)
// get O(1)
// Heapify

type Pair[K cmp.Ordered, T any] struct {
	key K
	val T
}

func (p *Pair[K, T]) Compare(pair *Pair[K, T]) int {
	return cmp.Compare(p.key, pair.key)
}

type MinHeap[K cmp.Ordered, T any] struct {
	heap     []*Pair[K, T]
	capacity int
}

func NewMinHeap[K cmp.Ordered, T any](capacity int) *MinHeap[K, T] {
	return &MinHeap[K, T]{
		heap:     make([]*Pair[K, T], 0, capacity),
		capacity: capacity,
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

func (h *MinHeap[K, T]) Insert(key K, val T) {
	pair := &Pair[K, T]{key, val}

	if len(h.heap) == h.capacity {
		if pair.Compare(h.Peek()) < 0 {
			return
		} else {
			h.heap[0] = pair
			currentIndex := 0
			for {
				leftIndex := left(currentIndex)
				rightIndex := right(currentIndex)
				smallest := currentIndex
				if leftIndex < len(h.heap) && h.heap[leftIndex].Compare(h.heap[smallest]) < 0 {
					smallest = leftIndex
				}

				if rightIndex < len(h.heap) && h.heap[rightIndex].Compare(h.heap[smallest]) < 0 {
					smallest = rightIndex
				}

				if smallest == currentIndex {
					break
				}
				h.heap[currentIndex], h.heap[smallest] = swap(h.heap[currentIndex], h.heap[smallest])
				currentIndex = smallest
			}
		}
	} else {
		h.heap = append(h.heap, pair)
		currentIndex := len(h.heap) - 1
		for currentIndex > 0 && pair.Compare(h.heap[parent(currentIndex)]) < 0 {
			h.heap[currentIndex], h.heap[parent(currentIndex)] = swap(h.heap[currentIndex], h.heap[parent(currentIndex)])
			currentIndex = parent(currentIndex)
		}
	}
}

func (h *MinHeap[K, T]) Peek() *Pair[K, T] {
	if len(h.heap) == 0 {
		return nil
	}
	return h.heap[0]
}

func (h *MinHeap[K, T]) Pop() *Pair[K, T] {
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
		smallest := currentIndex
		if leftIndex < len(h.heap) && h.heap[leftIndex].Compare(h.heap[smallest]) < 0 {
			smallest = leftIndex
		}

		if rightIndex < len(h.heap) && h.heap[rightIndex].Compare(h.heap[smallest]) < 0 {
			smallest = rightIndex
		}

		if smallest == currentIndex {
			break
		}
		h.heap[currentIndex], h.heap[smallest] = swap(h.heap[currentIndex], h.heap[smallest])
		currentIndex = smallest
	}
	return pair
}

func (h *MinHeap[K, T]) PrintHeap() {
	for _, pair := range h.heap {
		fmt.Print(*pair)
	}
	fmt.Println()
}
