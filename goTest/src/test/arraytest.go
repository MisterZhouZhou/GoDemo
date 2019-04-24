package test

import "fmt"

func arraytest()  {
	// numbers := make([]int, 3, 5)
	// fmt.Println(numbers)


	// numbers := []int{0,1,2,3,4,5,6,7,8}
	var numbers []int
	numbers = append(numbers, 0)
	numbers = append(numbers, 2,3,4)
	fmt.Println(numbers)
}

