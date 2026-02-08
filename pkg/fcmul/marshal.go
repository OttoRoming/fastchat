package fcmul

import (
	"fmt"
	"reflect"

	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
)

func marshalStruct(v reflect.Value) (element.Map, error) {
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("go value is not of kind struct")
	}

	vt := v.Type()
	elMap := make(element.Map)
	for i := 0; i < v.NumField(); i++ {
		typeField := vt.Field(i)
		field := v.Field(i)

		elKey := element.String(typeField.Name)
		elValue, err := marshalElement(field)
		if err != nil {
			return nil, err
		}

		elMap[elKey] = elValue
	}

	return elMap, nil
}

func marshalMap(v reflect.Value) (element.Map, error) {
	if v.Kind() != reflect.Map {
		return nil, fmt.Errorf("go value is not of kind map")
	}

	elMap := make(element.Map)
	iter := v.MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		elKey, err := marshalElement(key)
		if err != nil {
			return nil, err
		}

		elValue, err := marshalElement(value)
		if err != nil {
			return nil, err
		}

		elMap[elKey] = elValue
	}

	return elMap, nil
}

func marshalSliceArray(v reflect.Value) (element.List, error) {
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, fmt.Errorf("go value is not of kind slice or array")
	}

	var list element.List

	for i := range v.Len() {
		value, err := marshalElement(v.Index(i))
		if err != nil {
			return list, err
		}

		list = append(list, value)
	}

	return list, nil
}

func marshalElement(v reflect.Value) (element.Element, error) {
	switch v.Kind() {
	case reflect.String:
		return element.String(v.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return element.Int(v.Int()), nil
	case reflect.Bool:
		return element.Bool(v.Bool()), nil
	case reflect.Struct:
		return marshalStruct(v)
	case reflect.Map:
		return marshalMap(v)
	case reflect.Slice, reflect.Array:
		return marshalSliceArray(v)
	default:
		return element.Bool(false), fmt.Errorf("unsupported value kind %s", v.Kind())
	}

}

func Marshal(v any) (string, error) {
	vv := reflect.ValueOf(v)

	// if it is a pointer just deref for convenience
	if vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}

	el, err := marshalElement(vv)
	if err != nil {
		return "", err
	}

	str := Stringify(el, false)
	return str, nil
}

func MarshalPretty(v any) (string, error) {
	vv := reflect.ValueOf(v)

	// if it is a pointer just deref for convenience
	if vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}

	el, err := marshalElement(vv)
	if err != nil {
		return "", err
	}

	str := Stringify(el, true)
	return str, nil
}
