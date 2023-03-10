package gen

import (
	renderer "github.com/CherkashinEvgeny/gorender"
	"github.com/pkg/errors"
	"go/types"
)

func Type(t types.Type) (code renderer.Code) {
	switch v := t.(type) {
	case *types.Interface:
		return Interface(v)
	case *types.Struct:
		return Struct(v)
	case *types.Signature:
		return FuncType(v)
	case *types.Tuple:
		return Tuple(v)
	case *types.Map:
		return Map(v)
	case *types.Chan:
		return Chan(v)
	case *types.Slice:
		return Slice(v)
	case *types.Array:
		return Array(v)
	case *types.Pointer:
		return Pointer(v)
	case *types.Named:
		return Named(v)
	case *types.Basic:
		return Basic(v)
	default:
		panic(errors.Errorf("unknown type = %v", t))
	}
}

func Interface(t *types.Interface) (code renderer.Code) {
	methods := renderer.Methods()
	n := t.NumMethods()
	for i := 0; i < n; i++ {
		method := t.Method(i)
		methodType := method.Type()
		nd, ok := methodType.(*types.Named)
		if !ok {
			continue
		}
		methods.Add(renderer.Embedded(renderer.Named(
			nd.Obj().Pkg().Path(),
			nd.Obj().Name(),
		)))
	}
	for i := 0; i < n; i++ {
		method := t.Method(i)
		methodType := method.Type()
		s, ok := methodType.(*types.Signature)
		if !ok {
			continue
		}
		methods.Add(renderer.MethodDecl(method.Name(), renderer.Sign(
			FuncIn(s),
			FuncOut(s),
		)))
	}
	return renderer.Iface(methods)
}

func Struct(t *types.Struct) (code renderer.Code) {
	fields := renderer.FieldDefs()
	n := t.NumFields()
	for i := 0; i < n; i++ {
		field := t.Field(i)
		fields.Add(renderer.FieldDef(field.Name(), Type(field.Type())))
	}
	return renderer.Struct(fields)
}

func FuncType(t *types.Signature) (code renderer.Code) {
	return renderer.FuncType(FuncSign(t))
}

func Tuple(t *types.Tuple) (code renderer.Code) {
	n := t.Len()
	params := make([]renderer.Code, 0, n)
	for i := 0; i < n; i++ {
		param := t.At(i)
		params = append(params, Type(param.Type()))
	}
	return renderer.Join(params...)
}

func Map(t *types.Map) (code renderer.Code) {
	return renderer.Map(Type(t.Key()), Type(t.Elem()))
}

func Chan(t *types.Chan) (code renderer.Code) {
	switch t.Dir() {
	case types.SendOnly:
		return renderer.Chan(renderer.Send, Type(t.Elem()))
	case types.RecvOnly:
		return renderer.Chan(renderer.Receive, Type(t.Elem()))
	default:
		return renderer.Chan(renderer.SendAndReceive, Type(t.Elem()))
	}
}

func Slice(t *types.Slice) (code renderer.Code) {
	return renderer.Slice(Type(t.Elem()))
}

func Array(t *types.Array) (code renderer.Code) {
	return renderer.Array(int(t.Len()), Type(t.Elem()))
}

func Pointer(t *types.Pointer) (code renderer.Code) {
	return renderer.Ptr(Type(t.Elem()))
}

func Named(t *types.Named) (code renderer.Code) {
	obj := t.Obj()
	pkg := obj.Pkg()
	return renderer.Named(pkg.Path(), obj.Name())
}

func Basic(t *types.Basic) (code renderer.Code) {
	switch t.String() {
	case "bool":
		return renderer.Bool
	case "int":
		return renderer.Int
	case "int8":
		return renderer.Int8
	case "int16":
		return renderer.Int16
	case "int32":
		return renderer.Int32
	case "int64":
		return renderer.Int64
	case "uint":
		return renderer.Uint
	case "uint8":
		return renderer.Uint8
	case "uint16":
		return renderer.Uint16
	case "uint32":
		return renderer.Uint32
	case "uint64":
		return renderer.Uint64
	case "uintptr":
		return renderer.Uintptr
	case "float32":
		return renderer.Float32
	case "float64":
		return renderer.Float64
	case "complex64":
		return renderer.Complex64
	case "complex128":
		return renderer.Complex128
	case "string":
		return renderer.String
	case "byte":
		return renderer.Byte
	case "rune":
		return renderer.Rune
	default:
		panic("unsupported type")
	}
}
