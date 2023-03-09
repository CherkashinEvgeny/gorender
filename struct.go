package renderer

type Struct struct {
	fields Code
	ctx    Code
}

func NewStruct(fields ...Code) *Struct {
	s := &Struct{}
	if len(fields) != 0 {
		s.SetFields(NewStructFieldList(fields...))
	}
	return s
}

func (s *Struct) GetFields() Code {
	return s.fields
}

func (s *Struct) SetFields(fields Code) {
	s.fields = fields
	if fields != nil {
		fields.setContext(s)
	}
}

func (s *Struct) getContext() Code {
	return s.ctx
}

func (s *Struct) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *Struct) render(w *Writer) {
	w.Write("struct {")
	w.Br()
	w.AddIndent()
	s.fields.render(w)
	w.RemoveIndent()
	w.Br()
	w.Write("}")
}

type StructFieldList struct {
	items []Code
	ctx   Code
}

func NewStructFieldList(items ...Code) *StructFieldList {
	l := &StructFieldList{}
	l.Add(items...)
	return l
}

func (l *StructFieldList) Len() int {
	return len(l.items)
}

func (l *StructFieldList) Get(i int) Code {
	return l.items[i]
}

func (l *StructFieldList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *StructFieldList) getContext() Code {
	return l.ctx
}

func (l *StructFieldList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *StructFieldList) render(w *Writer) {
	for i, im := range l.items {
		if i != 0 {
			w.Br()
		}
		im.render(w)
	}
}

type StructField struct {
	name  string
	ftype Code
	ctx   Code
}

func NewStructField(name string, ftype Code) *StructField {
	f := &StructField{}
	f.SetName(name)
	f.SetType(ftype)
	return f
}

func (f *StructField) GetName() string {
	return f.name
}

func (f *StructField) SetName(name string) {
	f.name = name
}

func (f *StructField) GetType() Code {
	return f.ftype
}

func (f *StructField) SetType(ftype Code) {
	f.ftype = ftype
	if ftype != nil {
		ftype.setContext(f)
	}
}

func (f *StructField) getContext() Code {
	return f.ctx
}

func (f *StructField) setContext(ctx Code) {
	f.ctx = ctx
}

func (f *StructField) render(w *Writer) {
	w.Write(f.name)
	w.Write(" ")
	f.ftype.render(w)
}
