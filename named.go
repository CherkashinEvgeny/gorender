package renderer

type NamedRenderer struct {
	path string
	name string
	ctx  Code
}

func Named(path string, name string) *NamedRenderer {
	n := &NamedRenderer{}
	n.SetPath(path)
	n.SetName(name)
	return n
}

func (n *NamedRenderer) SetPath(path string) {
	n.path = path
}

func (n *NamedRenderer) SetName(name string) {
	n.name = name
}

func (n *NamedRenderer) getContext() Code {
	return n.ctx
}

func (n *NamedRenderer) setContext(ctx Code) {
	n.ctx = ctx
}

func (n *NamedRenderer) render(w *Writer) {
	w.Write(n.path)
	w.Write(".")
	w.Write(n.name)
}
