package gen

import (
	"fmt"
	renderer "github.com/CherkashinEvgeny/gorender"
	"go/importer"
	"go/types"
	"testing"
)

func TestInterface(t *testing.T) {
	pkg, err := importer.For("source", nil).Import("github.com/CherkashinEvgeny/gorender/types/test")
	if err != nil {
		panic(err)
	}
	imports := renderer.Imports()
	code := renderer.Blocks()
	interfaces := FindNamedInterfaces(pkg)
	for _, i := range interfaces {
		name := i.Name + "Impl"
		imports.Add(InterfaceImports(i.Value)...)
		code.Add(
			renderer.Type(name, renderer.Struct(nil)),
			ImplementInterface(
				renderer.Param("t", renderer.Id(name), false),
				i.Value,
				func(name string, signature *types.Signature) renderer.Code {
					return renderer.Raw(`fmt.Println("Test")`)
				},
			),
		)
	}
	fmt.Println(renderer.Render(renderer.Pkg("", "test", imports, code)))
}
