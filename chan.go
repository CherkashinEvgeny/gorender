package renderer

type ChanDir int

const (
	Receive ChanDir = 1
	Send    ChanDir = 2
)

type Chan struct {
	dir  ChanDir
	elem Code
	ctx  Code
}

func NewChan(dir ChanDir, elem Code) *Chan {
	c := &Chan{}
	c.SetDir(dir)
	c.SetElem(elem)
	return c
}

func (c *Chan) GetDir() ChanDir {
	return c.dir
}

func (c *Chan) SetDir(dir ChanDir) {
	c.dir = dir
}

func (c *Chan) GetElem() Code {
	return c.elem
}

func (c *Chan) SetElem(elem Code) {
	c.elem = elem
	if elem != nil {
		elem.setContext(c)
	}
}

func (c *Chan) getContext() Code {
	return c.ctx
}

func (c *Chan) setContext(ctx Code) {
	c.ctx = ctx
}

func (c *Chan) render(w *Writer) {
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
