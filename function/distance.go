package function

import "math"

// this function gets the distance of two coordinates
func Distance(x1, y1, x2, y2 int) int {
	var result float64
	x := (x2 - x1)
	y := (y2 - y1)
	result = math.Sqrt(float64((x * x) + (y * y)))
	return int(result)
}
