package hal

import (
	"github.com/axellse/cyclos/config"
	"github.com/axellse/cyclos/drivers/display"
	"github.com/axellse/cyclos/ihal"
)


type HAL struct {
	Display ihal.Display
}

func InitalizeHardware() *HAL {
	return &HAL{
		Display: display.Init(config.GlobalConfig.Display),
	}
}
