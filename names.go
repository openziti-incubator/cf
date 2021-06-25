package cf

import "reflect"

func NilNameConverter(f reflect.StructField) string {
	return f.Name
}