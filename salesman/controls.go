package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var currenGeneration = widget.NewLabel("-")
var scoreLabel = widget.NewLabel("-")
var spacer = widget.NewLabel("")
var citiesLabel, populationLabel, maxGenLabel, convergenceDelayLabel, crossoverRateLabel, mutationRateLabel, elitismRateLabel *widget.Label
var cities, population, maxGen, convergenceDelay, crossoverRate, mutationRate, elitismRate *widget.Slider

func createSlider(text string, min, max, initial int) (*widget.Label, *widget.Slider) {
	label := widget.NewLabel(fmt.Sprintf("%s\n%d", text, initial))
	label.Alignment = fyne.TextAlignTrailing
	slider := widget.NewSlider(float64(min), float64(max))
	slider.OnChanged = func(v float64) { label.SetText(fmt.Sprintf("%s\n%d", text, int(v))) }
	slider.Value = float64(initial)
	return label, slider
}

func buildControls() *fyne.Container {
	citiesLabel, cities = createSlider("Cities", 10, 100, 10)
	populationLabel, population = createSlider("Population", 100, 100000, 1000)
	// todo: remove max generation when the stop feature is implemented
	maxGenLabel, maxGen = createSlider("Max Generations", 10, 100000, 1000)
	convergenceDelayLabel, convergenceDelay = createSlider("Convergence delay", 20, 1000, 50)
	crossoverRateLabel, crossoverRate = createSlider("Crossover Rate", 0, 100, 20)
	mutationRateLabel, mutationRate = createSlider("Mutation Rate", 0, 100, 20)
	elitismRateLabel, elitismRate = createSlider("Elitism Rate", 0, 100, 10)

	return fyne.NewContainerWithLayout(layout.NewFormLayout(),

		citiesLabel, cities,

		widget.NewButton("New circular", onNewCircularProblemPressed),
		widget.NewButton("New random", onNewRandomProblemPressed),

		spacer, spacer,

		spacer, spacer,

		widget.NewLabelWithStyle("Genetic algorithm", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		spacer,

		populationLabel, population,

		convergenceDelayLabel, convergenceDelay,

		maxGenLabel, maxGen,

		crossoverRateLabel, crossoverRate,

		mutationRateLabel, mutationRate,

		elitismRateLabel, elitismRate,

		widget.NewButton("Solve", onGeneticSolveClick),
		scoreLabel,

		widget.NewButton("Stop", onStopClick),
		spacer,

		spacer,
		currenGeneration,

		spacer, spacer,

		widget.NewButton("Clear", clear),
		spacer,
	)
}
