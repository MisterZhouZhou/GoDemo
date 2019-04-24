package main

import "fmt"

func main() {
	// 创建切片
	numbers := []int{9, 1, 2, 3, 4, 5, 6, 7, 8}
	printSclice(numbers)

	// 打印原始切片
	fmt.Println("numbers=", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	sum := 0
	for index, num := range numbers {
		sum += num
		fmt.Print(index, num, "\n")
	}
	fmt.Println("sum:", sum)

	kvs := map[string]string{"a": "apple", "b": "baba"}
	for k, v := range kvs {
		fmt.Printf("%s->%s\n", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country, value := range countryCapitalMap {
		fmt.Println(country, "首都是", value)
	}
}

func printSclice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
