package engine

import (
	"errors"
	"fmt"
	"reflect"
)

var valueOfFhirOption = reflect.TypeOf([]map[string]interface{}{})

type DynamicValue struct {
	Value   interface{}
	typeOf  reflect.Type
	valueOf reflect.Value
}

func NewDynamicValue(value interface{}) *DynamicValue {
	return &DynamicValue{
		Value:   value,
		typeOf:  reflect.TypeOf(value),
		valueOf: reflect.ValueOf(value),
	}
}

func CastSliceValueAtIndexOfDynamicValue[T any](value *DynamicValue, index int) (T, error) {
	interfaceValue, err := value.SliceValueAtIndex(index)
	if err != nil {
		var zeroValue T
		return zeroValue, err
	}

	realValue, ok := interfaceValue.(T)
	if !ok {
		var zeroValue T
		return zeroValue, fmt.Errorf("value at index %d from the slice of a dynamic value is not of the requested type", index)
	}

	return realValue, nil
}

func (receiver *DynamicValue) IsSlice() bool {
	return receiver.typeOf.Kind() == reflect.Slice
}

func (receiver *DynamicValue) SliceSize() (int, error) {
	if !receiver.IsSlice() {
		return 0, errors.New("value of dynamicValue is not a slice")
	}

	return receiver.valueOf.Len(), nil
}

func (receiver *DynamicValue) SliceValueAtIndex(index int) (interface{}, error) {
	if !receiver.IsSlice() {
		return nil, errors.New("value of dynamicValue is not a slice")
	}

	return receiver.valueOf.Index(index).Interface(), nil
}

func (receiver *DynamicValue) IsSliceOfFhirOptions() bool {
	//is the value a []map[string]interface{}?
	return receiver.typeOf == valueOfFhirOption
}
