package set

type Set[T comparable] map[T]struct{}

// NewSet takes a variable number of arguments and returns a set containing those values.
func NewSet[T comparable](items ...T) Set[T] {
	set := make(Set[T])

	for _, item := range items {
		set[item] = struct{}{}
	}

	return set
}

// ArrayToSet takes an array and returns a set containing the array's elements.
func ArrayToSet[T comparable](items []T) Set[T] {
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

// Insert takes a value and inserts it into the set.
func (set *Set[T]) Insert(item T) {
	(*set)[item] = struct{}{}
}

// Remove deletes a given value from the set, if present.
func (set *Set[T]) Remove(item T) {
	delete(*set, item)
}
