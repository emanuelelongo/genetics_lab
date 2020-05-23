package genetic

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type fakeIndividual struct {
	fitnessScore float64
}

func (i *fakeIndividual) FitnessScore() float64 {
	return i.fitnessScore
}

func (i *fakeIndividual) Crossover(other Individual) (Individual, Individual) {
	return i, i
}

func (i *fakeIndividual) Mutate() Individual {
	return i
}

func TestSelectWhenHighScoreIsBetter(t *testing.T) {
	const EPSILON = 1
	const N = 10000

	// create a generation with 71 individuals:
	// 	70 with a score of  1
	//   1 with a score of 30
	gen := Generation{
		LowScoreBetter: false,
		Population:     make([]Individual, 71),
	}

	s30 := fakeIndividual{fitnessScore: 30}
	s1 := fakeIndividual{fitnessScore: 1}
	gen.Population[0] = &s30
	for i := 1; i < 71; i++ {
		gen.Population[i] = &s1
	}

	gen.SetComplete()

	selectS30Count := .0
	for i := 0; i < N; i++ {
		sel := gen.Select()
		if math.Abs(sel.FitnessScore()) == 30 {
			selectS30Count++
		}
	}
	selectS30Rate := (selectS30Count / N) * 100

	// in such population the s30 individual is expected to be selected with a probabilty of 30%
	if math.Abs(selectS30Rate-30) > EPSILON {
		t.Error(fmt.Sprintf("Expected a rate of 30%% but detected %f.", selectS30Rate))
	}
}

func TestSelectWhenLowScoreIsBetter(t *testing.T) {
	const EPSILON = 1
	const N = 10000

	// create a generation with 6 individuals:
	// 	1 with a score of  5
	//  2 with a score of 10
	//  3 with a score of 25
	gen := Generation{
		LowScoreBetter: true,
		Population:     make([]Individual, 6),
	}
	s5 := fakeIndividual{fitnessScore: 5}
	s10 := fakeIndividual{fitnessScore: 10}
	s25 := fakeIndividual{fitnessScore: 25}
	gen.Population[0] = &s5
	gen.Population[1] = &s10
	gen.Population[2] = &s10
	gen.Population[3] = &s25
	gen.Population[4] = &s25
	gen.Population[5] = &s25

	gen.SetComplete()

	selectS5Count := .0
	for i := 0; i < N; i++ {
		sel := gen.Select()
		if math.Abs(sel.FitnessScore()) == 5 {
			selectS5Count++
		}
	}
	selectS5Rate := (selectS5Count / N) * 100

	// in such population the s5 individual is expected to be selected with a probabilty of 38%
	if math.Abs(selectS5Rate-38) > EPSILON {
		t.Error(fmt.Sprintf("Expected a rate of 38%% but detected %f.", selectS5Rate))
	}
}

func TestSelectElite(t *testing.T) {
	gen := Generation{
		LowScoreBetter: false,
		Population:     make([]Individual, 20),
	}
	for i := 0; i < 20; i++ {
		gen.Population[i] = &fakeIndividual{fitnessScore: rand.Float64()}
	}

	gen.SetComplete()
	eliteSize := 6
	elite := gen.SelectElite(eliteSize)

	if len(elite) != eliteSize {
		t.Error(fmt.Sprintf("Expected an elite of %d individuals but detected %d.", eliteSize, len(elite)))
	}

	for i := range elite {
		if elite[i] != gen.Population[i] {
			t.Error(fmt.Sprintf("Expected individual n. %d to be included in the elite", i))
		}
	}
}
