package parser

import (
	"fmt"
	"strconv"

	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
	"github.com/OttoRoming/fastchat/pkg/fcmul/lexer"
	"github.com/OttoRoming/fastchat/pkg/fcmul/token"
)

type parser struct {
	tokens []token.Token
	position int
}

func newParser(input string) (parser, error) {
	tokens, err := lexer.Lex(input)
	if err != nil {
		return parser{}, err
	}

	return parser {
		tokens: tokens,
		position: 0,
	}, nil
}

func (l *parser)current() token.Token {
	// avoid out of bounds with a white lie
	if l.isDone() {
		return token.New(token.EndOfFile, "EOF")
	}
	return l.tokens[l.position]
}

func (l *parser)advance() {
	l.position ++
}

func (l *parser)isDone() bool {
	return l.position >= len(l.tokens)
}

func (l *parser)parseMap() (element.Map, error) {
	result := element.Map(make(map[element.Element]element.Element))

	if l.current().Kind != token.OpenBrace {
		return result, fmt.Errorf("expected open brace at start of map, found %s", l.current())
	}
	l.advance()

	for l.current().Kind != token.CloseBrace {
		key, err := l.parseElement()
		if err != nil {
			return result, err
		}

		if l.current().Kind != token.Arrow {
			return result, fmt.Errorf("expected arrow after key in map, found %s", l.current())
		}
		l.advance()

		value, err := l.parseElement()
		if err != nil {
			return result, err
		}
		result[key] = value
	}
	l.advance() // skip }

	return result, nil
}

func (l *parser)parseList() (element.List, error) {
	result := element.List([]element.Element{})

	if l.current().Kind != token.OpenBracket {
		return result, fmt.Errorf("expected open bracket at start of list, found %s", l.current())
	}
	l.advance()

	for l.current().Kind != token.CloseBracket {
		value, err := l.parseElement()
		if err != nil {
			return result, err
		}

		result = append(result, value)
	}
	l.advance() // skip ]

	return result, nil
}

func (l *parser)parseString() (element.String, error) {
	result := element.String(l.current().Literal)

	if l.current().Kind != token.String {
		return result, fmt.Errorf("expected string token at string element, found %s", l.current())
	}
	l.advance()

	return result, nil
}

func (l *parser)parseInt() (element.Int, error) {
	if l.current().Kind != token.Int {
		return element.Int(-1), fmt.Errorf("expected int token at int element, found %s", l.current())
	}

	value, err := strconv.ParseInt(l.current().Literal, 10, 64)
	if err != nil {
		return element.Int(-1), fmt.Errorf("int token could not be parsed %s", err)
	}
	l.advance()

	return element.Int(value), nil
}

func (l *parser)parseBool() (element.Bool, error) {
	result := false

	if l.current().Kind != token.True && l.current().Kind != token.False {
		return element.Bool(result), fmt.Errorf("expected true or false token at bool element, found %s", l.current())
	}

	if l.current().Kind == token.True {
		result = true
	}
	l.advance()

	return element.Bool(result), nil
}

func (l *parser)parseElement() (element.Element, error) {
	switch l.current().Kind {
		case token.OpenBrace:
			return l.parseMap()
		case token.OpenBracket:
			return l.parseList()
		case token.String:
			return l.parseString()
		case token.Int:
			return l.parseInt()
		case token.True, token.False:
			return l.parseBool()
		case token.EndOfFile:
			return element.Bool(false), fmt.Errorf("unexpected end of file")
		default:
			return element.Bool(false), fmt.Errorf("unexpected token %s", l.current())
	}
}

func Parse(source string) (element.Element, error) {
	parser, err := newParser(source)
	if err != nil {
		return element.Bool(false), err
	}

	el, err := parser.parseElement()
	if err != nil {
		return el, err
	}
	if !parser.isDone() {
		return el, fmt.Errorf("garbage tokens found after the first element of mcul source")
	}

	return el, nil
}
