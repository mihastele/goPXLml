package swatch

import (
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
	return s
}
