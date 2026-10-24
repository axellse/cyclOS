package main

import (
	"fmt"
	"time"

	"github.com/axellse/cyclos/cyclos/config"
	"github.com/axellse/cyclos/cyclos/drivers/display"
	"github.com/axellse/cyclos/cyclos/shared"
	"github.com/axellse/cyclos/cyclos/ui"
)

type System struct {
	Config config.Config
	HAL *shared.HAL
	PushState chan shared.SystemState
	UIListener chan shared.SystemState
	state shared.SystemState
}

func UIRenderer(system *System) {
	for {
		state := <- system.UIListener
		ui.BuildUI(state, system.HAL.Display)
	}
}

func main() {
	system := &System{}
	system.Config = config.LoadConfig()
	system.HAL = &shared.HAL{
		Display: display.Init(system.Config.Display),
	}

	system.PushState = make(chan shared.SystemState)
	system.UIListener = make(chan shared.SystemState)

	go UIRenderer(system)
	go func() {
		hour, min, _ := time.Now().Clock()
		clockStr := fmt.Sprintf("%02d", hour) + ":" + fmt.Sprintf("%02d", min)

		system.PushState <- shared.SystemState{
			Input: shared.InputState{
				UIPage: shared.Home,
			},
			Misc: shared.Misc{
				Clock: shared.Clock{
					Available: true,
					Current: clockStr,
				},
				Battery: shared.Battery{
					Status: shared.ThirdsPrecision,
					Value: 1,
				},
				DisplayWidth: 320,
				DisplayHeight: 170,
			},
			RTMetrics: shared.RTMetrics{
				Speed: 0,
				SpeedUnit: shared.KmH,
			},
		}
	}()
	for {
		newState := <- system.PushState
		system.state = newState

		system.UIListener <- system.state
	}
}
