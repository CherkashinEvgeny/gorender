package renderer

type LinesRenderer struct {
	items []Code
	ctx   Code
}

func Lines(codes ...Code) *LinesRenderer {
	s := &LinesRenderer{}
	s.Add(codes...)
	return s
}

func (s *LinesRenderer) Len() int {
	return len(s.items)
}

func (s *LinesRenderer) At(i int) Code {
	return s.items[i]
}

func (s *LinesRenderer) Add(items ...Code) {
	s.items = append(s.items, items...)
	for _, item := range items {
		item.setContext(s)
	}
}

func (s *LinesRenderer) getContext() Code {
	return s.ctx
}

func (s *LinesRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *LinesRenderer) render(w *Writer) {
	for i, item := range s.items {
		if i != 0 {
			w.Br()
		}
		item.render(w)
	}
}
