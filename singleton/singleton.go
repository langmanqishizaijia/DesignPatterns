package singleton

import "fmt"

type Singleton interface{
	SaySomething()
}
type ASingleton struct{
}
func (this *ASingleton) SaySomething(){
	fmt.Println("Singleton say say")
}
var singletonInstance Singleton

func NewSingletonInstance() Singleton{
	if nil == singletonInstance{
		singletonInstance = &ASingleton{}
	}
	return singletonInstance
}
