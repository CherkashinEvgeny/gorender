package renderer

type PkgRenderer struct {
	comment string
	name    string
	imports Code
	code    Code
}

func Pkg(comment string, name string, imports Code, code Code) *PkgRenderer {
	p := &PkgRenderer{}
	p.SetComment(comment)
	p.SetName(name)
	p.SetImports(imports)
	p.SetCode(code)
	return p
}

func (p *PkgRenderer) GetComment() string {
	return p.comment
}

func (p *PkgRenderer) SetComment(comment string) {
	p.comment = comment
}

func (p *PkgRenderer) GetName() string {
	return p.name
}

func (p *PkgRenderer) SetName(name string) {
	p.name = name
}

func (p *PkgRenderer) GetImports() Code {
	return p.imports
}

func (p *PkgRenderer) SetImports(imports Code) {
	p.imports = imports
	if imports != nil {
		imports.setContext(p)
	}
}

func (p *PkgRenderer) GetCode() Code {
	return p.code
}

func (p *PkgRenderer) SetCode(code Code) {
	p.code = code
	if code != nil {
		code.setContext(p)
	}
}

func (p *PkgRenderer) getContext() Code {
	return nil
}

func (p *PkgRenderer) setContext(_ Code) {
}

func (p *PkgRenderer) render(w *Writer) {
	if p.hasComment() {
		p.renderComment(w)
		w.Br()
	}
	w.Write("package ")
	w.Write(p.name)
	if p.hasImports() {
		w.Br()
		w.Br()
		p.renderImports(w)
	}
	if p.hasCode() {
		w.Br()
		w.Br()
		p.renderCode(w)
	}
}

func (p *PkgRenderer) hasComment() bool {
	return p.comment != ""
}

func (p *PkgRenderer) renderComment(w *Writer) {
	CommentRenderer(p.comment).render(w)
}

func (p *PkgRenderer) hasImports() bool {
	if p.imports == nil {
		return false
	}
	imports, ok := p.imports.(*ImportsRenderer)
	if !ok {
		return true
	}
	return imports.Len() != 0
}

func (p *PkgRenderer) renderImports(w *Writer) {
	p.imports.render(w)
}

func (p *PkgRenderer) hasCode() bool {
	if p.code == nil {
		return false
	}
	seq, ok := p.code.(Sequence)
	if !ok {
		return true
	}
	return seq.Len() != 0
}

func (p *PkgRenderer) renderCode(w *Writer) {
	p.code.render(w)
}
