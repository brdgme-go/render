package render

import (
	"errors"
	"fmt"
	"image/color"
	"strings"
)

type style struct {
	fg, bg color.Color
	bold   bool
}

func (s style) String() string {
	boldStr := ""
	if s.bold {
		boldStr = "1;"
	}
	fr, fg, fb, _ := s.fg.RGBA()
	br, bg, bb, _ := s.bg.RGBA()
	return fmt.Sprintf(
		"\x1b[0;%s38;2;%d;%d;%d;48;2;%d;%d;%dm",
		boldStr,
		fr/256,
		fg/256,
		fb/256,
		br/256,
		bg/256,
		bb/256,
	)
}

var defaultStyle = style{
	fg:   Colors[Black],
	bg:   Colors[White],
	bold: false,
}

func TermMarkupFuncs(playerNames []string) map[string]MarkupFunc {
	mf := DefaultMarkupFuncs(playerNames)
	styleStack := []style{defaultStyle}
	pushStyle := func(f func(style) style) (string, error) {
		l := len(styleStack)
		if l == 0 {
			return "", errors.New("styleStack is empty")
		}
		oldStyle := styleStack[l-1]
		newStyle := f(oldStyle)
		styleStack = append(styleStack, newStyle)
		return newStyle.String(), nil
	}
	popStyle := func() (string, error) {
		l := len(styleStack)
		if l <= 1 {
			return "", errors.New("styleStack is empty")
		}
		styleStack = styleStack[:l-1]
		return styleStack[l-2].String(), nil
	}
	mf["#b"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{#b}} takes no arguments")
		}
		return pushStyle(func(s style) style {
			s.bold = true
			return s
		})
	}
	mf["/b"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{#b}} takes no arguments")
		}
		return popStyle()
	}
	mf["#fg"] = func(args []string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("{{#fg}} expects one argument")
		}
		c, err := ParseColor(args[0])
		if err != nil {
			return "", err
		}
		return pushStyle(func(s style) style {
			s.fg = c
			return s
		})
	}
	mf["/fg"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{/fg}} takes no arguments")
		}
		return popStyle()
	}
	mf["#bg"] = func(args []string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("{{#bg}} expects one argument")
		}
		c, err := ParseColor(args[0])
		if err != nil {
			return "", err
		}
		return pushStyle(func(s style) style {
			s.bg = c
			return s
		})
	}
	mf["/bg"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{/bg}} takes no arguments")
		}
		return popStyle()
	}
	return mf
}

func RenderTerm(content string, playerNames []string) (string, error) {
	output, err := ParseMarkup(content, TermMarkupFuncs(playerNames))
	if err != nil {
		return "", err
	}
	lines := strings.Split(output, "\n")
	for i, l := range lines {
		lines[i] = fmt.Sprintf("%s\x1b[K", l)
	}
	return fmt.Sprintf(
		"%s%s%s",
		defaultStyle.String(),
		strings.Join(lines, "\n"),
		"\x1b[0m",
	), nil
}
