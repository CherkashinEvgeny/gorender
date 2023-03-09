package renderer

type Sequence struct {
	items []Code
	ctx   Code
}

func NewSequence(codes ...Code) *Sequence {
	s := &Sequence{}
	s.Add(codes...)
	return s
}

func (s *Sequence) Len() int {
	return len(s.items)
}

func (s *Sequence) Get(i int) Code {
	return s.items[i]
}

func (s *Sequence) Add(items ...Code) {
	s.items = append(s.items, items...)
	for _, item := range items {
		item.setContext(s)
	}
}

func (s *Sequence) getContext() Code {
	return s.ctx
}

func (s *Sequence) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *Sequence) render(w *Writer) {
	for _, code := range s.items {
		code.render(w)
	}
}
