package cf

import (
	"github.com/pkg/errors"
	"reflect"
)

type Instantiator func() interface{}

var globalInstantiators = make(map[reflect.Type]Instantiator)

func SetInstantiator(t reflect.Type, i Instantiator) {
	globalInstantiators[t] = i
}

type Setter func(v interface{}, f reflect.Value) error

func SetSetter(t reflect.Type, h Setter) {
	globalSetters[t] = h
}

var globalSetters = map[reflect.Type]Setter{
	reflect.TypeOf(0):          intHandler,
	reflect.TypeOf(float64(0)): float64Handler,
	reflect.TypeOf(true):       boolHandler,
	reflect.TypeOf(""):         stringHandler,
	reflect.TypeOf([]string{}): stringArrayHandler,
}

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
		f.SetString(vt)
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}

func stringArrayHandler(v interface{}, f reflect.Value) error {
	if vt, ok := v.([]string); ok {
		f.Set(reflect.ValueOf(vt))
		return nil
	}
	return errors.Errorf("got [%s], expected [%s]", reflect.TypeOf(v), f.Type())
}
