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
	IDSLowerBar
	IDSMiddle
	IDNLeftVert
	IDNBar
	IDNRightVert
	IDARisingStick
	IDABar
	IDKUpper
	IDKLower
	IDELower
	IDEMiddle
	IDIMiddle
	IDITop
	IDIBottom
	IDDCurve
)

var letterMap = map[byte]func() Letter{
	' ': LetterSpace,
	'S': LetterS,
	'N': LetterN,
	'A': LetterA,
	'K': LetterK,
	'E': LetterE,
	'I': LetterI,
	'D': LetterD,
}

// GetLetterMap returns a map to get the functions corresponding to each letter
func GetLetterMap() map[byte]func() Letter {
	newMap := make(map[byte]func() Letter)
	for key, val := range letterMap {
		newMap[key] = val
	}
	return newMap
}

// Letter defines a custom type for full letters. A letter is defined
// as a slice of segments.
type Letter []Segment

/* Space */

// LetterSpace returns a blank collection of segments
func LetterSpace() Letter {
	return Letter{}
}

/* Letter S */

// SStickLength is the length of the sticks on the tops and bottoms of
// an S
const SStickLength = 7

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
	gc.LineTo(width-1, stickThickness+SStickLength)
	gc.LineTo(width-1-stickThickness, stickThickness+SStickLength)
	gc.LineTo(width-1-stickThickness, stickThickness)
	gc.LineTo(0, stickThickness)
	gc.LineTo(0, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentSUpperBar) ID() SegmentID { return IDSUpperBar }

// SegmentSLowerBar is the bottom bar on an S
type SegmentSLowerBar struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentSLowerBar) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(0, height-1)
	gc.LineTo(width-1, height-1)
	gc.LineTo(width-1, height-stickThickness-1)
	gc.LineTo(stickThickness, height-stickThickness-1)
	gc.LineTo(stickThickness, height-stickThickness-SStickLength-1)
	gc.LineTo(0, height-stickThickness-SStickLength-1)
	gc.LineTo(0, height-stickThickness-1)
	gc.LineTo(0, height-1)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentSLowerBar) ID() SegmentID { return IDSLowerBar }

// SegmentSMiddle is the middle section of an S
type SegmentSMiddle struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentSMiddle) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(0, stickThickness-0.5)
	gc.LineTo(stickThickness, stickThickness-0.5)
	gc.LineTo(stickThickness, stickThickness+9)
	gc.LineTo(width-1, height-stickThickness-16)
	gc.LineTo(width-1, height-stickThickness-0.5)
	gc.LineTo(width-stickThickness-1, height-stickThickness-0.5)
	gc.LineTo(width-stickThickness-1, height-stickThickness-10)
	gc.LineTo(0, stickThickness+15)
	gc.LineTo(0, stickThickness-0.5)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentSMiddle) ID() SegmentID { return IDSMiddle }

// LetterS returns all the segments for the letter S
func LetterS() Letter {
	return Letter{&SegmentSLowerBar{}, &SegmentSMiddle{}, &SegmentSUpperBar{}}
}

/* Letter N */

// SegmentNLeftVert is the vertical line on the right hand side of an N
type SegmentNLeftVert struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentNLeftVert) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	height := seg.Height()
	gc.MoveTo(0, 0)
	gc.LineTo(stickThickness, 0)
	gc.LineTo(stickThickness, height-1)
	gc.LineTo(0, height-1)
	gc.LineTo(0, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentNLeftVert) ID() SegmentID { return IDNLeftVert }

// SegmentNBar is the vertical bar across the top of an N
type SegmentNBar struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentNBar) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width := seg.Width()
	gc.MoveTo(0, 0)
	gc.LineTo(width-1, 0)
	gc.LineTo(width-1, stickThickness-1)
	gc.LineTo(0, stickThickness-1)
	gc.LineTo(0, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentNBar) ID() SegmentID { return IDNBar }

