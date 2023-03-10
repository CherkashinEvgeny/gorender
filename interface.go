package renderer

type InterfaceRenderer struct {
	fields Code
	ctx    Code
}

func Interface(methods ...Code) *InterfaceRenderer {
	i := &InterfaceRenderer{}
	i.SetFields(Methods(methods...))
	return i
}

func (i *InterfaceRenderer) GetFields() Code {
	return i.fields
}

func (i *InterfaceRenderer) SetFields(fields Code) {
	i.fields = fields
	if fields != nil {
		fields.setContext(i)
	}
}

func (i *InterfaceRenderer) getContext() Code {
	return i.ctx
}

func (i *InterfaceRenderer) setContext(ctx Code) {
	i.ctx = ctx
}

func (i *InterfaceRenderer) render(w *Writer) {
	w.Write("interface {")
	w.Br()
	w.AddIndent()
	i.fields.render(w)
	w.RemoveIndent()
	w.Br()
	w.Write("}")
}

type MethodDeclsRenderer struct {
	items []Code
	ctx   Code
}

func Methods(methods ...Code) *MethodDeclsRenderer {
	l := &MethodDeclsRenderer{}
	l.Add(methods...)
	return l
}

func (l *MethodDeclsRenderer) Len() int {
	return len(l.items)
}

func (l *MethodDeclsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *MethodDeclsRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *MethodDeclsRenderer) getContext() Code {
	return l.ctx
}

func (l *MethodDeclsRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *MethodDeclsRenderer) render(w *Writer) {
	for i, m := range l.items {
		if i != 0 {
			w.Br()
		}
		m.render(w)
	}
}

type MethodDeclRenderer struct {
	name      string
	signature Code
	ctx       Code
}

func MethodDecl(name string, signature Code) *MethodDeclRenderer {
	m := &MethodDeclRenderer{}
	m.SetName(name)
	m.SetSignature(signature)
	return m
}

func (m *MethodDeclRenderer) GetName() string {
	return m.name
}

func (m *MethodDeclRenderer) SetName(name string) {
	m.name = name
}

func (m *MethodDeclRenderer) GetSignature() Code {
	return m.signature
}

func (m *MethodDeclRenderer) SetSignature(signature Code) {
	m.signature = signature
	if signature != nil {
		signature.setContext(m)
	}
}

func (m *MethodDeclRenderer) getContext() Code {
	return m.ctx
}

func (m *MethodDeclRenderer) setContext(ctx Code) {
	m.ctx = ctx
}

func (m *MethodDeclRenderer) render(w *Writer) {
	w.Write(m.name)
	m.signature.render(w)
}
