package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var simpleScoreOut = widget.NewLabel("-")
var currenGeneration = widget.NewLabel("-")
var geneticScoreOut = widget.NewLabel("-")
var spacer = widget.NewLabel("")
var citiesLabel, populationLabel, maxGenLabel, crossoverRateLabel, mutationRateLabel, elitismRateLabel *widget.Label
var cities, population, maxGen, crossoverRate, mutationRate, elitismRate *widget.Slider

func createSlider(text string, min, max, initial float64) (*widget.Label, *widget.Slider) {
	label := widget.NewLabel(text)
	slider := widget.NewSlider(min, max)
	slider.OnChanged = func(v float64) { label.SetText(fmt.Sprintf("%s [%d]", text, int(v))) }
	slider.Value = initial
	return label, slider
}

func buildControls() *fyne.Container {
	citiesLabel, cities = createSlider("Cities", 10, 100, 10)
	populationLabel, population = createSlider("Population", 100, 100000, 1000)
	maxGenLabel, maxGen = createSlider("Max Generations", 10, 100000, 1000)
	crossoverRateLabel, crossoverRate = createSlider("Crossover Rate", 0, 100, 20)
	mutationRateLabel, mutationRate = createSlider("Mutation Rate", 0, 100, 20)
	elitismRateLabel, elitismRate = createSlider("Elitism Rate", 0, 100, 10)

	return fyne.NewContainerWithLayout(layout.NewFormLayout(),

		citiesLabel, cities,

		widget.NewButton("New circular", onNewCircularProblemPressed),
		widget.NewButton("New random", onNewRandomProblemPressed),

		spacer, spacer,

		widget.NewLabelWithStyle("Simple algorithm", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		spacer,

		widget.NewButton("Solve", onSimpleSolveClick),
		simpleScoreOut,

		spacer, spacer,

		widget.NewLabelWithStyle("Genetic algorithm", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		spacer,

		populationLabel, population,

		maxGenLabel, maxGen,

		crossoverRateLabel, crossoverRate,

		mutationRateLabel, mutationRate,

		elitismRateLabel, elitismRate,

		widget.NewButton("Solve", onGeneticSolveClick),
		geneticScoreOut,

		spacer,
		currenGeneration,

		spacer, spacer,

		widget.NewButton("Clear", onClearPressed),
		spacer,
	)
}
