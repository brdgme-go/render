package render

import "io"

const (
	TemplateNodeTypeText = iota
	TemplateNodeTypeTag
)

type TemplateNode struct {
	Type     int
	Tag      string
	Args     []string
	Content  string
	Children []TemplateNode
}

func ParseTemplate(r io.Reader) ([]TemplateNode, error) {
	return nil, nil
}
