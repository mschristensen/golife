package golife

import (
	"fmt"
	_ "image/png"
	"math/rand"
	"runtime"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/golife/pkg/draw"
	"github.com/mschristensen/golife/pkg/life"
	"github.com/pkg/errors"
)

const (
	// WindowWidth is the width of the window in pixels
	WindowWidth = 1200
	// WindowHeight is the height of the window in pixels
	WindowHeight = 800
	// WorldWidth is the width of the world in # cells.
	WorldWidth = 1200
	// WorldHeight is the height of the world in # cells.
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

// Run runs the game.
func Run() {
	window, err := createWindow("GoLife", WindowWidth, WindowHeight)
	if err != nil {
		panic(errors.Wrap(err, "create window failed"))
	}
	frames := 0
	second := time.Tick(time.Second)
	world := life.NewWorld(WorldWidth, WorldHeight)
	drawer := draw.NewDrawer()
	for !window.Closed() {
		drawer.DrawFrame(window, world)
		world.Update(runtime.NumCPU())
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
