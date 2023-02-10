package cutils

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

// Str2Bytes string转[]byte无拷贝
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2Str byte数组直接转成string对象，不发生内存copy
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Uint82Str uint8数组直接转成string对象,不发生内存copy
func Uint82Str(u []uint8) string {
	return *(*string)(unsafe.Pointer(&u))
}

// Obj2Json 将对象序列化成字符串
func Obj2Json(obj interface{}) string {
	bdata, err := jsoniter.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bdata)
}
