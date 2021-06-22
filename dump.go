package cf

import (
	"fmt"
	"reflect"
)

func Dump(label string, cf interface{}) string {
	cfV := reflect.ValueOf(cf)
	if cfV.Kind() == reflect.Ptr {
		cfV = cfV.Elem()
	}
	if cfV.Kind() != reflect.Struct {
		return ""
	}
	out := label + " {\n"
	format := fmt.Sprintf("\t%%-%ds %%v\n", maxKeyLength(cfV))
	for i := 0; i < cfV.NumField(); i++ {
		if cfV.Field(i).CanInterface() {
			key := keyName(cfV.Type().Field(i))
			out += fmt.Sprintf(format, key, cfV.Field(i).Interface())
		}
	}
	out += "}\n"
	return out
}
