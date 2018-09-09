package factor

import "fmt"

/*
简单工厂模式，使用传类型来创建
*/
type Pen interface{
	//写字
	Write()
}
type Pencil struct{
}
func (this *Pencil) Write(){
	fmt.Println("I am writting with pencil")
}

type BrushPen struct{
}
func (this *BrushPen) Write(){
	fmt.Println("I am writing with brush pen")
}

type PenFactory struct{

}
func (this *PenFactory) ProducePen(typ string) Pen{

	switch typ {
	case "pencil":
		return &Pencil{}
	case "brush" :
		return &BrushPen{}
	default:
		panic("pen type is error")
	}

}

