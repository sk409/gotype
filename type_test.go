package gotype

import "testing"

func TestKind(t *testing.T) {
	check := func(i interface{}, expected string) {
		kind := Kind(i)
		if expected != kind {
			t.Errorf("Expected %s but got %s", expected, kind)
		}
	}
	var i interface{}
	i = 0
	check(i, "int")
	i = int8(0)
	check(i, "int8")
	i = int16(0)
	check(i, "int16")
	i = int32(0)
	check(i, "int32")
	i = int64(0)
	check(i, "int64")
	i = uint(0)
	check(i, "uint")
	i = uint8(0)
	check(i, "uint8")
	i = uint16(0)
	check(i, "uint16")
	i = uint32(0)
	check(i, "uint32")
	i = uint64(0)
	check(i, "uint64")
	i = ""
	check(i, "string")
	i = map[string]string{}
	check(i, "map")
	i = struct{}{}
	check(i, "struct")
	i = &struct{}{}
	check(i, "ptr")
}

func TestIsInt(t *testing.T) {
	var i interface{}
	i = 0
	isInt := IsInt(i)
	if !isInt {
		t.Error("false positive")
	}
	i = ""
	isInt = IsInt(i)
	if isInt {
		t.Error("false negative")
	}
}

func TestIsInt8(t *testing.T) {
	var i interface{}
	i = int8(0)
	isInt8 := IsInt8(i)
	if !isInt8 {
		t.Error("false positive")
	}
	i = ""
	isInt8 = IsInt8(i)
	if isInt8 {
		t.Error("false negative")
	}
}

func TestIsInt16(t *testing.T) {
	var i interface{}
	i = int16(0)
	isInt16 := IsInt16(i)
	if !isInt16 {
		t.Error("false positive")
	}
	i = ""
	isInt16 = IsInt16(i)
	if isInt16 {
		t.Error("false negative")
	}
}

func TestIsInt32(t *testing.T) {
	var i interface{}
	i = int32(0)
	isInt32 := IsInt32(i)
	if !isInt32 {
		t.Error("false positive")
	}
	i = ""
	isInt32 = IsInt32(i)
	if isInt32 {
		t.Error("false negative")
	}
}

func TestIsInt64(t *testing.T) {
	var i interface{}
	i = int64(0)
	isInt64 := IsInt64(i)
	if !isInt64 {
		t.Error("false positive")
	}
	i = ""
	isInt64 = IsInt64(i)
	if isInt64 {
		t.Error("false negative")
	}
}

func TestIsUint(t *testing.T) {
	var i interface{}
	i = uint(0)
	isUint := IsUint(i)
	if !isUint {
		t.Error("false positive")
	}
	i = ""
	isUint = IsUint(i)
	if isUint {
		t.Error("false negative")
	}
}

func TestIsUint8(t *testing.T) {
	var i interface{}
	i = uint8(0)
	isUint8 := IsUint8(i)
	if !isUint8 {
		t.Error("false positive")
	}
	i = ""
	isUint8 = IsUint8(i)
	if isUint8 {
		t.Error("false negative")
	}
}

func TestIsUint16(t *testing.T) {
	var i interface{}
	i = uint16(0)
	isUint16 := IsUint16(i)
	if !isUint16 {
		t.Error("false positive")
	}
	i = ""
	isUint16 = IsUint16(i)
	if isUint16 {
		t.Error("false negative")
	}
}

func TestIsUint32(t *testing.T) {
	var i interface{}
	i = uint32(0)
	isUint32 := IsUint32(i)
	if !isUint32 {
		t.Error("false positive")
	}
	i = ""
	isUint32 = IsUint32(i)
	if isUint32 {
		t.Error("false negative")
	}
}

func TestIsUint64(t *testing.T) {
	var i interface{}
	i = uint64(0)
	isUint64 := IsUint64(i)
	if !isUint64 {
		t.Error("false positive")
	}
	i = ""
	isUint64 = IsUint64(i)
	if isUint64 {
		t.Error("false negative")
	}
}

func TestIsString(t *testing.T) {
	var i interface{}
	i = ""
	isString := IsString(i)
	if !isString {
		t.Error("false positive")
	}
	i = 0
	isString = IsString(i)
	if isString {
		t.Error("false negative")
	}
}

func TestIsMap(t *testing.T) {
	var i interface{}
	i = map[string]string{}
	isMap := IsMap(i)
	if !isMap {
		t.Error("false positive")
	}
	i = ""
	isMap = IsMap(i)
	if isMap {
		t.Error("false negative")
	}
}

func TestIsStruct(t *testing.T) {
	var i interface{}
	i = struct{}{}
	isStruct := IsStruct(i)
	if !isStruct {
		t.Error("false positive")
	}
	i = ""
	isStruct = IsStruct(i)
	if isStruct {
		t.Error("false negative")
	}
}

func TestIsPointer(t *testing.T) {
	var i interface{}
	i = &struct{}{}
	isPointer := IsPointer(i)
	if !isPointer {
		t.Error("false positive")
	}
	i = ""
	isPointer = IsPointer(i)
	if isPointer {
		t.Error("false negative")
	}
}

func TestIsSlice(t *testing.T) {
	var i interface{}
	i = []string{}
	isSlice := IsSlice(i)
	if !isSlice {
		t.Error("false positive")
	}
	i = ""
	isSlice = IsSlice(i)
	if isSlice {
		t.Error("false negative")
	}
}
