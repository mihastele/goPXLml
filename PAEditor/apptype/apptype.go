package apptype

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type BrushType = int

type PxCanvasConfig struct {
	DawingArea     fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSiye         int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (state *State) setFilePath(path string) {
	state.FilePath = path
}
