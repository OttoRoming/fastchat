package fcmul

import (
	"testing"

	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
	"github.com/OttoRoming/fastchat/pkg/fcmul/parser"
	"github.com/stretchr/testify/assert"
)

func TestStringifyString(t *testing.T) {
	str := Stringify(element.String("123"), false)
	assert.Equal(t, `"123"`, str)
}

func TestStringifyStringPretty(t *testing.T) {
	str := Stringify(element.String("123"), true)
	assert.Equal(t, `"123"`, str)
}

func TestStringifyMap(t *testing.T) {
	// The way a map will get stringified is undefined and can thus not be tested directly compared to a string

	el := element.Map{
		element.String("id"):       element.Int(10),
		element.String("username"): element.String("Otto Roming"),
		element.String("password"): element.String("passw0rd"),
		element.String("comments"): element.List{
			element.String("first"),
			element.String("second"),
		},
		element.Int(69): element.Int(420),
	}

	str := Stringify(el, false)
	el2, err := parser.Parse(str)

	if assert.NoError(t, err) {
		assert.Equal(t, el, el2)
	}
}

func TestStringifyMapPretty(t *testing.T) {
	// The way a map will get stringified is undefined and can thus not be tested directly compared to a string

	el := element.Map{
		element.String("id"):       element.Int(10),
		element.String("username"): element.String("Otto Roming"),
		element.String("password"): element.String("passw0rd"),
		element.String("comments"): element.List{
			element.String("first"),
			element.String("second"),
		},
		element.Int(69): element.Int(420),
	}

	str := Stringify(el, true)
	el2, err := parser.Parse(str)

	if assert.NoError(t, err) {
		assert.Equal(t, el, el2)
	}
}
