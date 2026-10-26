package libinput

import (
	"slices"

	"github.com/axellse/cyclos/cyclos/shared"
)

//generic enum string type. See the method you're calling for possible values.
type EnumString string

//utility function that cycles a UI page in a given direction. direction EnumString takes "up" and "down". It panics on error.
func CycleUIPage(start shared.UIPage, direction EnumString) shared.UIPage {
	ci := slices.Index(shared.UIPages, start)
	if ci == -1 {
		panic("invalid start page")
	}

	dir := -1
	if direction == "down" {
		dir = 1
	}

	ni := ci + dir
	if ni == -1 {
		return shared.UIPages[len(shared.UIPages)-1]
	} else if ni == len(shared.UIPages) {
		return shared.UIPages[0]
	}
	return shared.UIPages[ni]
}