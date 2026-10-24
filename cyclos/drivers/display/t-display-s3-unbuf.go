//go:build d_tdisplays3_unbuf

package display

import (
	"image/color"
	"machine"

	"github.com/axellse/cyclos/cyclos/config"
	spii8080bitbang "github.com/axellse/cyclos/cyclos/drivers/lib/spi-i8080-bitbang"
	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/st7789"
)

type TDisplayS3Screen struct {
	disp st7789.Device
	conf config.DisplayConfig
}

func (d *TDisplayS3Screen) Size() (width, height int16) {
	return 320, 170
}

func (d *TDisplayS3Screen) BeginDisplay() error {
	if d.conf.HideDrawing {
		d.disp.EnableBacklight(false)
	}

	return nil
}

func (d *TDisplayS3Screen) Display() error {
	if d.conf.HideDrawing {
		d.disp.EnableBacklight(true)
	}

	return nil
}

func (d *TDisplayS3Screen) SetPixel(x, y int16, c color.RGBA) {
	d.disp.SetPixel(x, y, c)
}

func Init(conf config.DisplayConfig) *TDisplayS3Screen {
	machine.GPIO15.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.GPIO15.High()
	iface := spii8080bitbang.SPII8080Bitbang{
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
	return &TDisplayS3Screen{
		disp: display,
		conf: conf,
	}
}
