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

func unmarshalString(el element.Element, vv reflect.Value) error {
	if vv.Kind() != reflect.String {
		return fmt.Errorf("go value is not of kind string")
	}

	switch s := el.(type) {
	case element.String:
		vv.SetString(string(s))
	default:
		return fmt.Errorf("fcmul element is not of type string")
	}
	return nil
}

func unmarshalInt(el element.Element, vv reflect.Value) error {
	switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			 reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		default:
			return fmt.Errorf("go value is not of an int kind")
	}

	value := 0
	switch i := el.(type) {
	case element.Int:
		value = int(i)
	default:
		return fmt.Errorf("fcmul element is not of type int")
	}

	switch vv.Kind() {
	case reflect.Int:
		vv.SetInt(int64(value))
	}

	return nil
}

func unmarshalStruct(el element.Element, vv reflect.Value) error {
	if vv.Kind() != reflect.Struct{
		return fmt.Errorf("go value is not of kind struct")
	}

	var elMap element.Map
	switch m := el.(type) {
	case element.Map:
		elMap = m
	default:
		return fmt.Errorf("fcmul element is not of type map")
	}

	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		typeField := vt.Field(i)
		field := vv.Field(i)

		valueEl, found := elMap[element.String(typeField.Name)]
		if !found {
			return fmt.Errorf("fcmul is missing field '%s'", typeField.Name)
		}

		err := unmarshalElement(valueEl, field)
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalElement(el element.Element, vv reflect.Value) error {
	switch vv.Kind() {
		case reflect.String:
			return unmarshalString(el, vv)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			 reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return unmarshalInt(el, vv)
		case reflect.Struct:
			return unmarshalStruct(el, vv)
		case reflect.Map:
			return fmt.Errorf("not implemented")
		case reflect.Array, reflect.Slice:
			return fmt.Errorf("not implemented")
		default:
			return fmt.Errorf("unsupported value type")
	}
}


func Unmarshal(source string, v any) error {
	el, err := Parse(source)
	if err != nil {
		return err
	}

	vvp := reflect.ValueOf(v)
	if vvp.Kind() != reflect.Pointer {
		return fmt.Errorf("v must be a pointer")
	}

	vv := vvp.Elem()

	return unmarshalElement(el, vv)
}
