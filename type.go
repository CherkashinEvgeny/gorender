package renderer

import (
	"strings"
)

type Type struct {
	name    string
	ttype   Code
	methods Code
	ctx     Code
}

func NewType(name string, ttype Code, methods ...Code) *Type {
	t := &Type{}
	t.SetName(name)
	t.SetType(ttype)
	if len(methods) != 0 {
		t.SetMethods(NewTypeMethodList(methods...))
	}
	return t
}

func (t *Type) GetName() string {
	return t.name
}

func (t *Type) SetName(name string) {
	t.name = name
}

func (t *Type) GetType() Code {
	return t.ttype
}

func (t *Type) SetType(ttype Code) {
	t.ttype = ttype
}

func (t *Type) GetMethods() Code {
	return t.methods
}

func (t *Type) SetMethods(methods Code) {
	t.methods = methods
	if methods != nil {
		methods.setContext(t)
	}
}

func (t *Type) getContext() Code {
	return t.ctx
}

func (t *Type) setContext(ctx Code) {
	t.ctx = ctx
}

func (t *Type) render(w *Writer) {
	w.Write("type")
	w.Write(" ")
	w.Write(t.name)
	w.Write(" ")
	t.ttype.render(w)
	if t.methods != nil {
		w.Br()
		w.Br()
		t.methods.render(w)
	}
}

type TypeMethodList struct {
	items []Code
	ctx   Code
}

func NewTypeMethodList(items ...Code) *TypeMethodList {
	l := &TypeMethodList{}
	l.Add(items...)
	return l
}

func (l *TypeMethodList) Len() int {
	return len(l.items)
}

func (l *TypeMethodList) Get(i int) Code {
	return l.items[i]
}

func (l *TypeMethodList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *TypeMethodList) getContext() Code {
	return l.ctx
}

func (l *TypeMethodList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *TypeMethodList) render(w *Writer) {
	for i, item := range l.items {
		if i != 0 {
			w.Br()
			w.Br()
		}
		item.render(w)
	}
}

type TypeMethod struct {
	receiver  Code
	name      string
	signature Code
	ctx       Code
}

func NewTypeMethod(name string, signature Code) *TypeMethod {
	m := &TypeMethod{}
	m.SetName(name)
	m.SetSignature(signature)
	return m
}

func (m *TypeMethod) GetReceiver() Code {
	return m.receiver
}

func (m *TypeMethod) SetReceiver(receiver Code) {
	m.receiver = receiver
	if receiver != nil {
		receiver.setContext(m)
	}
}

func (m *TypeMethod) GetName() string {
	return m.name
}

func (m *TypeMethod) SetName(name string) {
	m.name = name
}

func (m *TypeMethod) GetSignature() Code {
	return m.signature
}

func (m *TypeMethod) SetSignature(signature Code) {
	m.signature = signature
	if signature != nil {
		signature.setContext(m)
	}
}

func (m *TypeMethod) getContext() Code {
	return m.ctx
}

func (m *TypeMethod) setContext(ctx Code) {
	m.ctx = ctx
	if m.receiver == nil {
		m.tryToSetDefaultReceiver()
	}
}

func (m *TypeMethod) tryToSetDefaultReceiver() {
	methodList, ok := m.ctx.(*TypeMethodList)
	if !ok {
		return
	}
	ttype, ok := methodList.getContext().(*Type)
	if !ok {
		return
	}
	typeName := ttype.GetName()
	if typeName == "" {
		return
	}
	receiverName := strings.ToLower(string([]rune(typeName)[1]))
	receiver := NewFunctionParam(receiverName, NewNamed("", typeName), false)
	m.SetReceiver(receiver)
}

func (m *TypeMethod) render(w *Writer) {
	w.Write("func ")
	m.receiver.render(w)
	w.Write(" ")
	w.Write(m.name)
	m.signature.render(w)
}
