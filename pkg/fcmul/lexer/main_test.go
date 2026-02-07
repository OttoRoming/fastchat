package lexer

import (
	"testing"
	"github.com/OttoRoming/fastchat/pkg/fcmul/token"
    "github.com/stretchr/testify/assert"
)

func TestLex(t *testing.T) {
	tokens, err := Lex("\"hello\"14{}[]->12")
	if assert.NoError(t, err) {
		assert.Equal(t, []token.Token{
			token.New(token.String, "hello"),
			token.New(token.Int, "14"),
			token.New(token.OpenBrace, "{"),
			token.New(token.CloseBrace, "}"),
			token.New(token.OpenBracket, "["),
			token.New(token.CloseBracket, "]"),
			token.New(token.Arrow, "->"),
			token.New(token.Int, "12"),
		}, tokens)
	}
}
