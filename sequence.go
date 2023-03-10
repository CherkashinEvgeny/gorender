package renderer

type SequenceRenderer struct {
	items []Code
	ctx   Code
}

func Cat(codes ...Code) *SequenceRenderer {
	s := &SequenceRenderer{}
	s.Add(codes...)
	return s
}

func (s *SequenceRenderer) Len() int {
	return len(s.items)
}

func (s *SequenceRenderer) At(i int) Code {
	return s.items[i]
}

func (s *SequenceRenderer) Add(items ...Code) {
	s.items = append(s.items, items...)
	for _, item := range items {
		item.setContext(s)
	}
}

func (s *SequenceRenderer) getContext() Code {
	return s.ctx
}

func (s *SequenceRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *SequenceRenderer) render(w *Writer) {
	for _, item := range s.items {
		item.render(w)
	}
}
