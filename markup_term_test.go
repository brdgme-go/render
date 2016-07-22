package render

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTerm(t *testing.T) {
	output, err := ParseMarkup(`
Actually, {{player 0}} and {{player 1}}, {{#b}}this text should be bold{{/b}}.
`, TermMarkupFuncs([]string{"Mick", "Steve"}))
	assert.NoError(t, err)
	fmt.Println(defaultStyle.String() + output + "\x1b[0m")
}
