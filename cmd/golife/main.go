package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/mschristensen/golife/pkg/app/golife"
)

func main() {
	pixelgl.Run(golife.Run)
}
