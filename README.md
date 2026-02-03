# Go Set

Homogeneous [set](https://en.wikipedia.org/wiki/Set_(abstract_data_type)) type for the Go language.

## Installation

```sh
go get github.com/pvargas/go-set
```

## Usage

This set implementation supports Go [comparable types](https://go.dev/ref/spec#Comparison_operators). Non-comparable types are not supported. 


### Creating a New Set

There are three ways to create a set. The first way is by using the `NewSet` method. 

`NewSet` is variadic, so any number of arguments can be provided. If no arguments are provided, the desired set type must be specified.

```go
// Empty set of strings
words := set.NewSet[string]()

// Set of integers
primes := set.NewSet(2, 3, 5, 7, 11, 13)
```

The second way to create a set is by using the `ArrayToSet` method. The resulting set will not contain any repeated elements found in the array.

```go
names := [...]string{"Alice", "Bob", "Bob", "Carol"}

nameSet := set.ArrayToSet(names)
fmt.Println(names)
```
`[Alice Bob Carol]`


The last way to create a set is by using Go's `make` function.

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
