package renderer

type ForRenderer struct {
	cond Code
	body Code
	ctx  Code
}

func For(cond Code, body Code) *ForRenderer {
	f := &ForRenderer{}
	f.SetCond(cond)
	f.SetBody(body)
	return f
}

func (f *ForRenderer) GetCond() Code {
	return f.cond
}

func (f *ForRenderer) SetCond(cond Code) {
	f.cond = cond
	if cond != nil {
		cond.setContext(f)
	}
}

func (f *ForRenderer) GetBody() Code {
	return f.body
}

func (f *ForRenderer) SetBody(body Code) {
	f.body = body
	if body != nil {
		body.setContext(f)
	}
}

func (f *ForRenderer) getContext() Code {
	return f.ctx
}

func (f *ForRenderer) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *ForRenderer) render(w *Writer) {
	w.Write("for ")
	if f.cond != nil {
		f.cond.render(w)
		w.Write(" ")
	}
	w.Write("{")
	if f.body != nil {
		w.Br()
		w.AddIndent()
		f.body.render(w)
		w.RemoveIndent()
		w.Br()
	}
	w.Write("}")
}
