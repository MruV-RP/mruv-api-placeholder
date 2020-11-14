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
		fieldCounter++
		switch field.Type.Kind() {
		case reflect.String:
			result.Field(i).SetString(fmt.Sprintf("Test%s", field.Name))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result.Field(i).SetInt(int64(fieldCounter))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result.Field(i).SetUint(uint64(fieldCounter))
		case reflect.Float32, reflect.Float64:
			result.Field(i).SetFloat(float64(fieldCounter) + 0.1488)
		}
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
