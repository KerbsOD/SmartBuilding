package app

import "math"

type Capacity interface {
	DaysToComplete(anArea float64) int
}

func FullAreaDivision(anArea float64, anotherArea float64) int {
	return int(math.Ceil(anArea / anotherArea))
}
