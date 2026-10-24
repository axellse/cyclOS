package libui

import (
	"image/color"

	"github.com/axellse/cyclos/cyclos/shared"
	"github.com/axellse/cyclos/cyclos/ui/icons"
)

type Point struct {
	X, Y int16
}

type Box struct {
	Width, Height int16
}

func DrawIcon(display shared.Display, start Point, icon icons.Icon, color color.RGBA) {
	bitmap := []byte(icon.Blob)
	for y := int16(0); y < icon.Height; y++ {
		for x := int16(0); x < icon.Width; x++ {
			byteIndex := ((y * icon.Width) + x) / 8
			bitIndex := ((y * icon.Width) + x) % 8

			wb := bitmap[byteIndex]
			mask := 0b10000000 >> bitIndex

			if wb & byte(mask) != 0 {
				display.SetPixel(start.X + x, start.Y + y, color)
			}
		}
	}
}

func FillRectangle(display shared.Display, start Point, end Point, color color.RGBA) {
	for y := start.Y; y < end.Y; y++ {
		for x := start.X; x < end.X; x++ {
			display.SetPixel(x, y, color)
		}
	}
}

func FillCircleQuadrant(display shared.Display, center Point, radius int16, quadrantX int16, quadrantY int16, color color.RGBA) {
	//quadrantX
	//0 = center.X - radius, center.X
	//1 = center.X, center.X + radius

	//quadrantY
	//0 = center.Y - radius, center.Y
	//1 = center.Y, center.Y + radius
	for x := center.X + (radius * (quadrantX-1)); x < center.X + (radius * quadrantX); x++ {
		for y := center.Y + (radius * (quadrantY-1)); y < center.Y + (radius * quadrantY); y++ {
			rx := x - center.X
			ry := y - center.Y
			if ((rx*rx) + (ry*ry)) <= (radius*radius) { //pythagoras thereom
				display.SetPixel(x, y, color)
			}
		}
	}
}

func FillWholeCircle(display shared.Display, center Point, radius int16, color color.RGBA) {
	FillCircleQuadrant(display, center, radius, 0, 0, color)
	FillCircleQuadrant(display, center, radius, 0, 1, color)
	FillCircleQuadrant(display, center, radius, 1, 0, color)
	FillCircleQuadrant(display, center, radius, 1, 1, color)
}

func FillRoundedRectangle(display shared.Display, start Point, end Point, radius int16, clr color.RGBA) {
	topLeft := Point{X: start.X + radius, Y: start.Y + radius}
	topRight := Point{X: end.X - radius, Y: start.Y + radius}
	bottomLeft := Point{X: start.X + radius, Y: end.Y - radius}
	bottomRight := Point{X: end.X - radius, Y: end.Y - radius}
	FillCircleQuadrant(display, topLeft, radius, 0, 0, clr)
	FillCircleQuadrant(display, topRight, radius, 1, 0, clr)
	FillCircleQuadrant(display, bottomLeft, radius, 0, 1, clr)
	FillCircleQuadrant(display, bottomRight, radius, 1, 1, clr)

	FillRectangle(display, Point{X: start.X + radius, Y: start.Y}, topRight, clr) //top
	FillRectangle(display, bottomLeft, Point{X: end.X - radius, Y: end.Y}, clr) //bottom
	FillRectangle(display, Point{X: start.X, Y: start.Y + radius}, bottomLeft, clr) //left
	FillRectangle(display, topRight, Point{X: end.X, Y: end.Y - radius}, clr) //right

	FillRectangle(display, topLeft, bottomRight, clr)
}