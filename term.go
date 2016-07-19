package render

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/aymerick/raymond"
)

const termEscape = "\x1b"

var termClearEOL = fmt.Sprintf("%s[K", termEscape)

func termHelpers() map[string]interface{} {
	attrStack := termAttrs{}
	return map[string]interface{}{
		"fg": func(color string, options *raymond.Options) raymond.SafeString {
			c, ok := Colors[color]
			if !ok {
				panic(fmt.Sprintf("'%s' is not a valid color name", color))
			}
			attrStack = attrStack.pushFg(c)
			output := attrStack.String() + options.Fn()
			attrStack = attrStack.pop()
			return raymond.SafeString(output + attrStack.String())
		},
		"bg": func(color string, options *raymond.Options) raymond.SafeString {
			c, ok := Colors[color]
			if !ok {
				panic(fmt.Sprintf("'%s' is not a valid color name", color))
			}
			attrStack = attrStack.pushBg(c)
			output := attrStack.String() + options.Fn()
			attrStack = attrStack.pop()
			return raymond.SafeString(output + attrStack.String())
		},
		"b": func(options *raymond.Options) raymond.SafeString {
			attrStack = attrStack.pushBold()
			output := attrStack.String() + options.Fn()
			attrStack = attrStack.pop()
			return raymond.SafeString(output + attrStack.String())
		},
	}
}

func termFg(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf(
		"%s[38;2;%d;%d;%dm",
		termEscape,
		r/256,
		g/256,
		b/256,
	)
}

func termBg(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf(
		"%s[48;2;%d;%d;%dm",
		termEscape,
		r/256,
		g/256,
		b/256,
	)
}

func Term(tpl *raymond.Template, context interface{}) (string, error) {
	cTpl := tpl.Clone()
	cTpl.RegisterHelpers(termHelpers())
	output, err := cTpl.Exec(context)
	if err != nil {
		return output, err
	}
	lines := strings.Split(output, "\n")
	for i, l := range lines {
		lines[i] = l + termClearEOL
	}
	return termAttrDefault.String() + strings.Join(lines, "\n"), nil
}
