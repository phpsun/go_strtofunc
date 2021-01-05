package main

import (
	"fmt"
	"reflect"
)

func main() {
	funcs := map[string]interface{}{"foo": foo, "bar": bar}
	Call(funcs, "foo")
	Call(funcs, "bar", 1, 2, 3)
}
func foo() {
	println("foo\r\n")
}
func bar(a, b, c int) {
	println("bar\r\n", a, b, c)
}
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		fmt.Println("参数个数：err")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
