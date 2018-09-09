package factor
import(
	"testing"
	"fmt"
)
func Test1PenFactory_ProducePen(t *testing.T) {
	typ := "pencil"
	penfactory := &PenFactory{}
	pen := penfactory.ProducePen(typ)
	pen.Write()
}

func TestABPen(t *testing.T){
	var abPenfactory ABFactory
	abPenfactory = new(ABPencilFactory)
	pen1 := abPenfactory.Produce()
	fmt.Println(pen1.Write())

	abPenfactory = new(ABBrushPenFactory)
	pen2 := abPenfactory.Produce()
	fmt.Println(pen2.Write())
}