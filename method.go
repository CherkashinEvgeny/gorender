package renderer

type MethodRenderer struct {
	receiver  Code
	name      string
	signature Code
	ctx       Code
}

func Method(receiver Code, name string, signature Code) *MethodRenderer {
	m := &MethodRenderer{}
	m.SetReceiver(receiver)
	m.SetName(name)
	m.SetSignature(signature)
	return m
}

func (m *MethodRenderer) GetReceiver() Code {
	return m.receiver
}

func (m *MethodRenderer) SetReceiver(receiver Code) {
	m.receiver = receiver
	if receiver != nil {
		receiver.setContext(m)
	}
}

func (m *MethodRenderer) GetName() string {
	return m.name
}

func (m *MethodRenderer) SetName(name string) {
	m.name = name
}

func (m *MethodRenderer) GetSignature() Code {
	return m.signature
}

func (m *MethodRenderer) SetSignature(signature Code) {
	m.signature = signature
	if signature != nil {
		signature.setContext(m)
	}
}

func (m *MethodRenderer) getContext() Code {
	return m.ctx
}

func (m *MethodRenderer) setContext(ctx Code) {
	m.ctx = ctx
}

func (m *MethodRenderer) render(w *Writer) {
	w.Write("func ")
	m.receiver.render(w)
	w.Write(" ")
	w.Write(m.name)
	m.signature.render(w)
}
