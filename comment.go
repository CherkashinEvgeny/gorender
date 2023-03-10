package renderer

import "strings"

type CommentRenderer string

func Comment(str string) CommentRenderer {
	return CommentRenderer(str)
}

func (c CommentRenderer) getContext() Code {
	return nil
}

func (c CommentRenderer) setContext(_ Code) {
}

func (c CommentRenderer) render(w *Writer) {
	lines := strings.Split(string(c), "\n")
	for i, line := range lines {
		if i != 0 {
			w.Br()
		}
		if !strings.HasPrefix(line, "//") {
			w.Write("// ")
		}
		w.Write(line)
	}
}
