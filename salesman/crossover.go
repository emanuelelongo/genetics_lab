package main

import (
	"emanuelelongo/genetics_lab/genetic"
)

func (s *solution) Crossover(other genetic.Individual) (genetic.Individual, genetic.Individual) {
	child1 := &solution{
		problem: s.problem,
		steps:   make([]int, len(s.steps)),
	}
	child2 := &solution{
		problem: s.problem,
		steps:   make([]int, len(s.steps)),
	}

	i := 0
	for child1.steps[i] == 0 {
		child1.steps[i] = s.steps[i]
		i = findIndex(s.steps, other.(*solution).steps[i])
	}
	for i := 0; i < len(child1.steps); i++ {
		if child1.steps[i] == 0 {
			child1.steps[i] = other.(*solution).steps[i]
		}
	}

	i = 0
	for child2.steps[i] == 0 {
		child2.steps[i] = other.(*solution).steps[i]
		i = findIndex(other.(*solution).steps, s.steps[i])
	}
	for i := 0; i < len(child2.steps); i++ {
		if child2.steps[i] == 0 {
			child2.steps[i] = s.steps[i]
		}
	}
	return child1, child2
}
