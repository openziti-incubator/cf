package cf

import (
	"github.com/pkg/errors"
	"reflect"
)

func intHandler(v interface{}, f reflect.Value) error {
	if vt, ok := v.(int); ok {
		f.SetInt(int64(vt))
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func float64Handler(v interface{}, f reflect.Value) error {
	if vt, ok := v.(float64); ok {
		f.SetFloat(vt)
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func boolHandler(v interface{}, f reflect.Value) error {
	if vt, ok := v.(bool); ok {
		f.SetBool(vt)
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func stringHandler(v interface{}, f reflect.Value) error {
	if vt, ok := v.(string); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetString(vt)
		} else {
			f.SetString(vt)
		}
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}
