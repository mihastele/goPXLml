package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"piksel.si/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}
	s.ExtendBaseWidget(s)
	s.Refresh()
	return s
}

func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{square: *square, objects: objects, parent: s}
}
