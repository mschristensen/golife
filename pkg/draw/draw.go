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
	scale := pixel.V(
		window.Bounds().Max.Y/float64(len(world.Grid)),
		window.Bounds().Max.X/float64(len(world.Grid[0])),
	)
	for y := range world.Grid {
		for x := range world.Grid[y] {
			if !world.Grid[y][x] {
				continue
			}
			imd.Push(pixel.V(float64(x)*scale.X, float64(y)*scale.Y))
			imd.Push(pixel.V(float64(x+1)*scale.X, float64(y+1)*scale.Y))
			imd.Rectangle(0)
		}
	}
	imd.Draw(d.Batch)

	d.Batch.Draw(window)
	return nil
}
