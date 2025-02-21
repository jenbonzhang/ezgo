# ezgo/maps Package

The `ezgo/maps` package provides utility functions for handling nested maps in Go. It simplifies working with multi-level maps by providing functions to set values at any depth while automatically initializing missing inner maps.

## Installation

To install the `ezgo/maps` package, run the following Go command:

```bash
go get github.com/jenbonzhang/ezgo/maps
```
## Usage
Here's a simple example of how to use the `ezgo/maps` package:

#### SetNestedMapValue
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
	ezm.SetNestedMapValue(nestedMap, "firstKey", 100, 42)

	// Output the result
	fmt.Println(nestedMap["firstKey"][100]) // Output: 42
}
```
### SetNestedMapValue3
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
	ezm.SetNestedMapValue3(nestedMap3, "firstKey", 100, 1.23, "Value A")

	// Output the result
	fmt.Println(nestedMap3["firstKey"][100][1.23]) // Output: Value A
}
```