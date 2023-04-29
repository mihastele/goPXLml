package main

import (
	"fyne.io/fyne/v2/app"
	"image/color"
	"piksel.si/apptype"
	"piksel.si/swatch"
	"piksel.si/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("PXL")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
