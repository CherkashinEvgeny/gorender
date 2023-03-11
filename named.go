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
	prefix := n.resolvePrefix()
	if prefix != "" {
		w.Write(prefix)
		w.Write(".")
	}
	w.Write(n.name)
}

func (n *NamedRenderer) resolvePrefix() string {
	imporrtsCode := n.getContext()
	for {
		if imporrtsCode == nil {
			break
		}
		_, ok := imporrtsCode.(*PkgRenderer)
		if ok {
			break
		}
		imporrtsCode = imporrtsCode.getContext()
	}
	if imporrtsCode == nil {
		return n.path
	}
	pkg, ok := imporrtsCode.(*PkgRenderer)
	if !ok {
		return n.path
	}
	imporrtsCode = pkg.GetImports()
	if imporrtsCode == nil {
		return n.path
	}
	imports, ok := pkg.GetImports().(*ImportsRenderer)
	if !ok {
		return n.path
	}
	var im *ImportRenderer
	c := imports.Len()
	for i := 0; i < c; i++ {
		importCode := imports.At(i)
		im, ok = importCode.(*ImportRenderer)
		if ok {
			break
		}
	}
	alias := im.GetAlias()
	if alias == "." || alias == "_" {
		return ""
	}
	if alias != "" {
		return alias
	}
	return im.GetName()
}
