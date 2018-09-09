package designPatterns

import "testing"

func TestPenFactory_ProducePen(t *testing.T) {
	typ := "pencil"
	penfactory := &PenFactory{}
	pen := penfactory.ProducePen(typ)
	pen.Write()
}
