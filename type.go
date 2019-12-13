package gotype

import (
	"fmt"
	"reflect"
)

func Kind(i interface{}) string {
	rt := reflect.TypeOf(i)
	return fmt.Sprint(rt.Kind())
}

func IsPointer(i interface{}) bool {
	return Kind(i) == "ptr"
}

func IsStruct(i interface{}) bool {
	return Kind(i) == "struct"
}
