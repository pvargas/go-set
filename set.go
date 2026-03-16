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

// Remove deletes a given value from the set.
// If the value is not present no operation is performed.
func (set *Set[T]) Remove(item T) {
	delete(*set, item)
}

// Union returns a set containing all elements from set A and set B.
func (setA *Set[T]) Union(setB Set[T]) Set[T] {
	setC := make(Set[T])

	maps.Copy(setC, *setA)
	maps.Copy(setC, setB)

	return setC
}

// Difference returns a set containing all elements from set A that are not members of set B.
func (setA *Set[T]) Difference(setB Set[T]) Set[T] {
	setC := make(Set[T])

	for element := range *setA {
		if !setB.Contains(element) {
			setC.Insert(element)
		}
	}

	return setC
}

// SymmetricDifference returns a set containing elements which are in either of the sets but not in their intersection.
func (setA *Set[T]) SymmetricDifference(setB Set[T]) Set[T] {
	setC := make(Set[T])

	for element := range *setA {
		if !setB.Contains(element) {
			setC.Insert(element)
		}
	}

	for element := range setB {
		if !setA.Contains(element) {
			setC.Insert(element)
		}
	}

	return setC
}

// Intersection returns a set containing only the elements from set A that are also elements of set B.
func (setA *Set[T]) Intersection(setB Set[T]) Set[T] {
	setC := make(Set[T])

	for element := range *setA {
		if setB.Contains(element) {
			setC.Insert(element)
		}
	}

	return setC
}

// IsSubset returns true if set A is a subset of set B.
func (setA *Set[T]) IsSubset(setB Set[T]) bool {
	for element := range *setA {
		if !setB.Contains(element) {
			return false
		}
	}

	return true
}

// IsProperSubset returns true if set A is a proper subset of set B.
func (setA *Set[T]) IsProperSubset(setB Set[T]) bool {
	return setA.IsSubset(setB) && len(*setA) < len(setB)
}

// IsDisjoint returns true if the two sets do not intersect.
func (setA *Set[T]) IsDisjoint(setB Set[T]) bool {
	return len(setA.Intersection(setB)) == 0
}
