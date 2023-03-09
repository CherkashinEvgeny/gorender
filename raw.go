package renderer

type Raw string

func (c Raw) getContext() Code {
	return nil
}

func (c Raw) setContext(_ Code) {
}

func (c Raw) render(w *Writer) {
	w.Write(string(c))
}
