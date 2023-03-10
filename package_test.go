package renderer

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	pkg := &PkgRenderer{}
	pkg.SetComment("heheh")
	pkg.SetName("test")
	imports := &ImportsRenderer{}
	imports.Add(Import("test", "test", "renderer"))
	pkg.SetImports(imports)
	code := &BlocksRenderer{}
	code.Add(
		Type(
			"Face",
			Iface(
				MethodDecl("String", Sign(In(), Out())),
				MethodDecl("Error", Sign(In(), Out())),
			),
		),
	)
	code.Add(Func("main", Sign(In(), Out()), Lines(
		For(Raw("a > 10"), Raw("fmt.Println(\"hehe\")")),
	)))
	pkg.SetCode(code)
	str := Render(pkg)
	fmt.Println(str)
}
