package lexer

import (
	"testing"
	"math"
	"fmt"
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

func TestTrailingWhitespace(t *testing.T) {
	tokens, err := Lex("  10  ")
	if assert.NoError(t, err) {
		assert.Equal(t, []token.Token{
			token.New(token.Int, "10"),
		}, tokens)
	}
}

func TestMap(t *testing.T) {
	tokens, err := Lex(`
		{
			"id" -> 10
			"username" -> "Otto Roming"
			"password" -> "passw0rd"
			"comments" -> ["first" "second"]
			69 -> 420
		}
	`)
	if assert.NoError(t, err) {
		assert.Equal(t, []token.Token{
			token.New(token.OpenBrace, "{"),
			token.New(token.String, "id"),
			token.New(token.Arrow, "->"),
			token.New(token.Int, "10"),
			token.New(token.String, "username"),
			token.New(token.Arrow, "->"),
			token.New(token.String, "Otto Roming"),
			token.New(token.String, "password"),
			token.New(token.Arrow, "->"),
			token.New(token.String, "passw0rd"),
			token.New(token.String, "comments"),
			token.New(token.Arrow, "->"),
			token.New(token.OpenBracket, "["),
			token.New(token.String, "first"),
			token.New(token.String, "second"),
			token.New(token.CloseBracket, "]"),
			token.New(token.Int, "69"),
			token.New(token.Arrow, "->"),
			token.New(token.Int, "420"),
			token.New(token.CloseBrace, "}"),
		}, tokens)
	}
}

func TestMaxInt(t *testing.T) {
	tokens, err := Lex(fmt.Sprint(math.MaxInt64))
	if assert.NoError(t, err) {
		assert.Equal(t, []token.Token{token.New(token.Int, fmt.Sprint(math.MaxInt64))}, tokens)
	}
}

func TestTrue(t *testing.T) {
	tokens, err := Lex("t")
	if assert.NoError(t, err) {
		assert.Equal(t, []token.Token{token.New(token.True, "t")}, tokens)
	}
}

func TestFalse(t *testing.T) {
	tokens, err := Lex("f")
	if assert.NoError(t, err) {
		assert.Equal(t, []token.Token{token.New(token.False, "f")}, tokens)
	}
}
