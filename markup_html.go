package render

import (
	"errors"
	"fmt"
)

func HtmlMarkupFuncs(playerNames []string) map[string]MarkupFunc {
	mf := DefaultMarkupFuncs(playerNames)
	mf["#b"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{#b}} takes no arguments")
		}
		return "<b>", nil
	}
	mf["/b"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{#b}} takes no arguments")
		}
		return "</b>", nil
	}
	mf["#fg"] = func(args []string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("{{#fg}} expects one argument")
		}
		c, err := ParseColor(args[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			`<span style="color:rgb(%s)">`,
			markupColorArg(c),
		), nil
	}
	mf["/fg"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{/fg}} takes no arguments")
		}
		return "</span>", nil
	}
	mf["#bg"] = func(args []string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("{{#bg}} expects one argument")
		}
		c, err := ParseColor(args[0])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(
			`<span style="background-color:rgb(%s)">`,
			markupColorArg(c),
		), nil
	}
	mf["/bg"] = func(args []string) (string, error) {
		if len(args) > 0 {
			return "", errors.New("{{/bg}} takes no arguments")
		}
		return "</span>", nil
	}
	return mf
}

func RenderHtml(content string, playerNames []string) (string, error) {
	output, err := ParseMarkup(content, HtmlMarkupFuncs(playerNames))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		`<pre style="color:rgb(%s);background-color:rgb(%s)">%s</pre>`,
		markupColorArg(Colors[Black]),
		markupColorArg(Colors[White]),
		output,
	), nil
}
