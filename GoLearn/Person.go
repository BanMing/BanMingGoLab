package main

import "fmt"

//结构体
type person struct {
	name string
	age  int
}

func main() {
	var P person
	P.name = "sss"
	P.age = 12
	fmt.Println(P.name)

	P1 := person{"ss2", 23}
	fmt.Println(P1.name, P1.age)

	P2 := person{name: "sw2", age: 25}
	fmt.Println(P2.name, P2.age)
}
