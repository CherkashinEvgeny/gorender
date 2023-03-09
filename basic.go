package renderer

const (
	Bool       = basic("bool")
	Int        = basic("int")
	Int8       = basic("int8")
	Int16      = basic("int16")
	Int32      = basic("int32")
	Int64      = basic("int64")
	Uint       = basic("uint")
	Uint8      = basic("uint8")
	Uint16     = basic("uint16")
	Uint32     = basic("uint32")
	Uint64     = basic("uint64")
	Uintptr    = basic("uintptr")
	Float32    = basic("float32")
	Float64    = basic("float64")
	Complex64  = basic("complex64")
	Complex128 = basic("complex128")
	String     = basic("string")
	Byte       = basic("byte")
	Rune       = basic("rune")
	Any        = basic("any")
	Error      = basic("error")
)

type basic string

func (t basic) getContext() Code {
	return nil
}

func (t basic) setContext(_ Code) {
}

func (t basic) render(w *Writer) {
	w.Write(string(t))
}
