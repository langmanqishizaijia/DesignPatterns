
/*
name:       工厂模式
charactor： 向用户隐藏了产品的生成细节， 用户只要发一个命令给工厂，工厂就会返回合适的产品，这样利于后台做类型扩展，做到服务的隐蔽性
*/
package main

import "fmt"

type Action interface{
	Move(int) int
}

type Animal struct{
	Name string
}


type Fish struct{
	Animal
}

func (this * Fish)Move( num int) int {
	fmt.Printf("I am is a Fish, I swiming %d meters \n", num)
	return num
}
type Dog struct{
	Animal
}
func (this * Dog) Move( num int) int {
	fmt.Printf("I am is a Dog, I runnig %d meters \n", num)
	return num
}
type Cat struct{
	Animal
}
func (this * Cat) Move( num int) int {
	fmt.Printf("I am is a cat, I jump %d meters \n", num)
	return num
}



type AnimalFactory struct {

}

func NewAnimalFactory ()*AnimalFactory{
	return &AnimalFactory{}
}


func (this *AnimalFactory)CreatAnimal(name string) Action{
	switch name {
	case "cat":
		return &Cat{}
	case "dog":
		return &Dog{}
	case "fish":
		return &Fish{}
	default:
		panic("error animal type")
	    return nil
	}
}






func main(){

	bird := NewAnimalFactory().CreatAnimal("cat")
	bird.Move(11)

	fish := NewAnimalFactory().CreatAnimal("fish")
	fish.Move(20)

	dog := NewAnimalFactory().CreatAnimal("dog")
	dog.Move(100)
}