// SegmentNRightVert is the vertical line on the right hand side of an N
type SegmentNRightVert struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentNRightVert) Draw(gc CellDrawer, cell *Cell) bool {
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
func (seg *SegmentNRightVert) ID() SegmentID { return IDNRightVert }

// LetterN returns all the segments for the letter N
func LetterN() Letter {
	return Letter{&SegmentNLeftVert{}, &SegmentNBar{}, &SegmentNRightVert{}}
}

/* Letter A */

// ABarUpperHeight is the height of the line through the centre of an
// A
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
	gc.MoveTo(stickThickness-0.5, ABarUpperHeight)
	gc.LineTo(width-stickThickness-0.5, ABarUpperHeight)
	gc.LineTo(width-stickThickness-0.5, ABarUpperHeight+stickThickness)
	gc.LineTo(stickThickness-0.5, ABarUpperHeight+stickThickness)
	gc.LineTo(stickThickness-0.5, ABarUpperHeight)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentABar) ID() SegmentID { return IDABar }

// LetterA returns all the segments for the letter A
func LetterA() Letter {
	return Letter{&SegmentARisingStick{}, &SegmentNRightVert{}, &SegmentABar{}}
}

/* Letter K */

// SegmentKUpper is the upper part of a K
type SegmentKUpper struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentKUpper) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(stickThickness-8.5, height/2-0.5)
	gc.LineTo(width-1-stickThickness, 0)
	gc.LineTo(width-1, 0)
	gc.LineTo(2*stickThickness-8.5, height/2-0.5)
	gc.LineTo(stickThickness-8.5, height/2-0.5)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentKUpper) ID() SegmentID { return IDKUpper }

// SegmentKLower is the upper part of a K
type SegmentKLower struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentKLower) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(stickThickness-8.5, height/2-1.5)
	gc.LineTo(width-1-stickThickness, height-1)
	gc.LineTo(width-1, height-1)
	gc.LineTo(2*stickThickness-8.5, height/2-1.5)
	gc.LineTo(stickThickness-8.5, height/2-1.5)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentKLower) ID() SegmentID { return IDKLower }

// LetterK returns all the segments for the letter K
func LetterK() Letter {
	return Letter{&SegmentNLeftVert{}, &SegmentKUpper{}, &SegmentKLower{}}
}

/* Letter E */

// EPointDepth is how far in the dip on the left of an E goes in
const EPointDepth = 13

// EEdgeDepth is how far in the centre line of the E is from the right
const EEdgeDepth = 5

// SegmentELower is the lower bar on an E
type SegmentELower struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentELower) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(0, height-1)
	gc.LineTo(width-1, height-1)
	gc.LineTo(width-1, height-1-stickThickness)
	gc.LineTo(0, height-1-stickThickness)
	gc.LineTo(0, height-1)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentELower) ID() SegmentID { return IDELower }

// SegmentEMiddle is the centre segment of an E
type SegmentEMiddle struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentEMiddle) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	// height := seg.Height()
	gc.MoveTo(0, height-0.5-stickThickness)
	gc.LineTo(EPointDepth-1, height/2-1)
	gc.LineTo(0, stickThickness-1)
	gc.LineTo(0, stickThickness-2)
	gc.LineTo(stickThickness-1, stickThickness-2)
	gc.LineTo(stickThickness-1, stickThickness-1)
	gc.LineTo(stickThickness+EPointDepth-6, height/2-stickThickness/2-1)
	gc.LineTo(width-1-EEdgeDepth, height/2-stickThickness/2-1)
	gc.LineTo(width-1-EEdgeDepth, height/2+stickThickness/2-1)
	gc.LineTo(stickThickness+EPointDepth-6, height/2+stickThickness/2-1)
	gc.LineTo(stickThickness-1, height-0.5-stickThickness)
	gc.LineTo(0, height-0.5-stickThickness)
	gc.Fill()
	gc.Close()

	return true
}

