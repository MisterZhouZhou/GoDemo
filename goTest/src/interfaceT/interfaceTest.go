package interfaceT

import "fmt"

//阶乘
func Factorial(n uint64)(result uint64)  {
	if n > 4 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

//斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func testFa()  {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func testFi()  {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {

}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

/**
  向外提供方法
 */
func TestPhone()  {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
