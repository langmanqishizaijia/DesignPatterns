package singleton

import "fmt"

/*
另一种单利模式
使用golang的init函数自动生成对象
*/
 type SingleTon2 interface{
 	SaySomething()
 }
 type singleTon2 struct{
 }
 func (this *singleTon2) SaySomething(){
 	fmt.Println("Singleton2")
 }

 var singletonInstance2 SingleTon2
 func init() {
 	singletonInstance2 = new(singleTon2)
 }
 func Singleton2SaySomething(){
 	singletonInstance2.SaySomething()
 }
