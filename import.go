package renderer

import (
	"strconv"
	"strings"
)

type ImportsRenderer struct {
	items   []Code
	pathMap map[string]struct{}
	nameMap map[string]struct{}
	ctx     Code
}

func Imports(items ...Code) *ImportsRenderer {
	i := &ImportsRenderer{}
	i.init()
	i.Add(items...)
	return i
}

func (l *ImportsRenderer) init() {
	if l.pathMap == nil {
		l.pathMap = map[string]struct{}{}
	}
	if l.nameMap == nil {
		l.nameMap = map[string]struct{}{}
	}
}

func (l *ImportsRenderer) Len() int {
	return len(l.items)
}

func (l *ImportsRenderer) At(i int) Code {
	return l.items[i]
}

func (l *ImportsRenderer) Add(items ...Code) {
	l.init()
	for _, item := range items {
		l.add(item)
	}
}

func (l *ImportsRenderer) add(item Code) {
	im, ok := item.(*ImportRenderer)
	if !ok {
		l.items = append(l.items, item)
		item.setContext(l)
		return
	}
	_, found := l.pathMap[im.GetPath()]
	if found {
		return
	}
	l.pathMap[im.GetPath()] = struct{}{}
	if im.GetAlias() == "." || im.GetAlias() == "_" {
		l.items = append(l.items, item)
		item.setContext(l)
		return
	}
	var name string
	if im.GetName() != "" {
		name = im.GetName()
	}
	if im.GetAlias() != "" {
		name = im.GetAlias()
	}
	if name == "" {
		l.items = append(l.items, item)
		item.setContext(l)
		return
	}
	alias := name
	counter := 1
	_, found = l.nameMap[alias]
	for found {
		counter++
		alias = name + strconv.Itoa(counter)
	}
	l.nameMap[alias] = struct{}{}
	if name != alias {
		im.SetAlias(alias)
	}
	l.items = append(l.items, item)
	item.setContext(l)
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
