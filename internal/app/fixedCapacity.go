package app

import (
	"math"
)

type FixedCapacity struct {
	dailyCapacity float64
}

func NewFixedCapacity(aCapacity float64) *FixedCapacity {
	fc := new(FixedCapacity)
	fc.dailyCapacity = aCapacity
	return fc
}

func (fc FixedCapacity) DaysToComplete(anArea float64) int {
	return int(math.Ceil(anArea / fc.dailyCapacity))
}
