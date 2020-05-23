package main

import (
	"fmt"
	"testing"
)

func TestFitnessScore(t *testing.T) {

	cities := []Point{{X: 0, Y: 0}, {X: 0, Y: 5}}
	problem := &Problem{
		cities: cities,
	}

	s := &solution{problem: problem, steps: []int{1}}

	score := s.FitnessScore()
	expected := 10.0
	if score != expected {
		t.Error(fmt.Sprintf("Expected a fitness score of %f but found %f", expected, score))
	}
}
