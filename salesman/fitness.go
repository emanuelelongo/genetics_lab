package main

func (s *solution) FitnessScore() float64 {
	if s.fitnessScore > 0 {
		return s.fitnessScore
	}
	n := len(s.steps)
	s.fitnessScore = distance(s.problem.cities[0], s.problem.cities[s.steps[0]])
	for i := 0; i < n-1; i++ {
		s.fitnessScore += distance(s.problem.cities[s.steps[i]], s.problem.cities[s.steps[i+1]])
	}
	s.fitnessScore += distance(s.problem.cities[s.steps[n-1]], s.problem.cities[0])
	return s.fitnessScore
}
