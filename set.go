/*
Package set defines a set type, a mutable, unordered collection of unique elements,
analogous to Python's set collection, as well as functions to perform typical mathematical set operations.

The underlying type	of this set implementation is a map.
This means that all operations that can be done with maps can also be done with sets.
*/
package set

import (
	"maps"
	"slices"
)

// Set is a map where the key type is generic
// and all values are empty structs that take up no space in memory.
type Set[T comparable] map[T]struct{}

// NewSet takes a variable number of arguments and returns a set containing those values.
func NewSet[T comparable](items ...T) Set[T] {
	set := make(Set[T])

	for _, item := range items {
		set[item] = struct{}{}
	}

	return set
}

// Contains returns true if the given value is found in the set.
// Otherwise, returns false.
func (set *Set[T]) Contains(item T) bool {
	_, ok := (*set)[item]
	return ok
}

// Elements returns a slice containing all members of the set.
func (set *Set[T]) Elements() []T {
	return slices.Collect(maps.Keys(*set))
}

// Insert takes a value and adds it to the set.
func (set *Set[T]) Insert(item T) {
	(*set)[item] = struct{}{}
}

// Remove deletes a given value from the set, if present.
func (set *Set[T]) Remove(item T) {
	delete(*set, item)
}

// Union returns a new set containing all elements that are in set A and set B, i.e., A ∪ B.
func (setA *Set[T]) Union(setB Set[T]) Set[T] {
	setC := make(Set[T])

	maps.Copy(setC, *setA)
	maps.Copy(setC, setB)

	return setC
}
