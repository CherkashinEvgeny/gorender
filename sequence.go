package renderer

type Sequence interface {
	Code
	Len() int
	At(i int) Code
	Add(items ...Code)
}
