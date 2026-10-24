//go:build d_simulator_unbuf

package display

import (
	"encoding/json"
	"fmt"
	"image/color"
	"syscall/js"

	"github.com/axellse/cyclos/cyclos/config"
)

type SimulatorScreen struct {
	conf config.DisplayConfig
}

func (d *SimulatorScreen) Size() (width, height int16) {
	val := js.Global().Call("dSize")
	return int16(val.Get("width").Int()), int16(val.Get("height").Int())
}

func (d *SimulatorScreen) BeginDisplay() error {
	return nil
}

func (d *SimulatorScreen) Display() error {
	return nil
}

func (d *SimulatorScreen) SetPixel(x, y int16, c color.RGBA) {
	w, h := d.Size()
	if x > w || y > h {
		fmt.Println("setPixel out of bounds!")
		return
	}

	ba, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	js.Global().Call("dSetPixel", x, y, string(ba))
}

func Init(conf config.DisplayConfig) *SimulatorScreen {
	return &SimulatorScreen{
		conf: conf,
	}
}
