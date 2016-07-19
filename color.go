package render

import "image/color"

const (
	Red        = "red"
	Pink       = "pink"
	Purple     = "purple"
	DeepPurple = "deeppurple"
	Indigo     = "indigo"
	Blue       = "blue"
	LightBlue  = "lightblue"
	Cyan       = "cyan"
	Teal       = "teal"
	Green      = "green"
	LightGreen = "lightgreen"
	Lime       = "lime"
	Yellow     = "yellow"
	Amber      = "amber"
	Orange     = "orange"
	DeepOrange = "deeporange"
	Brown      = "brown"
	Grey       = "grey"
	BlueGrey   = "bluegrey"
	White      = "white"
	Black      = "black"
)

// Colors for use in brdgme games.
var Colors = map[string]color.Color{
	Red:        color.RGBA{244, 67, 54, 255},
	Pink:       color.RGBA{233, 30, 99, 255},
	Purple:     color.RGBA{156, 39, 176, 255},
	DeepPurple: color.RGBA{103, 58, 183, 255},
	Indigo:     color.RGBA{63, 81, 181, 255},
	Blue:       color.RGBA{33, 150, 243, 255},
	LightBlue:  color.RGBA{3, 169, 244, 255},
	Cyan:       color.RGBA{0, 188, 212, 255},
	Teal:       color.RGBA{0, 150, 136, 255},
	Green:      color.RGBA{76, 175, 80, 255},
	LightGreen: color.RGBA{139, 195, 74, 255},
	Lime:       color.RGBA{205, 220, 57, 255},
	Yellow:     color.RGBA{255, 235, 59, 255},
	Amber:      color.RGBA{255, 193, 7, 255},
	Orange:     color.RGBA{255, 152, 0, 255},
	DeepOrange: color.RGBA{255, 87, 34, 255},
	Brown:      color.RGBA{121, 85, 72, 255},
	Grey:       color.RGBA{158, 158, 158, 255},
	BlueGrey:   color.RGBA{96, 125, 139, 255},
	White:      color.RGBA{255, 255, 255, 255},
	Black:      color.RGBA{0, 0, 0, 255},
}

// PlayerColors are a subset of the default colours suitable for player
// coloring.  These should be used in correct order to match log rendering.
var PlayerColors = []string{
	Green,
	Red,
	Blue,
	Orange,
	Purple,
	Brown,
	BlueGrey,
}

// PlayerColor gets the player colour for a given player number.
func PlayerColor(p int) string {
	return PlayerColors[p%len(PlayerColors)]
}

const contrastThreshold = 0x18000

func ContrastMono(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	if r+g+b > contrastThreshold {
		return Colors[White]
	}
	return Colors[Black]
}
