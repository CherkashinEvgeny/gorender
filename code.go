package renderer

type Code interface {
	getContext() Code
	setContext(ctx Code)
	render(w *Writer)
}
