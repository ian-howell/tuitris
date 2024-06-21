package ring

import "errors"

type Ring[T any] interface {
	Next()
	Prev()
	Get() T
	Values() []T
	Cursor() int
	Reset()
}

type ring[T any] struct {
	values []T
	cursor int
}

func New[T any](values ...T) (Ring[T], error) {
	if len(values) == 0 {
		return nil, errors.New("cannot create empty ring buffer")
	}
	return &ring[T]{values: values}, nil
}

func (r *ring[T]) Next() {
	r.cursor = (r.cursor + 1) % len(r.values)
}

func (r *ring[T]) Prev() {
	r.cursor = (r.cursor + len(r.values) - 1) % len(r.values)
}

func (r *ring[T]) Get() T {
	return r.values[r.cursor]
}

func (r *ring[T]) Values() []T {
	return r.values
}

func (r *ring[T]) Cursor() int {
	return r.cursor
}

func (r *ring[T]) Reset() {
	r.cursor = 0
}
