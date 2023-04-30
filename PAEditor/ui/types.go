package ui

import (
	"fyne.io/fyne/v2"
	"piksel.si/apptype"
	"piksel.si/pxcanvas"
	"piksel.si/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
