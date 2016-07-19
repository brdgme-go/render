package render

import (
	"fmt"
	"image/color"
)

type termAttr struct {
	fg, bg color.Color
	bold   bool
}

func (t termAttr) String() string {
	b := 0
	if t.bold {
		b = 1
	}
	fg := "39"
	if t.fg != nil {
		r, g, b, _ := t.fg.RGBA()
		fg = fmt.Sprintf(
			"38;2;%d;%d;%d",
			r/256,
			g/256,
			b/256,
		)
	}
	bg := "49"
	if t.bg != nil {
		r, g, b, _ := t.bg.RGBA()
		bg = fmt.Sprintf(
			"48;2;%d;%d;%d",
			r/256,
			g/256,
			b/256,
		)
	}
	return fmt.Sprintf(
		"%s[%d;%s;%sm",
		termEscape,
		b,
		fg,
		bg,
	)
}

var termAttrDefault = termAttr{
	fg: Colors[Black],
	bg: Colors[White],
}

type termAttrs []termAttr

func (t termAttrs) current() termAttr {
	if l := len(t); l > 0 {
		return t[l-1]
	}
	return termAttrDefault
}
func (t termAttrs) pushFg(c color.Color) termAttrs {
	nt := t.current()
	nt.fg = c
	return append(t, nt)
}
func (t termAttrs) pushBg(c color.Color) termAttrs {
	nt := t.current()
	nt.bg = c
	return append(t, nt)
}
func (t termAttrs) pushBold() termAttrs {
	nt := t.current()
	nt.bold = true
	return append(t, nt)
}
func (t termAttrs) pop() termAttrs {
	if l := len(t); l > 0 {
		return t[:l-1]
	}
	return t
}
func (t termAttrs) String() string {
	return t.current().String()
}
