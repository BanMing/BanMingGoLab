package main

import "fmt"

func Print() {
	fmt.Errorf("testIm print")
}

type TS struct {
	str string
}

func (ts TS) PrintTs() {
	fmt.Println("Ts print" + ts.str)
}
