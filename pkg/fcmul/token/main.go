package token

import "fmt"

type TokenType uint8

const (
	OpenBrace TokenType = iota // {
	CloseBrace // }
	OpenBracket // [
	CloseBracket // ]
	True // t
	False // f

	Arrow // ->

	EndOfFile // EOF

	String // "Hello, World"
	Int // 14
)

type Token struct {
	Kind TokenType
	Literal string
}

func New(kind TokenType, literal string) Token {
	return Token {
		Kind: kind,
		Literal: literal,
	}
}

func (t Token)String() string {
	switch t.Kind {
	case OpenBrace:
		return fmt.Sprintf("OpenBrace(%s)", t.Literal)
	case CloseBrace:
		return fmt.Sprintf("CloseBrace(%s)", t.Literal)
	case OpenBracket:
		return fmt.Sprintf("OpenBracket(%s)", t.Literal)
	case CloseBracket:
		return fmt.Sprintf("CloseBracket(%s)", t.Literal)
	case True:
		return fmt.Sprintf("True(%s)", t.Literal)
	case False:
		return fmt.Sprintf("False(%s)", t.Literal)
	case Arrow:
		return fmt.Sprintf("Arrow(%s)", t.Literal)
	case EndOfFile:
		return "EndOfFile"
	case String:
		return fmt.Sprintf("String(%s)", t.Literal)
	case Int:
		return fmt.Sprintf("Int(%s)", t.Literal)
	default:
		return "Unknown"
	}
}
