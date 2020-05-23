package genetic

// Individual is a candidate solution
type Individual interface {
	FitnessScore() float64
	Crossover(other Individual) (Individual, Individual)
	Mutate() Individual
}

// RandomIndividualBuilder creates a random individual
type RandomIndividualBuilder func() Individual
