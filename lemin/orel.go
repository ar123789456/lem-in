package lemin

import (
	"math"
)

func Orel(end, start *Honeycomb) float64 {
	x := float64(end.Coor[0]) - float64(start.Coor[0])
	y := float64(end.Coor[1]) - float64(start.Coor[1])
	return math.Sqrt(x*x + y*y)
}
