package render

import (
	"errors"
	"fmt"
	"image/color"
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
	styleChange := func(
		content string,
		mp MarkupParser,
		f func(style) style,
	) (string, error) {
		l := len(styleStack)
		if l == 0 {
			return "", errors.New("styleStack is empty")
		}
		oldStyle := styleStack[l-1]
		newStyle := f(oldStyle)
		styleStack = append(styleStack, newStyle)
		parsedContent, err := mp(content)
		if err != nil {
			return "", err
		}
		styleStack = styleStack[:l]
		return fmt.Sprintf(
			"%s%s%s",
			newStyle,
			parsedContent,
			oldStyle,
		), nil
	}
	mf["b"] = func(content string, args []string, mp MarkupParser) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{#b}} takes no arguments")
		}
		return styleChange(content, mp, func(s style) style {
			s.bold = true
			return s
		})
	}
	mf["fg"] = func(content string, args []string, mp MarkupParser) (string, error) {
		if len(args) != 1 {
			return "", errors.New("{{#fg}} expects one argument")
		}
		c, err := ParseColor(args[0])
		if err != nil {
			return "", err
		}
		return styleChange(content, mp, func(s style) style {
			s.fg = c
			return s
		})
	}
	mf["bg"] = func(content string, args []string, mp MarkupParser) (string, error) {
		if len(args) != 1 {
			return "", errors.New("{{#bg}} expects one argument")
		}
		c, err := ParseColor(args[0])
		if err != nil {
			return "", err
		}
		return styleChange(content, mp, func(s style) style {
			s.bg = c
			return s
		})
	}
	return mf
}
