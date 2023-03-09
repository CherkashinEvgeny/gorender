package renderer

func Render(code Code) string {
	w := &Writer{}
	code.render(w)
	return w.string()
}
