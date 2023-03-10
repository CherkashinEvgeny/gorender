package renderer

type JoinRenderer struct {
	items []Code
	ctx   Code
}

func Join(codes ...Code) *JoinRenderer {
	i := &JoinRenderer{}
	i.Add(codes...)
	return i
}

func (s *JoinRenderer) Len() int {
	return len(s.items)
}

func (s *JoinRenderer) At(i int) Code {
	return s.items[i]
}

func (s *JoinRenderer) Add(items ...Code) {
	s.items = append(s.items, items...)
	for _, item := range items {
		item.setContext(s)
	}
}

func (s *JoinRenderer) getContext() Code {
	return s.ctx
}

func (s *JoinRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *JoinRenderer) render(w *Writer) {
	for _, item := range s.items {
		item.render(w)
	}
}
