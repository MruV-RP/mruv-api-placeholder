package generator

import "reflect"

type IGenerator interface {
	GenerateData(t reflect.Type) interface{}
}
