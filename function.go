package renderer

type FunctionType struct {
	signature Code
	ctx       Code
}

func NewFunctionType(in Code, out Code) *FunctionType {
	f := &FunctionType{}
	f.SetSignature(NewFunctionSignature(in, out))
	return f
}

func (f *FunctionType) GetSignature() Code {
	return f.signature
}

func (f *FunctionType) SetSignature(signature Code) {
	f.signature = signature
	if signature != nil {
		signature.setContext(f)
	}
}

func (f *FunctionType) getContext() Code {
	return f.ctx
}

func (f *FunctionType) setContext(ctx Code) {
	f.ctx = ctx
	if f.signature != nil {
		f.signature.setContext(ctx)
	}
}

func (f *FunctionType) render(w *Writer) {
	w.Write("func")
	f.signature.render(w)
}

type FunctionSignature struct {
	in  Code
	out Code
	ctx Code
}

func NewFunctionSignature(in Code, out Code) *FunctionSignature {
	s := &FunctionSignature{}
	s.SetIn(in)
	s.SetOut(out)
	return s
}

func (s *FunctionSignature) GetIn() Code {
	return s.in
}

func (s *FunctionSignature) SetIn(in Code) {
	s.in = in
	if in != nil {
		in.setContext(s)
	}
}

func (s *FunctionSignature) GetOut() Code {
	return s.out
}

func (s *FunctionSignature) SetOut(out Code) {
	s.out = out
	if out != nil {
		out.setContext(s)
	}
}

func (s *FunctionSignature) getContext() Code {
	return s.ctx
}

func (s *FunctionSignature) setContext(ctx Code) {
	s.ctx = ctx
}

func (s *FunctionSignature) render(w *Writer) {
	w.Write("(")
	s.in.render(w)
	w.Write(")")
	w.Write(" ")
	w.Write("(")
	s.out.render(w)
	w.Write(")")
}

type FunctionParamList struct {
	items []Code
	ctx   Code
}

func NewFunctionParamList(items ...Code) *FunctionParamList {
	l := &FunctionParamList{}
	l.Add(items...)
	return l
}

func (l *FunctionParamList) Len() int {
	return len(l.items)
}

func (l *FunctionParamList) Get(i int) Code {
	return l.items[i]
}

func (l *FunctionParamList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *FunctionParamList) getContext() Code {
	return l.ctx
}

func (l *FunctionParamList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *FunctionParamList) render(w *Writer) {
	for i, param := range l.items {
		if i != 0 {
			w.Write(", ")
		}
		param.render(w)
	}
}

type FunctionParam struct {
	name     string
	ptype    Code
	variadic bool
	ctx      Code
}

func NewFunctionParam(name string, ptype Code, variadic bool) *FunctionParam {
	p := &FunctionParam{}
	p.SetName(name)
	p.SetType(ptype)
	p.SetVariadic(variadic)
	return p
}

func (p *FunctionParam) GetName() string {
	return p.name
}

func (p *FunctionParam) SetName(name string) {
	p.name = name
}

func (p *FunctionParam) GetType() Code {
	return p.ptype
}

func (p *FunctionParam) SetType(ptype Code) {
	p.ptype = ptype
	if ptype != nil {
		ptype.setContext(p)
	}
}

func (p *FunctionParam) GetVariadic() bool {
	return p.variadic
}

func (p *FunctionParam) SetVariadic(variadic bool) {
	p.variadic = variadic
}

func (p *FunctionParam) getContext() Code {
	return p.ctx
}

func (p *FunctionParam) setContext(ctx Code) {
	p.ctx = ctx
}

func (p *FunctionParam) render(w *Writer) {
	if p.name != "" {
		w.Write(p.name)
		w.Write(" ")
	}
	if p.variadic {
		w.Write("...")
	}
	p.ptype.render(w)
}
