package shared

import "image/color"
type HAL struct {
	Display Display
	Input Input
}

//The Display interface is a flexible interface for display drivers. It allows UI built on top to adapt to different displays and their capabilities (size and color depth).
type Display interface {
	//Size return the width and height of the display in pixels
	Size() (width, height int16)
	//BeginDisplay prepares the display for one or more SetPixel calls. The implementation of this may vary from display to display. After calling BeginDisplay and drawing is complete, Display needs to be called for the changes to propagate.
	BeginDisplay() error
	//Display instructs the display that all drawing is complete and that it may update the Display. Display always needs to be called after BeginDisplay.
	Display() error
	//SetPixel sets the pixel at x,y to the color c. BeginDisplay needs to be called before SetPixel is called.
	SetPixel(x, y int16, c color.RGBA)
}

//The Input interface is a flexible interface for human interface devices and gives the input driver full control over the Input system state object. Most drivers will fully or partially use a libinput input scheme.
type Input interface {
	//Channel returns a channel that new input state mutations will be pushed to. The input driver will call LatestState whenever an input event has occured, mutate the return value, and push into onto the channel. Channel will/can only be called once.
	Channel(LatestState func() InputState) chan InputState
}