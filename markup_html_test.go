package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderHtml(t *testing.T) {
	_, err := RenderHtml(`
Actually, {{player 0}} and {{player 1}}, {{#b}}this text should be bold{{/b}}.

Here is some {{#fg purple}}purple text{{/fg}} and here is a {{#bgc green}}green background{{/bgc}}.
`, []string{"Mick", "Steve"})
	assert.NoError(t, err)
}
