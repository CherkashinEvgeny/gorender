package renderer

type ParamRenderer struct {
	name     string
	ptype    Code
	variadic bool
	ctx      Code
}

func Param(name string, ptype Code, variadic bool) *ParamRenderer {
	p := &ParamRenderer{}
	p.SetName(name)
	p.SetType(ptype)
	p.SetVariadic(variadic)
	return p
}

func (p *ParamRenderer) GetName() string {
	return p.name
}

func (p *ParamRenderer) SetName(name string) {
	p.name = name
}

func (p *ParamRenderer) GetType() Code {
	return p.ptype
}

func (p *ParamRenderer) SetType(ptype Code) {
	p.ptype = ptype
	if ptype != nil {
		ptype.setContext(p)
	}
}

func (p *ParamRenderer) GetVariadic() bool {
	return p.variadic
}

func (p *ParamRenderer) SetVariadic(variadic bool) {
	p.variadic = variadic
}

func (p *ParamRenderer) getContext() Code {
	return p.ctx
}

func (p *ParamRenderer) setContext(ctx Code) {
	p.ctx = ctx
}

func (p *ParamRenderer) render(w *Writer) {
	if p.name != "" {
		w.Write(p.name)
		w.Write(" ")
	}
	if p.variadic {
		w.Write("...")
	}
	p.ptype.render(w)
}
