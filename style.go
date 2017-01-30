package render

import "image/color"

var (
	defaultFg = Black
	defaultBg = White
)

func withDefaultStyle(nodes []Node) []Node {
	return []Node{Fg{defaultFg, []Node{Bg{defaultBg, nodes}}}}
}

type style struct {
	fg   color.Color
	bg   color.Color
	bold bool
}
