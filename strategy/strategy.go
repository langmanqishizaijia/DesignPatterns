package strategy

import "fmt"

type Strate interface{
	Compute(num1, num2 int)int
}
type Devision struct{
}
func (this *Devision) Compute(num1, num2 int) int{
	defer func(){
		if f := recover() ;f != nil{
			fmt.Println(f)
			return
		}
	}()
	if  num2 == 0{
		panic("num2 should not 0")
	}
	return num1/num2
}
type Add struct{}
func (this *Add)Compute(num1,num2 int) int{
	return num1 + num2
}
type Sub struct{}
func (this *Sub)Compute(num1, num2 int)int{
	return num1-num2
}

type Mutile struct{}
func (this *Mutile)Compute(num1, num2 int)int{
	return num1 * num2
}

func NewStrate (typ string) Strate{
	switch typ{
	case "a":
		return &Add{}
	case "s":
		return &Sub{}
	case "m":
		return &Mutile{}
	case "d":
		return &Devision{}
	default:
		return nil
	}
}
