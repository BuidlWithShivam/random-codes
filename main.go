package main

import "fmt"

func main() {
	lruCacheTest()
}

func lruCacheTest() {
	type Test struct {
		name string
	}
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
