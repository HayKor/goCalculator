package main

import (
	"fmt"

	. "github.com/HayKor/goCalculator/pkg/calculator"
)

func main() {
	input := "5 + 3 * (2 - 1)"
	result, err := Calc(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // Output: 8
}
