package letters

import (
	"image/color"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

// CellDrawer represents a type that contains all of the drawing
// functions needed by a segment in order to draw within the cell
type CellDrawer interface {
	Fill(paths ...*draw2d.Path)
	FillStroke(paths ...*draw2d.Path)
	Stroke(paths ...*draw2d.Path)
	Close()
	SetFillColor(c color.Color)
	SetStrokeColor(c color.Color)
	MoveTo(x, y float64)
	LineTo(x, y float64)
}

// cellDrawer is the actual struct that implements the CellDrawer interface
type cellDrawer struct {
	GraphicContext *draw2dimg.GraphicContext
	CellTopLeft    [2]float64
	CellWidth      float64
	CellHeight     float64
	SegmentWidth   float64
	SegmentHeight  float64
}

// NewCellDrawer creates a CellDrawer for allowing a segment to draw
// within a cell. It allows the Cell and the segment's reference sizes
// to be different, and will translate between one and the other to
// simplify the construction of segment drawing procedures.
func NewCellDrawer(gc *draw2dimg.GraphicContext, cell *Cell, segment Segment) CellDrawer {
	return &cellDrawer{
		GraphicContext: gc,
		CellTopLeft:    cell.TopLeft,
		CellWidth:      cell.Width(),
		CellHeight:     cell.Height(),
		SegmentWidth:   segment.Width(),
		SegmentHeight:  segment.Height(),
	}
}

// Fill wraps GraphicContext.Fill
func (cd *cellDrawer) Fill(paths ...*draw2d.Path) {
	cd.GraphicContext.Fill(paths...)
}

// FillStroke wraps GraphicContext.FillStroke
func (cd *cellDrawer) FillStroke(paths ...*draw2d.Path) {
	cd.GraphicContext.FillStroke(paths...)
}

// Stroke wraps GraphicContext.Stroke
func (cd *cellDrawer) Stroke(paths ...*draw2d.Path) {
	cd.GraphicContext.Stroke(paths...)
}

// Close wraps GraphicContext.Close
func (cd *cellDrawer) Close() {
	cd.GraphicContext.Close()
}

// SetFillColor wraps GraphicContext.SetFillColor
func (cd *cellDrawer) SetFillColor(c color.Color) {
	cd.GraphicContext.SetFillColor(c)
}

// SetStrokeColor wraps GraphicContext.SetStrokeColor
func (cd *cellDrawer) SetStrokeColor(c color.Color) {
	cd.GraphicContext.SetStrokeColor(c)
}

// ConvertCoords converts coordinates from the segment system to the global cell system
func (cd *cellDrawer) ConvertCoords(x, y float64) (newX, newY float64) {
	// (x, y) is in the segment coordinate system
	// segment top left = (0,0), segment bottom right = (width-1, height-1)
	u := x / (cd.SegmentWidth - 1)
	v := y / (cd.SegmentHeight - 1)
	newX = cd.CellTopLeft[0] + u*(cd.CellWidth-1)
	newY = cd.CellTopLeft[1] + v*(cd.CellHeight-1)
	return
}

// MoveTo wraps GraphicContext.MoveTo, converting between coordinate systems in the process
func (cd *cellDrawer) MoveTo(x, y float64) {
	newX, newY := cd.ConvertCoords(x, y)
	cd.GraphicContext.MoveTo(newX, newY)
}

// LineTo wraps GraphicContext.LineTo, converting between coordinate systems in the process
func (cd *cellDrawer) LineTo(x, y float64) {
	newX, newY := cd.ConvertCoords(x, y)
	cd.GraphicContext.LineTo(newX, newY)
}
