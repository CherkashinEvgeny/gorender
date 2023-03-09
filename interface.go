package renderer

type Interface struct {
	methods Code
	ctx     Code
}

func NewInterface(methods ...Code) *Interface {
	i := &Interface{}
	i.SetMethods(NewInterfaceMethodList(methods...))
	return i
}

func (i *Interface) GetMethods() Code {
	return i.methods
}

func (i *Interface) SetMethods(methods Code) {
	i.methods = methods
	if methods != nil {
		methods.setContext(i)
	}
}

func (i *Interface) getContext() Code {
	return i.ctx
}

func (i *Interface) setContext(ctx Code) {
	i.ctx = ctx
}

func (i *Interface) render(w *Writer) {
	w.Write("interface {")
	w.Br()
	w.AddIndent()
	i.methods.render(w)
	w.RemoveIndent()
	w.Br()
	w.Write("}")
}

type InterfaceMethodList struct {
	items []Code
	ctx   Code
}

func NewInterfaceMethodList(methods ...Code) *InterfaceMethodList {
	l := &InterfaceMethodList{}
	l.Add(methods...)
	return l
}

func (l *InterfaceMethodList) Len() int {
	return len(l.items)
}

func (l *InterfaceMethodList) Get(i int) Code {
	return l.items[i]
}

func (l *InterfaceMethodList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *InterfaceMethodList) getContext() Code {
	return l.ctx
}

func (l *InterfaceMethodList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *InterfaceMethodList) render(w *Writer) {
	for i, m := range l.items {
		if i != 0 {
			w.Br()
		}
		m.render(w)
	}
}

type InterfaceMethod struct {
	name      string
	signature Code
	ctx       Code
}

func NewInterfaceMethod(name string, signature Code) *InterfaceMethod {
	m := &InterfaceMethod{}
	m.SetName(name)
	m.SetSignature(signature)
	return m
}

func (m *InterfaceMethod) GetName() string {
	return m.name
}

func (m *InterfaceMethod) SetName(name string) {
	m.name = name
}

func (m *InterfaceMethod) GetSignature() Code {
	return m.signature
}

func (m *InterfaceMethod) SetSignature(signature Code) {
	m.signature = signature
	if signature != nil {
		signature.setContext(m)
	}
}

func (m *InterfaceMethod) getContext() Code {
	return m.ctx
}

func (m *InterfaceMethod) setContext(ctx Code) {
	m.ctx = ctx
}

func (m *InterfaceMethod) render(w *Writer) {
	w.Write(m.name)
	m.signature.render(w)
}
