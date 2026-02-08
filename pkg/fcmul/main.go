/*
Fast Chat Mark Up Language

Provides a serializer and deserializer for the fcmul markup language
*/

package fcmul

import (
	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
	"github.com/OttoRoming/fastchat/pkg/fcmul/parser"
)

func Parse(source string) (element.Element, error) {
	return parser.Parse(source)
}
