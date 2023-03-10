package renderer

type CodeSequence interface {
	Code
	Len() int
	At(i int) Code
	Add(items ...Code)
}
