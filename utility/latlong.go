package utility

func DegMinSecToDecimal(degrees, minutes, seconds float64) float64 {
	return degrees + minutes/60.0 + seconds/3600.0
}
