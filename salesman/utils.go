package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Point is a 2D point with X and Y
type Point struct {
	X float64
	Y float64
}

func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func shuffleSlice(s []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

func findIndex(slice []int, value int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return i
		}
	}
	return -1
}

func rotateAbout(point Point, center Point, angle float64) Point {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	point.X -= center.X
	point.Y -= center.Y

	rotated := Point{
		X: point.X*cos - point.Y*sin,
		Y: point.X*sin + point.Y*cos,
	}

	rotated.X += center.X
	rotated.Y += center.Y
	return rotated
}

func logSolution(n string, s solution) {
	fmt.Printf("%s: ", n)
	for i := 0; i < len(s.steps); i++ {
		fmt.Printf("%d, ", s.steps[i])
	}
	fmt.Println()
}
