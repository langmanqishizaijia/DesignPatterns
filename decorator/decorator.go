package decorator

import (
	"strings"
)

/*
装饰者模式： 装饰对象中含有一个被装饰者的对象
*/
type StringsBuider interface {
	Build(messages...string ) string
}

type BaseStringsBuilder struct{}
func (this *BaseStringsBuilder) Build(messages... string )string{
	return strings.Join(messages, ",")
}

//大括号分割
type BraceStringBuilderDecorator struct{
	StringsBuider
}
func (this *BraceStringBuilderDecorator)Build(messages...string)string{
	return "{" + this.StringsBuider.Build(messages ...) + "}"
}
//引号分割
type QuoteStringBuilderDecorator struct{
	StringsBuider
}
func (this *QuoteStringBuilderDecorator)Build(messages ... string)string{
	return "\"" + this.StringsBuider.Build(messages...)+"\""
}