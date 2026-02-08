package fcmul

import (
	"fmt"
	"strings"

	"github.com/OttoRoming/fastchat/pkg/fcmul/element"
)

const (
	indentSize = 2
)

func stringifyString(el element.String) string {
	return fmt.Sprintf(`"%s"`, string(el))
}

func stringifyBool(el element.Bool) string {
	if el {
		return "t"
	}
	return "n"
}

func stringifyList(el element.List, pretty bool, indent uint) string {
	var builder strings.Builder

	multiline := pretty && len(el) >= 5

	builder.WriteByte('[')
	if multiline {
		builder.WriteByte('\n')
	}

	for i, value := range el {
		if multiline {
			for range indent * indentSize {
				builder.WriteByte(' ')
			}
		}

		valueString := ""
		if multiline {
			valueString = stringifyElement(value, pretty, indent+1)
		} else {
			valueString = (stringifyElement(value, pretty, indent))
		}
		builder.WriteString(valueString)

		if multiline {
			builder.WriteByte('\n')
		} else if i != len(el)-1 {
			builder.WriteByte(' ')
		}
	}

	if multiline {
		for range (indent - 1) * indentSize {
			builder.WriteByte(' ')
		}
	}
	builder.WriteByte(']')

	return builder.String()
}

func stringifyMap(el element.Map, pretty bool, indent uint) string {
	var builder strings.Builder

	builder.WriteByte('{')
	if pretty {
		builder.WriteByte('\n')
	}

	i := 0
	for key, value := range el {
		if pretty {
			for range indent * indentSize {
				builder.WriteByte(' ')
			}
		}

		builder.WriteString(stringifyElement(key, pretty, indent+1))
		if pretty {
			builder.WriteString(" -> ")
		} else {
			builder.WriteString("->")
		}
		builder.WriteString(stringifyElement(value, pretty, indent+1))

		if i != len(el)-1 {
			if pretty {
				builder.WriteByte('\n')
			} else {
				builder.WriteByte(' ')
			}
		}

		i++
	}

	if pretty {
		builder.WriteByte('\n')
		for range (indent - 1) * indentSize {
			builder.WriteByte(' ')
		}
	}
	builder.WriteByte('}')

	return builder.String()
}

func stringifyElement(el element.Element, pretty bool, indent uint) string {
	switch value := el.(type) {
	case element.Int:
		return fmt.Sprint(value)
	case element.String:
		return stringifyString(value)
	case element.Bool:
		return stringifyBool(value)
	case element.List:
		return stringifyList(value, pretty, indent)
	case element.Map:
		return stringifyMap(value, pretty, indent)
	}

	return "unknown element"
}

func Stringify(el element.Element, pretty bool) string {
	return stringifyElement(el, pretty, 1)
}
