package main

import (
	"fmt"
	"log"

	"github.com/joshbarrass/SnakeIsDead/pkg/letters"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

const LetterSpacing = 100

func getCells(text string, topLeft [2]float64) ([]*letters.Cell, error) {
	letterMap := letters.GetLetterMap()
	cells := []*letters.Cell{}
	for i, char := range text {
		letterFunc, ok := letterMap[byte(char)]
		if !ok {
			return nil, fmt.Errorf("character '%s' not available", char)
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

var cells, _ = getCells("SNA", [2]float64{20, 20})

func testDraw(gc *draw2dimg.GraphicContext) bool {
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
}

func main() {
	c := make(chan struct{})
	fmt.Println("WASM Go Initialised")

	// create a canvas
	cvs, err := canvas.NewCanvas2d(true)
	if err != nil {
		log.Fatalf("Could not create canvas")
	}

	cvs.Start(30, testDraw)

	// channel is unused, so this prevents main from terminating and killing the WASM
	<-c
}
