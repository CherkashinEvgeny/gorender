package renderer

type Named struct {
	path string
	name string
	ctx  Code
}

func NewNamed(path string, name string) *Named {
	n := &Named{}
	n.SetPath(path)
	n.SetName(name)
	return n
}

func (n *Named) SetPath(path string) {
	n.path = path
}

func (n *Named) SetName(name string) {
	n.name = name
}

func (n *Named) getContext() Code {
	return n.ctx
}

func (n *Named) setContext(ctx Code) {
	n.ctx = ctx
}

func (n *Named) render(w *Writer) {
	w.Write(n.path)
	w.Write(".")
	w.Write(n.name)
}
