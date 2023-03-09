package renderer

type FrameList struct {
	items []Code
	ctx   Code
}

func NewFrameList(items ...Code) *FrameList {
	l := &FrameList{}
	l.Add(items...)
	return l
}

func (l *FrameList) Len() int {
	return len(l.items)
}

func (l *FrameList) Get(i int) Code {
	return l.items[i]
}

func (l *FrameList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *FrameList) getContext() Code {
	return l.ctx
}

func (l *FrameList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *FrameList) render(w *Writer) {
	for i, item := range l.items {
		if i != 0 {
			w.Br()
			w.Br()
		}
		item.render(w)
	}
}

type Frame struct {
	value Code
	ctx   Code
}

func NewFrame(value Code) *Frame {
	f := &Frame{}
	f.SetValue(value)
	return f
}

func (f *Frame) GetValue() Code {
	return f.value
}

func (f *Frame) SetValue(code Code) {
	f.value = code
	if code != nil {
		code.setContext(f)
	}
}

func (f *Frame) getContext() Code {
	return f.ctx
}

func (f *Frame) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *Frame) render(w *Writer) {
	f.value.render(w)
}
