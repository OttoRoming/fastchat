package parser

import (
	"testing"

	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
	"github.com/stretchr/testify/assert"
)

func TestParseGarbage(t *testing.T) {
	_, err := Parse("\"hello\"14{}[]->12")
	assert.Error(t, err)
}

func TestParseString(t *testing.T) {
	el, err := Parse("\"123\"")
	if assert.NoError(t, err) {
		assert.Equal(t, element.String("123"), el)
	}
}

func TestParseMap(t *testing.T){
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
