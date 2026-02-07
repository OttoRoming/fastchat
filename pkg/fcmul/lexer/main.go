package lexer

import (
	"github.com/OttoRoming/fastchat/pkg/fcmul/token"
	"fmt"
	"strings"
)

type lexer struct {
	source string
	position int
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func newLexer(input string) lexer {
	return lexer {
		source: input,
		position: 0,
	}
}

func (l *lexer)current() byte {
	// avoid out of bounds with a white lie
	if l.isDone() {
		return ' '
	}
	return l.source[l.position]
}

func (l *lexer)advance() {
	l.position ++
}

func (l *lexer)isDone() bool {
	return l.position >= len(l.source)
}

func (l *lexer)nextToken() (token.Token, error) {
	var tok token.Token

	switch l.current() {
	// ignore whitespace
	case ' ', '\t', '\r', '\n':
		l.advance()
	// single char tokens
	case '{':
		tok = token.New(token.OpenBrace, "{")
		l.advance()
	case '}':
		tok = token.New(token.CloseBrace, "}")
		l.advance()
	case '[':
		tok = token.New(token.OpenBracket, "[")
		l.advance()
	case ']':
		tok = token.New(token.CloseBracket, "]")
		l.advance()
	// 2 char tokens
	case '-':
		l.advance()
		if l.current() == '>' {
			tok = token.New(token.Arrow, "->")
			l.advance()
		} else {
			return tok, fmt.Errorf("unexpected character after dash '%c'", l.current())
		}
	// special tokens
	case '"':
		l.advance()
		var builder strings.Builder
		for l.current() != '"' {
			builder.WriteByte(l.current())
			l.advance()
		}
		l.advance()
		tok = token.New(token.String, builder.String())
	default:
		if isDigit(l.current()) {
			var builder strings.Builder
			for isDigit(l.current()) {
				builder.WriteByte(l.current())
				l.advance()
			}
			tok = token.New(token.Int, builder.String())
		} else {
			return tok, fmt.Errorf("unexpected character '%c'", l.current())
		}
	}

	return tok, nil
}

func Lex(source string) ([]token.Token, error) {
	var tokens []token.Token

	lexer := newLexer(source)
	for !lexer.isDone() {
		fmt.Printf("lexer.position: %v\n", lexer.position)
		tok, err := lexer.nextToken()
		if err != nil {
			return tokens, err
		}

		tokens = append(tokens, tok)
	}

	return tokens, nil
}
