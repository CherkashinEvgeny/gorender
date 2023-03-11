package renderer

type SignRenderer struct {
	in  Code
	out Code
	ctx Code
}

func Sign(in Code, out Code) *SignRenderer {
	s := &SignRenderer{}
	s.SetIn(in)
	s.SetOut(out)
	return s
}

func (s *SignRenderer) GetIn() Code {
	return s.in
}

func (s *SignRenderer) SetIn(in Code) {
	s.in = in
	if in != nil {
		in.setContext(s)
	}
}

func (s *SignRenderer) GetOut() Code {
	return s.out
}

func (s *SignRenderer) SetOut(out Code) {
	s.out = out
	if out != nil {
		out.setContext(s)
	}
}

func (s *SignRenderer) getContext() Code {
	return s.ctx
}

func (s *SignRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *SignRenderer) render(w *Writer) {
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
	i := &InRenderer{}
	i.Add(items...)
	return i
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
	i := &OutRenderer{}
	i.Add(items...)
	return i
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
	brackets := l.needBrackets()
	if brackets {
		w.Write("(")
	}
	for i, param := range l.items {
		if i != 0 {
			w.Write(", ")
		}
		param.render(w)
	}
	if brackets {
		w.Write(")")
	}
}

func (l *OutRenderer) needBrackets() bool {
	if len(l.items) == 0 {
		return false
	}
	for _, item := range l.items {
		param, ok := item.(*ParamRenderer)
		if ok && param.name != "" {
			return true
		}
	}
	return false
}
