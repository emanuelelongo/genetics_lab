package genetic

import (
	"math"
	"math/rand"
)

// Config contains algorithm parameters
type Config struct {
	PopulationSize   int
	CrossoverRate    float64
	MutationRate     float64
	ElitismRate      float64
	MaxGenerations   int
	ConvergenceDelay int
	LowScoreIsBetter bool
}

func coinFlip(p float64) bool {
	return rand.Float64() < p
}

// Solve function run the genetic algorithm
func Solve(config Config, random RandomIndividualBuilder, progress chan Individual) {
	betterThan := func(a, b Individual) bool {
		if config.LowScoreIsBetter {
			return a.FitnessScore() < b.FitnessScore()
		}
		return a.FitnessScore() > b.FitnessScore()
	}

	generation := 1
	notImproving := 0
	prevGen := newGeneration(config.PopulationSize, config.LowScoreIsBetter)
	eliteSize := int(math.Round(config.ElitismRate * float64(config.PopulationSize)))

	for i := 0; i < config.PopulationSize; i++ {
		prevGen.Population[i] = random()
	}
	prevGen.SetComplete()

	bestOverall := prevGen.SelectBest()
	progress <- bestOverall

	for notImproving < config.ConvergenceDelay && generation < config.MaxGenerations {
		generation++
		newGen := newGeneration(config.PopulationSize, config.LowScoreIsBetter)
		elite := prevGen.SelectElite(eliteSize)

		for i, e := range elite {
			newGen.Population[i] = e
		}

		for i := len(elite); i < config.PopulationSize; i += 2 {
			child1 := prevGen.Select()
			child2 := prevGen.Select()

			if coinFlip(config.CrossoverRate) {
				child1, child2 = child1.Crossover(child2)
			}

			if coinFlip(config.MutationRate) {
				child1 = child1.Mutate()
				child2 = child2.Mutate()
			}

			newGen.Population[i] = child1

			if i+1 < config.PopulationSize {
				newGen.Population[i+1] = child2
			}
		}
		newGen.SetComplete()

		bestInGen := newGen.SelectBest()
		if betterThan(bestInGen, bestOverall) {
			bestOverall = bestInGen
			notImproving = 0
			progress <- bestOverall
		} else {
			notImproving++
		}
	}
}
