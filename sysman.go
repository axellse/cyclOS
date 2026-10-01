package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/axellse/cyclos/config"
	"github.com/axellse/cyclos/hal"
	lexenddeca "github.com/axellse/cyclos/ui/fonts/lexend-deca"
	"github.com/axellse/cyclos/ui/fonts/tratex"
	"tinygo.org/x/tinyfont"
)



func main() {
	time.Sleep(5 * time.Second)
	config.LoadConfig()
	hInterface := hal.InitalizeHardware()

	for speed := 0; speed < 100; speed++ {
		hInterface.Display.BeginDisplay()
		for y := int16(0); y < 170; y++ {
			for x := int16(0); x < 320; x++ {
				hInterface.Display.SetPixel(x,y,color.RGBA{21,21,21,255})
			}
		}
		fmt.Println("set base color")

		tinyfont.WriteLine(hInterface.Display, &tratex.Tratex, 5,160, strconv.Itoa(speed), color.RGBA{255,255,255,255})
		fmt.Println("wrote big text")
		_, w2 := tinyfont.LineWidth(&tratex.Tratex, strconv.Itoa(speed))
		fmt.Print("writing smaller text")
		tinyfont.WriteLine(hInterface.Display, &lexenddeca.LexendDeca, int16(w2) + 10,160, "km/h", color.RGBA{137,137,137,255})
		tinyfont.WriteLine(hInterface.Display, &lexenddeca.LexendDeca, 230, 23, "22:01", color.RGBA{137,137,137,255})
		fmt.Println("calling display")
		hInterface.Display.Display()
		time.Sleep(1 * time.Second)
	}
}
