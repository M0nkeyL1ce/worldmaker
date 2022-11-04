package plates

import (
	"fmt"
	"worldmaker/field"
	"worldmaker/world"
)

type Plate struct {
}

type Plates struct {
	plate []Plate
}

// Factory function
func NewPlates(world World) *Plates {
	return &Plates{}
}

// Getters
func Plates() *plates.Plates {
	// optionally do some bounds checking here
	y_norm := (w.LatSpan / 2) + y
	return &w.field[x*w.LatSpan+y_norm]
}

// Setter
func (w *world.World) SetPlate(x int, y int, value field.Field) {
	// optionally do some bounds checking here
	y_norm := (w.LatSpan / 2) + y
	w.Field.field[x*w.LatSpan+y_norm] = value
}

func (w *Plates) ToGIF() {
	fmt.Println("Plates GIF Output Here")
}
