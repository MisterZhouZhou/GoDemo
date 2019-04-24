package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {
	/* 声明局部变量 */
	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	g = a + b

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)

	var n [10]int
	var i, j int

	for i = 0; i < len(n); i++ {
		n[i] = i + 100
	}

	for j = 0; j < len(n); j++ {
		fmt.Printf("Element[%d]= %d\n", j, n[j])
	}
}
