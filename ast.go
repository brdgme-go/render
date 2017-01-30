package render

import (
	"image/color"
	"strings"
	"unicode/utf8"
)

// Alignment designates the horizontal alignment of a block.
type Alignment byte

// Alignment constants.
const (
	Left Alignment = iota + 1
	Center
	Right
)

// A Node is a node in a render tree.
type Node interface {
	transform() []Node
	toLines() [][]Node
}

func transform(nodes []Node) []Node {
	transformed := []Node{}
	for _, n := range nodes {
		transformed = append(transformed, n.transform()...)
	}
	return transformed
}

func toLines(nodes []Node) [][]Node {
	lines := [][]Node{}
	transformed := transform(nodes)
	var curLine []Node
	for _, t := range transformed {
		nLines := t.toLines()
		nLinesLen := len(nLines)
		if nLinesLen == 0 {
			continue
		}
		curLine = append(curLine, nLines[0]...)
		if nLinesLen > 1 {
			lines = append(lines, curLine)
			lines = append(lines, nLines[1:nLinesLen-1]...)
			curLine = nLines[nLinesLen-1]
		}
	}
	if curLine != nil {
		lines = append(lines, curLine)
	}
	return lines
}

func size(nodes []Node) (w, h int) {
	for _, l := range toLines(nodes) {
		lineW := 0
		h++
		for _, n := range l {
			if t, ok := n.(Text); ok {
				lineW += utf8.RuneCountInString(string(t))
			}
		}
		if lineW > w {
			w = lineW
		}
	}
	return
}

// A Text node is just text with no children.
type Text string

func (t Text) transform() []Node {
	return []Node{t}
}

func (t Text) toLines() [][]Node {
	lines := [][]Node{}
	for _, v := range strings.Split(string(t), "\n") {
		lines = append(lines, []Node{Text(v)})
	}
	return lines
}

type Table [][]Node

type Bold struct {
	Ch []Node
}

func (b Bold) transform() []Node {
	return []Node{Bold{transform(b.Ch)}}
}

func (b Bold) toLines() [][]Node {
	childLines := toLines(b.Ch)
	lines := make([][]Node, len(childLines))
	for _, l := range childLines {
		lines = append(lines, []Node{Bold{l}})
	}
	return lines
}

type Fg struct {
	Col color.Color
	Ch  []Node
}

func (fg Fg) transform() []Node {
	return []Node{Fg{Col: fg.Col, Ch: transform(fg.Ch)}}
}

func (fg Fg) toLines() [][]Node {
	childLines := toLines(fg.Ch)
	lines := make([][]Node, len(childLines))
	for _, l := range childLines {
		lines = append(lines, []Node{Fg{Col: fg.Col, Ch: l}})
	}
	return lines
}

type Bg struct {
	Col color.Color
	Ch  []Node
}

func (bg Bg) transform() []Node {
	return []Node{Bg{Col: bg.Col, Ch: transform(bg.Ch)}}
}

func (bg Bg) toLines() [][]Node {
	childLines := toLines(bg.Ch)
	lines := make([][]Node, len(childLines))
	for _, l := range childLines {
		lines = append(lines, []Node{Bg{Col: bg.Col, Ch: l}})
	}
	return lines
}

type Align struct {
	Align Alignment
	Width int
	Ch    []Node
}

func (a Align) transform() []Node {
	aligned := []Node{}
	for k := range toLines(a.Ch) {
		if k != 0 {
			aligned = append(aligned, Text("\n"))
		}
	}
	return aligned
}
