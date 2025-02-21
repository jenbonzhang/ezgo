# ezgo Package

The `ezgo` package provides a collection of utility functions to simplify common tasks in Go programming. It includes
various sub-packages that offer specific functionalities to enhance productivity and code readability.

## Installation

To install the `ezgo` package, run the following Go command:

```bash
go get github.com/jenbonzhang/ezgo
```

## Sub-packages

### ezgo/maps

 The ezgo/maps package provides utility functions for handling nested maps in Go. It simplifies working with multi-level maps by providing functions to set values at any depth while automatically initializing missing inner maps.

#### SetMapValue

Sets a value in a two-level nested map. If the inner map does not exist, it will be initialized automatically.

```go
package main

import (
	"fmt"
	ezm "github.com/jenbonzhang/ezgo/maps"
)

func main() {
	// Creating a two-level nested map
	nestedMap := make(map[string]map[int]int)

	// Setting a value
	ezm.SetMapValue(nestedMap, "firstKey", 100, 42)

	// Output the result
	fmt.Println(nestedMap["firstKey"][100]) // Output: 42
}

```

#### SetMapValue3

Sets a value in a three-level nested map with different key types at each level.

```go
package main

import (
	"fmt"
	ezm "github.com/jenbonzhang/ezgo/maps"
)

func main() {
	// Creating a three-level nested map
	nestedMap3 := make(map[string]map[int]map[float64]string)

	// Setting a value
	ezm.SetMapValue3(nestedMap3, "firstKey", 100, 1.23, "Value A")

	// Output the result
	fmt.Println(nestedMap3["firstKey"][100][1.23]) // Output: Value A
}
```
