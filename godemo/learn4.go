package main

func main() {
	var a int = 4
	var ptr *int

	ptr = &a
	println("a的值为：%d", a)
	println("a的指针地址：%d", ptr)
	println("*ptr为 %d", *ptr)

	var xa int = 100
	var xb int = 200
	var ret int
	ret = max(xa, xb)
	print("最大值为: ", ret)
	ya, yb := swap(xa, xb)
	print("交换后的结果", "a=", ya, "b=", yb)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y int) (int, int) {
	return y, x
}
