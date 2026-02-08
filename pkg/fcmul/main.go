/*
Fast Chat Mark Up Language

Provides a serializer and deserializer for the fcmul markup language
*/

package fcmul

import (
	"github.com/OttoRoming/fastchat/pkg/fcmul/parser"
	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
)

func Parse(source string) (element.Element, error) {
	return parser.Parse(source)
}
