package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	arch := runtime.GOARCH
	fmt.Printf("arch:%s\n", arch)

	intSize := strconv.IntSize
	fmt.Printf("int size:%d\n", intSize)

	// fmt.Printf("cpu number:%s\n", runtime.NumCPU)

	var b1 byte = 127
	var i1 int = 123
	var i2 int32 = 123
	var i3 int64 = 123
	var f1 float32 = 0.123
	var f2 float64 = 1987623.432
	fmt.Printf("b1:%T, i1:%T, i2:%T , i3:%T, f1:%T, f2:%T\n", b1, i1, i2, i3, f1, f2)

	var v1 int = 321
	var v2 int32
	v2 = int32(v1)
	fmt.Printf("v1:%d, v2:%d\n", v1, v2)

	// &指针，*取指针的值
	str := "test"
	ptr := &str
	fmt.Printf("ptr:%p, ptr type:%T, ptr val:%s\n", ptr, ptr, *ptr)
}
