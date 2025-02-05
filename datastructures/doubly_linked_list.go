package datastructures

import "fmt"

type Node[K, T any] struct {
	key  K
	data T
	prev *Node[K, T]
	next *Node[K, T]
}
type DoublyLinkedList[K, T any] struct {
	head *Node[K, T]
	tail *Node[K, T]
}

func (list *DoublyLinkedList[K, T]) Add(key K, value T) *Node[K, T] {
	newNode := &Node[K, T]{key, value, nil, nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
	return newNode
}

func (list *DoublyLinkedList[K, T]) Remove(node *Node[K, T]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}
}

func (list *DoublyLinkedList[K, T]) MoveToFront(node *Node[K, T]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		return
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}
	node.next = list.head
	list.head.prev = node
	list.head = node
	node.prev = nil
}

func (list *DoublyLinkedList[K, T]) RemoveLast() *Node[K, T] {
	node := list.tail
	if list.tail != nil {
		if list.tail.prev != nil {
			list.tail.prev.next = nil
			list.tail = list.tail.prev
		} else {
			list.head = nil
			list.tail = nil
		}
	} else {
		panic("Doubly linked list is empty")
	}
	return node
}

func (list *DoublyLinkedList[K, T]) PrintList() {
	for node := list.head; node != nil; node = node.next {
		fmt.Print(node.key, "->", node.data, ", ")
	}
	fmt.Println()
}
