package renderer

import (
	"strconv"
)

func Vals(items ...Code) *ValsRenderer {
	i := &ValsRenderer{}
	i.Add(items...)
	return i
}

func (v *ValsRenderer) Len() int {
	return len(v.items)
}

func (v *ValsRenderer) At(i int) Code {
	return v.items[i]
}

func (v *ValsRenderer) Add(items ...Code) {
	v.items = append(v.items, items...)
	for _, item := range items {
		item.setContext(v)
	}
}

func (v *ValsRenderer) getContext() Code {
	return v.ctx
}

func (v *ValsRenderer) setContext(ctx Code) {
	v.ctx = ctx
}

func (v *ValsRenderer) render(w *Writer) {
	for i, item := range v.items {
		if i != 0 {
			w.Br()
		}
		item.render(w)
		w.Write(",")
	}
}

type ValRenderer struct {
	value any
}

func val(value any) ValRenderer {
	return ValRenderer{value}
}

func (l ValRenderer) getContext() Code {
	return nil
}

func (l ValRenderer) setContext(ctx Code) {
}

func (l ValRenderer) render(w *Writer) {
	switch v := l.value.(type) {
	case bool:
		renderBoolLiteral(v, w)
	case int:
		renderIntLiteral(v, w)
	case int8:
		renderInt8Literal(v, w)
	case int16:
		renderInt16Literal(v, w)
	case int32:
		renderInt32Literal(v, w)
	case int64:
		renderInt64Literal(v, w)
	case uint:
		renderUintLiteral(v, w)
	case uint8:
		renderUint8Literal(v, w)
	case uint16:
		renderUint16Literal(v, w)
	case uint32:
		renderUint32Literal(v, w)
	case uint64:
		renderUint64Literal(v, w)
	case float32:
		renderFloat32Literal(v, w)
	case float64:
		renderFloat64Literal(v, w)
	case string:
		renderStringLiteral(v, w)
	}
}

func renderBoolLiteral(v bool, w *Writer) {
	w.Write(strconv.FormatBool(v))
}

func renderIntLiteral(v int, w *Writer) {
	w.Write(strconv.Itoa(v))
}

func renderInt8Literal(v int8, w *Writer) {
	w.Write(strconv.FormatInt(int64(v), 10))
}

func renderInt16Literal(v int16, w *Writer) {
	w.Write(strconv.FormatInt(int64(v), 10))
}

func renderInt32Literal(v int32, w *Writer) {
	w.Write(strconv.FormatInt(int64(v), 10))
}

func renderInt64Literal(v int64, w *Writer) {
	w.Write(strconv.FormatInt(v, 10))
}

func renderUintLiteral(v uint, w *Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint8Literal(v uint8, w *Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint16Literal(v uint16, w *Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint32Literal(v uint32, w *Writer) {
	w.Write(strconv.FormatUint(uint64(v), 10))
}

func renderUint64Literal(v uint64, w *Writer) {
	w.Write(strconv.FormatUint(v, 10))
}

func renderFloat32Literal(v float32, w *Writer) {
	w.Write(strconv.FormatFloat(float64(v), 'f', -1, 10))
}

func renderFloat64Literal(v float64, w *Writer) {
	w.Write(strconv.FormatFloat(v, 'f', -1, 10))
}

func renderStringLiteral(str string, w *Writer) {
	w.Write("\"")
	for _, c := range str {
		if c == '"' || c == '\\' {
			w.Write("\\")
		}
		w.Write(string(c))
	}
	w.Write("\"")
}
