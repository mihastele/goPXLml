package pxcanvas

import "fyne.io/fyne/v2"

func (pxCanvas *PxCanvas) Pan(previousCoord, curentCoord fyne.PointEvent) {
	xDiff := curentCoord.Position.X - previousCoord.Position.X
	yDiff := curentCoord.Position.Y - previousCoord.Position.Y
	pxCanvas.CanvasOffset.X += xDiff
	pxCanvas.CanvasOffset.Y += yDiff
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) Scale(direction int) {
	switch {
	case direction > 0:
		pxCanvas.PxSize++
	case direction < 0:
		if pxCanvas.PxSize > 2 {
			pxCanvas.PxSize--
		}
	default:
		pxCanvas.PxSize = 10
	}
}
