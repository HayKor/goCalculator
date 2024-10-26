package main

import (
	"fmt"

	. "github.com/HayKor/goCalculator/pkg/parsing"
)

func main() {
	input := "5 + 3 * (2 - 1)"
	parser := NewParser(NewLexer(input))
	result, err := parser.ParseExpression()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // Output: 8
}
