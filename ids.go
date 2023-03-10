package renderer

type IdsRenderer struct {
	items []Code
	ctx   Code
}

func Ids(ids ...string) *IdsRenderer {
	l := &IdsRenderer{}
	for _, id := range ids {
		l.Add(Id(id))
	}
	return l
}

func (l *IdsRenderer) Len() int {
	return len(l.items)
}

func (l *IdsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *IdsRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *IdsRenderer) getContext() Code {
	return l.ctx
}

func (l *IdsRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *IdsRenderer) render(w *Writer) {
	for i, item := range l.items {
		if i != 0 {
			w.Write(", ")
		}
		item.render(w)
	}
}
