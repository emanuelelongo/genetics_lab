package main

import "emanuelelongo/genetics_lab/genetic"

type solution struct {
	problem      *Problem
	steps        []int
	fitnessScore float64
}

func randomSolution(p *Problem) func() genetic.Individual {
	// the real steps count would be len(p.cities)
	// but for convenience the last step - which is always
	// the initial city - is omitted from representation
	stepsCount := len(p.cities) - 1
	return func() genetic.Individual {
		ind := solution{
			problem: p,
			steps:   make([]int, stepsCount),
		}
		for i := 0; i < stepsCount; i++ {
			ind.steps[i] = i + 1
		}

		shuffleSlice(ind.steps)
		return &ind
	}
}
