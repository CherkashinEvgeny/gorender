package renderer

import (
	"strconv"
)

type LitRenderer struct {
	value any
}

func Lit(value any) LitRenderer {
	return LitRenderer{value}
}

func (l LitRenderer) getContext() Code {
	return nil
}

func (l LitRenderer) setContext(ctx Code) {
}

func (l LitRenderer) render(w *Writer) {
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
