package token

type TokenType uint8

const (
	OpenBrace TokenType = iota // {
	CloseBrace // }
	OpenBracket // [
	CloseBracket // ]
	Arrow // ->

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
