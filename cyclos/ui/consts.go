package ui

import (
	"image/color"

)

var BasicMargin = int16(5)
var SidebarWidth = int16(30)
const TopbarHeight = int16(25) //this constant is not used by baseUi to calculate anything, and can not be changed without modification to baseUi. 

var AccentColor = color.RGBA{60, 77, 255, 255}
var IconColor = color.RGBA{219, 219, 219, 255}

var BackgroundColor = color.RGBA{21, 21, 21, 255}
var SecondaryBackgroundColor = color.RGBA{7, 7, 8, 255}

var PrimaryTextColor = color.RGBA{255, 255, 255, 255}
var SecondaryTextColor = color.RGBA{137, 137, 137, 255}