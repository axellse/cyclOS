package spii8080bitbang

import (
	"errors"
	"machine"
)

//SPII8080Bitbang implements drivers.SPI but transfers the data using 8 bit parallel. It is "bit-banged", eg. written in software without taking advantage of special hardware available on certain processors. It does not support reading for the moment.
type SPII8080Bitbang struct {
	DataPins []machine.Pin
	WRPin    machine.Pin
	RDPin    machine.Pin
}

func (b SPII8080Bitbang) Init() error {
	b.WRPin.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	b.RDPin.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	b.WRPin.High()
	b.RDPin.High()

	if len(b.DataPins) != 8 {
		return errors.New("invalid datapin length, must be 8")
	}

	for _, dp := range b.DataPins {
		dp.Configure(machine.PinConfig{
			Mode: machine.PinOutput,
		})
	}
	return nil

}

func (i SPII8080Bitbang) Tx(w, r []byte) error {
	for _, v := range w {
		i.WRPin.Low()
		for idx, dp := range i.DataPins {
			mask := uint8(1) << idx //00000001 >> idx
			dp.Set((v & mask) != 0)
		}
		i.WRPin.High()
	}

	if r != nil {
		return errors.New("read not supported")
	}
	return nil
}

func (i SPII8080Bitbang) Transfer(b byte) (byte, error) {
	err := i.Tx([]byte{b}, nil)
	return 0, err
}
