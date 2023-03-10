package renderer

type SliceRenderer struct {
	ctx  Code
	elem Code
}

func Slice(elem Code) *SliceRenderer {
	s := &SliceRenderer{}
	s.SetElem(elem)
	return s
}

func (s *SliceRenderer) GetElem() Code {
	return s.elem
}

func (s *SliceRenderer) SetElem(elem Code) {
	s.elem = elem
	if elem != nil {
		elem.setContext(s)
	}
}

func (s *SliceRenderer) getContext() Code {
	return s.ctx
}

func (s *SliceRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *SliceRenderer) render(w *Writer) {
	w.Write("[]")
	s.elem.render(w)
}
