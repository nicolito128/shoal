package shoal

import "iter"

/*
Map transforms a slice of one type into a slice of another type.

It uses Morph2 transformations, allowing for operations like:

	Map([]Type{...}, func(index int, value Type) NewType {...})
*/
func Map[S ~[]T, T, U any](collection S, mapping ...Morph2[int, T, U]) []U {
	result := make([]U, len(collection))
	for i, v := range MapSeq[S, T, U](collection, mapping...) {
		result[i] = v
	}
	return result
}

// MapSeq returns an iterator (Seq2) that applies a series of transformations
// to each element in the input slice.
func MapSeq[S ~[]T, T, U any](collection S, mapping ...Morph2[int, T, U]) iter.Seq2[int, U] {
	return func(yield func(int, U) bool) {
		for i, v := range collection {
			for _, morph := range mapping {
				if !yield(i, morph(i, v)) {
					return
				}
			}
		}
	}
}

/*
Filter returns a new slice containing only the elements that satisfy the given predicate(s).

It uses MorphCond transformations, allowing for operations like:

	Filter([]Type{...}, func(value Type) bool {...})
*/
func Filter[S ~[]T, T comparable](collection S, mapping ...MorphCond2[int, T]) S {
	filtered := make(S, 0, len(collection))
	for _, v := range FilterSeq(collection, mapping...) {
		filtered = append(filtered, v)
	}
	return filtered
}

// FilterSeq returns an iterator (Seq) that yields elements from the input slice that
// satisfy the given predicate(s).
func FilterSeq[S ~[]T, T comparable](collection S, mapping ...MorphCond2[int, T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range collection {
			ok := true
			for _, morph := range mapping {
				if !morph(i, v) {
					ok = false
					break
				}
			}

			if ok && !yield(i, v) {
				return
			}
		}
	}
}

/*
ForEach applies a function to each element in the input slice.

It uses Morph0 transformations, allowing for operations like:

	ForEach([]Type{...}, func(value Type) {...})
*/
func ForEach[S ~[]T, T any](collection S, morph Morph0[T]) {
	for v := range ForEachSeq(collection) {
		morph(v)
	}
}

// ForEachSeq returns an iterator (Seq) that yields each element in the input slice,
// allowing for operations to be performed on each element.
func ForEachSeq[S ~[]T, T any](collection S) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range collection {
			if !yield(v) {
				return
			}
		}
	}
}

/*
Unique returns a new slice containing only the unique elements from the input slice,
with optional transformations.
*/
func Unique[S ~[]T, T comparable](collection S, mapping ...Morph[T, T]) S {
	result := make(S, 0, len(collection))
	for uniq := range UniqueSeq(collection, mapping...) {
		result = append(result, uniq)
	}
	return result
}

// UniqueSeq returns an iterator (Seq) that yields unique elements from the input slice,
// with optional transformations.
func UniqueSeq[S ~[]T, T comparable](collection S, mapping ...Morph[T, T]) iter.Seq[T] {
	catched := make(map[T]struct{}, len(collection))
	return func(yield func(T) bool) {
		for i := range collection {
			if _, ok := catched[collection[i]]; ok {
				continue
			}
			catched[collection[i]] = struct{}{}

			ax := collection[i]
			for _, morph := range mapping {
				ax = morph(ax)
			}

			if !yield(ax) {
				return
			}
		}
	}
}

// Count returns the number of elements in the slice that are equal to the specified value.
func Count[S ~[]T, T comparable](collection S, value T) (count int) {
	for i := range collection {
		if collection[i] == value {
			count++
		}
	}
	return
}

/*
CountBy returns the number of elements in the slice that satisfy the given predicate(s).

It uses MorphCond transformations, allowing for operations like:

	CountBy([]Type{...}, func(value Type) bool {...})
*/
func CountBy[S ~[]T, T comparable](collection S, predicates ...MorphCond2[int, T]) (count int) {
	for range FilterSeq(collection, predicates...) {
		count++
	}
	return
}
