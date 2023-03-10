package renderer

import "strings"

type ImportsRenderer struct {
	items []Code
	ctx   Code
}

func Imports(items ...Code) *ImportsRenderer {
	l := &ImportsRenderer{}
	l.Add(items...)
	return l
}

func (l *ImportsRenderer) Len() int {
	return len(l.items)
}

func (l *ImportsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *ImportsRenderer) Add(items ...Code) {
	l.items = append(l.items, items...)
	for _, item := range items {
		item.setContext(l)
	}
}

func (l *ImportsRenderer) getContext() Code {
	return l.ctx
}

func (l *ImportsRenderer) setContext(ctx Code) {
	l.ctx = ctx
}

func (l *ImportsRenderer) render(w *Writer) {
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

type ImportRenderer struct {
	alias string
	name  string
	path  string
}

func Import(name string, alias string, path string) *ImportRenderer {
	i := &ImportRenderer{}
	i.SetName(name)
	i.SetAlias(alias)
	i.SetPath(path)
	return i
}

func (i *ImportRenderer) GetName() string {
	if i.name == "" {
		chunks := strings.Split(i.path, "/")
		return chunks[len(chunks)-1]
	}
	return i.name
}

func (i *ImportRenderer) SetName(name string) {
	i.name = name
}

func (i *ImportRenderer) GetAlias() string {
	return i.alias
}

func (i *ImportRenderer) SetAlias(alias string) {
	i.alias = alias
}

func (i *ImportRenderer) GetPath() string {
	return i.path
}

func (i *ImportRenderer) SetPath(path string) {
	i.path = path
}

func (i *ImportRenderer) getContext() Code {
	return nil
}

func (i *ImportRenderer) setContext(_ Code) {
}

func (i *ImportRenderer) render(w *Writer) {
	if i.alias != "" {
		w.Write(i.alias)
		w.Write(" ")
	}
	w.Write("\"")
	w.Write(i.path)
	w.Write("\"")
}
