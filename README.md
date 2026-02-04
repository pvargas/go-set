# Go Set

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
primes := set.NewSet[any](2, "three", 5.0, '7', 11i)

names := []string{"Alice", "Bob", "Bob", "Carol"}

// Set of three elements (Alice, Bob, and Carol) created from a slice of 4 elements. 
nameSet := set.NewSet(names...)

```

The second way to create a set is by using Go's `make` function.

```go
someSet := make(set.Set[string])
```

### Complete Example

```go
package example

import	"github.com/pvargas/go-set"


func (hp *htmlParser) domains(urls []string) set.Set[string] {
	domainSet := set.NewSet[string]() // or make(set.Set[string])

	for url := range urls {
		domain, err := domainFromUrl(url)

		if err != nil {
    	domainSet.Insert(domain)
		}
	}

	return domainSet
}
```
