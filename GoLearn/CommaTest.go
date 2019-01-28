package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Element interface{}
type List []Element
type Person struct {
	name string
	age  int
}

//定义String方法，实现fmt.stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years"
}
func main() {
	list := make(List, 3)
	list[0] = 1 //int
	list[1] = "Hello"
	list[2] = Person{"BanMing", 25}
	//sort.Interface()
	for index, element := range list {
		fmt.Println(reflect.TypeOf(element))
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is an person and its value is %s\n", index, value)
		}
		//可以使用switch来判断
		//switch vaule1 := element.(type) {
		//case int:
		//	fmt.Println("list[%d] is an int and its value is %d\n", index, vaule1)
		//}
	}

}
