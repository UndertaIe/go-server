package main

import (
	"fmt"
	"reflect"
	"testing"
)

type S struct {
	A string
}

func test_main(Tt testing.T) {
	var arr []S
	if arr == nil {
		fmt.Println("arr is nil")
	}
	arr = append(arr, S{"A"}, S{"B"}, S{"A"}, S{"B"})
	var x any = arr
	fmt.Println(x)
	v := tran(x)
	fmt.Println(v)
}

func tran(x any) any {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Slice {
		panic("x type must be slice")
	}
	data := make([]any, 0)
	len := v.Len()
	for i := 0; i < len; i++ {
		data = append(data, v.Index(i).Interface())
	}
	return data
}
