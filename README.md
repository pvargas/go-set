# Go Set

[![](https://godoc.org/github.compvargas/go-set?status.svg)](http://godoc.org/github.com/pvargas/go-set)
[![Build](https://github.com/pvargas/go-set/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/pvargas/go-set/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pvargas/go-set)](https://goreportcard.com/report/github.com/pvargas/go-set)

[Set](https://en.wikipedia.org/wiki/Set_(abstract_data_type)) type for the Go language.

## Features 

### Basic Operations

| Method   | Description                 |
|----------|-----------------------------|
| NewSet   | Takes a variable number of arguments and returns a set containing those values |
| Insert   | Takes a value and adds it to the set                                           |
| Contains | Returns true if the given value is found in the set                            |
| Elements | Returns a slice containing all members of the set                              |
| Remove   | Deletes a given value from the set                                             |

### Mathematical Set Operations

| Method         | Description            |
|----------------|------------------------|
| Union               | Returns a set containing all elements from set A and set B                   |
| Difference          | Returns a set containing only the elements from set A that are not in set B  |
| SymmetricDifference | Returns a set containing only the elements from set A that are not in set B  |
| Intersection        | Returns a set containing only the elements from set A that are also in set B |
| IsSubset            | Returns `true` if set A is a subset of set B                                 |
| IsProperSubset      | Returns `true` if set A is a proper subset of set B                          |
| IsDisjoint          | Returns `true` if the two sets do not intersect                              |


## Installation

```sh
go get github.com/pvargas/go-set
```

## Usage

This set implementation supports Go hashable [comparable types](https://go.dev/ref/spec#Comparison_operators). Non-comparable types are not supported and cannot be used as set elements. 

### Creating a New Set

To create a new set, use the `NewSet` method. `NewSet` is variadic, so any number of arguments can be provided. If no arguments are provided, the desired set type must be specified. The resulting set will not contain any repeated values.

```go
package example

import "github.com/pvargas/go-set"

// Empty set of runes
runeSet := set.NewSet[rune]()

// Heterogenous set
numbers := set.NewSet[any](2, "three", 5.0, '7', 11i)

names := []string{"Alice", "Bob", "Bob", "Carol"}

// Set of three elements (Alice, Bob, and Carol) 
// created from a slice of four elements. 
uniqueNames := set.NewSet(names...)

```