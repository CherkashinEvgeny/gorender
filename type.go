package renderer

type TypeRenderer struct {
	name    string
	ttype   Code
	methods Code
	ctx     Code
}

func Type(name string, ttype Code) *TypeRenderer {
	t := &TypeRenderer{}
	t.SetName(name)
	t.SetType(ttype)
	return t
}

func (t *TypeRenderer) GetName() string {
	return t.name
}

func (t *TypeRenderer) SetName(name string) {
	t.name = name
}

func (t *TypeRenderer) GetType() Code {
	return t.ttype
}

func (t *TypeRenderer) SetType(ttype Code) {
	t.ttype = ttype
}

func (t *TypeRenderer) getContext() Code {
	return t.ctx
}

func (t *TypeRenderer) setContext(ctx Code) {
	t.ctx = ctx
}

func (t *TypeRenderer) render(w *Writer) {
	w.Write("type")
	w.Write(" ")
	w.Write(t.name)
	w.Write(" ")
	t.ttype.render(w)
	if t.methods != nil {
		w.Br()
		w.Br()
		t.methods.render(w)
	}
}
