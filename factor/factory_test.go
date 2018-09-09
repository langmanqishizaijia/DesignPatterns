package factor
import(
	"testing"
)
func Test1PenFactory_ProducePen(t *testing.T) {
	typ := "pencil"
	penfactory := &PenFactory{}
	pen := penfactory.ProducePen(typ)
	pen.Write()
}
