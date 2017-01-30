package render

import (
	"bytes"
	"fmt"
	"image/color"
)

// ANSI renders the node using ANSI escape codes.
func ANSI(nodes []Node) string {
	return ansi(withDefaultStyle(transform(nodes)), defaultFg, defaultBg, false)
}

func rgb256SepSC(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%d;%d;%d", r/256, g/256, b/256)
}

func ansiEscape(fg, bg color.Color, bold bool) string {
	boldFlag := 0
	if bold {
		boldFlag = 1
	}
	return fmt.Sprintf(
		"\x1b[%d;38;2;%s;48;2;%sm",
		boldFlag,
		rgb256SepSC(fg),
		rgb256SepSC(bg),
	)
}

func ansi(nodes []Node, fg, bg color.Color, bold bool) string {
	output := bytes.NewBuffer(nil)
	for _, n := range nodes {
		switch n := n.(type) {
		case Bold:
			if !bold {
				output.WriteString(ansiEscape(fg, bg, true))
			}
			output.WriteString(ansi(n.Ch, fg, bg, true))
			if !bold {
				output.WriteString(ansiEscape(fg, bg, false))
			}
		case Fg:
			if n.Col != fg {
				output.WriteString(ansiEscape(n.Col, bg, bold))
			}
			output.WriteString(ansi(n.Ch, n.Col, bg, bold))
			if n.Col != fg {
				output.WriteString(ansiEscape(fg, bg, bold))
			}
		case Bg:
			if n.Col != bg {
				output.WriteString(ansiEscape(fg, n.Col, bold))
			}
			output.WriteString(ansi(n.Ch, fg, n.Col, bold))
			if n.Col != bg {
				output.WriteString(ansiEscape(fg, bg, bold))
			}
		case Text:
			output.WriteString(string(n))
		default:
			panic(fmt.Sprintf("Unknown node: %#v", n))
		}
	}
	return output.String()
}
