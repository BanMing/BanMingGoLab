package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	//fmt.Printf("Hello world")
	//FixStr()
	//ErrorFunc()
	//MapTest()
	//MapTestTemp()
	//ArrTestTemp()
	//SliceTestTemp()
	//ForTest()
	ArgParamTest(1, 3, 56, 54, 7)
}

var frenchHello string
var emptyString = ""

//多行字符串
var mutilLine = `hhh
ssss`

//集合申明变量
var (
	a int
	b string
)

// 修改字符串
func FixStr() {
	var s = "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
}

//切片字符串
func SpitStr() {
	s := "hello"
	c := "c" + s[1:]
	fmt.Println(c)
}

//错误类型
func ErrorFunc() {
	err := errors.New("emit macho dwarf :elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

//枚举设置
const (
	x = iota
	y = iota
	z = iota
	w //常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w=iota，因此w==3。
)

//数组定义定长
var arr [10]int

func ArrTest() {
	//a := [3]int{1, 2, 3}   //给数组初始化
	//b := [10]int{1, 2, 3}  //给数组初始化前三个值，后面值为默认
	//c := [...]int{4, 5, 6} //根据数组初始长度设置数组长度
}

//值类型
func ArrTestTemp() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 2
	fmt.Println(a[0])
}

//动态数组
var slice []int

func SliceTest() {
	//var ar = [10]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'}
	//var a, b []byte
	////切片表示
	//a = ar[2:5]
	//b = ar[3:5]
}
func SliceTestTemp() {
	a := make(map[int]string)
	b := a

	a1 := [2]int{1, 2}
	b1 := a1
	//t := reflect.TypeOf(a)
	//t1 := reflect.TypeOf(a1)
	fmt.Printf("%p  %p", a, b)
	fmt.Printf("%p  %p", a1, b1)

}

//map key在括号里 value为后面的值
var numbers map[string]int

func MapTest() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers["three"])
}
func MapTestTemp() {
	rating := map[string]float32{"C": 5, "GO": 4.5, "Python": 4.5, "C++": 2}
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C")
}

//在判断语句中可以声明变量
func IfTest() {
	if x := 10; x > 10 {
		fmt.Println("x>10")
	} else {
		fmt.Println("x<10")
	}
}

//goto测试
func GotoTest() {
	i := 0
	//大小写区别
Here:
	fmt.Println("Run Here Code " + string(i))
	i++
	goto Here
}

//For测试
func ForTest() {
	//for i := 1; i < 3; i++ {
	//	//fmt.Print("%{i} sss \n")
	//
	//	fmt.Println(strconv.Itoa(i))
	//}
	//
	//sum := 1
	//for ; sum < 1000; {
	//	sum += 1;
	//	fmt.Println(sum)
	//}
	////当while使用
	//sum = 1
	//for sum < 1000 {
	//sum += 1
	//	fmt.Println(sum)
	//}
	//for k, v := range map{
	//	fmt.Println("map's val :", v)
	//}
}

//多参数
func ArgParamTest(arg ...int) {
	//range 相当于是迭代器
	for _, n := range arg {
		fmt.Printf("And the number is :%d\n", n)
	}
}

//传入指针
func AddPoint(a *int) int {
	*a = *a + 1
	return *a
}

//defer 是在一个函数返回后执行，
// 倒序依次执行
// func ReadWrite() bool {
// 	//file.Open("file")
// 	//defer file.Close()
// }

//打印出来4，3，2，1
func DeferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
