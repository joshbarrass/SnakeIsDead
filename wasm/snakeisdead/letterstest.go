package main

import (
	"fmt"
	"log"
	"strings"
	"syscall/js"

	"github.com/joshbarrass/SnakeIsDead/pkg/letters"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

// LetterSpacing is the space in pixels between the left edge of
// letters
const LetterSpacing = 100

func getCells(text string, topLeft [2]float64) ([]*letters.Cell, error) {
	letterMap := letters.GetLetterMap()
	cells := []*letters.Cell{}
	for i, char := range text {
		letterFunc, ok := letterMap[byte(char)]
		if !ok {
			return nil, fmt.Errorf("character '%s' not available", string(char))
		}
		cells = append(cells, letters.NewCell(
			[2]float64{topLeft[0] + float64(LetterSpacing*i), topLeft[1]},
			[2]float64{topLeft[0] + float64(LetterSpacing*i) + letters.DefaultWidth, topLeft[1] + letters.DefaultHeight},
			letterFunc(),
			letters.ColorsDeath,
			letters.ColorsParadox,
		))
	}
	return cells, nil
}

// getDrawingFunction returns a function for drawing the cells for a
// phrase
func getDrawingFunction(phrase string) (func(*draw2dimg.GraphicContext) bool, error) {
	cells, err := getCells(phrase, [2]float64{20, 20})
	if err != nil {
		return nil, err
	}
	return func(gc *draw2dimg.GraphicContext) bool {
		// fill background
		gc.SetFillColor(cells[0].DeathColors[0])
		gc.MoveTo(0, 0)
		gc.LineTo(0, 10000)
		gc.LineTo(10000, 10000)
		gc.LineTo(10000, 0)
		gc.LineTo(0, 0)
		gc.Fill()
		// draw cell
		for _, cell := range cells {
			cell.Draw(gc)
		}
		return true
	}, nil
}

// Updater is a struct responsible for storing the drawing function
// and wrapping it, allowing it to be updated more easily
type Updater struct {
	Func func(*draw2dimg.GraphicContext) bool
}

// Draw wraps the drawing function
func (updater *Updater) Draw(gc *draw2dimg.GraphicContext) bool {
	return updater.Func(gc)
}

// SetFunc updates the drawing function
func (updater *Updater) SetFunc(f func(*draw2dimg.GraphicContext) bool) {
	updater.Func = f
}

func main() {
	c := make(chan struct{})
	fmt.Println("WASM Go Initialised")

	// create a canvas
	cvs, err := canvas.NewCanvas2d(true)
	if err != nil {
		log.Fatalf("Could not create canvas")
	}

	drawFunc, err := getDrawingFunction("SNAKE IS DEAD")
	if err != nil {
		panic(fmt.Sprintf("failed to get drawing function: %s", err))
	}
	updater := Updater{
		Func: drawFunc,
	}

	js.Global().Set("UpdatePhrase", js.FuncOf(
		func(this js.Value, i []js.Value) interface{} {
			if len(i) != 1 {
				return map[string]interface{}{
					"error": "wrong number of arguments",
				}
			}
			phrase := i[0].String()
			newFunc, err := getDrawingFunction(strings.ToUpper(phrase))
			if err != nil {
				return map[string]interface{}{
					"error": err.Error(),
				}
			}
			updater.Func = newFunc
			return map[string]interface{}{}
		},
	))

	cvs.Start(30, updater.Draw)

	// channel is unused, so this prevents main from terminating and killing the WASM
	<-c
}
