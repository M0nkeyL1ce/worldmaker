package plates

import (
	"fmt"
	"worldmaker/field"
	"worldmaker/world"
)

type Plate struct {
	plateId
}

var Plates []Plate

// Factory function
func NewPlates(world world.World) *Plates {
	return &Plates
}

// Getters
// Get the plates for a world at a location.
func (w *world.World) Plates() *Plates {
	// optionally do some bounds checking here
	return &Plates
}

// Setter
func (w *World) SetField(x int, y int, value field.Field) {
	// optionally do some bounds checking here
	w.TileData[x][y] = value
}

func (w *Plates) ToGIF() {
	fmt.Println("Plates GIF Output Here")
}
