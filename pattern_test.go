package main

import (
	"testing"
	fac "designPatterns/factor"
	"fmt"
	dec "designPatterns/decorator"
)
func TestPenFactory_ProducePen(t *testing.T) {
	typ := "pencil"
	penfactory := &fac.PenFactory{}
	pen := penfactory.ProducePen(typ)
	pen.Write()
}

func TestStringBuilder(t *testing.T){
var MB dec.StringsBuider
MB = &dec.BaseStringsBuilder{}
fmt.Println(MB.Build("hello wang shuchao"))

MB = &dec.BraceStringBuilderDecorator{MB}
fmt.Println(MB.Build("hello wang shuchao"))

MB = &dec.QuoteStringBuilderDecorator{MB}
fmt.Println(MB.Build("hello wang shuchao"))

}

