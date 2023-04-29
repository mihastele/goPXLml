package ui

import (
	"fyne.io/fyne/v2"
	"piksel.si/apptype"
	"piksel.si/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
