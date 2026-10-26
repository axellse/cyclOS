//go:build i_simulator

package input

import (
	"syscall/js"
	"time"

	"github.com/axellse/cyclos/cyclos/lib/libinput"
	"github.com/axellse/cyclos/cyclos/shared"
)

type SimulatorInput struct{}

func Init() SimulatorInput {
	return SimulatorInput{}
}

func (SimulatorInput) Channel(LatestState func() shared.InputState) chan shared.InputState {
	channel := make(chan shared.InputState)
	scheme := js.Global().Call("iGetScheme").String()

	buttonsPressed := map[string]bool{}
	go func() {
		for {
			if scheme == "3button" {
				for _, btn := range []string{"up", "confirm", "down"} {
					isPressed := js.Global().Call("iBoolStatus", btn).Bool()
					ok, wasPressed := buttonsPressed[btn]
					
					if isPressed && (!ok || !wasPressed) {
						state := LatestState()
						state = libinput.Mutate_ThreeButtonInput(state, libinput.EnumString(btn))
						channel <- state
					} else if !js.Global().Call("iBoolStatus", btn).Bool() {
						buttonsPressed[btn] = false
					}

					buttonsPressed[btn] = isPressed
				}
			} else {
				panic("unsupported scheme")
			}

			time.Sleep(5 * time.Millisecond)
		}
	}()

	return channel
}