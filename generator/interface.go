package generator

import "reflect"

type IGenerator interface {
	FillWithTestData(object interface{}) interface{}
	GenerateData(t reflect.Type) interface{}
}
