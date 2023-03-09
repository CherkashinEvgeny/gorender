package renderer

type Map struct {
	key   Code
	value Code
	ctx   Code
}

func NewMap(key Code, value Code) *Map {
	m := &Map{}
	m.SetKey(key)
	m.SetValue(value)
	return m
}

func (m *Map) GetKey() Code {
	return m.key
}

func (m *Map) SetKey(key Code) {
	m.key = key
	if key != nil {
		key.setContext(m)
	}
}

func (m *Map) GetValue() Code {
	return m.value
}

func (m *Map) SetValue(value Code) {
	m.value = value
	if value != nil {
		value.setContext(m)
	}
}

func (m *Map) getContext() Code {
	return m.ctx
}

func (m *Map) setContext(ctx Code) {
	m.ctx = ctx
}

func (m *Map) render(w *Writer) {
	w.Write("map[")
	m.key.render(w)
	w.Write("]")
	m.value.render(w)
}
