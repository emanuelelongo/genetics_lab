package main

import (
	"emanuelelongo/genetics_lab/genetic"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/fogleman/gg"
)

var ctx *gg.Context
var raster *canvas.Raster
var currentProblem Problem

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Drawing")

	ctx = gg.NewContext(1000, 1000)
	ctx.SetRGB(.1, .1, .1)
	raster = canvas.NewRasterFromImage(ctx.Image())
	controls := buildControls()
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(nil, nil, controls, nil),
			controls,
			raster,
		),
	)

	window.Resize(fyne.NewSize(1000, 1000))
	window.ShowAndRun()
}

func onNewCircularProblemPressed() {
	ctx.SetRGB(.1, .1, .1)
	ctx.Clear()
	currentProblem = newCircularProblem(ctx.Width(), ctx.Height(), int(cities.Value))
	drawProblem(currentProblem)
	raster.Refresh()
}

func onNewRandomProblemPressed() {
	ctx.SetRGB(.1, .1, .1)
	ctx.Clear()
	currentProblem = newRandomProblem(ctx.Width(), ctx.Height(), int(cities.Value))
	drawProblem(currentProblem)
	raster.Refresh()
}

func onSimpleSolveClick() {
	s := randomSolution(&currentProblem)()
	drawSolution(currentProblem, s.(*solution))
	raster.Refresh()
}

func onGeneticSolveClick() {

	config := genetic.Config{
		CrossoverRate:         crossoverRate.Value / 100,
		ElitismRate:           elitismRate.Value / 100,
		LowScoreIsBetter:      true,
		MaxGenerations:        int(maxGen.Value),
		MutationRate:          mutationRate.Value / 100,
		NotImprovingThreshold: 5,
		PopulationSize:        int(population.Value),
	}

	s := genetic.Solve(config, randomSolution(&currentProblem))
	sol := s.(*solution)
	drawSolution(currentProblem, sol)
	raster.Refresh()
}

func onClearPressed() {
	ctx.SetRGB(.1, .1, .1)
	ctx.Clear()
	drawProblem(currentProblem)
	raster.Refresh()
}

func drawProblem(problem Problem) {
	ctx.SetRGB(.8, 0, 0)
	ctx.SetLineWidth(16)
	for i, city := range problem.cities {
		if i == 0 {
			ctx.SetRGB(.8, .8, 0)
			ctx.DrawPoint(city.X, city.Y, 2)
			ctx.Stroke()
			ctx.SetRGB(.8, 0, 0)
			continue
		}
		ctx.DrawPoint(city.X, city.Y, 2)
		ctx.Stroke()
	}
}

func drawSolution(p Problem, s *solution) {
	logSolution("Final solution", *s)

	ctx.SetRGB(0, .5, 0)
	ctx.SetLineWidth(4)
	n := len(s.steps)
	from := p.cities[0]
	for i := 0; i < n; i++ {
		to := p.cities[s.steps[i]]
		ctx.DrawLine(from.X, from.Y, to.X, to.Y)
		ctx.Stroke()
		from = to
	}
	ctx.DrawLine(p.cities[s.steps[n-1]].X, p.cities[s.steps[n-1]].Y, p.cities[0].X, p.cities[0].Y)
	ctx.Stroke()
}
