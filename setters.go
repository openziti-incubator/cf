package cf

import (
	"github.com/pkg/errors"
	"reflect"
)

func intSetter(v interface{}, f reflect.Value) error {
	if vt, ok := v.(int); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetInt(int64(vt))
		} else {
			f.SetInt(int64(vt))
		}
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func int8Setter(v interface{}, f reflect.Value) error {
	if vt, ok := v.(int8); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetInt(int64(vt))
		} else {
			f.SetInt(int64(vt))
		}
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func int16Setter(v interface{}, f reflect.Value) error {
	if vt, ok := v.(int16); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetInt(int64(vt))
		} else {
			f.SetInt(int64(vt))
		}
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func int32Setter(v interface{}, f reflect.Value) error {
	if vt, ok := v.(int32); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetInt(int64(vt))
		} else {
			f.SetInt(int64(vt))
		}
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func float64Setter(v interface{}, f reflect.Value) error {
	if vt, ok := v.(float64); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetFloat(vt)
		} else {
			f.SetFloat(vt)
		}
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func boolSetter(v interface{}, f reflect.Value) error {
	if vt, ok := v.(bool); ok {
		if f.Kind() == reflect.Ptr {
			f.Elem().SetBool(vt)
		} else {
			f.SetBool(vt)
		}
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func stringSetter(v interface{}, f reflect.Value) error {
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
