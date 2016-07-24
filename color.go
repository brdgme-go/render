package render

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

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
	Red:        color.RGBA{211, 47, 47, 255},
	Pink:       color.RGBA{194, 24, 91, 255},
	Purple:     color.RGBA{123, 31, 162, 255},
	DeepPurple: color.RGBA{81, 45, 168, 255},
	Indigo:     color.RGBA{48, 63, 159, 255},
	Blue:       color.RGBA{25, 118, 210, 255},
	LightBlue:  color.RGBA{2, 136, 209, 255},
	Cyan:       color.RGBA{0, 151, 167, 255},
	Teal:       color.RGBA{0, 121, 107, 255},
	Green:      color.RGBA{56, 142, 60, 255},
	LightGreen: color.RGBA{104, 159, 56, 255},
	Lime:       color.RGBA{175, 180, 43, 255},
	Yellow:     color.RGBA{251, 192, 45, 255},
	Amber:      color.RGBA{255, 160, 0, 255},
	Orange:     color.RGBA{245, 124, 0, 255},
	DeepOrange: color.RGBA{230, 74, 25, 255},
	Brown:      color.RGBA{93, 64, 55, 255},
	Grey:       color.RGBA{97, 97, 97, 255},
	BlueGrey:   color.RGBA{69, 90, 100, 255},
	White:      color.RGBA{255, 255, 255, 255},
	Black:      color.RGBA{0, 0, 0, 255},
}

// PlayerColors are a subset of the default colours suitable for player
// coloring.  These should be used in correct order to match log rendering.
var PlayerColors = []string{
	Green,
	Red,
	Blue,
	Yellow,
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
		return Colors[Black]
	}
	return Colors[White]
}

func parseUint8(s string) (uint8, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("unable to parse int, %s", err)
	}
	if i < 0 || i > 255 {
		return 0, fmt.Errorf("expect between 0 and 255, got %d", i)
	}
	return uint8(i), nil
}

func parseUint8s(strs []string) ([]uint8, error) {
	l := len(strs)
	is := make([]uint8, len(strs))
	if l == 0 {
		return is, nil
	}
	var err error
	for k, s := range strs {
		is[k], err = parseUint8(s)
		if err != nil {
			return is, fmt.Errorf("unable to parse int at index %d, %s", k, err)
		}
	}
	return is, nil
}

func ParseColor(s string) (color.Color, error) {
	// Use color names before RGB.
	if c, ok := Colors[s]; ok {
		return c, nil
	}
	// Fall back to RGB.
	parts := strings.Split(s, ",")
	if len(parts) == 3 {
		is, err := parseUint8s(parts)
		if err == nil {
			return color.RGBA{is[0], is[1], is[2], 255}, nil
		}
	}
	return nil, fmt.Errorf(
		"'%s' is not a valid color, please use a valid color string or R,G,B in decimal bytes",
		s,
	)
}
