package main

import "fmt"

//定义一个枚举
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

//color作为byte的别名
type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

//一个列表的box设置为一个类型
type BoxList []Box //a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	//这里自动转化原例子
	//go语言会自动识别如果传入的是指针
	//*b.color=c
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PainItBlack() {
	for i, _ := range bl {
		//这里会自动识别接收者中调用的方法是否是传入值类型，也就是指针
		//（&bl[i]）.SetColor(BLACK)
		//先取到地址再调用
		bl[i].SetColor(BLACK)
	}
}

//重写字符串
func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		{4, 4, 4, RED},
		{10, 10, 1, YELLOW},
		{2, 5, 20, BLACK},
		{8, 20, 1, WHITE},
		{20, 20, 20, YELLOW},
	}
	fmt.Printf("we have %d boxes in our set \n", len(boxes))
	fmt.Println("first one is ", boxes[0])
	fmt.Println("first one volume:", boxes[0].Volume())
	fmt.Println("boxes biggest color:", boxes.BiggestsColor())
}
