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
	if p.comment != "" {
		CommentRenderer(p.comment).render(w)
		w.Br()
	}
	w.Write("package ")
	w.Write(p.name)
	if p.imports != nil {
		w.Br()
		w.Br()
		p.imports.render(w)
	}
	if p.code != nil {
		w.Br()
		w.Br()
		p.code.render(w)
	}
}
