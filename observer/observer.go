package observer

import (
	"fmt"
)

/*
观察者模式
订阅者   报社
所有订阅者给报社留一个通知窗口，一旦报社有信息发生改变，就会用窗口统一去通知各个订阅者
*/
type Obsever interface{
	Notify(interface{})
}
type Subject struct{
	obsevers []Obsever
	state string
}
func (s *Subject)SetState(state string ){
	s.state = state
	s.NotfiyAllObservers()
}
func (s *Subject)Attach(observer... Obsever){
	s.obsevers = append(s.obsevers, observer...)
}
func (s *Subject)NotfiyAllObservers(){
	for _, obs := range s.obsevers{
		obs.Notify(s)
	}
}

type AObserver struct {
	Id string
}
func (this *AObserver) Notify(sub interface{}){
	fmt.Printf("obserId : %s recive state %s \n",this.Id,sub.(*Subject).state)
}