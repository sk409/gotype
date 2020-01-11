package gotype

import (
	"fmt"
	"reflect"
)

func Kind(i interface{}) string {
	rt := reflect.TypeOf(i)
	return fmt.Sprint(rt.Kind())
}

func IsInt(i interface{}) bool {
	return Kind(i) == "int"
}

func IsInt8(i interface{}) bool {
	return Kind(i) == "int8"
}

func IsInt16(i interface{}) bool {
	return Kind(i) == "int16"
}

func IsInt32(i interface{}) bool {
	return Kind(i) == "int32"
}

func IsInt64(i interface{}) bool {
	return Kind(i) == "int64"
}

func IsUint(i interface{}) bool {
	return Kind(i) == "uint"
}

func IsUint8(i interface{}) bool {
	return Kind(i) == "uint8"
}

func IsUint16(i interface{}) bool {
	return Kind(i) == "uint16"
}

func IsUint32(i interface{}) bool {
	return Kind(i) == "uint32"
}

func IsUint64(i interface{}) bool {
	return Kind(i) == "uint64"
}

func IsString(i interface{}) bool {
	return Kind(i) == "string"
}

func IsMap(i interface{}) bool {
	return Kind(i) == "map"
}

func IsStruct(i interface{}) bool {
	return Kind(i) == "struct"
}

func IsPointer(i interface{}) bool {
	return Kind(i) == "ptr"
}

func IsSlice(i interface{}) bool {
	return Kind(i) == "slice"
}
