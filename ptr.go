package renderer

type PtrRenderer struct {
	elem Code
	ctx  Code
}

func Ptr(elem Code) *PtrRenderer {
	p := &PtrRenderer{}
	p.SetElem(elem)
	return p
}

func (p *PtrRenderer) GetElem() Code {
	return p.elem
}

func (p *PtrRenderer) SetElem(elem Code) {
	p.elem = elem
	if elem != nil {
		elem.setContext(p)
	}
}

func (p *PtrRenderer) getContext() Code {
	return p.ctx
}

func (p *PtrRenderer) setContext(ctx Code) {
	p.ctx = ctx
}

func (p *PtrRenderer) render(w *Writer) {
	w.Write("*")
	p.elem.render(w)
}
