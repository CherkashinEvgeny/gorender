package renderer

type StructRenderer struct {
	fields Code
	ctx    Code
}

func Struct(fields Code) *StructRenderer {
	s := &StructRenderer{}
	s.SetFields(fields)
	return s
}

func (s *StructRenderer) GetFields() Code {
	return s.fields
}

func (s *StructRenderer) SetFields(fields Code) {
	s.fields = fields
	if fields != nil {
		fields.setContext(s)
	}
}

func (s *StructRenderer) getContext() Code {
	return s.ctx
}

func (s *StructRenderer) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *StructRenderer) render(w *Writer) {
	w.Write("struct {")
	w.Br()
	w.AddIndent()
	s.fields.render(w)
	w.RemoveIndent()
	w.Br()
	w.Write("}")
}

type FieldDefsRenderer struct {
	items []Code
	ctx   Code
}

func FieldDefs(items ...Code) *FieldDefsRenderer {
	i := &FieldDefsRenderer{}
	i.Add(items...)
	return i
}

func (l *FieldDefsRenderer) Len() int {
	return len(l.items)
}

func (l *FieldDefsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *FieldDefsRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *FieldDefsRenderer) getContext() Code {
	return l.ctx
}

func (l *FieldDefsRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *FieldDefsRenderer) render(w *Writer) {
	for i, im := range l.items {
		if i != 0 {
			w.Br()
		}
		im.render(w)
	}
}

type FieldDefRenderer struct {
	name  string
	ftype Code
	ctx   Code
}

func FieldDef(name string, ftype Code) *FieldDefRenderer {
	f := &FieldDefRenderer{}
	f.SetName(name)
	f.SetType(ftype)
	return f
}

func (f *FieldDefRenderer) GetName() string {
	return f.name
}

func (f *FieldDefRenderer) SetName(name string) {
	f.name = name
}

func (f *FieldDefRenderer) GetType() Code {
	return f.ftype
}

func (f *FieldDefRenderer) SetType(ftype Code) {
	f.ftype = ftype
	if ftype != nil {
		ftype.setContext(f)
	}
}

func (f *FieldDefRenderer) getContext() Code {
	return f.ctx
}

func (f *FieldDefRenderer) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *FieldDefRenderer) render(w *Writer) {
	w.Write(f.name)
	w.Write(" ")
	indent := f.indent()
	for i := 0; i < indent; i++ {
		w.Write(" ")
	}
	f.ftype.render(w)
}

func (f *FieldDefRenderer) indent() int {
	fields, ok := f.ctx.(*FieldDefsRenderer)
	if !ok {
		return 0
	}
	selfIndex := 0
	for selfIndex < fields.Len() {
		field := fields.At(selfIndex)
		self, ok := field.(*FieldDefRenderer)
		if !ok {
			continue
		}
		if f == self {
			break
		}
	}
	maxGap := 0
	for i := selfIndex - 1; i >= 0; i-- {
		field := fields.At(selfIndex)
		previousField, ok := field.(*FieldDefRenderer)
		if !ok {
			break
		}
		gap := len(previousField.GetName()) - len(f.GetName())
		if gap > maxGap {
			maxGap = gap
		}
	}
	for i := selfIndex + 1; i < fields.Len(); i++ {
		field := fields.At(selfIndex)
		nextField, ok := field.(*FieldDefRenderer)
		if !ok {
			break
		}
		gap := len(nextField.GetName()) - len(f.GetName())
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}
