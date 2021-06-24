package cf

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func Load(data map[string]interface{}, cf interface{}) error {
	return load(data, cf, globalTypeHandlers)
}

func LoadCustom(data map[string]interface{}, cf interface{}, typeHandlers map[reflect.Type]TypeHandler) error {
	localTypeHandlers := make(map[reflect.Type]TypeHandler)
	for k, v := range globalTypeHandlers {
		localTypeHandlers[k] = v
	}
	for k, v := range typeHandlers {
		localTypeHandlers[k] = v
	}
	return load(data, cf, localTypeHandlers)
}

func load(data map[string]interface{}, cf interface{}, typeHandlers map[reflect.Type]TypeHandler) error {
	cfV := reflect.ValueOf(cf)
	if cfV.Kind() == reflect.Ptr {
		cfV = cfV.Elem()
	}
	if cfV.Kind() != reflect.Struct {
		return errors.Errorf("cf type [%s] not struct", cfV.Type())
	}
	for i := 0; i < cfV.NumField(); i++ {
		if cfV.Field(i).CanInterface() {
			fd := parseFieldData(cfV.Type().Field(i))
			if !fd.skip {
				if v, found := data[fd.name]; found {
					if cfV.Field(i).CanSet() {
						if handler, found := typeHandlers[cfV.Type().Field(i).Type]; found {
							// handler-based
							if err := handler(v, cfV.Field(i)); err != nil {
								return errors.Wrapf(err, "field '%s'", fd.name)
							}
						} else {
							nestedType := cfV.Type().Field(i).Type

							if nestedType.Kind() == reflect.Struct || (nestedType.Kind() == reflect.Ptr && nestedType.Elem().Kind() == reflect.Struct) {
								nested := instantiatePtr(nestedType)
								if submap, ok := v.(map[string]interface{}); ok {
									err := load(submap, nested, typeHandlers)
									if err != nil {
										return errors.Wrapf(err, "field '%s'", fd.name)
									}
								} else {
									return errors.Errorf("invalid submap for field '%s'", fd.name)
								}

								if nestedType.Kind() == reflect.Ptr {
									// by reference
									cfV.Field(i).Set(reflect.ValueOf(nested))
								} else {
									// by value
									cfV.Field(i).Set(reflect.ValueOf(nested).Elem())
								}

							} else {
								return errors.Errorf("no type handler for field '%s' of type [%s]", fd.name, cfV.Field(i).Type())
							}
						}
					}
				} else {
					if fd.required {
						return errors.Errorf("no data found for required field '%s'", fd.name)
					}
				}
			}
		}
	}
	return nil
}

type fieldData struct {
	name     string
	skip     bool
	required bool
}

func parseFieldData(v reflect.StructField) fieldData {
	fd := fieldData{name: v.Name, skip: false, required: false}
	data := v.Tag.Get("cf")
	if data != "" {
		tokens := strings.Split(data, ",")
		for _, token := range tokens {
			if token == "+required" {
				fd.required = true
			} else if token == "+skip" {
				fd.skip = true
			} else {
				fd.name = token
			}
		}
	}
	return fd
}

func instantiatePtr(t reflect.Type) interface{} {
	var it = t
	if it.Kind() == reflect.Ptr {
		it = it.Elem()
	}
	if i, found := globalInstantiators[it]; found {
		return i()
	} else {
		return reflect.New(it).Interface()
	}
}
