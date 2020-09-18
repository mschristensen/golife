package golife

import (
	"fmt"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/mschristensen/golife/pkg/draw"
	"github.com/mschristensen/golife/pkg/life"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pkg/errors"
)

const (
	WorldWidth  = 1400
	WorldHeight = 800
)

func createWindow(title string, width, height float64) (*pixelgl.Window, error) {
	window, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "new window failed")
	}
	window.SetSmooth(true)
	return window, nil
}

func Run() {
	window, err := createWindow("GoLife", WorldWidth, WorldHeight)
	if err != nil {
		panic(errors.Wrap(err, "create window failed"))
	}
	frames := 0
	second := time.Tick(time.Second)
	world := life.NewWorld(WorldWidth, WorldHeight)
	drawer := draw.NewDrawer()
	for !window.Closed() {
		drawer.DrawFrame(window, world)
		window.Update()
		frames++
		select {
		case <-second:
			window.SetTitle(fmt.Sprintf("%s | FPS: %d", "GoLife", frames))
			frames = 0
		default:
		}
	}
}

func init() {
	rand.Seed(time.Now().Unix())
}
