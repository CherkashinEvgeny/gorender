package renderer

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	pkg := &Package{}
	pkg.SetComment("heheh")
	pkg.SetName("test")
	imports := &ImportList{}
	imports.Add(NewImport("test", "renderer"))
	pkg.SetImports(imports)
	code := &FrameList{}
	code.Add(
		NewType(
			"Face",
			NewInterface(
				Raw("String() string"),
				Raw("Error() string"),
			),
		),
	)
	pkg.SetCode(code)
	str := Render(pkg)
	fmt.Println(str)
}
