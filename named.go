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
	im := n.findImport()
	if im == nil {
		return ""
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

func (n *NamedRenderer) findImport() *ImportRenderer {
	imports := n.findImports()
	if imports == nil {
		return nil
	}
	c := imports.Len()
	for i := 0; i < c; i++ {
		importCandidate := imports.At(i)
		im, ok := importCandidate.(*ImportRenderer)
		if ok && n.path == im.GetPath() {
			return im
		}
	}
	return nil
}

func (n *NamedRenderer) findImports() *ImportsRenderer {
	pkg := n.findPackage()
	if pkg == nil {
		return nil
	}
	importsCandidate := pkg.GetImports()
	imports, ok := importsCandidate.(*ImportsRenderer)
	if !ok {
		return nil
	}
	return imports
}

func (n *NamedRenderer) findPackage() *PkgRenderer {
	ctx := n.getContext()
	for ctx != nil {
		pkg, ok := ctx.(*PkgRenderer)
		if ok {
			return pkg
		}
		ctx = ctx.getContext()
	}
	return nil
}
