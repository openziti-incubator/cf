package cf

import (
	"github.com/iancoleman/strcase"
	"reflect"
)

func NilNameConverter(f reflect.StructField) string {
	return f.Name
}

func SnakeCaseNameConverter(f reflect.StructField) string {
	return strcase.ToSnake(f.Name)
}
