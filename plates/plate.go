package plates

import (
	"fmt"
	"worldmaker/world"
)

type plate struct {
	plateId uint
}

type plateCollection struct {
	index  int
	plates []*plate
}

func (p *plateCollection) createIterator() utility.iterator {
	return &plateIterator{
		plates: p.plates,
	}
}

// Factory function
func NewPlates(world world.World) *Plates {
	ps = new(Plates)
	return ps
}

func (w *Plates) ToGIF() {
	fmt.Println("Plates GIF Output Here")
}
