# goCalculator

Calculator which can evaluate basic mathematical expressions.
It has lexer and parser.

## installation

```shell
go install github.com/HayKor/goCalculator
```

## Usage

```go
package main

import (
	"fmt"

	calc "github.com/HayKor/goCalculator/pkg/calculator"
)

func main() {
	input := "5 + 3 * (2 - 1)"
	result, err := calc.Calc(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // Output: 8
}
```
