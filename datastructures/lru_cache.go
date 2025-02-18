package datastructures

import (
	"errors"
	"fmt"
	"random-codes/utils"
	"sync"
)

type LRUCache[K KeyConstraint, T any] struct {
	capacity int
	size     int
	List     *DoublyLinkedList[K, T]
	values   map[K]*Node[K, T]
	mutex    sync.Mutex
}

func NewLRUCache[K KeyConstraint, T any](capacity int) *LRUCache[K, T] {
	return &LRUCache[K, T]{
		capacity: capacity,
		size:     0,
		List:     &DoublyLinkedList[K, T]{nil, nil},
		values:   make(map[K]*Node[K, T]),
	}
}

func (c *LRUCache[K, T]) Put(key K, data T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.size == c.capacity {
		removedNode := c.List.RemoveLast()
		fmt.Println("Replaced from cache, key : ", removedNode.key)
		c.size--
	}
	node := c.List.Add(key, data)
	c.values[key] = node
	c.size++
	fmt.Println("Key, value added: ", key, ":", data)
}

func (c *LRUCache[K, T]) Get(key K) (T, error) {
	node, ok := c.values[key]
	if !ok {
		return utils.GetZero[T](), errors.New(fmt.Sprintf("Key not found in cache : ", key))
	}
	c.mutex.Lock()
	c.List.MoveToFront(node)
	c.mutex.Unlock()
	fmt.Println("Getting from cache, key : ", key)
	return node.data, nil
}

func (c *LRUCache[K, T]) Remove(key K) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	node, ok := c.values[key]
	if ok {
		fmt.Println("Key removed from cache : ", key)
		c.List.Remove(node)
		delete(c.values, key)
		c.size--
		return nil
	} else {
		return errors.New(fmt.Sprintf("Key not found in cache : ", key))
	}
}
