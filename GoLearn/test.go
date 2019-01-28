package m

import (
	"fmt"
	"github.com/BanMing/BanMingGoLab/GoLearn/TestPackage"
)

func main() {
	fmt.Errorf("sssss\n")
	//var ts = TS{str: "sss"}
	//ts.PrintTs()
	tp := TestPackage.TP{
		Str: "dsfsef",
	}
	tp.Print()
	ts := TS{
		str: "dddd",
	}
	ts.PrintTs()
	//ts := TS{str: "s"}
	//ts.PrintTs()
}
