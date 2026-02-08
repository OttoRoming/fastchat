package parser

import (
	"fmt"
	"math"
	"testing"

	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
	"github.com/stretchr/testify/assert"
)

func TestGarbage(t *testing.T) {
	_, err := Parse("\"hello\"14{}[]->12")
	assert.Error(t, err)
}

func TestString(t *testing.T) {
	el, err := Parse("\"123\"")
	if assert.NoError(t, err) {
		assert.Equal(t, element.String("123"), el)
	}
}

func TestMap(t *testing.T) {
	el, err := Parse(`
		{
			"id" -> 10
			"username" -> "Otto Roming"
			"password" -> "passw0rd"
			"comments" -> ["first" "second"]
			69 -> 420
		}
	`)
	if assert.NoError(t, err) {
		assert.Equal(t, element.Map(map[element.Element]element.Element{
			element.String("id"):       element.Int(10),
			element.String("username"): element.String("Otto Roming"),
			element.String("password"): element.String("passw0rd"),
			element.String("comments"): element.List([]element.Element{
				element.String("first"),
				element.String("second"),
			}),
			element.Int(69): element.Int(420),
		}), el)
	}
}

func TestMaxInt(t *testing.T) {
	el, err := Parse(fmt.Sprint(math.MaxInt64))
	if assert.NoError(t, err) {
		assert.Equal(t, element.Int(math.MaxInt64), el)
	}
}

func TestTrue(t *testing.T) {
	el, err := Parse("t")
	if assert.NoError(t, err) {
		assert.Equal(t, element.Bool(true), el)
	}
}

func TestFalse(t *testing.T) {
	el, err := Parse("f")
	if assert.NoError(t, err) {
		assert.Equal(t, element.Bool(false), el)
	}
}
