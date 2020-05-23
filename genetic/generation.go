package genetic

import (
	"math/rand"
	"sort"
)

// Generation is a single iteration of the evolution
type Generation struct {
	LowScoreBetter  bool
	Population      []Individual
	Complete        bool
	scoreSum        float64
	inverseScoreSum float64
}

func newGeneration(popSize int, lowScoreBetter bool) *Generation {
	return &Generation{
		LowScoreBetter:  lowScoreBetter,
		Population:      make([]Individual, popSize),
		Complete:        false,
		scoreSum:        0,
		inverseScoreSum: 0,
	}
}

// SelectBest return the best individual in this generation
func (g *Generation) SelectBest() Individual {
	if !g.Complete {
		panic("Can't select on uncompleted generation")
	}
	return g.Population[0]
}

// Select return a random individual to be included in the next generation
func (g *Generation) Select() Individual {
	if g.LowScoreBetter {
		return g.selectWhenLowScoreBetter()
	}
	return g.selectWhenHigherBetter()
}

func (g *Generation) selectWhenHigherBetter() Individual {
	random := rand.Float64() * g.scoreSum
	for i := 0; i < len(g.Population); i++ {
		random -= g.Population[i].FitnessScore()
		if random <= 0 {
			return g.Population[i]
		}
	}
	return g.Population[len(g.Population)-1]
}

func (g *Generation) selectWhenLowScoreBetter() Individual {
	random := rand.Float64() * g.inverseScoreSum
	for i := 0; i < len(g.Population); i++ {
		random -= g.scoreSum / g.Population[i].FitnessScore()
		if random <= 0 {
			return g.Population[i]
		}
	}
	return g.Population[len(g.Population)-1]
}

// SelectElite return the elite of best individuals to be included in the next generation
func (g *Generation) SelectElite(eliteSize int) []Individual {
	if !g.Complete {
		panic("Can't select on uncompleted generation")
	}
	return g.Population[:eliteSize]
}

// SetComplete set the generation as complete
func (g *Generation) SetComplete() {
	if g.Complete {
		panic("Generation has already been completed")
	}

	for _, i := range g.Population {
		g.scoreSum += i.FitnessScore()
	}

	if g.LowScoreBetter {
		sort.Slice(g.Population, func(i, j int) bool {
			return g.Population[i].FitnessScore() < g.Population[j].FitnessScore()
		})
		g.inverseScoreSum = 0
		for _, i := range g.Population {
			g.inverseScoreSum += g.scoreSum / i.FitnessScore()
		}
	} else {
		sort.Slice(g.Population, func(i, j int) bool {
			return g.Population[i].FitnessScore() > g.Population[j].FitnessScore()
		})
	}

	g.Complete = true
}
