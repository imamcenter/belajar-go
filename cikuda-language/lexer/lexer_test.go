package lexer

import (
	"cikuda-language/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = def(x, y) {
	x + y;
	};
	let result = add(five, ten);`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "def"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)
	fmt.Printf("l.ch: %v\n", string(l.ch))

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Printf("tok: %v\n", tok)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestStr(t *testing.T) {
	// var val int
	var name = "=+(){},;"
	var ch byte = name[0]
	fmt.Println(len(name))
	fmt.Printf("ch: %v\n", ch)

	var ln int32 = 61
	var d byte = '='
	fmt.Println(string(ln))
	// fmt.Println('=' == ln)
	fmt.Printf("d: %v\n", d)
	switch ln {
	case '=':
		fmt.Println("true")
	case ')':
		fmt.Println("false")
	}
}
