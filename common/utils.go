package common

import (
	"errors"
	"reflect"
)

func CopyStruct(src, dst interface{}) error {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()
	for i := 0; i < sval.NumField(); i++ {
		val := sval.Field(i)
		name := sval.Type().Field(i).Name
		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			return errors.New("field cannot find")
		}
		dvalue.Set(val)
	}
	return nil
}
