package ui

import (

	"github.com/axellse/cyclos/cyclos/lib/libui"
	"github.com/axellse/cyclos/cyclos/shared"
	lexenddeca "github.com/axellse/cyclos/cyclos/ui/fonts/lexend-deca"
	"github.com/axellse/cyclos/cyclos/ui/icons"
	"tinygo.org/x/tinyfont"
)

func BuildBase(state shared.SystemState, display shared.Display) {
	//background
	libui.FillRectangle(display, libui.Point{X: 0, Y: 0}, libui.Point{X: state.Misc.DisplayWidth-1, Y: state.Misc.DisplayHeight-1}, BackgroundColor)

	//sidebar background
	start := libui.Point{
		X: state.Misc.DisplayWidth - SidebarWidth,
		Y: 0,
	}
	end := libui.Point{
		X: state.Misc.DisplayWidth -1,
		Y: state.Misc.DisplayHeight - 1,
	}

	libui.FillRectangle(display, start, end, SecondaryBackgroundColor)
	if state.Misc.Clock.Available {
		/*hour, min, _ := state.Misc.Clock.Current.Clock()
		clockStr := fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", min)*/

		_, w2 := tinyfont.LineWidth(&lexenddeca.LexendDeca, state.Misc.Clock.Current)
		tinyfont.WriteLine(display, &lexenddeca.LexendDeca, state.Misc.DisplayWidth - SidebarWidth - BasicMargin - int16(w2), BasicMargin + 18, state.Misc.Clock.Current, SecondaryTextColor) //18 for font height
	}

	if state.Misc.Battery.Status != shared.BatNotAvailable {
		libui.DrawIcon(display, libui.Point{X: BasicMargin, Y: BasicMargin}, icons.Icons["batteryBase"], SecondaryTextColor)
	}

	totalHeight := BasicMargin * int16(len(shared.UIPages) -1) 
	for _, icon := range shared.UIPages {
		totalHeight += icons.Icons[string(icon)].Height
	}

	workingY := (state.Misc.DisplayHeight/2) - (totalHeight/2)
	for _, icon := range shared.UIPages {
		iconObj := icons.Icons[string(icon)]
		col := IconColor
		if state.Input.UIPage == shared.UIPage(icon) {
			col = AccentColor
		}
		
		libui.DrawIcon(display, libui.Point{X: state.Misc.DisplayWidth - iconObj.Width - 6, Y: workingY}, iconObj, col)
		workingY += BasicMargin + iconObj.Height
	}
}

func BuildUI(state shared.SystemState, display shared.Display) {
	display.BeginDisplay()

	BuildBase(state, display)
	switch state.Input.UIPage {
	case shared.Home:
		BuildHome(state, display)
	case shared.Stat:
		BuildStat(state, display)
	}

	display.Display()
}