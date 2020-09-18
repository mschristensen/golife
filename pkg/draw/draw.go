package draw

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/golife/pkg/life"
	"golang.org/x/image/colornames"
)

// Drawer enables batch drawing sprites from a given strip.
type Drawer struct {
	Batch *pixel.Batch
}

// NewDrawer creates a new Drawer for the given sprite strip.
func NewDrawer() *Drawer {
	return &Drawer{
		Batch: pixel.NewBatch(&pixel.TrianglesData{}, nil),
	}
}

// DrawFrame draws the world on the window.
func (d *Drawer) DrawFrame(window *pixelgl.Window, world *life.World) error {
	window.Clear(colornames.Aliceblue)
	d.Batch.Clear()

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(pixel.V(float64(10), float64(10)))
	imd.Push(pixel.V(float64(20), float64(20)))
	// cellWidth := 1.0
	for i := range world.Grid {
		for j := range world.Grid[i] {
			if !world.Grid[i][j] {
				continue
			}
			// imd.Push(pixel.V(float64(i)*cellWidth, float64(j)*cellWidth))
			// imd.Push(pixel.V(float64(i+10)*cellWidth, float64(j+10)*cellWidth))
			// imd.Rectangle(0)
		}
		break
	}
	// imd.Color = pixel.RGB(0, 1, 0)
	// imd.Color = pixel.RGB(0, 0, 1)
	// imd.Push(pixel.V(500, 700))
	// imd.Polygon(0)
	imd.Draw(d.Batch)
	// d.Batch.SetMatrix()

	d.Batch.Draw(window)
	return nil
}
