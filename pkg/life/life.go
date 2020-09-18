package life

import "math/rand"

// World describes the world grid.
type World struct {
	Width, Height int
	// Grid describes the set of cells in the world and their alive/dead state.
	// Note that the first row in the array corresponds to the bottom of the screen,
	// and the first element of a row corresponds to the left side of the screen.
	Grid [][]bool
}

// NewWorld returns a World of the given dimensions with a randomly initialised grid.
func NewWorld(width, height int) *World {
	world := &World{
		Width:  width,
		Height: height,
	}
	world.Grid = make([][]bool, height)
	for y := range world.Grid {
		world.Grid[y] = make([]bool, width)
		for x := range world.Grid[y] {
			world.Grid[y][x] = rand.Float64() < 0.2
		}
	}
	return world
}

// WrapCoords accepts an (x, y) coordinate and returns a coordinate which is
// within the bounds of the World, using a toroidal wrapping.
func (w *World) WrapCoords(x, y int) (int, int) {
	if x < 0 {
		x = w.Width - 1
	}
	if x > w.Width-1 {
		x = 0
	}
	if y < 0 {
		y = w.Height - 1
	}
	if y > w.Height-1 {
		y = 0
	}
	return x, y
}

// CountLiveNeighbours returns the number of live neighbours
// around the given cell position.
func (w *World) CountLiveNeighbours(x, y int) int {
	count := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == 0 && j == 0 {
				continue
			}
			_x, _y := w.WrapCoords(x+i, y+j)
			if w.Grid[_y][_x] {
				count++
			}
		}
	}
	return count
}

// Update applies the rules to the world.
func (w *World) Update() {
	// Create a deep copy of the world grid
	nextGrid := make([][]bool, len(w.Grid))
	for y := range w.Grid {
		nextGrid[y] = make([]bool, len(w.Grid[y]))
		copy(nextGrid[y], w.Grid[y])
	}

	// Apply Conway's rules
	for y := range w.Grid {
		for x := range w.Grid[y] {
			n := w.CountLiveNeighbours(x, y)
			if w.Grid[y][x] {
				// Any live cell with fewer than two or more than three live neighbours dies
				if n < 2 || n > 3 {
					nextGrid[y][x] = false
				}
			} else {
				// Any dead cell with exactly three live neighbours becomes a live cell
				if n == 3 {
					nextGrid[y][x] = true
				}
			}
		}
	}

	// Update the world grid
	w.Grid = nextGrid
}
