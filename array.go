package renderer

import "strconv"

type ArrayRenderer struct {
	size int
	elem Code
	ctx  Code
}

func Array(size int, elem Code) *ArrayRenderer {
	a := &ArrayRenderer{}
	a.SetSize(size)
	a.SetElem(elem)
	return a
}

func (a *ArrayRenderer) GetSize() int {
	return a.size
}

func (a *ArrayRenderer) SetSize(size int) {
	a.size = size
}

func (a *ArrayRenderer) GetElem() Code {
	return a.elem
}

func (a *ArrayRenderer) SetElem(elem Code) {
	a.elem = elem
	if elem != nil {
		elem.setContext(a)
	}
}

func (a *ArrayRenderer) getContext() Code {
	return a.ctx
}

func (a *ArrayRenderer) setContext(ctx Code) {
	a.ctx = ctx
}

func (a *ArrayRenderer) render(w *Writer) {
	w.Write("[]")
	w.Write(strconv.Itoa(a.size))
	a.elem.render(w)
}
