package gen

import (
	renderer "github.com/CherkashinEvgeny/gorender"
	"go/types"
)

type NamedInterface struct {
	Name  string
	Value *types.Interface
}

func FindNamedInterfaces(pkg *types.Package) (items []NamedInterface) {
	pkgScope := pkg.Scope()
	names := pkgScope.Names()
	for _, name := range names {
		obj := pkgScope.Lookup(name)
		t := obj.Type()
		named, ok := t.(*types.Named)
		if !ok {
			continue
		}
		iface, ok := named.Underlying().(*types.Interface)
		if !ok {
			continue
		}
		items = append(items, NamedInterface{Name: name, Value: iface})
	}
	return items
}

func FilterNamedInterfaces(items []NamedInterface, names []string) (filteredItems []NamedInterface) {
	namesMap := make(map[string]struct{}, len(names))
	for _, name := range names {
		namesMap[name] = struct{}{}
	}
	filteredItems = make([]NamedInterface, 0, len(names))
	for _, item := range items {
		_, found := namesMap[item.Name]
		if found {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

func ImplementInterface(
	receiver renderer.Code,
	iface *types.Interface,
	f func(name string, signature *types.Signature) renderer.Code,
) renderer.Code {
	n := iface.NumMethods()
	methods := make([]renderer.Code, 0, n)
	for i := 0; i < n; i++ {
		method := iface.Method(i)
		methodType := method.Type()
		signature, ok := methodType.(*types.Signature)
		if !ok {
			continue
		}
		methods = append(methods, Method(
			receiver,
			method.Name(),
			signature,
			f(method.Name(), signature),
		))
	}
	return renderer.Blocks(methods...)
}

func Method(
	receiver renderer.Code,
	method string,
	signature *types.Signature,
	body renderer.Code,
) renderer.Code {
	return renderer.Method(
		receiver,
		method,
		FuncSign(signature),
		body,
	)
}
