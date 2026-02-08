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

	var elInt element.Int
	switch i := el.(type) {
	case element.Int:
		elInt = i
	default:
		return fmt.Errorf("fcmul element is not of type int")
	}

	switch vv.Kind() {
	case reflect.Int:
		vv.SetInt(int64(elInt))
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

func unmarshalMap(el element.Element, vv reflect.Value) error {
	if vv.Kind() != reflect.Map {
		return fmt.Errorf("go value is not of kind map")
	}

	vt := vv.Type()

	// the map might be nil or might be filled with old stuff, remove that
	vv.Set(reflect.MakeMap(vt))

	var elMap element.Map
	switch m := el.(type) {
	case element.Map:
		elMap = m
	default:
		return fmt.Errorf("fcmul element is not of type map")
	}

	for elKey, elValue  := range elMap {
		key := reflect.New(vt.Key()).Elem()
		err := unmarshalElement(elKey, key)
		if err != nil {
			return err
		}

		value := reflect.New(vt.Elem()).Elem()
		err = unmarshalElement(elValue, value)
		if err != nil {
			return err
		}

		vv.SetMapIndex(key, value)
	}

	return nil
}

func unmarshalSlice(el element.Element, vv reflect.Value) error {
	if vv.Kind() != reflect.Slice {
		return fmt.Errorf("go value is not of kind slice")
	}
	vt := vv.Type()

	var elList element.List
	switch l := el.(type) {
	case element.List:
		elList = l
	default:
		return fmt.Errorf("fcmul element is not of type list")
	}

	length := len(elList)
	vv.Set(reflect.MakeSlice(vt, length, length))
	for i := range length {
		err := unmarshalElement(elList[i], vv.Index(i))
		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalArray(el element.Element, vv reflect.Value) error {
	if vv.Kind() != reflect.Array {
		return fmt.Errorf("go value is not of kind array")
	}

	var elList element.List
	switch l := el.(type) {
	case element.List:
		elList = l
	default:
		return fmt.Errorf("fcmul element is not of type list")
	}

	length := len(elList)

	if vv.Len() != length {
		return fmt.Errorf("fcmul list length is different from go array length (fcmul: %d, go: %d)", len(elList), vv.Len())
	}

	for i := range length {
		err := unmarshalElement(elList[i], vv.Index(i))
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
			return unmarshalMap(el, vv)
		case reflect.Slice:
			return unmarshalSlice(el, vv)
		case reflect.Array:
			return unmarshalArray(el, vv)
		default:
			return fmt.Errorf("unsupported value type %s", vv.Kind())
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
