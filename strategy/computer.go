package strategy

import "fmt"

type Computer struct{
	Num1 int
	Num2 int
	strate Strate
}
func (this *Computer) SetStrate(strate Strate) {
	this.strate = strate
}
func (p Computer) Do() int {
	defer func() {
		if f := recover(); f != nil {
			fmt.Println(f)
		}
	}()

	if p.strate == nil {
		panic("Strategier is null")
	}

	return p.strate.Compute(p.Num1, p.Num2)
}
