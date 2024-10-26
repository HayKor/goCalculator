package parsing

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input       string
		expectValue float64
	}{
		{"5+5", 10},
		{"5+5*2", 15},
		{"(5+5)*2", 20},
		{"(5+5)*2+2", 22},
	}

	for _, tc := range tests {
		lexer := NewLexer(tc.input)
		parser := NewParser(lexer)
		gotValue, err := parser.ParseExpression()
		if err != nil {
			t.Errorf("Got error %v", err.Error())
		}
		if tc.expectValue != gotValue {
			t.Errorf("Expected %v, got %v", tc.expectValue, gotValue)
		}
	}
}
