package main

import (
	"fmt"
	"log"

	"github.com/joshbarrass/SnakeIsDead/pkg/letters"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

var cell = letters.NewCell([2]float64{20, 20}, [2]float64{20 + letters.DefaultWidth, 20 + letters.DefaultHeight}, []letters.Segment{
	&letters.SegmentARisingStick{},
	&letters.SegmentAVert{},
	&letters.SegmentABar{},
}, letters.ColorsDeath, letters.ColorsParadox)

func testDraw(gc *draw2dimg.GraphicContext) bool {
	// fill background
	gc.SetFillColor(cell.DeathColors[0])
	gc.MoveTo(0, 0)
	gc.LineTo(0, 10000)
	gc.LineTo(10000, 10000)
	gc.LineTo(10000, 0)
	gc.LineTo(0, 0)
	gc.Fill()
	// draw cell
	cell.Draw(gc)
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
