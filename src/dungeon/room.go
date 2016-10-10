package dungeon

import (
	"fmt"
	"geometry"
	"math"
	"math/rand"
)

// Room is a room in the dungeon
type Room struct {
	Width, Height        float64
	TopLeft, BottomRight geometry.Point
	locs                 []Loc
}

// NewRoom creates a Room with specified dimensions
func NewRoom(x, y, width, height int) *Room {
	return &Room{
		Width:       float64(width),
		Height:      float64(height),
		TopLeft:     geometry.Point{float64(x), float64(y)},
		BottomRight: geometry.Point{float64(x + width), float64(y + height)},
		locs:        []Loc{},
	}
}

// Center calculates the center of a Room
func (r *Room) Center() geometry.Point {
	diff := r.BottomRight.Diff(r.TopLeft)
	return geometry.Point{X: r.TopLeft.X + diff.X/2, Y: r.TopLeft.Y + diff.Y/2}
}

// Point implements the geometry.Pointer interface
func (pointer *Room) Point() geometry.Point {
	return pointer.Center()
}

// IsInside determines whether a geometry.Pointer is inside the Room
func (r *Room) IsInside(p geometry.Pointer) bool {
	point := p.Point()
	return r.TopLeft.X <= point.X &&
		r.BottomRight.X >= point.X &&
		r.TopLeft.Y <= point.Y &&
		r.BottomRight.Y >= point.Y
}

// Overlap determines whether or not another room (including spacing tolerance) overlaps this Room
func (r *Room) Overlap(o *Room, spacing float64) bool {
	return r.TopLeft.X-spacing <= o.BottomRight.X &&
		r.BottomRight.X+spacing >= o.TopLeft.X &&
		r.TopLeft.Y-spacing <= o.BottomRight.Y &&
		r.BottomRight.Y+spacing >= o.TopLeft.Y
}

func (room *Room) PlaceLoc(r rune) {
	room.locs = append(room.locs, Loc{room.randomPoint(), r})
}

func (r *Room) randomPoint() geometry.Pointer {
	return geometry.Point{
		X: r.TopLeft.X + float64(rand.Intn(int(r.Width))),
		Y: r.TopLeft.Y + float64(rand.Intn(int(r.Height))),
	}
}

func (r *Room) String() string {
	diff := r.TopLeft.Diff(r.BottomRight)
	return fmt.Sprintf("%v (%vx%v)", r.TopLeft, math.Abs(float64(diff.X)), math.Abs(float64(diff.Y)))
}

// Draw writes the Room to a buffer
func (r *Room) Draw(b [][]rune) {
	for col := r.TopLeft.Y; col < r.BottomRight.Y; col++ {
		for row := r.TopLeft.X; row < r.BottomRight.X; row++ {
			b[int(col)][int(row)] = empty
		}
	}
	for _, l := range r.locs {
		p := l.Point()
		b[int(p.Y)][int(p.X)] = l.Symbol
	}
}

type Loc struct {
	geometry.Pointer
	Symbol rune
}
