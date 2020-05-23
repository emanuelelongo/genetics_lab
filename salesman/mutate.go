package main

import (
	"emanuelelongo/genetics_lab/genetic"
	"math/rand"
)

func (s *solution) Mutate() genetic.Individual {
	mut := &solution{
		problem:      s.problem,
		steps:        make([]int, len(s.steps)),
		fitnessScore: s.fitnessScore,
	}
	copy(mut.steps, s.steps)

	i := rand.Intn(len(mut.steps))
	j := rand.Intn(len(mut.steps))
	mut.steps[i], mut.steps[j] = mut.steps[j], mut.steps[i]
	return mut
}
