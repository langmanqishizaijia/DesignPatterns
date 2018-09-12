package main

import (

	"flag"
	"designPatterns/strategy"
	"fmt"
)

var num1 *int = 	flag.Int("num1",1,"please input num1 ")
var num2 *int = 	flag.Int("num2",1,"please input num2 ")

var typ *string  = 	flag.String("type","a","please a cal type ")

func init(){
	flag.Parse()
}
func main(){
	strate := strategy.NewStrate(*typ)
	computer := &strategy.Computer{Num1:*num1,Num2:*num2}
	computer.SetStrate(strate)
	fmt.Println(computer.Do())
}
