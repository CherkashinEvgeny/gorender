package renderer

type Package struct {
	comment string
	name    string
	imports Code
	code    Code
}

func NewPackage(comment string, name string, imports Code, code Code) *Package {
	p := &Package{}
	p.SetComment(comment)
	p.SetName(name)
	p.SetImports(imports)
	p.SetCode(code)
	return p
}

func (p *Package) GetComment() string {
	return p.comment
}

func (p *Package) SetComment(comment string) {
	p.comment = comment
}

func (p *Package) GetName() string {
	return p.name
}

func (p *Package) SetName(name string) {
	p.name = name
}

func (p *Package) GetImports() Code {
	return p.imports
}

func (p *Package) SetImports(imports Code) {
	p.imports = imports
	if imports != nil {
		imports.setContext(p)
	}
}

func (p *Package) GetCode() Code {
	return p.code
}

func (p *Package) SetCode(code Code) {
	p.code = code
	if code != nil {
		code.setContext(p)
	}
}

func (p *Package) getContext() Code {
	return nil
}

func (p *Package) setContext(_ Code) {
}

func (p *Package) render(w *Writer) {
	if p.comment != "" {
		Comment(p.comment).render(w)
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
