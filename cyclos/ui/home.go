package ui

import (
	"strconv"

	"github.com/axellse/cyclos/cyclos/shared"
	lexenddeca "github.com/axellse/cyclos/cyclos/ui/fonts/lexend-deca"
	"github.com/axellse/cyclos/cyclos/ui/fonts/tratex"
	"tinygo.org/x/tinyfont"
)

func BuildHome(state shared.SystemState, display shared.Display) {
	tinyfont.WriteLine(display, &tratex.Tratex, 5, state.Misc.DisplayHeight-10, strconv.Itoa(state.RTMetrics.Speed), PrimaryTextColor)
	_, w2 := tinyfont.LineWidth(&tratex.Tratex, strconv.Itoa(state.RTMetrics.Speed))
	tinyfont.WriteLine(display, &lexenddeca.LexendDeca, int16(w2)+10, state.Misc.DisplayHeight-10, string(state.RTMetrics.SpeedUnit), SecondaryTextColor)
}