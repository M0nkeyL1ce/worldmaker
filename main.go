package main

import (
	"flag"
	"fmt"
	"worldmaker/world"
)

const DefaultCircumference = 24901
const DefaultResolution = 1
const DefaultUnit = "mile"
const DefaultDebug = true

func main() {
	var circumference = *flag.Int("c", DefaultCircumference, "size of the world in user-defined units")
	var resolution = *flag.Int("r", DefaultResolution, "resolution for rendering world")
	var unit = *flag.String("u", DefaultUnit, "unit for the circumference")
	var debug = *flag.Bool("d", DefaultDebug, "print debugging information")
	flag.Parse()

	fmt.Printf("circumference: %d %s(s) with resolution of %d %s(s)\n", circumference, unit, resolution, unit)

	world := world.NewWorld(circumference, resolution, debug)
	if debug {
		fmt.Printf("HTMDepth: %d\n", world.HtmDepth)
		fmt.Printf("Circumference: %d\n", world.Circumference)
		fmt.Printf("Resolution: %d\n", world.Resolution)
	}

	world.ToGIF()
}
