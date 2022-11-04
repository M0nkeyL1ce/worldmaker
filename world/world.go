package world

import (
	"fmt"
	"math"

	"worldmaker/field"
	"worldmaker/utility"
)

type World struct {
	TileData      [][]field.Field
	Circumference int
	Resolution    int
	HtmDepth      int
	LongSpan      int
	LatSpan       int
}

// Factory function
func NewWorld(circumference, resolution int, debug bool) *World {
	// Calculate the world geometry based on resolution
	long_span := int(float32(circumference) / float32(resolution))
	lat_span := int((float32(circumference)) / 2.0 / float32(resolution))
	// Calculate the triangle depth needed.
	// 24,901 needs 8*4^6 (32,768) data points for a total depth of 7.
	// logarithm of x base b = log(x)/log(b)
	htm_depth := int(math.Ceil(math.Log(float64(circumference)/float64(resolution)/8.0)/math.Log(4.0))) + 1
	fields := utility.Make2D[field.Field](long_span, lat_span)
	world := &World{TileData: fields,
		Circumference: circumference, Resolution: resolution, HtmDepth: htm_depth,
		LongSpan: long_span, LatSpan: lat_span}
	SetMeanAltitude(world, debug)
	return world
}

func SetMeanAltitude(world *World, debug bool) *World {
	// Set the altitude field to a perfect sphere
	mean_alt := float32(world.Circumference) / (2.0 * math.Pi)
	for lng := 0; lng < world.LongSpan; lng++ {
		for lat := 0; lat < world.LatSpan; lat++ {
			world.TileData[lng][lat].Altitude = mean_alt
		}
	}
	if debug {
		fmt.Printf("%+v\n", world.TileData[0][0])
		fmt.Printf("Mean Altitude set to %f (= %f)\n", world.TileData[0][0].Altitude, mean_alt)
	}
	return world
}

// Getters
func (w *World) Field(x, y int) *field.Field {
	// optionally do some bounds checking here
	return &w.TileData[x][y]
}

// Setter
func (w *World) SetField(x int, y int, value field.Field) {
	// optionally do some bounds checking here
	w.TileData[x][y] = value
}

func (w *World) ToGIF() {
	fmt.Println("GIF Output Here")
}
