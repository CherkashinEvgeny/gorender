package renderer

type Pointer struct {
	elem Code
	ctx  Code
}

func NewPointer(elem Code) *Pointer {
	p := &Pointer{}
	p.SetElem(elem)
	return p
}

func (p *Pointer) GetElem() Code {
	return p.elem
}

func (p *Pointer) SetElem(elem Code) {
	p.elem = elem
	if elem != nil {
		elem.setContext(p)
	}
}

func (p *Pointer) getContext() Code {
	return p.ctx
}

func (p *Pointer) setContext(ctx Code) {
	p.ctx = ctx
}

func (p *Pointer) render(w *Writer) {
	w.Write("*")
	p.elem.render(w)
}
