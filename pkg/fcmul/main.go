/*
Fast Chat Mark Up Language

Provides a serializer and deserializer for the fcmul markup language
*/

package fcmul

import (
	"reflect"
	"fmt"
	"github.com/OttoRoming/fastchat/pkg/fcmul/parser"
	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
)

func Parse(source string) (element.Element, error) {
	return parser.Parse(source)
}

func Unmarshal(source string, v any) {
	vValue := reflect.ValueOf(v)
	vType := vValue.Type()

	fmt.Println("arst:")
	for i := 0; i < vValue.NumField(); i++ {
		field := vType.Field(i)

		fmt.Printf("field.Name: %v\n", field.Name)
		if field.Type.Kind() == reflect.String {
			fmt.Printf("field.Type: %v\n", field.Type)
		}
		fmt.Printf("field: %v\n", field)
	}
}
