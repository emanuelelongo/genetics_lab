package main

import (
	"emanuelelongo/genetics_lab/genetic"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"github.com/fogleman/gg"
)

var ctx *gg.Context
var raster *canvas.Raster
var currentProblem Problem
var progress chan genetic.Individual

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Drawing")
	progress = make(chan genetic.Individual, 20)

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
	go drawer()
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
	progress <- s
}

func drawer() {
	for s := range progress {
		clear()
		drawSolution(currentProblem, s.(*solution))
		raster.Refresh()
	}
	// doesn't draw more than a solution in 100ms (this is unlikely to happen)
	// ticker := time.NewTicker(time.Millisecond * 100)
	// defer ticker.Stop()
	// canDraw := false

	// for {
	// 	select {
	// 	case s := <-progress:
	// 		if canDraw {
	// 			canDraw = false
	// 			onClearPressed()
	// 			drawSolution(currentProblem, s.(*solution))
	// 			raster.Refresh()
	// 		}
	// 	case <-ticker.C:
	// 		canDraw = true
	// 	default:
	// 	}
	// }
}

func onGeneticSolveClick() {
	config := genetic.Config{
		CrossoverRate:    crossoverRate.Value / 100,
		ElitismRate:      elitismRate.Value / 100,
		LowScoreIsBetter: true,
		MaxGenerations:   int(maxGen.Value),
		MutationRate:     mutationRate.Value / 100,
		ConvergenceDelay: int(convergenceDelay.Value),
		PopulationSize:   int(population.Value),
	}

	go genetic.Solve(config, randomSolution(&currentProblem), progress)
}

func onStopClick() {

}

func clear() {
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
	logSolution("Drawing solution:", *s)
	scoreLabel.SetText(fmt.Sprintf("%.2f", s.FitnessScore()))

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
