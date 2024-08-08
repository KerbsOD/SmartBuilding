package app

type Capacity interface {
	DaysToComplete(anArea float64) int
}
