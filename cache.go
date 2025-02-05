package main

import "cmp"

type KeyConstraint interface {
	cmp.Ordered
}

type Cache[K KeyConstraint, T any] interface {
	Get(key K) (T, error)
	Put(key K, value T)
	Remove(key K) error
}
