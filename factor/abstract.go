package factor

import "fmt"

type ABPen interface{
	Write()string
}
type ABPencil struct {
}
func (this *ABPencil) Write() string{
	return fmt.Sprintf("I am writing with ab pencil")
}
type ABBrushPen struct{
}
func (this *ABBrushPen) Write()string{
	return fmt.Sprintf("I am writing with ab Brush Pen")
}

type ABFactory interface{
	Produce() ABPen
}
type ABPencilFactory struct{
}
func (this *ABPencilFactory) Produce () ABPen{
	return new(ABPencil)
}
type ABBrushPenFactory struct{
}
func (this *ABBrushPenFactory)Produce() ABPen{
	return new(ABBrushPen)
}

