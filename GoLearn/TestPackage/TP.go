package TestPackage

import "fmt"

type TP struct {
	Str string
}

func (tp TP) Print() {
	fmt.Println("TP Print:" + tp.Str)
}
