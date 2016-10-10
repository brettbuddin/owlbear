package dungeon

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestForExplosions(t *testing.T) {
	for i := 0; i < 100; i++ {
		config := Config{
			Width:          100,
			Height:         100,
			Rooms:          10,
			MinRoomDim:     5,
			MinRoomSpacing: 1,
			MaxAttempts:    10,
		}
		buffer := bytes.NewBuffer(nil)
		New(config).Draw(buffer)

		if buffer.Len() == 0 {
			t.Errorf("map expected to be written, but was found empty")
		}
	}
}
