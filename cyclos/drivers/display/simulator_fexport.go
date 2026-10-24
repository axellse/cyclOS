//go:build d_simulator_fexport

package display

import (
	"encoding/json"
	"fmt"
	"image/color"

	"github.com/axellse/cyclos/cyclos/cyclos/config"
	"tinygo.org/x/drivers/pixel"
)

type SimulatorScreen struct {
	conf config.DisplayConfig
	buf  pixel.Image[pixel.RGB888]
}

type widthHeightObject struct {
	Width  int
	Height int
}

//export dSize
func (d *SimulatorScreen) size() (sizeObj string)

func (d *SimulatorScreen) Size() (width, height int16) {
	who := widthHeightObject{}
	err := json.Unmarshal([]byte(d.size()), &who)
	if err != nil {
		panic(err)
	}

	return int16(who.Width), int16(who.Width)
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
			fmt.Println("buf is", d.buf)
			fmt.Println("get is", d.buf.Get(int(x), int(y)))
			ba, err := json.Marshal(d.buf.Get(int(x), int(y)).RGBA())
			if err != nil {
				panic(err)
			}

			setPixel(int(x), int(y), string(ba))
		}
	}
	d.buf = pixel.Image[pixel.RGB888]{}

	return nil
}

//export dSetPixel
func setPixel(x, y int, colorObject string)

func (d *SimulatorScreen) SetPixel(x, y int16, c color.RGBA) {
	fmt.Println("setting buf")
	d.buf.Set(int(x), int(y), pixel.NewColor[pixel.RGB888](c.R, c.G, c.B))
}

func Init(conf config.DisplayConfig) *SimulatorScreen {
	return &SimulatorScreen{
		conf: conf,
	}
}
