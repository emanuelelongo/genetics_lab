package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCrossover(t *testing.T) {
	cities := make([]Point, 9)
	problem := &Problem{
		cities: cities,
	}
	parent1 := &solution{problem: problem, steps: []int{1, 2, 3, 4, 5, 6, 7, 8}}
	parent2 := &solution{problem: problem, steps: []int{8, 5, 2, 1, 3, 6, 4, 7}}

	child1, child2 := parent1.Crossover(parent2)

	if !reflect.DeepEqual(child1.(*solution).steps, []int{1, 5, 2, 4, 3, 6, 7, 8}) {
		t.Error(fmt.Sprintf("Crossover failed, child1 is %d", child1.(*solution).steps))
	}
	if !reflect.DeepEqual(child2.(*solution).steps, []int{8, 2, 3, 1, 5, 6, 4, 7}) {
		t.Error(fmt.Sprintf("Crossover failed, child1 is %d", child2.(*solution).steps))
	}
}
