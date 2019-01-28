package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
	phone  int
}

//基类方法
func (human Human) SayHi() {
	fmt.Println("my name:" + human.name + " hi~")
}

type Student struct {
	Human      //匿名字段，那么默认student就包含了human字段
	speciality string
	phone      int
}

//重写方法
func (student Student) SayHi() {
	//student.Human.SayHi()//实现基类方法
	fmt.Println("student my name:" + student.name + " Hi~")
}
func main() {
	mark := Student{Human{"Mark", 24, 123, 123123}, "student", 9999}
	fmt.Println(mark.name, mark.speciality)
	//默认调用最上层名字相同的字段
	fmt.Println(mark.phone)
	fmt.Println(mark.Human.phone)
	mark.SayHi()
	mark.Human.SayHi()
}
