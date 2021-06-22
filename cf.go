package cf

import (
	"github.com/pkg/errors"
	"reflect"
)

func Load(data map[string]interface{}, cf interface{}) error {
	cfV := reflect.ValueOf(cf)
	if cfV.Kind() == reflect.Ptr {
		cfV = cfV.Elem()
	}
	if cfV.Kind() != reflect.Struct {
		return errors.Errorf("cf type [%s] not struct", cfV.Type())
	}
	for i := 0; i < cfV.NumField(); i++ {
		if cfV.Field(i).CanInterface() {
			key := keyName(cfV.Type().Field(i))
			if v, found := data[key]; found {
				if cfV.Field(i).CanSet() {
					if handler, found := typeHandlers[reflect.TypeOf(cfV.Field(i).Interface())]; found {
						if err := handler(v, cfV.Field(i)); err != nil {
							return errors.Wrapf(err, "field '%s'", key)
						}
					} else {
						return errors.Errorf("no type handler for field '%s' of type [%s]", key, cfV.Field(i).Type())
					}
				}
			}
		}
	}
	return nil
}

func keyName(v reflect.StructField) string {
	key := v.Name
	tag := v.Tag.Get("cf")
	if tag != "" {
		key = tag
	}
	return key
}
