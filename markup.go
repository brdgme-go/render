package render

import (
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

const (
	MarkupTagNormal = iota
	MarkupTagOpen
	MarkupTagClose

	MarkupTagStart      = "{{"
	MarkupTagEnd        = "}}"
	MarkupTagBlockOpen  = "#"
	MarkupTagBlockClose = "/"
)

var (
	MarkupTagStartLen        = len(MarkupTagStart)
	MarkupTagEndLen          = len(MarkupTagEnd)
	MarkupTagBlockOpenLen    = len(MarkupTagBlockOpen)
	MarkupTagBlockCloseLen   = len(MarkupTagBlockClose)
	MarkupTagStartBlockClose = MarkupTagStart + MarkupTagBlockClose
)

type MarkupParser func(string) (string, error)

type MarkupFunc func(args []string) (string, error)

func DefaultMarkupFuncs(playerNames []string) map[string]MarkupFunc {
	return map[string]MarkupFunc{
		"player": func(args []string) (string, error) {
			if len(args) != 1 {
				return "", errors.New("{{player}} must take one argument")
			}
			p, err := strconv.Atoi(args[0])
			if err != nil || p < 0 {
				return "", fmt.Errorf(
					"{{player}} must take a positive integer as an argument, got %s",
					args[0],
				)
			}
			if l := len(playerNames); p >= l {
				return "", fmt.Errorf(
					"Invalid player number for {{player}}, expect 0-%d, got %d",
					l-1,
					p,
				)
			}
			return Bold(Fg(
				fmt.Sprintf("• %s", playerNames[p]),
				Colors[PlayerColor(p)],
			)), nil
		},
		"#bgc": func(args []string) (string, error) {
			if len(args) != 1 {
				return "", errors.New("{{#bgc}} requires the color as an argument")
			}
			c, err := ParseColor(args[0])
			if err != nil {
				return "", err
			}
			fg := ContrastMono(c)
			return fmt.Sprintf(
				"{{#bg %s}}{{#fg %s}}",
				markupColorArg(c),
				markupColorArg(fg),
			), nil
		},
		"/bgc": func(args []string) (string, error) {
			return "{{/fg}}{{/bg}}", nil
		},
	}
}

func ParseTag(s string) (
	matched string,
	tagType int,
	tag string,
	args []string,
	err error,
) {
	if !strings.HasPrefix(s, MarkupTagStart) {
		err = errors.New("does not start with correct tag open")
		return
	}
	endOffset := strings.Index(s, MarkupTagEnd)
	if endOffset == -1 {
		err = errors.New("unterminated tag")
		return
	}
	fields := strings.Fields(s[MarkupTagStartLen:endOffset])
	if len(fields) == 0 {
		err = errors.New("empty tag")
		return
	}
	tagType = MarkupTagNormal
	tag = fields[0]
	if strings.HasPrefix(tag, MarkupTagBlockOpen) {
		tagType = MarkupTagOpen
	} else if strings.HasPrefix(tag, MarkupTagBlockClose) {
		tagType = MarkupTagClose
	}
	if tag == "" {
		err = errors.New("blank tag")
		return
	}
	args = fields[1:]
	matched = s[:endOffset+MarkupTagEndLen]
	return
}

func ParseMarkup(in string, f map[string]MarkupFunc) (string, error) {
	l := len(in)
	if l == 0 {
		return in, nil
	}
	offset := 0
	output := &bytes.Buffer{}
	for {
		nextTag := strings.Index(in[offset:], MarkupTagStart)
		if nextTag == -1 {
			// No more tags, finish
			output.WriteString(in[offset:])
			return output.String(), nil
		}
		output.WriteString(in[offset : offset+nextTag])
		matched, _, tag, args, err := ParseTag(in[offset+nextTag:])
		if err != nil {
			return "", err
		}
		ff, ok := f[tag]
		if !ok {
			return "", fmt.Errorf("no function for %s", tag)
		}
		funcOutput, err := ff(args)
		if err != nil {
			return "", err
		}
		offset += nextTag + len(matched)
		tail := in[offset:]
		in = in[:offset] + funcOutput + tail
	}
}

func Bold(s string) string {
	return fmt.Sprintf("{{#b}}%s{{/b}}", s)
}

func markupColorArg(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("%d,%d,%d", r/256, g/256, b/256)
}

func Fg(s string, c color.Color) string {
	return fmt.Sprintf("{{#fg %s}}%s{{/fg}}", markupColorArg(c), s)
}

func Bg(s string, c color.Color) string {
	return fmt.Sprintf("{{#bg %s}}%s{{/bg}}", markupColorArg(c), s)
}
