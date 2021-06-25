package cf

import (
	"github.com/iancoleman/strcase"
	"reflect"
)

func PassthroughNameConverter(f reflect.StructField) string {
	return f.Name
}

func SnakeCaseNameConverter(f reflect.StructField) string {
	return strcase.ToSnake(f.Name)
}
