mr
===

`mr` is a Go library based on generics that provides some common collection operation functions and utility functions.

## Installation

Use the go get command to install:

```
go get -u github.com/lyuangg/mr
```

## Usage

```go
import "github.com/lyuangg/mr"
```

## Function List

The mr library provides the following functions:

- `Reduce`: Reduce the elements in a slice.
- `Contains`: Check if a slice contains a specific element.
- `Map`: Map the elements in a slice.
- `ToMap`: Convert a slice to a map.
- `Filter`: Filter the elements in a slice based on a specified condition.
- `Diff`: Get the difference between two slices.
- `Intersect`: Get the intersection of two slices.
- `Unique`: Remove duplicate elements from a slice.
- `Merge`: Merge two slices.

## Examples

Here are some usage examples:

```go
package main

import (
    "fmt"
    "github.com/lyuangg/mr"
)

func main() {
    // Reduce
    numbers := []int{1, 2, 3, 4, 5}
    sum := mr.Reduce(numbers, func(a, b int) int {
        return a + b
    }, 0)
    fmt.Println("Sum:", sum)

    // Contains
    names := []string{"Alice", "Bob", "Charlie"}
    contains := mr.Contains(names, "Bob", func(a string) string { return a })
    fmt.Println("Contains Bob:", contains)

    // Map
    squares := mr.Map(numbers, func(n int) int {
        return n * n
    })
    fmt.Println("Squares:", squares)

    // ToMap
    persons := []Person{
        {ID: 1, Name: "Alice"},
        {ID: 2, Name: "Bob"},
    }
    personsMap := mr.ToMap(persons, func(p Person) int {
        return p.ID
    })
    fmt.Println("Persons Map:", personsMap)

    // Filter
    adults := mr.Filter(persons, func(p Person) bool {
        return p.Age >= 18
    })
    fmt.Println("Adults:", adults)
}

type Person struct {
    ID   int
    Name string
    Age  int
}
```

## License

This library is released under the MIT License. Please see the license file for more information.