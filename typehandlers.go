package cf

import (
	"github.com/pkg/errors"
	"reflect"
)

type TypeHandler func(v interface{}, f reflect.Value) error

func SetTypeHandler(t reflect.Type, h TypeHandler) {
	typeHandlers[t] = h
}

var typeHandlers = map[reflect.Type]TypeHandler {
	reflect.TypeOf(0): func(v interface{}, f reflect.Value) error {
		if vt, ok := v.(int); ok {
			f.SetInt(int64(vt))
			return nil
		}
		return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
	},
	reflect.TypeOf(float64(0)): func(v interface{}, f reflect.Value) error {
		if vt, ok := v.(float64); ok {
			f.SetFloat(vt)
			return nil
		}
		return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
	},
	reflect.TypeOf(true): func(v interface{}, f reflect.Value) error {
		if vt, ok := v.(bool); ok {
			f.SetBool(vt)
			return nil
		}
		return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
	},
	reflect.TypeOf(""): func(v interface{}, f reflect.Value) error {
		if vt, ok := v.(string); ok {
			f.SetString(vt)
			return nil
		}
		return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
	},
}
