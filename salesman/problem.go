package main

import (
	"math"
	"math/rand"
)

// Problem is an instance of the Traveling Salesman Problem
// the starting city is the first in the array by convention
type Problem struct {
	width  int
	height int
	cities []Point
}

func newCircularProblem(width, height, citiesCount int) Problem {
	const margin = 80
	prob := Problem{
		width:  width,
		height: height,
	}

	center := Point{X: float64(width / 2), Y: float64(height / 2)}
	angle := 2 * math.Pi / float64(citiesCount)

	p := Point{X: margin, Y: center.Y}
	for i := 0; i < citiesCount; i++ {
		prob.cities = append(prob.cities, p)
		p = rotateAbout(p, center, angle)
	}
	return prob
}

func newRandomProblem(width, height, citiesCount int) Problem {
	const margin = 80

	prob := Problem{
		width:  width,
		height: height,
	}
	for i := 0; i < citiesCount; i++ {
		prob.cities = append(prob.cities, Point{X: float64(rand.Intn(width - margin)), Y: float64(rand.Intn(height - margin))})
	}
	return prob
}
