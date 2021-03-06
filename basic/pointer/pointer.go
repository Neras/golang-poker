package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1) // 取地址符 &
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	stringPointer()
}

func stringPointer() {
	s := "good bye"
	var p *string = &s // p是一个指针，存储的s的地址
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
	fmt.Printf("Here is the string %v\n", &s)
}