package renderer

type StructRenderer struct {
	body Code
	ctx  Code
}

func Struct(fields ...Code) *StructRenderer {
	s := &StructRenderer{}
	if len(fields) != 0 {
		s.SetFields(Fields(fields...))
	}
	return s
}

func (s *StructRenderer) GetFields() Code {
	return s.body
}

func (s *StructRenderer) SetFields(fields Code) {
	s.body = fields
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
	s.body.render(w)
	w.RemoveIndent()
	w.Br()
	w.Write("}")
}

type FieldsRenderer struct {
	items []Code
	ctx   Code
}

func Fields(items ...Code) *FieldsRenderer {
	l := &FieldsRenderer{}
	l.Add(items...)
	return l
}

func (l *FieldsRenderer) Len() int {
	return len(l.items)
}

func (l *FieldsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *FieldsRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *FieldsRenderer) getContext() Code {
	return l.ctx
}

func (l *FieldsRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *FieldsRenderer) render(w *Writer) {
	for i, im := range l.items {
		if i != 0 {
			w.Br()
		}
		im.render(w)
	}
}

type FieldRenderer struct {
	name  string
	ftype Code
	ctx   Code
}

func Field(name string, ftype Code) *FieldRenderer {
	f := &FieldRenderer{}
	f.SetName(name)
	f.SetType(ftype)
	return f
}

func (f *FieldRenderer) GetName() string {
	return f.name
}

func (f *FieldRenderer) SetName(name string) {
	f.name = name
}

func (f *FieldRenderer) GetType() Code {
	return f.ftype
}

func (f *FieldRenderer) SetType(ftype Code) {
	f.ftype = ftype
	if ftype != nil {
		ftype.setContext(f)
	}
}

func (f *FieldRenderer) getContext() Code {
	return f.ctx
}

func (f *FieldRenderer) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *FieldRenderer) render(w *Writer) {
	w.Write(f.name)
	w.Write(" ")
	indent := f.indent()
	for i := 0; i < indent; i++ {
		w.Write(" ")
	}
	f.ftype.render(w)
}

func (f *FieldRenderer) indent() int {
	fields, ok := f.ctx.(*FieldsRenderer)
	if !ok {
		return 0
	}
	selfIndex := 0
	for selfIndex < fields.Len() {
		field := fields.At(selfIndex)
		self, ok := field.(*FieldRenderer)
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
		previousField, ok := field.(*FieldRenderer)
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
		nextField, ok := field.(*FieldRenderer)
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
