package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	redius float64
}

//这种用法像似c# this扩展的写法
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.redius * c.redius * math.Pi
}

//func (i Circle) area() float64 {
//	return i * i
//}
func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	fmt.Println(r1.area())
	fmt.Println(r2.area())
	fmt.Println(c1.area())
	//fmt.Println(r1.area())

}

type ages int
type money float32
type mouths map[string]int
