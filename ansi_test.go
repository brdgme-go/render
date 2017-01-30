package render

import (
	"fmt"
	"testing"
)

func TestAnsi(t *testing.T) {
	fmt.Println(ANSI([]Node{
		Bold{[]Node{
			Text("bold"),
			Fg{Blue, []Node{
				Text("blue"),
				Fg{Amber, []Node{
					Text("amber"),
					Bg{Red, []Node{Text("redbg")}},
				}},
				Text("blue\nwith new line"),
			}},
		}},
		Text("not bold"),
	}))
}
