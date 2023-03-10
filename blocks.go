package renderer

type BlocksRenderer struct {
	items []Code
	ctx   Code
}

func Blocks(items ...Code) *BlocksRenderer {
	i := &BlocksRenderer{}
	i.Add(items...)
	return i
}

func (l *BlocksRenderer) Len() int {
	return len(l.items)
}

func (l *BlocksRenderer) At(i int) Code {
	return l.items[i]
}

func (l *BlocksRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *BlocksRenderer) getContext() Code {
	return l.ctx
}

func (l *BlocksRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *BlocksRenderer) render(w *Writer) {
	for i, item := range l.items {
		if i != 0 {
			w.Br()
			w.Br()
		}
		item.render(w)
	}
}
