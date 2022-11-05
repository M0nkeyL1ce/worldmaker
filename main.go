package main

import (
	"flag"
	"fmt"
	"math"
	"worldmaker/utility"
	"worldmaker/world"

	s2 "github.com/golang/geo/s2"
)

const DefaultCircumference = 24901
const DefaultResolution = 1
const DefaultUnit = "mile"
const DefaultDebug = true
const DefaultMaxLevel = 12

func main() {
	var circumference = *flag.Int("c", DefaultCircumference, "size of the world in user-defined units")
	var resolution = *flag.Int("r", DefaultResolution, "resolution for rendering world")
	//var unit = *flag.String("u", DefaultUnit, "unit for the circumference")
	var max_level = *flag.Int("l", DefaultMaxLevel, "S2 Maximum Level")
	var debug = *flag.Bool("d", DefaultDebug, "print debugging information")
	flag.Parse()

	world := world.NewWorld(circumference, resolution, max_level, debug)
	if debug {
		fmt.Printf("MaxLevel: %d\n", world.RegionCoverage.MaxLevel)
		fmt.Printf("Circumference: %d\n", world.Circumference)
		fmt.Printf("Resolution: %d\n", world.Resolution)
		fmt.Printf("MaxLevel: %d\n", world.RegionCoverage.MaxLevel)
	}

	if debug {
		// Bloomington Indiana = 39°09′44″N 86°31′45″W
		latitude := utility.DegMinSecToDecimal(39, 9, 44)      // North = Positive
		longitude := utility.DegMinSecToDecimal(-86, -31, -45) // West = Negative

		ll := s2.LatLngFromDegrees(latitude, longitude)
		cid := s2.CellIDFromLatLng(ll)
		fmt.Printf("Latlng: %s\n\n", ll)
		fmt.Printf("CellID: %s (Level %d, Latlng %s)\n", cid, cid.Level(), cid.LatLng())
		pcid := cid.Parent(world.RegionCoverage.MaxLevel)
		fmt.Printf("CellID: %s (Level %d, Latlng %s)\n", pcid, pcid.Level(), pcid.LatLng())
		pl := make(s2.Polyline, 2)
		pl[0] = cid.Point()
		pl[1] = pcid.Point()
		arc_d := float64(pl.Length()) * float64(circumference) / (2.0 * math.Pi)
		deg := float64(pl.Length()) / (2.0 * math.Pi) * 360.0
		fmt.Printf("Distance Between: %f radians\n", pl.Length())
		fmt.Printf("Distance Between: %f degrees\n", deg)
		fmt.Printf("Distance Between: %f miles\n", arc_d)
	}

	world.ToGIF()
}
