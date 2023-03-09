package renderer

import "strings"

type ImportList struct {
	items []Code
	ctx   Code
}

func NewImportList(items ...Code) *ImportList {
	l := &ImportList{}
	l.Add(items...)
	return l
}

func (l *ImportList) Len() int {
	return len(l.items)
}

func (l *ImportList) Get(i int) Code {
	return l.items[i]
}

func (l *ImportList) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *ImportList) getContext() Code {
	return l.ctx
}

func (l *ImportList) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *ImportList) render(w *Writer) {
	w.Write("import (")
	w.Br()
	w.AddIndent()
	for i, im := range l.items {
		if i != 0 {
			w.Br()
		}
		im.render(w)
	}
	w.RemoveIndent()
	w.Br()
	w.Write(")")
}

type Import struct {
	alias string
	name  string
	path  string
}

func NewImport(alias string, path string) *Import {
	i := &Import{}
	i.SetAlias(alias)
	i.SetPath(path)
	return i
}

func (i *Import) GetAlias() string {
	return i.alias
}

func (i *Import) SetAlias(alias string) {
	i.alias = alias
}

func (i *Import) GetName() string {
	if i.name == "" {
		chunks := strings.Split(i.path, "/")
		return chunks[len(chunks)-1]
	}
	return i.name
}

func (i *Import) SetName(alias string) {
	i.alias = alias
}

func (i *Import) GetPath() string {
	return i.path
}

func (i *Import) SetPath(path string) {
	i.path = path
}

func (i *Import) getContext() Code {
	return nil
}

func (i *Import) setContext(_ Code) {
}

func (i *Import) render(w *Writer) {
	if i.alias != "" {
		w.Write(i.alias)
		w.Write(" ")
	}
	w.Write("\"")
	w.Write(i.path)
	w.Write("\"")
}
