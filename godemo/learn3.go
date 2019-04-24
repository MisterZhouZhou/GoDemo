package main

import "unsafe"

// 枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)
}
