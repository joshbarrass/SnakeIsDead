package letters

import (
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

// Cell represents a single cell containing LetterSegments
type Cell struct {
	TopLeft       [2]float64
	BottomRight   [2]float64
	Segments      []Segment
	DeathColors   [2]color.RGBA
	ParadoxColors [2]color.RGBA
	LetterColor   color.Color
}

// NewCell creates a new Cell. The color arguments take an array of
// two colors, the first is the background and the second is the
// foreground.
func NewCell(topLeft, bottomRight [2]float64, segments []Segment, deathColors, paradoxColors [2]color.RGBA) *Cell {
	// fix the order if necessary
	if topLeft[0] > bottomRight[0] {
		topLeft[0], bottomRight[0] = bottomRight[0], topLeft[0]
	}
	if topLeft[1] > bottomRight[1] {
		topLeft[1], bottomRight[1] = bottomRight[1], topLeft[1]
	}

	// return the cell
	return &Cell{
		TopLeft:       topLeft,
		BottomRight:   bottomRight,
		Segments:      segments,
		DeathColors:   deathColors,
		ParadoxColors: paradoxColors,
		LetterColor:   deathColors[0],
	}
}

// Width returns the cell's width
func (cell *Cell) Width() float64 {
	return cell.BottomRight[0] - cell.TopLeft[0]
}

// Height returns the cell's height
func (cell *Cell) Height() float64 {
	return cell.BottomRight[1] - cell.TopLeft[1]
}

// Scale increases the size of the cell, keeping the centre the same
func (cell *Cell) Scale(factor float64) {
	width := cell.Width()
	height := cell.Height()
	newWidth := width * factor
	newHeight := height * factor

	cell.TopLeft[0] = cell.TopLeft[0] + width/2 - newWidth/2
	cell.TopLeft[1] = cell.TopLeft[1] + height/2 - newHeight/2
	cell.BottomRight[0] = cell.BottomRight[0] - width/2 + newWidth/2
	cell.BottomRight[1] = cell.BottomRight[1] - height/2 + newHeight/2
}

// Translate moves the cell on the screen
func (cell *Cell) Translate(x, y float64) {
	cell.TopLeft[0] += x
	cell.TopLeft[1] += y
	cell.BottomRight[0] += x
	cell.BottomRight[1] += y
}

// Draw calls the Draw method for all of its segments for them to
// render onto it. The image is then rendered onto the main context.
func (cell *Cell) Draw(gc *draw2dimg.GraphicContext) bool {
	hasChanged := false
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0x00})
	for _, seg := range cell.Segments {
		if seg.Draw(NewCellDrawer(gc, cell, seg), cell) {
			hasChanged = true
		}
	}
	return hasChanged
}
