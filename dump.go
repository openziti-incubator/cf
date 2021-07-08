package cf

import (
	"fmt"
	"reflect"
)

func Dump(indent int, cf interface{}, opt *Options) string {
	cfV := reflect.ValueOf(cf)
	if cfV.Kind() == reflect.Ptr {
		cfV = cfV.Elem()
	}
	if cfV.Kind() != reflect.Struct {
		return ""
	}
	out := "{\n"
	format := fmt.Sprintf("%s%%-%ds %%v\n", nLevels(indent), maxKeyLength(cfV, opt))
	for i := 0; i < cfV.NumField(); i++ {
		if cfV.Field(i).CanInterface() {
			fd := parseFieldData(cfV.Type().Field(i), opt)
			fieldType := cfV.Type().Field(i).Type
			if fieldType.Kind() == reflect.Slice {
				for j := 0; j < cfV.Field(i).Len(); j++ {
					if j > 0 {
						out += nLevels(indent)
					}
					out += "{"

				}

			} else if fieldType.Kind() == reflect.Struct || (fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct) {
				out += fmt.Sprintf(format, fd.name, "struct")
			} else {
				out += fmt.Sprintf(format, fd.name, cfV.Field(i).Interface())
			}
		}
	}
	out += "}\n"
	return out
}

func nLevels(n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += "\t"
	}
}

func maxKeyLength(cfV reflect.Value, opt *Options) int {
	maxKeyLength := 0
	for i := 0; i < cfV.NumField(); i++ {
		fd := parseFieldData(cfV.Type().Field(i), opt)
		keyLength := len(fd.name)
		if keyLength > maxKeyLength {
			maxKeyLength = keyLength
		}
	}
	return maxKeyLength
}
