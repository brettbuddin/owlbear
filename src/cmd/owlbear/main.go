package main

import (
	"bytes"
	"dungeon"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		config dungeon.Config
		seed   int64
	)
	flag.IntVar(&config.Width, "width", 100, "width of the dungeon")
	flag.IntVar(&config.Height, "height", 50, "height of the dungeon")
	flag.IntVar(&config.Rooms, "rooms", 20, "number of rooms")
	flag.IntVar(&config.MaxAttempts, "max-attempts", 20, "maximum number of attempts at creating a room")
	flag.IntVar(&config.MinRoomDim, "min-room-dim", 3, "minimum room diminsion")
	flag.IntVar(&config.MinRoomSpacing, "min-room-spacing", 5, "minimum spacing between rooms")
	flag.Int64Var(&seed, "seed", 0, "random seed (useful in debugging)")
	flag.Parse()

	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	fmt.Printf("Seed: %v\n", seed)
	rand.Seed(seed)

	buffer := bytes.NewBuffer(nil)
	dungeon.New(config).Draw(buffer)
	fmt.Print(buffer.String())
}
