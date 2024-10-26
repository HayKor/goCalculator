package calculator

import (
	. "github.com/HayKor/goCalculator/internal/parsing"
)

// Реализовать функцию func Calc(expression string) (float64, error)
// expression - строка-выражение состоящее из односимвольных идентификаторов и знаков арифметических действий
// Входящие данные - цифры(рациональные), операции +, -, *, /,
// операции приоритезации ( и ) В случае ошибки записи выражения функция выдает ошибку.

func Calc(expression string) (float64, error) {
	lexer := NewLexer(expression)
	parser := NewParser(lexer)

	result, err := parser.ParseExpression()
	if err != nil {
		return 0, err
	}
	return result, nil
}
