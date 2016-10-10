package dungeon

import (
	"fmt"
	"geometry"
	"io"
	"math/rand"
)

const (
	earth    = '#'
	empty    = ' '
	entrance = '!'
	exit     = '$'
)

// Config establishes constraints for the process of generating a Dungeon
type Config struct {
	Rooms                      int
	Width, Height              int
	MaxAttempts                int
	MinRoomSpacing, MinRoomDim int
}

// Dungeon is the main arena
type Dungeon struct {
	Config Config
}

// New returns a Dungeon
func New(c Config) *Dungeon {
	return &Dungeon{
		Config: c,
	}
}

// Draw draws the dungeon to an output buffer
func (d *Dungeon) Draw(w io.Writer) {
	// Fill the dungeon with earth
	area := make([][]rune, d.Config.Height)
	for i := 0; i < d.Config.Height; i++ {
		area[i] = make([]rune, d.Config.Width)
		for j := 0; j < d.Config.Width; j++ {
			area[i][j] = earth
		}
	}

	// Generate rooms
	rooms := d.generateRooms()
	pointers := []geometry.Pointer{}
	for _, r := range rooms {
		pointers = append(pointers, r)
	}

	// Calculate the Delaunay triangulation (using room centers) and then calculate the MST from the graph.
	graph := geometry.Delaunay(pointers)
	mst := geometry.Kruskal(pointers, graph)

	// Draw corridors and rooms to the dungeon map
	for _, e := range mst {
		d.connect(e.A.(*Room), e.B.(*Room), area)
	}

	// Place the entrance and exit. This could be improved by locating the longest contiguous edge of the MST placing
	// the entrance and exit at either end of the segment. Boring and random placement for now...
	entranceIdx := rand.Intn(len(rooms))
	exitIdx := rand.Intn(len(rooms))
	for i, r := range rooms {
		if i == entranceIdx {
			r.PlaceLoc(entrance)
		}
		if i == exitIdx {
			r.PlaceLoc(exit)
		}
		r.Draw(area)
	}
	for row := 0; row < d.Config.Height; row++ {
		for col := 0; col < d.Config.Width; col++ {
			fmt.Fprintf(w, "%s", string(area[row][col]))
		}
		fmt.Fprintf(w, "\n")
	}
}

func (d *Dungeon) connect(from, to *Room, area [][]rune) {
	fc, tc := from.Center(), to.Center()

	x, y := fc.X, fc.Y
	diff := fc.Diff(tc)
	dirX, dirY := pole(diff.X), pole(diff.Y)

	// Move on the Y axis until we align with the target's center
	for y != tc.Y && ((dirY == -1 && y >= tc.Y) || (dirY == 1 && y <= tc.Y)) {
		area[int(y)][int(x)] = empty
		y += dirY
	}

	// Move on the X axis until we align with the target's center
	for x != tc.X && ((dirX == -1 && x >= tc.X) || (dirX == 1 && x <= tc.X)) {
		area[int(y)][int(x)] = empty
		x += dirX
	}
}

func (d *Dungeon) generateRooms() []*Room {
	rooms := []*Room{}
	for i := 0; i < d.Config.Rooms; i++ {
		var attempts int
		for attempts < d.Config.MaxAttempts {
			attempts++

			width := rand.Intn(d.Config.Width / 2)
			height := rand.Intn(d.Config.Height / 2)
			min := d.Config.MinRoomDim
			if width < min || height < min {
				continue
			}

			x := rand.Intn(d.Config.Width)
			y := rand.Intn(d.Config.Height)
			r := NewRoom(x, y, width, height)
			if err := d.reviewCandidate(r, rooms); err != nil {
				continue
			}

			rooms = append(rooms, r)
			break
		}
	}
	return rooms
}

func (d *Dungeon) reviewCandidate(r *Room, existing []*Room) error {
	if r.TopLeft.X <= 0 || r.TopLeft.Y <= 0 {
		return fmt.Errorf("top left too close to edge")
	}
	if r.BottomRight.X >= float64(d.Config.Width)-1 || r.BottomRight.Y >= float64(d.Config.Height)-1 {
		return fmt.Errorf("bottom right too close to edge")
	}
	for _, existingR := range existing {
		if r.Overlap(existingR, float64(d.Config.MinRoomSpacing)) {
			return fmt.Errorf("collision")
		}
	}
	return nil
}

func pole(v float64) float64 {
	if v < 0 {
		return 1
	}
	return -1
}
