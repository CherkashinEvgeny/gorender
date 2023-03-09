package renderer

import "strconv"

type Array struct {
	size int
	elem Code
	ctx  Code
}

func NewArray(size int, elem Code) *Array {
	a := &Array{}
	a.SetSize(size)
	a.SetElem(elem)
	return a
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) SetSize(size int) {
	a.size = size
}

func (a *Array) GetElem() Code {
	return a.elem
}

func (a *Array) SetElem(elem Code) {
	a.elem = elem
	if elem != nil {
		elem.setContext(a)
	}
}

func (a *Array) getContext() Code {
	return a.ctx
}

func (a *Array) setContext(ctx Code) {
	a.ctx = ctx
}

func (a *Array) render(w *Writer) {
	w.Write("[]")
	w.Write(strconv.Itoa(a.size))
	a.elem.render(w)
}
