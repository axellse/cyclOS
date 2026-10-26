//This package contains a collection of input schemes, each with a "mutate function" that takes a starting state and some sort of event, and returns a mutated state.
package libinput

import "github.com/axellse/cyclos/cyclos/shared"

//A simple, 3 button up, down and confirm button scheme. button EnumString takes "up", "down" and "confirm".
func Mutate_ThreeButtonInput(state shared.InputState, button EnumString) shared.InputState {
	uipage := state.UIPage
	if button == "up" || button == "down" {
		uipage = CycleUIPage(state.UIPage, button)
	}

	return shared.InputState{
		UIPage: uipage,
	}
}