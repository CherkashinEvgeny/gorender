package renderer

type IdList struct {
	items []Code
	ctx   Code
}

func (l *IdList) Len() int {
	return len(l.items)
}

func (l *IdList) Get(i int) Code {
	return l.items[i]
}

func (l *IdList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *IdList) getContext() Code {
	return l.ctx
}

func (l *IdList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *IdList) render(w *Writer) {
	for i, item := range l.items {
		if i != 0 {
			w.Write(", ")
		}
		item.render(w)
	}
}

type Id string

func (i Id) getContext() Code {
	return nil
}

func (i Id) setContext(ctx Code) {
}

func (i Id) render(w *Writer) {
	w.Write(string(i))
}
