package renderer

type FuncTypeRenderer struct {
	signature Code
	ctx       Code
}

func FuncType(in Code, out Code) *FuncTypeRenderer {
	f := &FuncTypeRenderer{}
	f.SetSignature(Sign(in, out))
	return f
}

func (f *FuncTypeRenderer) GetSignature() Code {
	return f.signature
}

func (f *FuncTypeRenderer) SetSignature(signature Code) {
	f.signature = signature
	if signature != nil {
		signature.setContext(f)
	}
}

func (f *FuncTypeRenderer) getContext() Code {
	return f.ctx
}

func (f *FuncTypeRenderer) setContext(ctx Code) {
	f.ctx = ctx
	if f.signature != nil {
		f.signature.setContext(ctx)
	}
}

func (f *FuncTypeRenderer) render(w *Writer) {
	w.Write("func")
	f.signature.render(w)
}
