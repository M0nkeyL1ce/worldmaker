package world

import (
	"fmt"
	"math"

	"worldmaker/field"
	"worldmaker/utility"
)

type World struct {
	TileData       [][]field.Field
	RegionCoverage RegionCoverer
	Circumference  int
	Resolution     int
	LongSpan       int
	LatSpan        int
}

type RegionCoverer struct {
	MinLevel int // the minimum cell level to be used.
	MaxLevel int // the maximum cell level to be used.
	LevelMod int // the LevelMod to be used.
	MaxCells int // the maximum desired number of cells in the approximation.
}

// Factory function
func NewWorld(circumference, resolution, max_level int, debug bool) *World {
	world := new(World)
	// Calculate the world geometry based on resolution
	long_span := int(float64(circumference) / float64(resolution))
	lat_span := int((float64(circumference)) / 2.0 / float64(resolution))
	fields := utility.Make2D[field.Field](long_span, lat_span)

	// Populate the world...
	world.TileData = fields
	world.RegionCoverage.MaxLevel = max_level
	world.Circumference = circumference
	world.Resolution = resolution
	world.LongSpan = long_span
	world.LatSpan = lat_span
	SetMeanAltitude(world, debug)
	return world
}

func SetMeanAltitude(world *World, debug bool) *World {
	// Set the altitude field to a perfect sphere
	mean_alt := float64(world.Circumference) / (2.0 * math.Pi)
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
