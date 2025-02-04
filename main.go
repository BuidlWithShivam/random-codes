package main

import "fmt"

type Test struct {
	name string
}

func main() {
	minHeapTest()
}

func minHeapTest() {
	minHeap := NewHeap[int, Test](5, false)
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
	lruCache := NewLRUCache[string, Test](3)
	lruCache.Put("A", Test{name: "1"})
	lruCache.Put("B", Test{name: "2"})
	lruCache.Put("C", Test{name: "3"})
	lruCache.Put("D", Test{name: "4"})
	lruCache.list.PrintList()

	value, err := lruCache.Get("B")
	if err != nil {
		panic(err)
	}
	fmt.Println("Value for B in cache: ", value)
	lruCache.list.PrintList()

	err = lruCache.Remove("C")

	lruCache.list.PrintList()

	lruCache.Put("E", Test{name: "5"})
	lruCache.list.PrintList()
}
