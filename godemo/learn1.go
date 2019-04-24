package main

import (
	"fmt"
)

// 基本数据类型
var a = "菜鸟教程"
var b string = "runoob.com"
var c bool

var x, y int
var (
	xa int
	xb bool
)
var xc, xd int = 1, 2
var xe, xf = 123, "hello"

func main() {
	// print(a, b, c)
	// // fmt输出方式
	fmt.Print(a, b, c)
	xg, xh := 123, "hello"
	println(x, y, xa, xb, xc, xd, xe, xf, xg, xh)
}
