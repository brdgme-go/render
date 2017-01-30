package render

import (
	"bytes"
	"fmt"
)

// HTML renders the node in HTML.
func HTML(nodes []Node) string {
	return html(withDefaultStyle(transform(nodes)))
}

func html(nodes []Node) string {
	output := bytes.NewBuffer(nil)
	for _, n := range nodes {
		switch n := n.(type) {
		case Bold:
			output.WriteString("<b>")
			output.WriteString(html(n.Ch))
			output.WriteString("</b>")
		case Text:
			output.WriteString(string(n))
		default:
			panic(fmt.Sprintf("Unknown node: %#v", n))
		}
	}
	return output.String()
}
