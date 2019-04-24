package main

import (
	"fmt"
	// "./src/cal"
	// "./src/cal/multi"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {
	// result := cal.Add(1,2)
	// fmt.Println(result)

	// result2 := cal.Subtract(110,2)
	// fmt.Println(result2)

	// result3 := multi.Multi(2,4)
	// fmt.Println(result3)

	// var a int = 10
	// fmt.Printf("地址%x\n", &a)
	// a= 11
	// fmt.Printf("地址%x1\n", &a)


	// 创建一个新的结构体
	// book := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	// fmt.Println(book.title)

	nums := []int{2,3,4}
	sum := 0
	for i,num := range(nums) {
		sum += num
		fmt.Println(i,"==", num)
	}
	fmt.Println(sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range(kvs) {
		fmt.Println(k," ", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
    }

	fmt.Printf("c".capitalize())


}	