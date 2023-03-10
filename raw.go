package renderer

type RawRenderer string

func Raw(code string) RawRenderer {
	return RawRenderer(code)
}

func (c RawRenderer) getContext() Code {
	return nil
}

func (c RawRenderer) setContext(_ Code) {
}

func (c RawRenderer) render(w *Writer) {
	w.Write(string(c))
}
