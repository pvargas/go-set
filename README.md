# Go Set

[![Build](https://github.com/pvargas/go-set/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/pvargas/go-set/actions/workflows/ci.yml)

[Set](https://en.wikipedia.org/wiki/Set_(abstract_data_type)) type for the Go language.

## Installation

```sh
go get github.com/pvargas/go-set
```

## Usage

This set implementation supports Go hashable [comparable types](https://go.dev/ref/spec#Comparison_operators). Non-comparable types are not supported and cannot be used as set elements. 


### Creating a New Set

There are two ways to create a set. The first way is by using the `NewSet` method. 

`NewSet` is variadic, so any number of arguments can be provided. If no arguments are provided, the desired set type must be specified. The resulting set will not contain any repeated argument values.

```go
package example

import "github.com/pvargas/go-set"

// Empty set of runes
words := set.NewSet[rune]()

// Heterogenous set
numbers := set.NewSet[any](2, "three", 5.0, '7', 11i)

names := []string{"Alice", "Bob", "Bob", "Carol"}

// Set of three elements (Alice, Bob, and Carol) 
// created from a slice of four elements. 
nameSet := set.NewSet(names...)

```

The second way to create a set is by using Go's `make` function.

```go
someSet := make(set.Set[string])
```

### Basic Operations

| Method   | Description                                                                    |
|----------|--------------------------------------------------------------------------------|
| Contains | Returns true if the given value is found in the set                            |
| Elements | Returns a slice containing all members of the set                              |
| Insert   | Takes a value and adds it to the set                                           |
| NewSet   | Takes a variable number of arguments and returns a set containing those values |
| Remove   | Deletes a given value from the set                                             |

### Mathematical Set Operations

| Method | Description                                                       |
|--------|-------------------------------------------------------------------|
| Union  | Returns a set containing all elements that are in set A and set B |
