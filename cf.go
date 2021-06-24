package cf

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func Bind(cf interface{}, data map[string]interface{}, opt *Options) error {
	cfV := reflect.ValueOf(cf)
	if cfV.Kind() == reflect.Ptr {
		cfV = cfV.Elem()
	}
	if cfV.Kind() != reflect.Struct {
		return errors.Errorf("provided type [%s] is not a struct", cfV.Type())
	}
	for i := 0; i < cfV.NumField(); i++ {
		if cfV.Field(i).CanInterface() {
			fd := parseFieldData(cfV.Type().Field(i))
			if !fd.skip {
				if v, found := data[fd.name]; found {
					if cfV.Field(i).CanSet() {
						if handler, found := opt.Setters[cfV.Type().Field(i).Type]; found {
							// handler-based type
							if err := handler(v, cfV.Field(i)); err != nil {
								return errors.Wrapf(err, "field '%s'", fd.name)
							}
						} else {
							// nested structure
							nestedType := cfV.Type().Field(i).Type
							if nestedType.Kind() == reflect.Struct || (nestedType.Kind() == reflect.Ptr && nestedType.Elem().Kind() == reflect.Struct) {
								nested := instantiateAsPtr(nestedType, opt)
								if subData, ok := v.(map[string]interface{}); ok {
									err := Bind(nested, subData, opt)
									if err != nil {
										return errors.Wrapf(err, "field '%s'", fd.name)
									}
								} else {
									return errors.Errorf("invalid submap for field '%s'", fd.name)
								}

								if nestedType.Kind() == reflect.Ptr {
									// by pointer
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
	// execute wirings for type
	if opt.TypeWirings != nil {
		if wirings, found := opt.TypeWirings[cfV.Elem().Type()]; found {
			for _, wiring := range wirings {
				if err := wiring(cfV.Elem()); err != nil {
					return errors.Wrapf(err, "error wiring [%s]", cfV.Elem().Type().Name())
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

func instantiateAsPtr(t reflect.Type, opt *Options) interface{} {
	var it = t
	if it.Kind() == reflect.Ptr {
		it = it.Elem()
	}
	if i, found := opt.Instantiators[it]; found {
		return i()
	} else {
		return reflect.New(it).Interface()
	}
}
