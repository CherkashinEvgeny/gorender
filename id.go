package renderer

type IdRenderer string

func Id(id string) IdRenderer {
	return IdRenderer(id)
}

func (i IdRenderer) getContext() Code {
	return nil
}

func (i IdRenderer) setContext(ctx Code) {
}

func (i IdRenderer) render(w *Writer) {
	w.Write(string(i))
}
