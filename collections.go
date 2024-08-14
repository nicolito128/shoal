package shoal

import (
	"slices"
	"sync"
)

func EqualSlice[S ~[]T, T comparable](s1 S, s2 S) bool {
	return slices.Equal(s1, s2)
}

func Equal[S ~[]T, T comparable](c1 *Collection[S, T], c2 *Collection[S, T]) bool {
	return slices.Equal(c1.Unwrap(), c2.Unwrap())
}

type Collection[S ~[]T, T any] struct {
	slice    S
	mu       sync.RWMutex
	readOnly bool
}

func NewCollection[S ~[]T, T any](slice S, readOnly bool) *Collection[S, T] {
	return &Collection[S, T]{
		slice:    slice,
		readOnly: readOnly,
	}
}

func (c *Collection[S, T]) Cap() int {
	return cap(c.slice)
}

func (c *Collection[S, T]) Len() int {
	return len(c.slice)
}

func (c *Collection[S, T]) ReadOnly() bool {
	return c.readOnly
}

func (c *Collection[S, T]) Get(ind int) (T, error) {
	if ind < 0 || ind >= c.Len() {
		var t T
		return t, ErrIndexOutOfRaange
	}

	c.mu.RLock()
	value := c.slice[ind]
	c.mu.RUnlock()
	return value, nil
}

func (c *Collection[S, T]) Set(ind int, value T) error {
	if c.readOnly {
		return ErrReadOnlyCollection
	}
	if ind < 0 || ind >= c.Len() {
		return ErrIndexOutOfRaange
	}

	c.mu.Lock()
	c.slice[ind] = value
	c.mu.Unlock()
	return nil
}

func (c *Collection[S, T]) Append(value ...T) error {
	if c.readOnly {
		return ErrReadOnlyCollection
	}

	c.mu.Lock()
	c.slice = append(c.slice, value...)
	c.mu.Unlock()
	return nil
}

func (c *Collection[S, T]) Unwrap() S {
	return c.slice
}
