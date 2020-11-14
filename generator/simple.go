package generator

import (
	"fmt"
	"reflect"
)

type SimpleGenerator struct {
}

func (s SimpleGenerator) FillWithTestData(object interface{}) interface{} {
	typ := reflect.TypeOf(object).Elem()
	result := reflect.ValueOf(object).Elem()
	fieldCounter := 0
	for i := 0; i < typ.NumField(); i++ {
		if !result.Field(i).CanInterface() {
			continue // continue on unexported fields
		}

		field := typ.Field(i)
		fmt.Printf("Field type: %v Field name: %v\n", field.Type, field.Name)
		switch field.Type.Kind() {
		case reflect.String:
			result.Field(i).SetString(fmt.Sprintf("Test%s", field.Name))
		case reflect.Int:
			result.Field(i).SetInt(int64(fieldCounter))
		}
		fieldCounter++
	}
	return object
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}
