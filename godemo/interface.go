package main

import "fmt"

type Person interface {
	name() string
	age() int
}

type Woman struct {
}

func (woman Woman) name() string {
	return "Jin Yawei"
}

func (woman Woman) age() int {
	return 23
}

type Men struct {
}

func (men Men) name() string {
	return "liweibin"
}
func (men Men) age() int {
	return 27
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

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	var person Person
	person = new(Woman)
	fmt.Println(person.name())
	fmt.Println(person.age())

	person = new(Men)
	fmt.Println(person.name())
	fmt.Println(person.age())

}
