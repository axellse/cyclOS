//go:build d_tdisplays3

package display

import (
	"image/color"
	"machine"

	"github.com/axellse/cyclos/cyclos/config"
	i8080 "github.com/axellse/cyclos/cyclos/lib/i8080"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

type TDisplayS3Screen struct {
	disp st7789.Device
	conf config.DisplayConfig
	buf pixel.Image[pixel.RGB565BE]
}

func (d *TDisplayS3Screen) Size() (width, height int16) {
	return d.disp.Size()
}

func (d *TDisplayS3Screen) BeginDisplay() error {
	if d.conf.HideDrawing {
		d.disp.EnableBacklight(false)
	}

	return nil
}

func (d *TDisplayS3Screen) Display() error {
	d.disp.DrawBitmap(0,0, d.buf)

	if d.conf.HideDrawing {
		d.disp.EnableBacklight(true)
	}
	return nil
}

func (d *TDisplayS3Screen) SetPixel(x, y int16, c color.RGBA) {
	w, h := d.Size()
	if x >= w || y >= h {
		return
	}

	d.buf.Set(int(x), int(y), pixel.NewColor[pixel.RGB565BE](c.R, c.G, c.B))
}

func Init(conf config.DisplayConfig) *TDisplayS3Screen {
	machine.GPIO15.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.GPIO15.High()
	iface := i8080.SoftSPI{
		WRPin:    machine.GPIO8,
		RDPin:    machine.GPIO9,
		DataPins: []machine.Pin{machine.GPIO39, machine.GPIO40, machine.GPIO41, machine.GPIO42, machine.GPIO45, machine.GPIO46, machine.GPIO47, machine.GPIO48},
	}

	err := iface.Init()
	if err != nil {
		panic(err)
	}

	display := st7789.New(iface, machine.GPIO5, machine.GPIO7, machine.GPIO6, machine.GPIO38)
	display.Configure(st7789.Config{
		ColumnOffset: 35, //(240 - 170) / 2 = 35, https://homeding.github.io/boards/esp32s3/lilygo-t-display-s3.htm
		Width:        170,
		Height:       320,
		Rotation:     drivers.Rotation90,
	})

	display.EnableBacklight(true)
	w, h := display.Size()
	return &TDisplayS3Screen{
		disp: display,
		conf: conf,
		buf: pixel.NewImage[pixel.RGB565BE](int(w), int(h)),
	}
}