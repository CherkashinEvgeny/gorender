package renderer

type MapRenderer struct {
	key   Code
	value Code
	ctx   Code
}

func Map(key Code, value Code) *MapRenderer {
	m := &MapRenderer{}
	m.SetKey(key)
	m.SetValue(value)
	return m
}

func (m *MapRenderer) GetKey() Code {
	return m.key
}

func (m *MapRenderer) SetKey(key Code) {
	m.key = key
	if key != nil {
		key.setContext(m)
	}
}

func (m *MapRenderer) GetValue() Code {
	return m.value
}

func (m *MapRenderer) SetValue(value Code) {
	m.value = value
	if value != nil {
		value.setContext(m)
	}
}

func (m *MapRenderer) getContext() Code {
	return m.ctx
}

func (m *MapRenderer) setContext(ctx Code) {
	m.ctx = ctx
}

func (m *MapRenderer) render(w *Writer) {
	w.Write("map[")
	m.key.render(w)
	w.Write("]")
	m.value.render(w)
}
