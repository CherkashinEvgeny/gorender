package gen

import (
	renderer "github.com/CherkashinEvgeny/gorender"
	"go/types"
)

func FuncSign(t *types.Signature) renderer.Code {
	return renderer.Sign(FuncIn(t), FuncOut(t))
}

func FuncIn(t *types.Signature) renderer.Code {
	in := renderer.In()
	params := t.Params()
	n := params.Len()
	if t.Variadic() {
		n--
	}
	for i := 0; i < n; i++ {
		param := params.At(i)
		in.Add(renderer.Param(param.Name(), Type(param.Type()), false))
	}
	if t.Variadic() {
		param := params.At(n)
		in.Add(renderer.Param(param.Name(), Type(param.Type()), true))
	}
	return in
}

func FuncOut(t *types.Signature) renderer.Code {
	out := renderer.Out()
	params := t.Results()
	n := params.Len()
	for i := 0; i < n; i++ {
		variable := params.At(i)
		out.Add(renderer.Param(variable.Name(), Type(variable.Type()), false))
	}
	return out
}
