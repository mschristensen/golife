package life

import "math/rand"

// World describes the world grid.
type World struct {
	Width, Height int64
	Grid          [][]bool
}

// NewWorld returns a World of the given dimensions with a randomly initialised grid.
func NewWorld(width, height int64) *World {
	world := &World{
		Width:  width,
		Height: height,
	}
	world.Grid = make([][]bool, height)
	for i := range world.Grid {
		world.Grid[i] = make([]bool, width)
		for j := range world.Grid[i] {
			world.Grid[i][j] = rand.Float64() > 0.5
		}
	}
	return world
}
