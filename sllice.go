package renderer

type Slice struct {
	ctx  Code
	elem Code
}

func NewSlice(elem Code) *Slice {
	s := &Slice{}
	s.SetElem(elem)
	return s
}

func (s *Slice) GetElem() Code {
	return s.elem
}

func (s *Slice) SetElem(elem Code) {
	s.elem = elem
	if elem != nil {
		elem.setContext(s)
	}
}

func (s *Slice) getContext() Code {
	return s.ctx
}

func (s *Slice) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *Slice) render(w *Writer) {
	w.Write("[]")
	s.elem.render(w)
}
