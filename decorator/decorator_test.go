package decorator

import (
	"fmt"
	"testing"
	)

func Test1StringBuilder(t *testing.T){
	var MB StringsBuider
	MB = &BaseStringsBuilder{}
	fmt.Println(MB.Build("hello wang shuchao"))

	MB = &BraceStringBuilderDecorator{MB}
	fmt.Println(MB.Build("hello wang shuchao"))

	MB = &QuoteStringBuilderDecorator{MB}
	fmt.Println(MB.Build("hello wang shuchao"))

}