// ID returns the ID of the segment
func (seg *SegmentEMiddle) ID() SegmentID { return IDEMiddle }

// LetterE returns all the segments for the letter E
func LetterE() Letter {
	return Letter{&SegmentNBar{}, &SegmentELower{}, &SegmentEMiddle{}}
}

/* Letter I */

// IOverhang is how far the top bar of an I hangs over
const IOverhang = 7

// SegmentIMiddle is the line through the centre of an I
type SegmentIMiddle struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentIMiddle) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(width/2-1-stickThickness/2, 0)
	gc.LineTo(width/2-1+stickThickness/2, 0)
	gc.LineTo(width/2-1+stickThickness/2, height-1)
	gc.LineTo(width/2-1-stickThickness/2, height-1)
	gc.LineTo(width/2-1-stickThickness/2, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentIMiddle) ID() SegmentID { return IDIMiddle }

// SegmentITop is the line atop an I
type SegmentITop struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentITop) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width := seg.Width()
	gc.MoveTo(width/2-1-stickThickness/2-IOverhang, 0)
	gc.LineTo(width/2-1+stickThickness/2+IOverhang, 0)
	gc.LineTo(width/2-1+stickThickness/2+IOverhang, stickThickness-1)
	gc.LineTo(width/2-1-stickThickness/2-IOverhang, stickThickness-1)
	gc.LineTo(width/2-1-stickThickness/2-IOverhang, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentITop) ID() SegmentID { return IDITop }

// SegmentIBottom is the line below an I
type SegmentIBottom struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentIBottom) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(width/2-1-stickThickness/2-IOverhang, height-1)
	gc.LineTo(width/2-1+stickThickness/2+IOverhang, height-1)
	gc.LineTo(width/2-1+stickThickness/2+IOverhang, height-stickThickness-1)
	gc.LineTo(width/2-1-stickThickness/2-IOverhang, height-stickThickness-1)
	gc.LineTo(width/2-1-stickThickness/2-IOverhang, height-1)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentIBottom) ID() SegmentID { return IDIBottom }

// LetterI returns all the segments for the letter I
func LetterI() Letter {
	return Letter{&SegmentIMiddle{}, &SegmentITop{}, &SegmentIBottom{}}
}

/* Letter D */

// Sizes a D
const (
	DGapThinnest    = 5
	DDiagonalHeight = DefaultHeight / 4
)

// SegmentDCurve is the curved line of a D
type SegmentDCurve struct {
	segment
}

// Draw defines the behaviour of the segment
func (seg *SegmentDCurve) Draw(gc CellDrawer, cell *Cell) bool {
	fillColor := cell.DeathColors[1]
	gc.SetFillColor(color.RGBA{fillColor.R, fillColor.G, fillColor.B, 0xff})
	width, height := seg.Width(), seg.Height()
	gc.MoveTo(0, 0)
	gc.LineTo(width-stickThickness-1, 0)
	gc.LineTo(width-1, DDiagonalHeight-1)
	gc.LineTo(width-1, height-DDiagonalHeight-1)
	gc.LineTo(width-stickThickness-1, height-1)
	gc.LineTo(0, height-1)
	gc.LineTo(0, height-stickThickness-1)
	gc.LineTo(stickThickness-1+DGapThinnest, height-stickThickness-1)
	gc.LineTo(width-1-stickThickness, height-11-stickThickness)
	gc.LineTo(width-1-stickThickness, stickThickness+9)
	gc.LineTo(stickThickness-1+DGapThinnest, stickThickness-1)
	gc.LineTo(0, stickThickness-1)
	gc.LineTo(0, 0)
	gc.Fill()
	gc.Close()
	return true
}

// ID returns the ID of the segment
func (seg *SegmentDCurve) ID() SegmentID { return IDDCurve }

// LetterD returns all the segments for the letter D
func LetterD() Letter {
	return Letter{&SegmentNLeftVert{}, &SegmentDCurve{}}
}
