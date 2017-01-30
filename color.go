package render

import "image/color"

// Standard colors recommended for games.
var (
	Red        = color.Color(color.RGBA{211, 47, 47, 255})
	Pink       = color.Color(color.RGBA{194, 24, 91, 255})
	Purple     = color.Color(color.RGBA{123, 31, 162, 255})
	DeepPurple = color.Color(color.RGBA{81, 45, 168, 255})
	Indigo     = color.Color(color.RGBA{48, 63, 159, 255})
	Blue       = color.Color(color.RGBA{25, 118, 210, 255})
	LightBlue  = color.Color(color.RGBA{2, 136, 209, 255})
	Cyan       = color.Color(color.RGBA{0, 151, 167, 255})
	Teal       = color.Color(color.RGBA{0, 121, 107, 255})
	Green      = color.Color(color.RGBA{56, 142, 60, 255})
	LightGreen = color.Color(color.RGBA{104, 159, 56, 255})
	Lime       = color.Color(color.RGBA{175, 180, 43, 255})
	Yellow     = color.Color(color.RGBA{251, 192, 45, 255})
	Amber      = color.Color(color.RGBA{255, 160, 0, 255})
	Orange     = color.Color(color.RGBA{245, 124, 0, 255})
	DeepOrange = color.Color(color.RGBA{230, 74, 25, 255})
	Brown      = color.Color(color.RGBA{93, 64, 55, 255})
	Grey       = color.Color(color.RGBA{97, 97, 97, 255})
	BlueGrey   = color.Color(color.RGBA{69, 90, 100, 255})
	White      = color.Color(color.RGBA{255, 255, 255, 255})
	Black      = color.Color(color.RGBA{0, 0, 0, 255})
)

// PlayerColors are a subset of the standard colours suitable for player
// coloring.  These should be used in correct order to match log rendering.
var PlayerColors = []color.Color{
	Green,
	Red,
	Blue,
	Yellow,
	Purple,
	Brown,
	BlueGrey,
}

// PlayerColor gets the player colour for a given player number.
func PlayerColor(p int) color.Color {
	return PlayerColors[p%len(PlayerColors)]
}

const contrastThreshold = 0x18000

// ContrastMono returns black or white to contrast with the given color.
func ContrastMono(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	if r+g+b > contrastThreshold {
		return Black
	}
	return White
}
