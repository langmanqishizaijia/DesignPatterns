package observer

import "testing"

func TestObserver (t *testing.T){
	//a := new(AObserver)
	//a.Id = "wangshuchao"
	a := &AObserver{Id:"wangshuchao"}

	//b := new(AObserver)
	//b.Id = "sunjing"
	b := &AObserver{Id:"sunjing"}
	s := new(Subject)
	s.Attach(a,b)

	s.SetState("kkkkkk")
	s.SetState("i know")
}
