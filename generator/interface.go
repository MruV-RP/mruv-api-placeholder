package generator

type IGenerator interface {
	FillWithTestData(object interface{}) interface{}
}
