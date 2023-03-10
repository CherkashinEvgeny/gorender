package renderer

type ChanDir int

const (
	SendAndReceive ChanDir = 0
	Receive        ChanDir = 1
	Send           ChanDir = 2
)

type ChanRenderer struct {
	dir  ChanDir
	elem Code
	ctx  Code
}

func Chan(dir ChanDir, elem Code) *ChanRenderer {
	c := &ChanRenderer{}
	c.SetDir(dir)
	c.SetElem(elem)
	return c
}

func (c *ChanRenderer) GetDir() ChanDir {
	return c.dir
}

func (c *ChanRenderer) SetDir(dir ChanDir) {
	c.dir = dir
}

func (c *ChanRenderer) GetElem() Code {
	return c.elem
}

func (c *ChanRenderer) SetElem(elem Code) {
	c.elem = elem
	if elem != nil {
		elem.setContext(c)
	}
}

func (c *ChanRenderer) getContext() Code {
	return c.ctx
}

func (c *ChanRenderer) setContext(ctx Code) {
	c.ctx = ctx
}

func (c *ChanRenderer) render(w *Writer) {
	switch c.dir {
	case Receive:
		w.Write("<-chan ")
	case Send:
		w.Write("chan<- ")
	default:
		w.Write("chan ")
	}
	c.elem.render(w)
}
