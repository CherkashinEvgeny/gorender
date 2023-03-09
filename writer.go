package renderer

import "strings"

type Writer struct {
	indent   int
	indented bool
	sb       strings.Builder
}

func (w *Writer) Write(str string) {
	if strings.Count(str, "\n") == 0 {
		w.writeLine(str)
		return
	}
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if i != 0 {
			w.Br()
		}
		w.writeLine(line)
	}
	return
}

func (w *Writer) writeLine(line string) {
	if line == "" {
		return
	}
	if !w.indented {
		w.writeIndent()
	}
	w.sb.WriteString(line)
}

func (w *Writer) writeIndent() {
	for i := 0; i < w.indent; i++ {
		w.sb.WriteString("\t")
	}
	w.indented = true
}

func (w *Writer) Br() {
	w.sb.WriteString("\n")
	w.indented = false
}

func (w *Writer) AddIndent() {
	w.Indent(1)
}

func (w *Writer) RemoveIndent() {
	w.Indent(-1)
}

func (w *Writer) Indent(n int) {
	w.indent += n
	if w.indent < 0 {
		w.indent = 0
	}
}

func (w *Writer) string() string {
	return w.sb.String()
}
