package renderer

type FuncRenderer struct {
	name      string
	signature Code
	body      Code
	ctx       Code
}

func Func(name string, signature Code, body Code) *FuncRenderer {
	f := &FuncRenderer{}
	f.SetName(name)
	f.SetSignature(signature)
	f.SetBody(body)
	return f
}

func (f *FuncRenderer) GetName() string {
	return f.name
}

func (f *FuncRenderer) SetName(name string) {
	f.name = name
}

func (f *FuncRenderer) GetSignature() Code {
	return f.signature
}

func (f *FuncRenderer) SetSignature(signature Code) {
	f.signature = signature
	if signature != nil {
		signature.setContext(f)
	}
}

func (f *FuncRenderer) GetBody() Code {
	return f.body
}

func (f *FuncRenderer) SetBody(body Code) {
	f.body = body
	if body != nil {
		body.setContext(f)
	}
}

func (f *FuncRenderer) getContext() Code {
	return f.ctx
}

func (f *FuncRenderer) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *FuncRenderer) render(w *Writer) {
	w.Write("func ")
	w.Write(f.name)
	f.signature.render(w)
	w.Write(" ")
	w.Write(" {")
	if f.body != nil {
		w.Br()
		w.AddIndent()
		f.body.render(w)
		w.RemoveIndent()
		w.Br()
	}
	w.Write("}")
}
