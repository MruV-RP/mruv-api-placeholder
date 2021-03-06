package generator

import (
	"reflect"
	"testing"
)

type TestStructure struct {
	somePrivateField string

	Name    string
	Surname string
	Age     int
}

func TestSimpleGenerator_FillWithTestData(t *testing.T) {
	type args struct {
		object interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"Simple test",
			args{
				object: &TestStructure{},
			},
			&TestStructure{
				Name:    "TestName",
				Surname: "TestSurname",
				Age:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SimpleGenerator{}
			if got := s.FillWithTestData(tt.args.object); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillWithTestData() = %v, want %v", got, tt.want)
			}
		})
	}
}
