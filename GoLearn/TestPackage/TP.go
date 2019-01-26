package TestPackage

import "fmt"

type TP struct {
	str string
}

func (tp TP) Print() {
	fmt.Println("TP Print:" + tp.str)
}
