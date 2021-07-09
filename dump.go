package cf

import (
	"fmt"
	"reflect"
)

func Tree(cf interface{}, opt *Options) string {
	return treeNode(reflect.ValueOf(cf), -1, opt)
}

func treeNode(v reflect.Value, indent int, opt *Options) string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Struct:
		return treeNodeStruct(v, indent + 1, opt)
	case reflect.Slice:
		return treeNodeSlice(v, indent + 1, opt)
	default:
		return treeNodeValue(v)
	}
}

func treeNodeStruct(v reflect.Value, indent int, opt *Options) string {
	format := fmt.Sprintf("%%-%ds", maxKeyLength(v, opt))
	fmt.Println(format)
	out := "{\n"
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			fd := parseFieldData(v.Type().Field(i), opt)
			out += nLevels(indent + 1) + fmt.Sprintf(format, fd.name) + " = " + treeNode(v.Field(i), indent, opt) + "\n"
		}
	}
	out += nLevels(indent) + "}"
	return out
}

func treeNodeSlice(v reflect.Value, indent int, opt *Options) string {
	out := "[\n"
	for i := 0; i < v.Len(); i++ {
		out += nLevels(indent + 1) + treeNode(v.Index(i), indent, opt) + "\n"
	}
	out += nLevels(indent) + "]"
	return out
}

func treeNodeValue(v reflect.Value) string {
	if v.Kind() == reflect.String {
		return fmt.Sprintf("\"%v\"", v.Interface())
	} else {
		return fmt.Sprintf("%v", v.Interface())
	}
}

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
	return out
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
