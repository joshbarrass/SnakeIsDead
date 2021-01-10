package letters

import (
	"image/color"
)

// Default dimensions of the original segment's cell
const (
	DefaultWidth  float64 = 51.
	DefaultHeight float64 = 85.
)

// Letter Thickness Guides
const (
	stickThickness float64 = 18
)

// Original Colours
var (
	ColorsDeath = [2]color.RGBA{
		{0x31, 0x34, 0x16, 0xff},
		{0x95, 0x97, 0x7E, 0xff},
	}
	ColorsParadox = [2]color.RGBA{
		{0x50, 0x41, 0x39, 0xff},
		{0x00, 0x00, 0x00, 0xff},
	}
)

// Segment IDs
const (
	IDSUpperBar SegmentID = iota
	IDARisingStick
	IDAVert
	IDABar
)

// Letter defines a custom type for full letters. A letter is defined
// as a slice of segments.
type Letter []Segment

/* Letter S */

// SegmentSUpperBar is the top bar on an S
type SegmentSUpperBar struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentSUpperBar) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width := seg.Width()
	gc.MoveTo(0, 0)
	gc.LineTo(width-1, 0)
	gc.LineTo(width-1, stickThickness)
	gc.LineTo(0, stickThickness)
	gc.LineTo(0, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentSUpperBar) ID() SegmentID { return IDSUpperBar }

/* Letter A */

const ABarUpperHeight = DefaultHeight - 35

// SegmentARisingStick is one of the parts of an A
//
//    / /
//   / /
//  | |
//  | |
type SegmentARisingStick struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentARisingStick) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	// gc.SetFillColor(color.RGBA{0xff, 0x99, 0x99, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(0, height-1)
	gc.LineTo(stickThickness, height-1)
	gc.LineTo(stickThickness, ABarUpperHeight)
	gradient := (40 - height) / (stickThickness - width + 1) // (y2-y1)/(x1-x2) because y axis is reversed
	gc.LineTo(width-1, height-35-gradient*(width-1-stickThickness))
	gc.LineTo(width-1, 0)
	gc.LineTo(width-stickThickness-1, 0) // (x2, y2)
	gc.LineTo(0, height-40)              // (x1, y1)
	gc.LineTo(0, height-1)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentARisingStick) ID() SegmentID { return IDARisingStick }

// SegmentAVert is the vertical line on the right hand side of an A
type SegmentAVert struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentAVert) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	// gc.SetFillColor(color.RGBA{0x99, 0xff, 0x99, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(width-stickThickness-1, 0)
	gc.LineTo(width-1, 0)
	gc.LineTo(width-1, height-1)
	gc.LineTo(width-stickThickness-1, height-1)
	gc.LineTo(width-stickThickness-1, 0)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentAVert) ID() SegmentID { return IDAVert }

// SegmentABar is the horizontal line through an A
type SegmentABar struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentABar) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	// gc.SetFillColor(color.RGBA{0x99, 0x99, 0xff, 0xff})
	width := seg.Width()
	gc.MoveTo(stickThickness, ABarUpperHeight)
	gc.LineTo(width-stickThickness-1, ABarUpperHeight)
	gc.LineTo(width-stickThickness-1, ABarUpperHeight+stickThickness)
	gc.LineTo(stickThickness, ABarUpperHeight+stickThickness)
	gc.LineTo(stickThickness, ABarUpperHeight)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentABar) ID() SegmentID { return IDABar }

// LetterA returns all the segments for the letter A
func LetterA() Letter {
	return Letter{&SegmentARisingStick{}, &SegmentAVert{}, &SegmentABar{}}
}
