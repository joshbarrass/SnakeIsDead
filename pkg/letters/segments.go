package letters

// SegmentState is a type used for storing various internal states of the segments
type SegmentState int

// Possible states of a segment
const (
	StateFadein = iota
	StateSolid
	StateFadeout
)

// SegmentID is a unique type for unique Segment IDs
type SegmentID int

// Segment is an interface representing a segment of a letter.
// Any coordinates used should be defined in terms of the cell
type Segment interface {
	Width() float64
	Height() float64
	Draw(CellDrawer, *Cell) bool // defined for each type
	ID() SegmentID               // defined for each type
}

// segment is a struct representing the actual segment. This is
// inherited from to create the actual segments
type segment struct {
	State SegmentState
	W     float64
	H     float64
}

// Width returns the Width of the segment
func (seg *segment) Width() float64 {
	if seg.W <= 0 {
		return DefaultWidth
	}
	return seg.W
}

// Height returns the Height of the segment
func (seg *segment) Height() float64 {
	if seg.H <= 0 {
		return DefaultHeight
	}
	return seg.H
}
