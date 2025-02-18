package main

import (
	"fmt"
	"random-codes/datastructures"
)

type Test struct {
	name string
}

func main() {
	lfuCacheTest()
}

func minHeapTest() {
	minHeap := datastructures.NewHeap[int, Test](5, false)
	minHeap.Insert(1, Test{name: "A"})
	minHeap.PrintHeap()
	minHeap.Insert(2, Test{name: "B"})
	minHeap.PrintHeap()
	minHeap.Insert(3, Test{name: "C"})
	minHeap.PrintHeap()
	minHeap.Insert(4, Test{name: "D"})
	minHeap.PrintHeap()
	minHeap.Insert(0, Test{name: "E"})
	minHeap.PrintHeap()

	pair := minHeap.Peek()
	fmt.Println(*pair)
	pair = minHeap.Pop()
	fmt.Println(*pair)
	minHeap.PrintHeap()
}

func lruCacheTest() {
	lruCache := datastructures.NewLRUCache[string, Test](3)
	lruCache.Put("A", Test{name: "1"})
	lruCache.Put("B", Test{name: "2"})
	lruCache.Put("C", Test{name: "3"})
	lruCache.Put("D", Test{name: "4"})
	lruCache.List.PrintList()

	value, err := lruCache.Get("B")
	if err != nil {
		panic(err)
	}
	fmt.Println("Value for B in cache: ", value)
	lruCache.List.PrintList()

	err = lruCache.Remove("C")

	lruCache.List.PrintList()

	lruCache.Put("E", Test{name: "5"})
	lruCache.List.PrintList()

}

func lfuCacheTest() {
	cache := datastructures.NewLFUCache[string, Test](3)
	cache.Put("A", Test{name: "1"})
	cache.Put("B", Test{name: "2"})
	cache.Put("C", Test{name: "3"})
	cache.Put("D", Test{name: "4"})

	cache.Order.PrintHeap()

	value, err := cache.Get("B")
	if err != nil {
		panic(err)
	}
	fmt.Println("Value for B in cache: ", value)

	err = cache.Remove("C")

	cache.Order.PrintHeap()

	cache.Put("E", Test{name: "5"})
	cache.Order.PrintHeap()

	cache.Get("D")
	cache.Get("D")

	cache.Order.PrintHeap()
}

// build a library with comparable can be used as generic
