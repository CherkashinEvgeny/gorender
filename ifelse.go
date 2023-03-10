package renderer

type IfElseRenderer struct {
	cond       Code
	ifBranch   Code
	elseBranch Code
	ctx        Code
}

func If(cond Code, code Code) *IfElseRenderer {
	i := &IfElseRenderer{}
	i.SetCond(cond)
	i.SetIfBranch(code)
	return i
}

func IfElse(cond Code, ifBranch Code, elseBranch Code) *IfElseRenderer {
	i := &IfElseRenderer{}
	i.SetCond(cond)
	i.SetIfBranch(ifBranch)
	i.SetElseBranch(elseBranch)
	return i
}

func (i *IfElseRenderer) GetCond() Code {
	return i.cond
}

func (i *IfElseRenderer) SetCond(cond Code) {
	i.cond = cond
	if cond != nil {
		cond.setContext(i)
	}
}

func (i *IfElseRenderer) GetIfBranch() Code {
	return i.ifBranch
}

func (i *IfElseRenderer) SetIfBranch(ifBranch Code) {
	i.ifBranch = ifBranch
	if ifBranch != nil {
		ifBranch.setContext(i)
	}
}

func (i *IfElseRenderer) GetElseBranch() Code {
	return i.elseBranch
}

func (i *IfElseRenderer) SetElseBranch(elseBranch Code) {
	i.elseBranch = elseBranch
	if elseBranch != nil {
		elseBranch.setContext(i)
	}
}

func (i *IfElseRenderer) getContext() Code {
	return i.ctx
}

func (i *IfElseRenderer) setContext(ctx Code) {
	i.ctx = ctx
}

func (i *IfElseRenderer) render(w *Writer) {
	w.Write("if ")
	i.cond.render(w)
	w.Write(" {")
	if i.ifBranch != nil {
		w.Br()
		w.AddIndent()
		i.ifBranch.render(w)
		w.RemoveIndent()
		w.Br()
	}
	w.Write("}")
	if i.elseBranch != nil {
		w.Write(" else ")
		w.Write("{")
		w.Br()
		w.AddIndent()
		i.ifBranch.render(w)
		w.RemoveIndent()
		w.Br()
		w.Write("}")
	}
}
