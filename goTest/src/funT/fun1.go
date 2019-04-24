package funcT

import "fmt"

func func1()  {
	_,numb,strs := numbers()
	fmt.Println(numb,strs)
	func2()
}

func numbers()(int, int, string) {
	a,b,c:=1,2,"str"
	return a,b,c
}