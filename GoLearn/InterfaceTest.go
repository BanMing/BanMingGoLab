package main

import "fmt"

type HumanTest struct {
	name  string
	ages  string
	phone int
}
type StudentTest struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (studentTest StudentTest) SayHi() {
	fmt.Println("student test say hi", studentTest.name)
}

//声明接口
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(money float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {

}
