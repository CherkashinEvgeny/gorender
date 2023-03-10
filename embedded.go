package renderer

type EmbeddedRenderer struct {
	ftype Code
	ctx   Code
}

func Embedded(ftype Code) *EmbeddedRenderer {
	f := &EmbeddedRenderer{}
	f.SetType(ftype)
	return f
}

func (f *EmbeddedRenderer) GetType() Code {
	return f.ftype
}

func (f *EmbeddedRenderer) SetType(ftype Code) {
	f.ftype = ftype
	if ftype != nil {
		ftype.setContext(f)
	}
}

func (f *EmbeddedRenderer) getContext() Code {
	return f.ctx
}

func (f *EmbeddedRenderer) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *EmbeddedRenderer) render(w *Writer) {
	f.ftype.render(w)
}
