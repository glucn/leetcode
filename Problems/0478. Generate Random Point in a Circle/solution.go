package _478__Generate_Random_Point_in_a_Circle

import "math/rand"
import "math"

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */

// Runtime 64 ms, faster than 83.33%
// Memory 16.8 MB, less than 100%

type Solution struct {
	radius  float64
	xCenter float64
	yCenter float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius:  radius,
		xCenter: x_center,
		yCenter: y_center,
	}
}

func (s *Solution) RandPoint() []float64 {
	// rand.Float64() * s.radius is wrong
	// Refer to http://www.anderswallin.net/2009/05/uniform-random-points-in-a-circle-using-polar-coordinates/
	r := math.Sqrt(rand.Float64()) * s.radius
	a := rand.Float64() * math.Pi * 2
	return []float64{s.xCenter + r*math.Sin(a), s.yCenter + r*math.Cos(a)}
}
