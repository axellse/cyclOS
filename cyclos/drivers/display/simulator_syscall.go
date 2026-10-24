//go:build d_simulator

package display

import (
	"encoding/json"
	"fmt"
	"image/color"
	"syscall/js"

	"github.com/axellse/cyclos/cyclos/config"
	"tinygo.org/x/drivers/pixel"
)

type SimulatorScreen struct {
	conf config.DisplayConfig
	buf  pixel.Image[pixel.RGB888]
}

func (d *SimulatorScreen) Size() (width, height int16) {
	val := js.Global().Call("dSize")
	return int16(val.Get("width").Int()), int16(val.Get("height").Int())
}

func (d *SimulatorScreen) BeginDisplay() error {
	w, h := d.Size()
	fmt.Println("making buf", w, h)
	d.buf = pixel.NewImage[pixel.RGB888](int(w), int(h))

	return nil
}

func (d *SimulatorScreen) Display() error {
	fmt.Println("displaying buf")
	w, h := d.Size()
	for x := range w {
		for y := range h {
			ba, err := json.Marshal(d.buf.Get(int(x), int(y)).RGBA())
			if err != nil {
				panic(err)
			}

			js.Global().Call("dSetPixel", x, y, string(ba))
		}
	}
	d.buf = pixel.Image[pixel.RGB888]{}

	return nil
}

func (d *SimulatorScreen) SetPixel(x, y int16, c color.RGBA) {
	w, h := d.Size()
	if x >= w || y >= h {
		fmt.Println("setPixel out of bounds!")
		return
	}
	d.buf.Set(int(x), int(y), pixel.NewColor[pixel.RGB888](c.R, c.G, c.B))
}

func Init(conf config.DisplayConfig) *SimulatorScreen {
	return &SimulatorScreen{
		conf: conf,
	}
}
