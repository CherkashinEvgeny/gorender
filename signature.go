package renderer

type SignatureRenderer struct {
	in  Code
	out Code
	ctx Code
}

func Signature(in Code, out Code) *SignatureRenderer {
	s := &SignatureRenderer{}
	s.SetIn(in)
	s.SetOut(out)
	return s
}

func (s *SignatureRenderer) GetIn() Code {
	return s.in
}

func (s *SignatureRenderer) SetIn(in Code) {
	s.in = in
	if in != nil {
		in.setContext(s)
	}
}

func (s *SignatureRenderer) GetOut() Code {
	return s.out
}

func (s *SignatureRenderer) SetOut(out Code) {
	s.out = out
	if out != nil {
		out.setContext(s)
	}
}

func (s *SignatureRenderer) getContext() Code {
	return s.ctx
}

func (s *SignatureRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *SignatureRenderer) render(w *Writer) {
	s.in.render(w)
	if s.out == nil {
		return
	}
	out, ok := s.out.(*OutRenderer)
	if !ok {
		return
	}
	if out.Len() == 0 {
		return
	}
	w.Write(" ")
	s.out.render(w)
}

type InRenderer struct {
	items []Code
	ctx   Code
}

func In(items ...Code) *InRenderer {
	l := &InRenderer{}
	l.Add(items...)
	return l
}

func (l *InRenderer) Len() int {
	return len(l.items)
}

func (l *InRenderer) At(i int) Code {
	return l.items[i]
}

func (l *InRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *InRenderer) getContext() Code {
	return l.ctx
}

func (l *InRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *InRenderer) render(w *Writer) {
	w.Write("(")
	for i, param := range l.items {
		if i != 0 {
			w.Write(", ")
		}
		param.render(w)
	}
	w.Write(")")
}

type OutRenderer struct {
	items []Code
	ctx   Code
}

func Out(items ...Code) *OutRenderer {
	l := &OutRenderer{}
	l.Add(items...)
	return l
}

func (l *OutRenderer) Len() int {
	return len(l.items)
}

func (l *OutRenderer) At(i int) Code {
	return l.items[i]
}

func (l *OutRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *OutRenderer) getContext() Code {
	return l.ctx
}

func (l *OutRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *OutRenderer) render(w *Writer) {
	if l.Len() > 1 {
		w.Write("(")
	}
	for i, param := range l.items {
		if i != 0 {
			w.Write(", ")
		}
		param.render(w)
	}
	if l.Len() > 1 {
		w.Write(")")
	}
}
