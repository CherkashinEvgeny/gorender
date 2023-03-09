package renderer

import "strings"

type Comment string

func (c Comment) getContext() Code {
	return nil
}

func (c Comment) setContext(_ Code) {
}

func (c Comment) render(w *Writer) {
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
