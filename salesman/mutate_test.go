package main

import (
	"reflect"
	"testing"
)

func TestMutate(t *testing.T) {
	steps := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sol := &solution{
		steps: make([]int, len(steps)),
	}
	copy(sol.steps, steps)

	mut := sol.Mutate()

	if reflect.DeepEqual(mut.(*solution).steps, sol.steps) {
		t.Error("Expected mutation to change something in solution")
	}
	if !reflect.DeepEqual(steps, sol.steps) {
		t.Error("Moutation shouldn't change initial solution")
	}
}
