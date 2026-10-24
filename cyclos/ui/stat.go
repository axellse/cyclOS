package ui

import (

	"github.com/axellse/cyclos/cyclos/lib/libui"
	"github.com/axellse/cyclos/cyclos/shared"
)

func CalcRowBoxes(state shared.SystemState, cols int16, rows int16) (width int16, height int16) {
	colWidth := (state.Misc.DisplayWidth - SidebarWidth - ((cols +1) * BasicMargin)) / cols
	colHeight := (state.Misc.DisplayHeight - TopbarHeight - ((rows +1) * BasicMargin)) / rows
	return colWidth, colHeight
}

func DrawBox(state shared.SystemState, display shared.Display, width int16, height int16, widthM int16, colI int16, rowI int16) {
	topLeft := libui.Point{X: ((colI + 1) * BasicMargin) + (colI * width), Y: TopbarHeight + BasicMargin + ((height + BasicMargin) * rowI)}
	bottomRight := libui.Point{X: topLeft.X + (width * widthM) + (BasicMargin * (widthM - 1)), Y: topLeft.Y + height}

	libui.FillRoundedRectangle(display, topLeft, bottomRight, 6, SecondaryBackgroundColor)
}

func BuildStat(state shared.SystemState, display shared.Display) {
	r1W, r1H := CalcRowBoxes(state, 3, 2)
	for i := range int16(3) {
		DrawBox(state, display, r1W, r1H, 1, i, 0)
	}

	r2W, r2H := CalcRowBoxes(state, 3, 2)
	DrawBox(state, display, r2W, r2H, 2, 0, 1)
	DrawBox(state, display, r2W, r2H, 1, 2, 1)
}

/*

*/